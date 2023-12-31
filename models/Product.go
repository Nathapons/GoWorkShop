package models

import "time"

type Product struct {
	ID        uint `gorm:"primary_key"; json:"id"`
	Name      string `json:"name"`
	Stock     int64 `json:"stock"`
	Price     float64 `json:"price"`
	Image     string `json:"image"`
	CreatedAt time.Time `json:time"`
}
