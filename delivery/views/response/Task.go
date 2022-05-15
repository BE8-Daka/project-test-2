package response

import "time"

type InsertTask struct {
	Name      string 	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	ProjectID 	uint 		`json:"project_id"`
}