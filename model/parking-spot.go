package model

import "gorm.io/gorm"

type ParkingSpot struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
    CarParkID    uint     
	Type         string `json:"type"`
	Size         string `json:"size"`
	Availability string `json:"availability"`
}
