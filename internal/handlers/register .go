package handlers

import (
	"errors"
	"finance/common"
	"finance/database"
	"finance/enums"
	"finance/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(data map[string]string, cost int) ([]byte, error) {
	if _, ok := data[enums.Password]; !ok {
		return nil, errors.New("password key not found in map")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data[enums.Password]), cost)
	if err != nil {
		log.Println("Error", err)
	}

	return hashedPassword, nil
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := common.ReadJSON(r, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	requiredFields := []string{enums.Name, enums.UserName, enums.Password}

	common.ValidateFields(w, requiredFields, data)

	password, _ := hashPassword(data, 14)

	user := models.User{
		Name:     data[enums.Name],
		Email:    data[enums.Email],
		Username: data[enums.UserName],
		Password: password,
	}
	result := database.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusCreated, &user)
}
