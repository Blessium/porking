package handler

import (
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/service"
	"github.com/labstack/echo/v4"
	"net/http"
    "strconv"
)

type CarParkController struct {
    carParkService *service.CarParkService `di.inject:"carParkService"` 
}

func (ca *CarParkController) AddCarPark(c echo.Context) error {
	car := new(model.CarPark)

	if err := c.Bind(car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    if err := ca.carParkService.AddCarPark(car); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusCreated, car)
}

func (ca *CarParkController) GetCarPark(c echo.Context) error {

	id := c.Param("id")
    p_id, err := strconv.Atoi(id)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    car, err := ca.carParkService.GetCarPark(uint(p_id))
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusOK, car)
}

func (ca *CarParkController) GetAllCarParks(c echo.Context) error {
    ps, err := ca.carParkService.GetAllCarParks()

    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusOK, ps)
}

func (ca *CarParkController) UpdatedCarPark(c echo.Context) error {
	car := new(model.CarPark)

	id := c.Param("id")

	if err := c.Bind(car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    p_id, err := strconv.Atoi(id)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    ca.carParkService.UpdatedCarPark(car, uint(p_id))

	return c.JSON(http.StatusOK, car)
}
