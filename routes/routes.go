package routes

import (
	"net/http"

	"github.com/sol-pay/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", root)
	e.GET("/hello", hello, middleware.LoggerMiddleware) // Add middleware here
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo!")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Echo with Middleware!")
}
