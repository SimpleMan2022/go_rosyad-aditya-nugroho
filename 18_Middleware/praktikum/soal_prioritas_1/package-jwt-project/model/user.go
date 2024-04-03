package model_package_jwt

import "time"

type User struct {
	Id        int       `gorm:"primaryKey;not null" json:"id"`
	Email     string    `gorm:"type:varchar(100);not null"json:"email"`
	Password  string    `gorm:"type:varchar(255);not null"json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

type UserResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
