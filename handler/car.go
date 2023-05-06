package handler

import (
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/utils"
	"github.com/blessium/porking/service"
	"github.com/labstack/echo/v4"
	"net/http"
    "github.com/goioc/di"
    "strconv"
)

type CarController struct {
    carService *service.CarService `di.inject:"carService"`
}

func (ca *CarController) AddCar(c echo.Context) error {

    car := new(model.Car)

	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	car.UserID = user_id

    if err := ca.carService.AddCar(car); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.String(http.StatusCreated, "Added car")
}

func (ca *CarController) UpdateCar(c echo.Context) error {

    car := new(model.Car)
    
    id := c.Param("id")

	user_id, err := utils.Extract_id_from_token(c)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }


	if err := c.Bind(&car); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
    
    car_id, err := strconv.Atoi(id)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    if err := ca.carService.UpdateCar(car, user_id, uint(car_id)); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }
    

	return c.String(http.StatusCreated, "Added car")
}

func (ca *CarController) GetCars(c echo.Context) error {
	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    cars, err := ca.carService.GetCarsById(user_id) 
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusFound, cars)
}

func (ca *CarController) GetCar(c echo.Context) error {
    car_id := c.Param("id")
    
	user_id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    car_id_i, err := strconv.Atoi(car_id)
    cars, err := ca.carService.GetCarById(user_id, uint(car_id_i)) 
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusFound, cars)
}

func (u CarController) GetInstance() *CarController {
    return di.GetInstance("carHandler").(*CarController)
}
