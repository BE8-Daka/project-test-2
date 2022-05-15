package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type TaskModel interface {
	Insert(task *entity.Task) response.InsertTask
	GetAll(user_id uint) []response.Task
	Update(id uint, task *entity.Task) (response.UpdateTask, error)
	Delete(id uint) response.DeleteTask
	UpdateStatus(id uint, task *map[string]interface{}) response.UpdateTask
	CheckExist(id, user_id uint) bool
}