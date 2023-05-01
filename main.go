package main

import (
	"github.com/labstack/echo/v4"
    "github.com/blessium/porking/utils"
)


func main() {
	e := echo.New()
    utils.InitValidator()
	registerRoutes(e)
	e.Logger.Fatal(e.Start(":1234"))
}
