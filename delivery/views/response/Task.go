package response

import (
	"time"

	"gorm.io/gorm"
)

type InsertTask struct {
	Name      string 	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	ProjectID 	uint 		`json:"project_id"`
}

type TaskID struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
}

type UpdateTask struct {
	Name 		string 		`json:"name"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type DeleteTask struct {
	Name		string			`json:"name"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at"`
}