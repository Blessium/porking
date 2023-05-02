package model

import (
	"strconv"
	"time"
)

type Reservation struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	StartTime  time.Time `gorm:"index:,composite:time"`
	EndTime    time.Time `gorm:"index:,composite:time"`
	Cost       float32
	QRCodePath string
	CarID      uint
	CarParkID  uint
	UserID     uint `gorm:"index:idx_user"`
}

type ReservationRequest struct {
	StartTime UnixTime `json:"start_time"`
	EndTime   UnixTime `json:"end_time"`
	Cost      float32  `json:"cost"`
	CarID     uint     `json:"car_id"`
	CarParkID uint     `json:"car_park_id"`
}

func (r *ReservationRequest) ConvertToReservation() Reservation {
	var res Reservation
	res.StartTime = r.StartTime.Time
	res.EndTime = r.EndTime.Time
	res.Cost = r.Cost
	res.CarID = r.CarID
	res.CarParkID = r.CarParkID
	return res
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	timestamp, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)

	return nil
}

func (u *UnixTime) MarshalJSON() ([]byte, error) {
	timestamp := u.Time.Unix()
	return []byte(strconv.FormatInt(timestamp, 10)), nil
}
