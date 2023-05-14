package main

import (
	"github.com/blessium/porking/handler"
	"github.com/blessium/porking/service"
	"github.com/goioc/di"
	"reflect"
)

func InitDi() {

	_, _ = di.RegisterBean("userService", reflect.TypeOf((*service.UserService)(nil)))
	_, _ = di.RegisterBean("carService", reflect.TypeOf((*service.CarService)(nil)))
	_, _ = di.RegisterBean("reservationService", reflect.TypeOf((*service.ReservationService)(nil)))
	_, _ = di.RegisterBean("carParkService", reflect.TypeOf((*service.CarParkService)(nil)))
	_, _ = di.RegisterBean("parkingSpotService", reflect.TypeOf((*service.ParkingSpotService)(nil)))

	_, _ = di.RegisterBean("keyHandler", reflect.TypeOf((*handler.KeysController)(nil)))
	_, _ = di.RegisterBean("userHandler", reflect.TypeOf((*handler.UserController)(nil)))
	_, _ = di.RegisterBean("carHandler", reflect.TypeOf((*handler.CarController)(nil)))
	_, _ = di.RegisterBean("reservationHandler", reflect.TypeOf((*handler.ReservationController)(nil)))
	_, _ = di.RegisterBean("carParkHandler", reflect.TypeOf((*handler.CarParkController)(nil)))
	_, _ = di.RegisterBean("parkingSpotHandler", reflect.TypeOf((*handler.ParkingSpotController)(nil)))

	_ = di.InitializeContainer()
}
