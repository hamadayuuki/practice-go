package repository

import "go-rest-api/model"

type IUserRepository interface {
//                 引数                            戻り値
	GetUserByEmail(user *model.User, email string) error
	CerateUser(user *model.User) error
}