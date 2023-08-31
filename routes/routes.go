package routes

import (
	"net/http"

	"github.com/gagliardetto/solana-go"
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

	return account.PrivateKey.String(), account.PublicKey().String()

	//// Create a new RPC client:
	//client := rpc.New(rpc.TestNet_RPC)
	//
	//// Airdrop 5 SOL to the new account:
	//out, err := client.RequestAirdrop(
	//  context.TODO(),
	//  account.PublicKey(),
	//  solana.LAMPORTS_PER_SOL*1,
	//  rpc.CommitmentFinalized,
	//)
	//if err != nil {
	//  panic(err)
	//}
	//fmt.Println("airdrop transaction signature:", out)
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
