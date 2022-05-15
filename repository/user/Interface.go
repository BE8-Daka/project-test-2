package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type UserModel interface {
	Insert(user *entity.User) (response.InsertUser, error)
	Login(username, password string) (response.InsertLogin, error)
	GetbyID(id uint) response.GetUser
	Update(user_id uint, user *entity.User) (response.UpdateUser, error)
	Delete(user_id uint) response.DeleteUser
}