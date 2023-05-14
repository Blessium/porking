package handler

import (
    "github.com/blessium/porking/utils"
    "github.com/labstack/echo/v4"
    "net/http"
)

type KeysController struct {

}

func (k *KeysController) GetPublicKey(c echo.Context) error {
    utils, err := utils.GetPublicKey()
    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }
    return c.String(http.StatusOK, utils)
}
