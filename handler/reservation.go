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

    var user model.User

	result := db.Limit(1).First(&user, user_id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")
	}

	if err := c.Bind(&r); err != nil {
		fmt.Println("Bruuh")
		return c.String(http.StatusBadRequest, err.Error())
	}

    qr_path, err := utils.GenerateQR(&r)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    re := r.ConvertToReservation()
	re.UserID = user_id
    re.QRCodePath = "http://localhost:1234/qr/" + qr_path
	db.Save(&re)

    if err := utils.SendEmail(user.Email,&user, &re, qr_path); err != nil {
        return c.String(http.StatusInternalServerError, err.Error());
    }
	return c.JSON(http.StatusCreated, re)
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
