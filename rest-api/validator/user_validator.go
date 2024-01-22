package validator

import (
	"go-rest-api/model"

	"github.com/go-ozzo/ozzo-validation/v4/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}
