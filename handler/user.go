package handler

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func GetUsers() {

}

func AddUser(e echo.Context) error {
    return e.String(http.StatusCreated, "Aggiunto con successo")
}
