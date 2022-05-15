package response

import "time"

type InsertProject struct {
	Name      string 	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}