package handler

import (
	"errors"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/utils"
	"github.com/golang-jwt/jwt/v4"
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

	id, err := extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

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

	return c.JSON(http.StatusFound, u.CleanUser())
}

func UpdateUser(c echo.Context) error {
	u := new(model.User)

	id, err := extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

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

	uInfo := new(model.UserInfo)
	if err := c.Bind(&uInfo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	db.Save(uInfo.UpdateUser(u))

	return c.JSON(http.StatusFound, u.CleanUser())
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

func AuthUser(c echo.Context) error {

	u := new(model.UserLogin)

	db, err := database.ConnectDatabase()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := utils.Validate.Struct(*u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user := new(model.User)
	result := db.Limit(1).Where(&model.User{Email: u.Email, Password: u.Password}).Find(&user)

	if result.Error != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected == 0 {
		return c.String(http.StatusNotFound, "Utente non trovato")
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, token)
}

func extract_id_from_token(c echo.Context) (uint, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("Jwt token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return 0, errors.New("failed to cast claims as jwt.MapClaims")
	}
	return uint(claims["user_id"].(float64)), nil
}
