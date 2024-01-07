package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	SignUp(context echo.Context) error
	LogIn(context echo.Context) error
	LogOut(context echo.Context) error
}