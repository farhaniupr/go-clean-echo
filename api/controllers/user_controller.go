package controllers

import (
	"clean-go-echo/library"
	"clean-go-echo/models"
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

func (u UserController) ListUser(c echo.Context) error {

	user, err := u.servicesMethod.ListUser()
	if err != nil {
		log.Println(err.Error())
	}

	return c.JSON(200, user)
}

func (u UserController) StoreUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		log.Println(err.Error())
		return err
	}

	user.SetTime()

	user, err := u.servicesMethod.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return c.JSON(200, user)
}
