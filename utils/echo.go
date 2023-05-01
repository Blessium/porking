package utils

import (
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v4"
    "errors"
)

func Extract_id_from_token(c echo.Context) (uint, error) {
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
