package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
    "errors"
    "encoding/json"
)

type QR struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	Image     string
}

type QRInfo struct {
	StartTime     UnixTime `json:"start_time"`
	EndTime       UnixTime `json:"end_time"`
	Cost          float32  `json:"cost"`
	CarID         uint     `json:"car_id"`
	CarParkID     uint     `json:"car_park_id"`
	ParkingSpotID uint     `json:"parking_spot_id"`
}

type QRCert struct {
	Model QRInfo `json:"reservation_info"`
	Hash  string `json:"hash"`
	Firma string `json:"fingerprint"`
}

func (q *QR) BeforeCreate(scope *gorm.DB) error {
	q.ID = uuid.NewV4().String()
	return nil
}

func (q *QRCert) MarsharlQR() ([]byte, error) {
	qr_json, err := json.Marshal(&q)
	if err != nil {
		return nil, errors.New("Could not parse the json reservation request")
	}
	return qr_json, nil
}
