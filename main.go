package main

import (
	"log"

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

	e := echo.New()

	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
