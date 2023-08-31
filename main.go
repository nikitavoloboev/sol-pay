package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwk"
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

// var tokenAuth *jwtauth.JWTAuth
//
//	func init() {
//		// You can use any string value as a "secret key". In production, use a more secure method to store and retrieve it.
//		tokenAuth = jwtauth.New("HS256", []byte("secret_key"), nil)
//	}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	r := chi.NewRouter()

	r.Use(SessionMiddleware)
	r.Get("/hello", helloHandler)

	// Public routes
	//r.Get("/auth", authHandler)

	//// Protected routes
	//r.Group(func(r chi.Router) {
	//	r.Use(jwtauth.Verifier(tokenAuth))
	//	r.Use(jwtauth.Authenticator)

	//	r.Get("/protected", protectedHandler)
	//})

	http.ListenAndServe(":8080", r)
}

//func authHandler(w http.ResponseWriter, r *http.Request) {
//	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
//
//	render.JSON(w, r, map[string]string{"token": tokenString})
//}
//
//func protectedHandler(w http.ResponseWriter, r *http.Request) {
//	_, claims, _ := jwtauth.FromContext(r.Context())
//	render.JSON(w, r, claims)
//}
//
