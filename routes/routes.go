package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
	e.POST("/pay", h.sendPayment, middleware.LoggerMiddleware)
	// e.PUT("/users/:id", updateUser)
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo!")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Echo with Middleware!")
}

func generateWallet() (string, string) {
	//return "walletAddress"

	// Create a new account:
	account := solana.NewWallet()
	//fmt.Println("account private key:", account.PrivateKey)
	//fmt.Println("account public key:", account.PublicKey())

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

func GetUserByID(db *gorm.DB, userID uint) (*model.User, error) {
	var user model.User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type UserIDInput struct {
	SourceUserID uint `json:"source_user_id"`
	TargetUserID uint `json:"target_user_id"`
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

	/*FIXME - create it from db details */
	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

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
