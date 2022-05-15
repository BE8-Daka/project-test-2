package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string 	`gorm:"type:varchar(35);not null"`
	Username string 	`gorm:"type:varchar(35);not null;unique"`
	NoHp     string 	`gorm:"type:varchar(15);not null;unique"`
	Email    string 	`gorm:"type:varchar(35);not null;unique"`
	Password string 	`gorm:"type:varchar(255);not null"`
	Tasks    []Task    	`gorm:"foreignkey:UserID"`
	Projects []Project 	`gorm:"foreignkey:UserID"`
}

type Project struct {
	gorm.Model
	Name   	string 	`gorm:"type:varchar(35);not null"`
	UserID 	uint 	`gorm:"not null"`
	Tasks 	[]Task 	`gorm:"foreignkey:ProjectID"`
}

type Task struct {
	gorm.Model
	Name   		string 	`gorm:"type:varchar(35);not null"`
	Status 		bool  	`gorm:"type:boolean;not null"`
	UserID 		uint 	`gorm:"not null"`
	ProjectID 	uint 	`gorm:"not null"`
}