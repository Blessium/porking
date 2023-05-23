package service

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/skip2/go-qrcode"
    "github.com/blessium/porking/utils"
)


type IQRService interface {
    GenerateQR(r *model.QRInfo) (*model.QR, error)
}

type QRService struct {

}

func (q *QRService) GenerateQR(r *model.QRInfo) (*model.QR, error) {
	json, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("Could not parse the json reservation request")
	}

	h := sha256.New()
	h.Write([]byte(json))
	hash := h.Sum(nil)

	var qr model.QRCert
	qr.Model = *r
	qr.Hash = hex.EncodeToString(hash)
	signed_hash, err := utils.SignMessage(hash)
	if err != nil {
		return nil, err
	}
	qr.Firma = signed_hash

	qr_json, err := qr.MarsharlQR()
	if err != nil {
		return nil, errors.New("Could not parse the json reservation request")
	}

	qr_image, err := qrcode.Encode(string(qr_json), qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	encoded_image := b64.StdEncoding.EncodeToString(qr_image)

	var qr_db model.QR
	qr_db.Image = encoded_image

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	if result := db.Save(&qr_db); result.Error != nil {
        return nil, result.Error
    }

	return &qr_db, nil
}

