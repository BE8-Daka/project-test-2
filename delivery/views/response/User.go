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

type GetUser struct {
	Name 		string 		`json:"name"`
	Username 	string 		`json:"username"`
	NoHp 		string 		`json:"no_hp"`
	Email 		string 		`json:"email"`
}

type UpdateUser struct {
	Name      string 	`json:"name"`
	NoHp      string 	`json:"no_hp"`
	Email     string 	`json:"email"`
	Password  string 	`json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}