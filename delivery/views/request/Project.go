package request

type Project struct {
	Name string `json:"name" validate:"required"`
}