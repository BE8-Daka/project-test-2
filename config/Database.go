package config

import (
	"fmt"
	"project-test/entity"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config AppConfig) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.DB_Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Project{}, &entity.Task{})

	// db.Create(&entity.User{
	// 	Name: "Admin Website",
	// 	Username: "admin",
	// 	NoHp: "081212121212",
	// 	Email: "admin@gmail.com",
	// 	Password: "$2a$14$fqChZ4CqMd9uvfE7MU6y4OvjTsBHoIBSbN/Iymyu9fCBJ9/VoCXum", // password
	// })

	// db.Create(&entity.Project{
	// 	Name: "Unassigned",
	// 	UserID: 1,
	// })
}
