package service

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
)

type ParkingSpotService struct {
}

func (p *ParkingSpotService) AddParkingSpot(id uint, m *model.ParkingSpot) error {

	car, err := new(CarParkService).GetCarPark(id)
	if err != nil {
		return err
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return err
	}

	db.Model(car).Association("ParkingSpots").Append(m)

	return nil
}

func (p *ParkingSpotService) GetAllParkingSpots(id uint) (*[]model.ParkingSpot, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	car, err := new(CarParkService).GetCarPark(id)
	if err != nil {
		return nil, err
	}

	pe := new([]model.ParkingSpot)
	db.Model(&car).Association("ParkingSpots").Find(pe)

	return pe, nil
}
