package main

import (
	"github.com/blessium/porking/handler"
	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo) {
	user := e.Group("users")
	user.POST("", handler.AddUser)
	user.GET("", handler.GetUsers)
	user.GET("/:id", handler.GetUser)
    user.PUT("/:id", handler.UpdateUser)
    user.POST("/auth", handler.AuthUser)

	car_park := e.Group("car_parks")
	car_park.POST("", handler.AddCarPark)
	car_park.GET("", handler.GetAllCarParks)
	car_park.GET("/:id", handler.GetCarPark)
	car_park.PUT("/:id", handler.UpdatedCarPark)

	parking_spot := car_park.Group("/:id/parking_spots")
    parking_spot.GET("", handler.GetAllParkingSpots)
	parking_spot.POST("", handler.AddParkingSpot)
}
