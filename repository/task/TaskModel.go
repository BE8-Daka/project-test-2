package repository

import (
	"errors"
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

func (m *taskModel) GetAll(user_id uint) []response.Task {
	var tasks []entity.Task
	m.DB.Where("user_id = ?", user_id).Find(&tasks)

	var results []response.Task
	for _, task := range tasks {
		results = append(results, response.Task {
			ID: task.ID,
			Name: task.Name,
			ProjectID: task.ProjectID,
		})
	}

	return results
}

func (m *taskModel) Update(id uint, task *entity.Task) (response.UpdateTask, error) {
	if task.Name == "" && task.ProjectID == 0 {
		return response.UpdateTask{}, errors.New("name or project_id is required")
	}

	m.DB.Where("id = ?", id).Updates(&task)

	return response.UpdateTask{
		Name: 	task.Name,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (m *taskModel) CheckExist(id, user_id uint) bool {
	var task entity.Task
	result := m.DB.Where("id = ? AND user_id = ?", id, user_id).First(&task)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func (m *taskModel) Delete(id uint) response.DeleteTask {
	var task *entity.Task
	
	m.DB.Where("id = ?", id).Find(&task)
	m.DB.Delete(&task)

	return response.DeleteTask{
		Name: 	task.Name,
		DeletedAt: task.DeletedAt,
	}
}