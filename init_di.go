package main

import (
    "github.com/goioc/di"
    "github.com/blessium/porking/service"
    "github.com/blessium/porking/handler"
    "reflect"
)

func InitDi() {
    _, _ = di.RegisterBean("userService", reflect.TypeOf((*service.UserService)(nil))) 
    _, _ = di.RegisterBean("userHandler", reflect.TypeOf((*handler.UserController)(nil))) 
    _ = di.InitializeContainer()
}
