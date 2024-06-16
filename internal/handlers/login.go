package handlers

import (
	"errors"
	"finance/common"
	"finance/database"
	"finance/enums"
	"finance/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func checkUser(data map[string]string, user models.User) error {

	if err := database.DB.Where("username = ?", data[enums.UserName]).First(&user).Error; err != nil {
		return fmt.Errorf("cannot find user")
	}

	return nil
}

func compareHashedPassword(data map[string]string, user models.User) error {
	if _, ok := data[enums.Password]; !ok {
		return errors.New("password key not found in map")
	}

	checkUser(data, user)

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data[enums.Password])); err != nil {

		return fmt.Errorf("incorrect login details")
	}

	return nil
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := common.ReadJSON(r, &data); err != nil {
		common.ErrorResponse(w, err.Error(), http.StatusBadRequest)

	}

	var user models.User

	if err := database.DB.Where("username = ?", data[enums.UserName]).First(&user).Error; err != nil {
		common.ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	if err := compareHashedPassword(data, user); err != nil {
		common.ErrorResponse(w, err.Error(), http.StatusNotFound)

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, _ := claims.SignedString([]byte(os.Getenv("API_SECRET")))

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 48),
		// HTTPOnly: true,
	})

	common.WriteJSON(w, http.StatusOK, &user)
}

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Path:    "/",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	})

	common.WriteJSON(w, http.StatusOK, "successfully logged out")

}
