package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	if err != nil {
		log.Fatal(err)
	}

	// Add CORS middleware - Allow requests from any origin with the methods needed by your application
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
