package handlers

import (
	"encoding/json"
	"errors"
	"finance/common"
	"finance/database"
	"finance/enums"
	"finance/models"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	UserID     string  `gorm:"not null" json:"user_id"`
	Name       string  `gorm:"not null" json:"name"`
	Amount     float64 `gorm:"not null" json:"amount"`
	Balance    float64 `gorm:"not null" json:"balance"`
	Bank       string  `gorm:"not null" json:"bank"`
	Currency   string  `gorm:"not null" json:"currency"`
	Type       string  `gorm:"not null" json:"type"`
	CreateDate string  `gorm:"autoCreateTime" json:"createdate"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewTransaction(data map[string]string, amount, balance float64) *models.Transaction {
	return &models.Transaction{
		UserID:   data["userId"],
		Name:     data["name"],
		Amount:   amount,
		Balance:  balance,
		Bank:     data["bank"],
		Currency: data["currency"],
		Type:     data["type"],
	}
}

func getBalance(data map[string]string, balance float64) error {
	if err := database.DB.Table("transactions").Select("balance").Where("user_Id = ?", data["userId"]).Scan(&balance).Error; err != nil {
		return err
	}
	return nil
}

func HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction
	id := r.PathValue("id")

	if err := database.DB.Delete(&transaction, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorResponse(w, "id not found", http.StatusNotFound)

		}
	}

}

func transactionType(transaction *models.Transaction, balance float64) ResponseMessage {

	var response ResponseMessage

	switch transaction.Type {

	case "deposit":

		transaction.Balance = balance + transaction.Amount

		log.Printf("user id: %v, name: %v, amount: %v", transaction.UserID, transaction.Name, transaction.Amount)

	case "withdrawal":

		if transaction.Balance < transaction.Amount {
			response = ResponseMessage{
				Message: "not enough funds",
			}
		}

		transaction.Balance = balance - transaction.Amount
		transaction.Amount = -transaction.Amount

		log.Printf("user id: %v, name: %v, amount: %v", transaction.UserID, transaction.Name, transaction.Amount)

	}

	return response

}

func HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var balance float64
	var user models.User
	var data map[string]string

	if err := common.ReadJSON(r, &data); err != nil {
		common.ErrorResponse(w, "Invalid", http.StatusBadRequest)
		return
	}

	requiredFields := []string{enums.UserId, enums.Amount, enums.Name, enums.Type, enums.Currency, enums.Bank}

	common.ValidateFields(w, requiredFields, data)

	if err := database.DB.First(&user, "id = ?", data[enums.UserId]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			common.ErrorResponse(w, "user does not exist", http.StatusNotFound)
		} else {
			common.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := database.DB.Table("transactions").Where("user_Id = ?", data[enums.UserId]).Group(enums.UserId).Pluck("SUM(amount)", &balance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorResponse(w, "balance not found", http.StatusNotFound)
		}
	}

	amount, err := strconv.ParseFloat(data[enums.Amount], 64)

	if err != nil {
		common.ErrorResponse(w, "invalid amount", http.StatusBadRequest)
		return
	}

	transaction := NewTransaction(data, amount, balance)
	response := transactionType(transaction, balance)

	if err := common.SaveToDB(&transaction); err != nil {
		common.ErrorResponse(w, "error saving transaction", http.StatusBadRequest)
	}
	common.WriteJSON(w, http.StatusOK, response)
}

func getAllTransactions() ([]Transaction, error) {

	var transaction []Transaction
	// user,_ := GetUser()

	if err := database.DB.Find(&transaction).Error; err != nil {
		return transaction, errors.New("no user found")
	}
	return transaction, nil

}

func HandleGetAllTransactions(w http.ResponseWriter, r *http.Request) {

	transaction, err := getAllTransactions()

	if err != nil {
		common.ErrorResponse(w, "transaction not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		log.Printf("Error encoding response %v", err)
		common.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
	}

}
