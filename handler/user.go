package handler

import (
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUsers(c echo.Context) error {
	var users []model.User

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := db.Find(&users)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")
	}

	return c.JSON(http.StatusFound, users)
}

func GetUser(c echo.Context) error {
	u := new(model.User)

	id := c.Param("id")

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result := db.Limit(1).First(&u, id)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")
	}

	return c.JSON(http.StatusFound, u)
}

func AddUser(c echo.Context) error {

	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db.Save(&u)

	return c.JSON(http.StatusCreated, u)
}
