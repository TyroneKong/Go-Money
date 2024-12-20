package validateFields

import "net/http"

func ErrorResponse(w http.ResponseWriter, message string, status int) {

	http.Error(w, message, status)

}

func ValidateFields(w http.ResponseWriter, fields []string, data map[string]string) {

	for _, field := range fields {

		if _, ok := data[field]; !ok {
			ErrorResponse(w, field+"is required", http.StatusBadRequest)
		}

	}

}
