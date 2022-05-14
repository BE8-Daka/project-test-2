package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
)

type UserModel interface {
	Insert(user *entity.User) (response.InsertUser, error)
	Login(username, password string) (response.InsertLogin, error)
}