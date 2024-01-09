package repository

import "go-rest-api/model"

type ITaskRepository interface {
	// CRUD
	CreateTask(task *model.Task) error
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}