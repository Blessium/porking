package handler

import (
	"fmt"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetQRCode(c echo.Context) error {

	id := string(c.Param("uuid"))

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var qr model.QR

	fmt.Println(id)

	result := db.Limit(1).Find(&qr, "id = ?", id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, result.Error.Error())
	}

	return c.String(http.StatusFound, qr.Image)
}
