package routes

import (
	"context"
	"encoding/json"
	"fmt"
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
	DB *gorm.DB
}

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	h := &Handler{DB: db}

	e.GET("/", root)
	e.GET("/hello", hello, middleware.LoggerMiddleware, middleware.SessionMiddleware()) // Add middleware here
	e.POST("/users", h.addUser, middleware.LoggerMiddleware)
	e.GET("/users/:id", h.showUser, middleware.LoggerMiddleware)
	e.POST("/pay", h.sendPayment, middleware.LoggerMiddleware)
	// e.PUT("/users/:id", updateUser)
	/*Goods*/
	e.POST("/goods", h.createGood)
	e.GET("/goods", h.getGoods)
	e.PUT("/goods/:id", h.updateGood)
	e.DELETE("/goods/:id", h.deleteGood)
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
	client := rpc.New(rpc.DevNet_RPC)

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
	ExternalWallet *string `json:"external_wallet,omitempty"`
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

	// Create a new RPC client:
	rpcClient := rpc.New(rpc.DevNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)

	}

	// Load the account that you will send funds FROM:
	// accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
	accountFrom, err := solana.PrivateKeyFromBase58(sourceUser.PrivateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("private key:", accountFrom.String())
	fmt.Println("public key:", accountFrom.PublicKey().String())

	// The public key of the account that you will send sol TO:
	accountTo := solana.MustPublicKeyFromBase58(targetUser.Wallet)

	recent, err := rpcClient.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}
	amount := uint64(1)
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

	// FIXME TODO
	// update BoughtProducts for those who buy

	return nil
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
