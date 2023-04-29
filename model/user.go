package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
    ID    uint `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
