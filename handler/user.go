package handler

import (
	"github.com/blessium/porking/model"
	"github.com/blessium/porking/service"
	"github.com/blessium/porking/utils"
	"github.com/labstack/echo/v4"
    "github.com/goioc/di"
	"net/http"
)

type UserController struct {
	userService *service.UserService `di.inject:"userService"`
}

func (u *UserController) GetUsers(c echo.Context) error {

	users, err := u.userService.GetUsers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusFound, users)
}

func (u *UserController) GetUser(c echo.Context) error {

	id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    user, err := u.userService.GetUserById(id)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusFound, user.CleanUser())
}

func (u *UserController) UpdateUser(c echo.Context) error {
	id, err := utils.Extract_id_from_token(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	uInfo := new(model.UserInfo)
	if err := c.Bind(&uInfo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
    
    if err := u.userService.UpdateUser(uInfo, id); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
    }

	return c.String(http.StatusOK, "User updated successfully")
}

func (u *UserController) AddUser(c echo.Context) error {

	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    if err := u.userService.AddUser(user); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusCreated, u)
}

func (u *UserController) AuthUser(c echo.Context) error {

	login := new(model.UserLogin)

	if err := c.Bind(&login); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := utils.Validate.Struct(*u); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

    user, err := u.userService.AuthUser(login)
    if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
    }

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, token)
}

func (u UserController) GetInstance() *UserController {
   return di.GetInstance("userHandler").(*UserController) 
}
