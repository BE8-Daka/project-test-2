package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type TaskModel interface {
	Insert(task *entity.Task) response.InsertTask
	GetAll(user_id uint) []response.Task
	Update(id uint, task *entity.Task) (response.UpdateTask, error)
	CheckExist(id, user_id uint) bool
}