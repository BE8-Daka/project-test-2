package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type TaskModel interface {
	Insert(task *entity.Task) response.InsertTask
}