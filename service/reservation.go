package service

import (
    "github.com/blessium/porking/model"
    "github.com/blessium/porking/database"
    "github.com/blessium/porking/utils"
)

type ReservationService struct {

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

    qr_res_model := res.ConvertToQRReservation()
    qr_res_model.ParkingSpotID = parking_spot.ID
    
    qr_path, err := utils.GenerateQR(qr_res_model)
    if err != nil {
        return nil, err
    }

    re := res.ConvertToReservation()
	re.UserID = user_id
    re.QRCodePath = "http://localhost:1234/qr/" + qr_path
    re.ParkingSpotID = parking_spot.ID

	db.Save(re)
    go utils.SendEmail(user.Email, user, re, qr_path)
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
