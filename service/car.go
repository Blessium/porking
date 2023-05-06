package service

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
)

type CarService struct {
}

func (u *CarService) GetCarById(user_id uint, id uint) (*model.Car, error) {

	car := new(model.Car)

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	if err := db.Model(&model.User{ID: user_id}).Association("Cars").Find(&car, id); err != nil {
         return nil, err
	}

	return car, nil
}

func (c *CarService) GetCarsById(user_id uint) (*[]model.Car, error) {
    cars := new([]model.Car)

	db, err := database.ConnectDatabase()
	if err != nil {
        return nil, err
	}

	if err := db.Model(&model.User{ID: user_id}).Association("Cars").Find(&cars); err != nil {
         return nil, err
	}

    return cars, nil
}

func (c *CarService) AddCar(car *model.Car) error {
	db, err := database.ConnectDatabase()

	if err != nil {
		return err
	}

	db.Save(car)
	return nil
}

func (u *CarService) UpdateCar(new_car *model.Car, user_id uint, id uint) error {
	_, err := u.GetCarById(user_id, id)
	if err != nil {
		return err
	}

	db, err := database.ConnectDatabase()

	if err != nil {
		return err
	}

    new_car.ID = id
    new_car.UserID = user_id

	db.Save(new_car)
	return nil
}
