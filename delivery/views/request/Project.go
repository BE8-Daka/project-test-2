package request

type InsertProject struct {
	Name string `json:"name" validate:"required"`
}