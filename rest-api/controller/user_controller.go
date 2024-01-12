package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(context echo.Context) error
	LogIn(context echo.Context) error
	LogOut(context echo.Context) error
}

// MARK: - User Controller の実体

type userController struct {
	userUsecase usecase.IUserUseCase
}

func NewUserController(userUsecase usecase.IUserUseCase) IUserController {
	return &userController{userUsecase}
}

func (uc *userController) SignUp(context echo.Context) error {
	user := model.User{}
	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.userUsecase.SignUp(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(context echo.Context) error {
	user := model.User{}
	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.userUsecase.Login(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	// Cookie の設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	context.SetCookie(cookie)
	return context.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(context echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	context.SetCookie(cookie)
	return context.NoContent(http.StatusOK)
}

