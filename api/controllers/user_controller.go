package controllers

import (
	"clean-go-echo/library"
	"clean-go-echo/models"
	"clean-go-echo/services"

	"github.com/albrow/forms"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	env            library.Env
	logger         library.LoggerZap
	servicesMethod services.User_MethodService
}

func ModuleUserController(userservice services.User_MethodService,
	env library.Env, log library.LoggerZap) UserController {
	return UserController{
		env:            env,
		servicesMethod: userservice,
		logger:         log,
	}
}

func (u UserController) ListUser(c echo.Context) error {

	data_req, _ := forms.Parse(c.Request())

	user, err := u.servicesMethod.ListUser(data_req.GetInt("limit"))
	if err != nil {
		library.Writelog(c, u.env, "err", err.Error())
	}

	return library.ResponseInterface(c, 200, user, "List User")
}

func (u UserController) StoreUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return err
	}

	user.SetTime()

	user, err := u.servicesMethod.CreateUser(user)
	if err != nil {
		library.Writelog(c, u.env, "err", err.Error())
		return err
	}

	return library.ResponseInterface(c, 201, user, "List User")
}
