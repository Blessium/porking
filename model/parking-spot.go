package model

type ParkingSpot struct {
	ID           uint `gorm:"primaryKey"`
	CarParkID    uint
	Type         string `json:"type"`
	Size         string `json:"size"`
	Availability string `json:"availability"`
}
