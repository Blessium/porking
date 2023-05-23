package service

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
)

type ParkingSpotService struct {
}

func (p *ParkingSpotService) AddParkingSpots(id uint, m *model.ParkingSpotRequest) error {

	car, err := new(CarParkService).GetCarPark(id)
	if err != nil {
		return err
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return err
	}

    spots, err := m.ToParkingSpotModel()
    if err != nil {
        return err
    }

	db.Model(car).Association("ParkingSpots").Append(spots)

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
    if err := db.Model(&car).Association("ParkingSpots").Find(pe); err != nil {
        return nil, err
    }

	return pe, nil
}

func (p *ParkingSpotService) GetReservationParkingSpot(car_park_id uint) (*model.ParkingSpot, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	car, err := new(CarParkService).GetCarPark(car_park_id)
	if err != nil {
		return nil, err
	}

	pe := new(model.ParkingSpot)

    if err := db.Model(&car).Association("ParkingSpots").Find(pe, "availability = true"); err != nil {
        return nil, err
    }
    return pe, nil
}
