package models

import "time"

type User struct {
	ID        uint   `json:"id"; gorm:"primary_key"`
	Username  string `json:"username"; gorm:"unique" form:"username" binding:"required"`
	Password  string `json:"password"; form:"password" binding:"required"`
	Level     string `json:"level"; gorm:"default:normal"`
	CreatedAt time.Time `json:"time"`
}
