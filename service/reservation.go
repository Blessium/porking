package service

import (
    "github.com/blessium/porking/model"
    "github.com/blessium/porking/database"
    "github.com/gookit/event"
    b64 "encoding/base64"
    "github.com/blessium/porking/events/types"
)

type ReservationService struct {
    qrService IQRService  `di.inject:"qrService"`
}


func (r *ReservationService) CreateReservation(res *model.ReservationRequest, user_id uint) (*model.Reservation, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
        return nil, err
	}

    user, err := new(UserService).GetUserById(user_id)
    if err != nil {
        return nil, err
    }

    parking_spot, err := new(ParkingSpotService).GetReservationParkingSpot(res.CarParkID)
    if err != nil {
        return nil, err
    }

    

    qrinfo := res.ConvertToQRInfo()
    qrinfo.ParkingSpotID = parking_spot.ID
    
    qr, err := r.qrService.GenerateQR(qrinfo)
    if err != nil {
        return nil, err
    }

    re := res.ConvertToReservation()
	re.UserID = user.ID
    re.QRCodePath = "http://localhost:1234/qr/" + qr.ID
    re.ParkingSpotID = parking_spot.ID

	db.Save(re)

    qr_bytes := make([]byte, b64.StdEncoding.DecodedLen(len(qr.Image)))
    _, err = b64.StdEncoding.Decode(qr_bytes, []byte(qr.Image))
    if err != nil {
        return nil, err
    }

    qr_event := &types.EmailQREvent {
        Qr_data:  qr_bytes,
        Receiver: user.Email,
    }
    qr_event.SetName("email.QR")

    event.AsyncFire(qr_event)
    return re, nil
}

func (r *ReservationService) GetAllReservations(user_id uint) (*[]model.Reservation, error) {
    res := new([]model.Reservation)

	db, err := database.ConnectDatabase()
	if err != nil {
        return nil, err
	}

	if err := db.Model(&model.User{ID: user_id}).Association("Reservations").Find(&res); err != nil {
        return nil, err
	}
    return res, nil
}
