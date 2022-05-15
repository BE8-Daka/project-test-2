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
	var projects []entity.Project
	m.DB.Where("user_id = ?", user_id).Find(&projects)

	var results []response.Project
	for _, project := range projects {
		results = append(results, response.Project{
			ID: project.ID,
			Name: project.Name,
		})
	}

	return results
}

func (m *projectModel) Update(id uint, project *entity.Project) response.UpdateProject {
	m.DB.Where("id = ?", id).Updates(&project)

	return response.UpdateProject{
		Name: 	project.Name,
		UpdatedAt: project.UpdatedAt,
	}
}

func (m *projectModel) CheckExist(id, user_id uint) bool {
	var project entity.Project
	result := m.DB.Model(&entity.Project{}).Where("id = ? AND user_id = ?", id, user_id).First(&project)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func (m *projectModel) Delete(id uint) response.DeleteProject {
	var project *entity.Project
	
	m.DB.Where("id = ?", id).Find(&project)
	m.DB.Delete(&project)

	return response.DeleteProject{
		Name: 	project.Name,
		DeletedAt: project.DeletedAt,
	}
}