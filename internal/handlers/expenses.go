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

type Expense struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	UserId    string  `gorm:"not null" json:"user_id"`
	Name      string  `gorm:"not null" json:"name"`
	Amount    float64 `gorm:"not null" json:"amount"`
	Total     float64 `gorm:"not null" json:"total"`
	Type      string  `gorm:"not null" json:"type"`
	CreatedAt string  `gorm:"autoCreateTime" json:"created_at"`
}

func NewExpense(data map[string]string, amount, total float64) *models.Expense {

	return &models.Expense{
		UserId: data["userId"],
		Name:   data["name"],
		Amount: amount,
		Total:  total,
		Type:   data["type"],
	}

}

func HandleCreateExpense(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var data map[string]string
	var total float64

	if err := common.ReadJSON(r, &data); err != nil {
		common.ErrorResponse(w, ("Invalid payload"), http.StatusBadRequest)

		return
	}

	requiredFields := []string{enums.UserId, enums.Name, enums.Amount, enums.Type}

	common.ValidateFields(w, requiredFields, data)

	if err := database.DB.Where("id = ?", data[enums.UserId]).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorResponse(w, "User not found", http.StatusNotFound)

			return
		}
		common.ErrorResponse(w, "database error", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Table("expenses").Where("user_Id = ?", data[enums.UserId]).Group(enums.UserId).Pluck("SUM(amount)", &total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	amount, _ := strconv.ParseFloat(data["amount"], 64)
	total += amount
	expense := NewExpense(data, amount, total)

	if err := database.DB.Create(&expense).Error; err != nil {
		common.ErrorResponse(w, "failed to create expense", http.StatusInternalServerError)
		return
	}
	common.WriteJSON(w, http.StatusOK, expense)

}

func getAllExpenses() ([]Expense, error) {

	var expense []Expense

	if err := database.DB.Find(&expense).Error; err != nil {
		return expense, errors.New("no expenses found")
	}

	return expense, nil
}

func HandleGetAllExpenses(w http.ResponseWriter, r *http.Request) {

	expense, err := getAllExpenses()

	log.Printf("the expenses are %v", expense)

	if err != nil {
		common.ErrorResponse(w, "expense not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&expense); err != nil {
		log.Printf("Error encoding response %v", err)
		common.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
	}

}
