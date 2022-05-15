package repository

import (
	"errors"
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

func (m *userModel) Login(username, password string) (response.InsertLogin, error) {
	var user entity.User
	result := m.DB.Where("username = ?", username).First(&user)

	if result.RowsAffected == 0 {
		return response.InsertLogin{}, errors.New("username or password is wrong")
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return response.InsertLogin{}, errors.New("username or password is wrong")
		} else {
			return response.InsertLogin{
				ID: 	user.ID,
				Name: 	user.Name,
				Token: 	"",
			}, nil
		}
	}
}

func (m *userModel) GetbyID(id uint) response.GetUser {
	var user entity.User
	m.DB.Where("id = ?", id).First(&user)

	return response.GetUser{
		Name: 	user.Name,
		Username: user.Username,
		NoHp: 	user.NoHp,
		Email: 	user.Email,
	}
}

func (m *userModel) Update(user_id uint, user *entity.User) (response.UpdateUser, error) {
	user.Name = strings.Title(strings.ToLower(user.Name))
	user.Email = strings.ToLower(user.Email)
	originalPassword := user.Password

	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	
	result := m.DB.Where("id = ?", user_id).Updates(&user)

	if result.RowsAffected == 0 {
		return response.UpdateUser{}, result.Error
	} else {
		return response.UpdateUser{
			Name: 	user.Name,
			NoHp: 	user.NoHp,
			Email: 	user.Email,
			Password: originalPassword,
			UpdatedAt: user.UpdatedAt,
		}, nil
	}
}

func (m *userModel) Delete(user_id uint) response.DeleteUser {
	var user *entity.User
	// m.DB.Clauses(clause.Returning{}).Where("id = ?", user_id).Delete(&users)

	m.DB.Where("id = ?", user_id).Find(&user)
	m.DB.Delete(&user)

	return response.DeleteUser{
		Name : user.Name,
		DeletedAt : user.DeletedAt,
	}
}