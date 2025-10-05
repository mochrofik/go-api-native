package models

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Register struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyProfile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
