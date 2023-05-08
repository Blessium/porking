package handler

import (
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ParkingSpotController struct {
	parkingSpotService *service.ParkingSpotService `di.inject:"parkingSpotService"`
}

func (p *ParkingSpotController) AddParkingSpot(c echo.Context) error {

	id := c.Param("id")

	pa := new(model.ParkingSpot)

	if err := c.Bind(pa); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	p_id, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := p.parkingSpotService.AddParkingSpot(uint(p_id), pa); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "created")
}
func (p *ParkingSpotController) GetAllParkingSpots(c echo.Context) error {

	id := c.Param("id")
	p_id, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    ps, err := p.parkingSpotService.GetAllParkingSpots(uint(p_id))
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusCreated, ps)
}
