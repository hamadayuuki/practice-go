package validator

import (
	"go-rest-api/model"

	"github.com/go-ozzo/ozzo-validation/v4/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

// MARK: - user validator の実体

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		// validate for Email
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 10 char"),
			is.Email.Error("is not valid email format"),   // Emailのフォーマットに準拠しているか
		),
		// validate for Password
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited 6〜30 char"),
		),
	)
}
