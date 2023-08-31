package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Logged %s from %s at %s", c.Request().Method, c.Request().RemoteAddr, time.Now().Format(time.RFC3339))
		return next(c)
	}
}

