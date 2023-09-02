package main

import (
	"log"
	"net/http"

	"github.com/jub0bs/fcors"
	"github.com/labstack/echo/v4"
	"github.com/nikitavoloboev/sol-pay/model"
	"github.com/nikitavoloboev/sol-pay/routes"
)

func main() {
	// Set up the database
	db, err := model.SetupDatabase()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	defer db.Close()

	// Migrate the database to ensure the user table exists
	model.MigrateDB(db)

	e := echo.New()
	cors, err := fcors.AllowAccess(
		fcors.FromOrigins("http://localhost:8080"),
		fcors.WithMethods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		),
		fcors.WithRequestHeaders("Authorization"),
	)
	if err != nil {
		log.Fatal(err)
	}
	e.Use(echo.WrapMiddleware(cors))

	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
