package handler

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddCarPark(c echo.Context) error {
	car := new(model.CarPark)

	if err := c.Bind(car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db.Save(&car)

	return c.JSON(http.StatusCreated, car)
}

func GetCarPark(c echo.Context) error {

	id := c.Param("id")

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
		return c.String(http.StatusNotFound, "Utente non trovato")
	}
    


	return c.JSON(http.StatusOK, car)
}

func GetAllCarParks(c echo.Context) error {
	var car_parks []model.CarPark

	db, err := database.ConnectDatabase()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := db.Find(&car_parks)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, car_parks)
}

func UpdatedCarPark(c echo.Context) error {
	car := new(model.CarPark)

	id := c.Param("id")

	db, err := database.ConnectDatabase()

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	old_car := new(model.CarPark)
	result := db.Limit(1).First(&old_car, id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")
	}

	if err := c.Bind(car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	car.ID = old_car.ID

	db.Save(&car)
	return c.JSON(http.StatusOK, car)
}
