package model_package

import "time"

type Package struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"`
	Sender         string    `gorm:"type:varchar(100);not null" json:"sender"`
	Receiver       string    `gorm:"type:varchar(100);not null" json:"receiver"`
	SenderLocation string    `gorm:"type:varchar(100);not null" json:"sender_location"`
	SenderReceiver string    `gorm:"type:varchar(100);not null" json:"sender_receiver"`
	Fee            int       `gorm:"type:int(11);not null" json:"fee"`
	Weight         float64   `json:"weight"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PackageRequest struct {
	Name           string  `json:"name"`
	Sender         string  `json:"sender"`
	Receiver       string  `json:"receiver"`
	SenderLocation string  `json:"sender_location"`
	SenderReceiver string  `json:"sender_receiver"`
	Fee            int     `json:"fee"`
	Weight         float64 `json:"weight"`
}

type PackageResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
