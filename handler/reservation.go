package handler

import (
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/utils"
	"github.com/blessium/porking/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReservationController struct {
    reservationService *service.ReservationService `di.inject:"reservationService"`    
}

func (r *ReservationController) CreateReservation(c echo.Context) error {
    res := new(model.ReservationRequest)

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    if err := c.Bind(res); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    f_res, err := r.reservationService.CreateReservation(res, user_id)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusCreated, f_res)
}

func (r *ReservationController) GetAllReservations(c echo.Context) error {
	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    res, err := r.reservationService.GetAllReservations(user_id) 
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusFound, res)
}
