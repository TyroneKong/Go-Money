package handlers

import (
	"encoding/json"
	"errors"
	"finance/common"
	"finance/database"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint   `json:"id" gorm:"PrimaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func HandleGetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve JWT from the cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		common.ErrorResponse(w, "Unauthorized: missing or invalid cookie", http.StatusUnauthorized)

		return
	}

	jwtToken := cookie.Value
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil || !token.Valid {
		common.ErrorResponse(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		log.Printf("Error parsing token: %v", err)
		return
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		common.ErrorResponse(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
		log.Printf("Error asserting claims")
		return
	}

	issuer := claims.Issuer

	log.Printf("issuer %v", issuer)

	if issuer == "" {
		common.ErrorResponse(w, "Unauthorized: missing issuer", http.StatusUnauthorized)
		log.Printf("Issuer is empty in claims")
		return
	}

	var user User
	if err := database.DB.Where("id = ?", issuer).First(&user).Error; err != nil {
		common.ErrorResponse(w, "User not found", http.StatusNotFound)
		log.Printf("Error fetching user from database: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		common.ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error encoding user response: %v", err)
	}
}
