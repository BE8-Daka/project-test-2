package request

type InsertTask struct {
	Name      string `json:"name" validate:"required"`
	ProjectID uint   `json:"project_id"`
}

type UpdateTask struct {
	Name      string `json:"name"`
	ProjectID uint   `json:"project_id"`
}