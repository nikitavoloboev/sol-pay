package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func SessionMiddleware() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("hanko")
			if err == http.ErrNoCookie {
				return c.Redirect(http.StatusTemporaryRedirect, "/unauthorized")
			}
			if err != nil {
				return err
			}
			hankoApiURL := "https://e879ccc9-285e-49d3-b37e-b569f0db4035.hanko.io"
			// replace "hankoApiURL" with your API URL
			set, err := jwk.Fetch(
				context.Background(),
				fmt.Sprintf("%v/.well-known/jwks.json", hankoApiURL),
			)
			if err != nil {
				return err
			}

			token, err := jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/unauthorized")
			}

			log.Printf("session for user '%s' verified successfully", token.Subject())

			c.Set("token", cookie.Value)
			c.Set("user", token.Subject())

			return next(c)
		}
	}
}
