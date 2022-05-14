package request

type InsertUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}