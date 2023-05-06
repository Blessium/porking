package service

import (
	"errors"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
)

type CarParkService struct {
}

func (c *CarParkService) AddCarPark(p *model.CarPark) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		return err
	}

	db.Save(p)
	return nil
}

func (c *CarParkService) GetCarPark(id uint) (*model.CarPark, error) {
	p := new(model.CarPark)

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	result := db.Limit(1).First(p, id)

	if result.Error != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Car park non trovato")
	}

	return p, nil
}

func (c *CarParkService) GetAllCarParks() (*[]model.CarPark, error) {

	ps := new([]model.CarPark)
	db, err := database.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	result := db.Find(ps)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Nessun car park trovato")
	}
	return ps, nil
}

func (c *CarParkService) UpdatedCarPark(p *model.CarPark, id uint) error {
    _, err := c.GetCarPark(id)  
    if err != nil {
        return err
    }
    p.ID = id
    
    if err := c.AddCarPark(p); err != nil {
        return err
    }
    return nil
}

