package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"

	"gorm.io/gorm"
)

type projectModel struct {
	DB *gorm.DB
}

func NewProjectModel(db *gorm.DB) *projectModel {
	return &projectModel{db}
}

func (m *projectModel) Insert(project *entity.Project) (response.InsertProject, error) {	
	result := m.DB.Create(&project)

	if result.RowsAffected == 0 {
		return response.InsertProject{}, result.Error
	} else {
		return response.InsertProject{
			Name: 	project.Name,
			CreatedAt: project.CreatedAt,
		}, nil
	}
}

func (m *projectModel) GetAll(user_id uint) []response.Project {
	var projects []response.Project
	result := m.DB.Where("user_id = ?", user_id).Find(&projects)

	if result.RowsAffected == 0 {
		return []response.Project{}
	} else {
		return projects
	}
}