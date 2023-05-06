package main

import (
	"github.com/blessium/porking/handler"
    "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo) {

    jwt_mid := echojwt.WithConfig(echojwt.Config{SigningKey: []byte("abracadabra"),})

    userController := handler.UserController{}.GetInstance()
    carController := handler.CarController{}.GetInstance()

    user := e.Group("users")

	user.POST("", userController.AddUser)
	user.GET("", userController.GetUsers)
	user.GET("/me", userController.GetUser, jwt_mid)
    user.PUT("/me", userController.UpdateUser, jwt_mid)
    user.POST("/auth", userController.AuthUser)

	car_park := e.Group("car_parks")
	car_park.POST("", handler.AddCarPark)
	car_park.GET("", handler.GetAllCarParks)
	car_park.GET("/:id", handler.GetCarPark)
	car_park.PUT("/:id", handler.UpdatedCarPark)

	parking_spot := car_park.Group("/:id/parking_spots")
    parking_spot.GET("", handler.GetAllParkingSpots)
	parking_spot.POST("", handler.AddParkingSpot)

    cars := e.Group("cars", jwt_mid)
    cars.GET("", carController.GetCars)
    cars.GET("/:id", carController.GetCar)
    cars.POST("", carController.AddCar)
    cars.PUT("/:id", carController.UpdateCar)

    res := e.Group("reservations", jwt_mid)
    res.GET("", handler.GetAllReservations)
    res.POST("", handler.CreateReservation)

    e.GET("qr/:uuid", handler.GetQRCode)
}
