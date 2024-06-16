package common

import (
	"encoding/json"
	"finance/database"
	"net/http"
)

func ReadJSON(r *http.Request, data any) error {

	return json.NewDecoder(r.Body).Decode(&data)
}

func WriteJSON(w http.ResponseWriter, status int, response any) error {

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, message string, status int) {

	http.Error(w, message, status)

}

func SaveToDB(data any) error {

	database.DB.Create(data)

	return nil
}

func ValidateFields(w http.ResponseWriter, fields []string, data map[string]string) {

	for _, field := range fields {

		if _, ok := data[field]; !ok {
			ErrorResponse(w, field+"is required", http.StatusBadRequest)
		}

	}

}
