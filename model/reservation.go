package model

import "time"

type Reservation struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
    StartTime     time.Time `json:"start_time"`
    EndTime       time.Time `json:"end_time"`
    Cost          float32 `json:"cost"`
	QRCodePath    string
    ParkingSpotID uint
	ParkingSpot   ParkingSpot
	UserID        uint
	User          User
}
