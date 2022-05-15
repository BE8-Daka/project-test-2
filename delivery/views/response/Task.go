package response

import "time"

type InsertTask struct {
	Name      string 	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}