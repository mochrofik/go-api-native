package authcontroller

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var register *models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Username == "" || register.Email == "" || register.Password == "" || register.PasswordConfirm == "" {
		helper.Response(w, 400, "Input correctly data", nil)
		return
	}

	if register.Password != register.PasswordConfirm {
		helper.Response(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := helper.HashPassword(register.Password)

	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Username:  register.Username,
		Email:     register.Email,
		Password:  passwordHash,
		Status:    1, //1 Aktif, 0 Tidak aktif
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Register successfully", nil)

}

func Login(w http.ResponseWriter, r *http.Request) {

	var login *models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var user *models.User

	if login.Username == "" || login.Password == "" {
		helper.Response(w, 400, "Input data correctly", nil)
		return
	}

	if err := config.DB.First(&user, "username = ?", login.Username).Error; err != nil {
		helper.Response(w, 404, "wrong email or password", nil)
		return
	}

	if err := helper.VerifyPassword(user.Password, login.Password); err != nil {
		helper.Response(w, 404, "wrong email or password", nil)
		return
	}

	helper.Response(w, 200, "login successfully", nil)

}
