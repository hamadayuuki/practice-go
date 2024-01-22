package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

// MARK: - task validator の実体

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		// validate for Title
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),   // 値が格納されているか
			validation.RuneLength(1, 10).Error("limited max 10 char"),   // 文字の長さが1〜10か
		),
	)
}
