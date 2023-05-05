package utils

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/skip2/go-qrcode"
    "encoding/hex"
)

type Qr struct{
    Model model.ReservationRequest `json:"reservation_info"`
	Hash  string `json:"checksum"`
}

func GenerateQR(r *model.ReservationRequest) (string, error) {
	json, err := json.Marshal(r)
	if err != nil {
		return "", errors.New("Could not parse the json reservation request")
	}

	fmt.Println(string(json))
	h := sha256.New()
	h.Write([]byte(json))
	hash := h.Sum(nil)

    var qr Qr
    qr.Model = *r
    qr.Hash = hex.EncodeToString(hash)

	qr_json, err := marsharlQR(qr)
	if err != nil {
		return "", errors.New("Could not parse the json reservation request")
	}


	qr_image, err := qrcode.Encode(string(qr_json), qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	encoded_image := b64.StdEncoding.EncodeToString(qr_image)

	var qr_db model.QR
	qr_db.Image = encoded_image

	db, err := database.ConnectDatabase()
	if err != nil {
		return "", err
	}

	db.Save(&qr_db)

	return qr_db.ID, nil
}

func marsharlQR(q Qr) ([]byte, error) {
	qr_json, err := json.Marshal(&q)
	if err != nil {
		return nil, errors.New("Could not parse the json reservation request")
	}
	return qr_json, nil
}
