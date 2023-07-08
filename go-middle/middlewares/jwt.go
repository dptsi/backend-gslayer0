package middlewares

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gslayer0/go-middle/config"
	"github.com/gslayer0/go-middle/helpers"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			response := map[string]string{"message": err.Error()}
			helpers.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		tokenString := c.Value
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrSignatureInvalid):
				response := map[string]string{"message": "Unauthorized"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case errors.Is(err, jwt.ErrTokenExpired):
				response := map[string]string{"message": "Token expired"}
				helpers.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helpers.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)

	})
}
