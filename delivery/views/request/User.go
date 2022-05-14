package request

type InsertUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InsertLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"required"`
	NoHp     string `json:"no_hp"`
	Email    string `json:"email"`
	Password string `json:"password"`
}