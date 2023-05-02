package database

import (
	"github.com/blessium/porking/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "postgres://blessium:blessium@porking-database:5432/postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.CarPark{}, &model.ParkingSpot{}, &model.Reservation{}, &model.Car{}, &model.QR{})

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	return db, nil
}
