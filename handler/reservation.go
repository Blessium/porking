package handler

import (
	"fmt"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateReservation(c echo.Context) error {
	var r model.ReservationRequest

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&r); err != nil {
		fmt.Println("Bruuh")
		return c.String(http.StatusBadRequest, err.Error())
	}

    re := r.ConvertToReservation()
	re.UserID = user_id
    fmt.Println("What the fuck")
	db.Save(&re)
	return c.String(http.StatusCreated, "Created")
}

func GetAllReservations(c echo.Context) error {
	var r []model.Reservation

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := db.Model(&model.User{ID: user_id}).Association("Reservations").Find(&r); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusFound, r)
}
