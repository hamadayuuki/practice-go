package validator

import (
	"go-rest-api/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}