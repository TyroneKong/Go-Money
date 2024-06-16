package middleware

import (
	"finance/common"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		if err != nil {
			common.ErrorResponse(w, "unauthorized", http.StatusUnauthorized)

			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("API_SECRET")), nil

		})

		if err != nil || !token.Valid {
			common.ErrorResponse(w, "unauthorized", http.StatusUnauthorized)

			log.Printf("token is invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	headers := map[string]string{"Access-Control-Allow-Origin": "http://localhost:3000", "Access-Control-Allow-Credentials": "true", "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS", "Access-Control-Allow-Headers": "Origin, Content-Type, Accept"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for k, v := range headers {

			w.Header().Set(k, v)

		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
