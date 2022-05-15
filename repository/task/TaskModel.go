package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"

	"gorm.io/gorm"
)

type taskModel struct {
	DB *gorm.DB
}

func NewTaskModel(db *gorm.DB) *taskModel {
	return &taskModel{db}
}

func (m *taskModel) Insert(task *entity.Task) response.InsertTask {
	if task.ProjectID == 0 {
		task.ProjectID = 1
	}
	
	m.DB.Create(&task)

	return response.InsertTask{
		Name: 	task.Name,
		CreatedAt: task.CreatedAt,
	}
}