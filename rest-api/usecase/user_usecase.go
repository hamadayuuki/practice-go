package usecase

import "go-rest-api/model"

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	// JWT token を返すため、string を返すようにしている
	Login(user model.User) (string, error)
}