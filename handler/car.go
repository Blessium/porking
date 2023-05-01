package handler

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddCar(c echo.Context) error {
	var car model.Car

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	car.UserID = user_id

	db.Save(&car)

	return c.String(http.StatusCreated, "Added car")
}

func UpdateCar(c echo.Context) error {
	var car model.Car

	id := c.Param("id")

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var old_car model.Car
	result := db.Limit(1).Find(&old_car, id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Macchina non trovata")
	}

	db.Model(&old_car).Updates(&car)

	return c.String(http.StatusCreated, "Added car")
}

func GetCars(c echo.Context) error {
	var cars []model.Car

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := db.Model(&model.User{ID: user_id}).Association("Cars").Find(&cars); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusFound, cars)
}
