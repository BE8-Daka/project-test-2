package response

import "time"

type InsertUser struct {
	Name      string 	`json:"name"`
	Username  string 	`json:"username"`
	NoHp      string 	`json:"no_hp"`
	Email     string 	`json:"email"`
	Password  string 	`json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type InsertLogin struct {
	ID 			uint 		`json:"id"`
	Name 		string 		`json:"name"`
	Token 		string 		`json:"token"`
}