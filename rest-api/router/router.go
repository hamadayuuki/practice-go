package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc contorller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	return e
}