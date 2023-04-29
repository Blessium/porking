package main

import (
    "github.com/labstack/echo/v4"
    "github.com/blessium/porking/handler"
)

func registerRoutes(e *echo.Echo) {
    user := e.Group("users")
    user.POST("", handler.AddUser)
    user.GET("", handler.GetUsers)
    user.GET("/:id", handler.GetUser)
}
