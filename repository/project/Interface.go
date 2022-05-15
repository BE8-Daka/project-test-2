package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type ProjectModel interface {
	Insert(project *entity.Project) (response.InsertProject, error)
	GetAll(user_id uint) []response.Project
	Update(id uint, project *entity.Project) response.UpdateProject
	CheckExist(id, user_id uint) bool
}