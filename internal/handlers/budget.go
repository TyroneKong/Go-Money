package handlers

import (
	"errors"
	"finance/common"
	"finance/database"
	"finance/enums"
	"finance/models"
	"log"
	"net/http"
	"strconv"
)

type Budget struct {
	ID        string  `gorm:"primary key;autoIncrement" json:"id"`
	UserID    string  `gorm:"not null" json:"userid"`
	Name      string  `gorm:"not null" json:"name"`
	Amount    float64 `gorm:"not nul" json:"amount"`
	Type      string  `gorm:"not null" json:"type"`
	CreatedAt string  `gorm:"autoCreateTime" json:"created_at"`
}

func NewBudget(data map[string]string, amount float64) *models.Budget {

	return &models.Budget{
		UserID: data["userId"],
		Name:   data["name"],
		Amount: amount,
		Type:   data["type"],
	}

}

func HandleCreateBudget(w http.ResponseWriter, r *http.Request) {
	var user User
	var data map[string]string

	var response ResponseMessage

	if err := common.ReadJSON(r, &data); err != nil {
		common.ErrorResponse(w, err.Error(), http.StatusNotFound)

	}

	requiredFields := []string{enums.UserId, enums.Name, enums.Amount, enums.Type}

	common.ValidateFields(w, requiredFields, data)

	amount, _ := strconv.ParseFloat(data[enums.Amount], 64)

	budget := NewBudget(data, amount)
	if err := common.SaveToDB(&budget); err != nil {
		common.ErrorResponse(w, "error saving to database", http.StatusInternalServerError)
	}

	response = ResponseMessage{
		Message: "budget created successfully",
	}

	log.Printf("budget has been created for user %v", user.Name)
	w.Header().Set("content-Type", "application/json")
	common.WriteJSON(w, http.StatusOK, response)

}

func getAllBudgets() ([]Budget, error) {
	var budget []Budget
	if err := database.DB.Find(&budget).Error; err != nil {
		return budget, errors.New("unable to retrieve budgets")
	}

	return budget, nil

}

func HandleGetAllBudgets(w http.ResponseWriter, r *http.Request) {
	budget, err := getAllBudgets()

	if err != nil {
		common.ErrorResponse(w, err.Error(), http.StatusNotFound)

		return
	}

	common.WriteJSON(w, http.StatusOK, budget)
}
