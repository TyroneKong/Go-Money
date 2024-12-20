package validateFields

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateFields(t *testing.T) {
	tests := []struct {
		name       string
		fields     []string
		data       map[string]string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "All fields present",
			fields:     []string{"username", "email"},
			data:       map[string]string{"username": "john_doe", "email": "john@example.com"},
			wantStatus: http.StatusOK,
			wantBody:   "",
		},
		{
			name:       "Missing email field",
			fields:     []string{"username", "email"},
			data:       map[string]string{"username": "john_doe"},
			wantStatus: http.StatusBadRequest,
			wantBody:   "",
		},
		{
			name:       "No fields provided",
			fields:     []string{},
			data:       map[string]string{"username": "john_doe", "email": "john@example.com"},
			wantStatus: http.StatusOK,
			wantBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			ValidateFields(rr, tt.fields, tt.data)

			if tt.wantStatus != http.StatusOK {
				if rr.Code != tt.wantStatus {
					t.Errorf("got status %v, want %v", rr.Code, tt.wantStatus)
				}
			} else if rr.Body.String() != "" {

				t.Errorf("expected no response body, got %q", rr.Body.String())
			}

			if tt.wantBody != "" && !strings.Contains(rr.Body.String(), tt.wantBody) {
				t.Errorf("got body %q, want %q", rr.Body.String(), tt.wantBody)
			}
		})
	}
}
