package controllers

import (
	"clean-go-echo/library"
	"clean-go-echo/services"
	"log"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	env            library.Env
	servicesMethod services.User_MethodService
}

func ModuleUserController(userservice services.User_MethodService,
	env library.Env) UserController {
	return UserController{
		env:            env,
		servicesMethod: userservice,
	}
}

func (u UserController) GetUser(c echo.Context) error {

	user, err := u.servicesMethod.ListUser()
	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(200, user)
}
