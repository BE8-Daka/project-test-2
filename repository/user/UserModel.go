package repository

import (
	"project-test/delivery/views/response"
	"project-test/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *userModel {
	return &userModel{db}
}

func (m *userModel) Insert(user *entity.User) (response.InsertUser, error) {
	user.Name = strings.Title(strings.ToLower(user.Name))
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	originalPassword := user.Password

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	
	result := m.DB.Create(&user)

	if result.RowsAffected == 0 {
		return response.InsertUser{}, result.Error
	} else {
		return response.InsertUser{
			Name: 	user.Name,
			Username: user.Username,
			NoHp: 	user.NoHp,
			Email: 	user.Email,
			Password: originalPassword,
			CreatedAt: user.CreatedAt,
		}, nil
	}
}