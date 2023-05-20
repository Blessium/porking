package model

import (
    "errors"
)

type ParkingSpot struct {
	ID           uint `gorm:"primaryKey"`
	CarParkID    uint
	Type         string 
	Size         string 
	Availability bool
}

type ParkingSpotRequest struct {
	Type         string `json:"type"`
	Size         string `json:"size"`
    Quantity     uint `json:"quantity"`
}

func (p * ParkingSpotRequest) ToParkingSpotModel() (*[]ParkingSpot, error) {
    if p.Quantity <= 0 {
        return nil, errors.New("Quantity not allowed")
    }

    spots := []ParkingSpot{}

    for i := 0; i < int(p.Quantity); i++ {
        var spot ParkingSpot
        spot.Type = p.Type
        spot.Size = p.Size
        spot.Availability = true
        spots = append(spots, spot)
    }

    return &spots, nil
}
