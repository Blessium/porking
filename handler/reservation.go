package handler

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateReservation(c echo.Context) error {

	return c.String(http.StatusCreated, "Created")
}

func GetAllReservations(c echo.Context) error {
	var r []model.Reservation

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := db.Find(&r)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")

	}
	return c.JSON(http.StatusFound, r)
}
