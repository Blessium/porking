package handler

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddParkingSpot(c echo.Context) error {

	id := c.Param("id")

	db, err := database.ConnectDatabase()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	p := new(model.ParkingSpot)
	car := new(model.CarPark)

	result := db.Limit(1).First(&car, id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Car park non trovato")
	}

	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db.Model(car).Association("ParkingSpots").Append(p)

	return c.JSON(http.StatusCreated, car)
}
func GetAllParkingSpots(c echo.Context) error {

	id := c.Param("id")

	var p []model.ParkingSpot

	db, err := database.ConnectDatabase()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	car := new(model.CarPark)

	result := db.Limit(1).First(&car, id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Car park non trovato")
	}

    
    db.Model(&car).Association("ParkingSpots").Find(&p)

	return c.JSON(http.StatusCreated, p)
}
