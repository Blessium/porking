package main

import (
	"github.com/blessium/porking/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading credentials.env")
	}
	InitDi()
    if err := utils.GenerateKeys(); err != nil {
        panic(err)
    }
	utils.InitValidator()
	registerRoutes(e)
	e.Logger.Fatal(e.Start(":1234"))
}
