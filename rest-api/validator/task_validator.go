package validator

import (
	"go-rest-api/model"
)

type ITaskValidator interface {
	TaskValidate(task model.task) error
}




