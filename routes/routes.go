package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nikitavoloboev/sol-pay/middleware"
	"github.com/nikitavoloboev/sol-pay/model"
)

type Handler struct {
	DB           *gorm.DB
	solanaClient *rpc.Client
}

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	h := &Handler{DB: db, solanaClient: rpc.New(rpc.DevNet_RPC)}

	e.GET("/", root)
	e.GET("/hello", hello, middleware.LoggerMiddleware, middleware.SessionMiddleware()) // Add middleware here
	/* Users */
	e.POST("/users", h.addUser, middleware.LoggerMiddleware)
	e.GET("/users/:id", h.showUser, middleware.LoggerMiddleware)
	// e.PUT("/users/:id", updateUser)
	/*Goods*/
	e.POST("/goods", h.createGood)
	e.GET("/goods", h.getGoods)
	e.GET("/goods/:id", h.getGoodDetails)
	e.PUT("/goods/:id", h.updateGood)
	e.DELETE("/goods/:id", h.deleteGood)
	/* Payment  */
	e.POST("/pay", h.sendPayment, middleware.LoggerMiddleware)
	e.POST("/balance", h.walletBalance, middleware.LoggerMiddleware)
	e.POST("/can-i-buy", h.canIBuyProduct, middleware.LoggerMiddleware)
	e.POST("/usdc-balance", h.getUsdcAccountBalance, middleware.LoggerMiddleware)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo!")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Echo with Middleware!")
}

func generateWallet() (string, string) {
	// Create a new account:
	account := solana.NewWallet()
	fmt.Println("account private key:", account.PrivateKey)
	fmt.Println("account public key:", account.PublicKey())

	// Create a new RPC client:
	client, _ := getRPCSolana()

	// Airdrop 5 SOL to the new account:
	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("airdrop transaction signature:", out)
	return account.PrivateKey.String(), account.PublicKey().String()
}

func (h *Handler) addUser(c echo.Context) error {
	user := model.User{}
	user.PrivateKey, user.Wallet = generateWallet()
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := model.CreateUser(h.DB, &user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) showUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := GetUserByID(h.DB, uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user details")
	}
	userResponse, _ := json.Marshal(user)
	userMap := make(map[string]interface{})
	json.Unmarshal(userResponse, &userMap)
	delete(userMap, "private_key")

	var products []model.ProductName
	products, err = model.GetProductsCountByUserID(h.DB, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch details"})
	}
	userMap["created_products"] = products

	var productsSold []model.ProductSoldCount
	productsSold, err = model.GetProductsSoldCountUserID(h.DB, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch details"})
	}
	userMap["sold_products"] = productsSold

	var productsBought []model.ProductName
	productsBought, err = model.GetProductsBoughtByUserID(h.DB, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch details"})
	}
	userMap["bought_products"] = productsBought
	return c.JSON(http.StatusOK, userMap)
}

func GetUserByID(db *gorm.DB, userID uint) (*model.User, error) {
	var user model.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type UserIDInput struct {
	SourceUserID   uint    `json:"source_user_id"`
	TargetUserID   uint    `json:"target_user_id"`
	ProductID      uint    `json:"product_id"`
	ExternalWallet *string `json:"external_wallet,omitempty"`
}

type UserIDAPIInput struct {
	UserID uint `json:"user_id"`
}
type UserIDProductIDAPIInput struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func getWalletBalance(wallet string) *big.Float {

	client, _ := getRPCSolana()

	pubKey := solana.MustPublicKeyFromBase58(wallet)
	out, err := client.GetBalance(
		context.TODO(),
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
	// Convert lamports to sol:
	var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))
	return solBalance
}

func (h *Handler) walletBalance(c echo.Context) error {
	var input UserIDAPIInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sourceUser, err := GetUserByID(h.DB, input.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Print(sourceUser.Wallet)
	solBalance := getWalletBalance(sourceUser.Wallet)
	balanceuint32, _ := solBalance.Float32()
	balanceInUsd := balanceuint32 * currentSolanaPriceInUSD()
	balanceInUsdStr := fmt.Sprintf("%.2f", balanceInUsd)

	return c.JSON(http.StatusOK, map[string]string{
		"balance":     solBalance.Text('f', 10),
		"balance_usd": balanceInUsdStr,
	})

}

func currentSolanaPriceInUSD() float32 {
	return 19.59
}

func currentUSDPriceInSOL() float32 {
	return 3.61
}

func (h *Handler) checkBalance(productId uint, userId uint) bool {
	sourceUser, _ := GetUserByID(h.DB, userId)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//}
	fmt.Print(sourceUser.Wallet)
	solBalance := getWalletBalance(sourceUser.Wallet)

	product, _ := model.GetGoodsByID(h.DB, productId)

	balance, _ := solBalance.Float32()

	balanceInUsd := balance * currentSolanaPriceInUSD()

	return balanceInUsd >= float32(product.Price)
}

func (h *Handler) getUsdcAccountBalance(c echo.Context) error {
	tokenAddress, _, _ := solana.FindAssociatedTokenAddress(solana.MustPublicKeyFromBase58("3GDCH1RDYFn96ZjPn3QL4BE52t53CmhfqEWUEd3wSTco"), solana.MustPublicKeyFromBase58("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU"))
	fmt.Println(tokenAddress)

	accountBalance, err := h.solanaClient.GetTokenAccountBalance(context.Background(), tokenAddress, rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	// h.getAccountBalance(c echo.Context)

	// 66.3
	return c.JSON(http.StatusOK, map[string]string{"usdc_balance": accountBalance.Value.UiAmountString})
}

func (h *Handler) canIBuyProduct(c echo.Context) error {
	var input UserIDProductIDAPIInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	productId := input.ProductID
	canIBuyIt := h.checkBalance(productId, input.UserID)

	return c.JSON(http.StatusOK, map[string]bool{"can_pay": canIBuyIt})
}

func getRPCSolana() (*rpc.Client, error) {
	return rpc.New(rpc.DevNet_RPC), nil
	// return rpc.New(rpc.MainNetBeta_RPC), nil // - this is a REAL ONE :scream:
}

func (h *Handler) sendPayment(c echo.Context) error {
	var input UserIDInput
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sourceUser, err := GetUserByID(h.DB, input.SourceUserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	targetUser, err := GetUserByID(h.DB, input.TargetUserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	productId := input.ProductID
	product, err := model.GetGoodsByID(h.DB, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	//sourceUserBalance := getWalletBalance(sourceUser.Wallet)

	/* do source user have enough balance to cover product cost and transaction fee */
	if !h.checkBalance(productId, input.SourceUserID) {
		return echo.NewHTTPError(http.StatusInternalServerError, "not enough balance")
	}

	// Create a new RPC client:
	rpcClient, _ := getRPCSolana()

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}

	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromBase58(sourceUser.PrivateKey)
	if err != nil {
		panic(err)
	}
	//fmt.Println("private key:", accountFrom.String())
	fmt.Println("SOURCE Wallet:", accountFrom.PublicKey().String())
	fmt.Println("TARGET Wallet:", targetUser.Wallet)

	/* Good to go - trying to do transaction */

	// The public key of the account that you will send sol TO:
	accountTo := solana.MustPublicKeyFromBase58(targetUser.Wallet)

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	// amount := uint64(1) FIXME NOTE - can be usefull for debug
	amount := uint64(float32(product.Price) / currentUSDPriceInSOL() * float32(math.Pow(10, 9)))
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				accountFrom.PublicKey(),
				accountTo,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)

	err = model.InsertUserGood(h.DB, sourceUser.ID, product.ID)
	if err != nil {
		fmt.Println("Error inserting record:", err)
	}
	return c.JSON(http.StatusOK, map[string]string{"tix": sig.String()})
}

/*
func updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic to update the user with the given ID.
	// For demonstration, just updating the ID with the one from the path.
	//user.ID = id

	return c.JSON(http.StatusOK, user)
}
*/

func (h *Handler) createGood(c echo.Context) error {
	good := &model.Product{}
	if err := c.Bind(good); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := model.CreateGood(h.DB, good); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, good)
}

func (h *Handler) getGoods(c echo.Context) error {
	// Get user ID from the query parameter
	userIDParam := c.QueryParam("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	goods, err := model.GetGoodsByUserID(h.DB, uint(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, goods)
}

func (h *Handler) updateGood(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	good := &model.Product{}
	if err := c.Bind(good); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	good.ID = uint(id)
	if err := h.DB.Save(&good).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, good)
}

func (h *Handler) getGoodDetails(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	productDetails, err := model.GetProductsBoughtCountByProductID(h.DB, uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// FIXME NOTE: bind error for marchalling from ORM to JSON
	// if err := c.Bind(productDetails); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	return c.JSON(http.StatusOK, productDetails)
}

func (h *Handler) deleteGood(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.DB.Delete(&model.Product{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
