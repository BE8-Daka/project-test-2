package response

import "time"

type InsertProject struct {
	Name      string 	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Project struct {
	ID		  uint      `json:"id"`
	Name      string 	`json:"name"`
}

type UpdateProject struct {
	Name      string 	`json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}