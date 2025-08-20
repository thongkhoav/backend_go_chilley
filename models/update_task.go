package models

type UpdateTaskRequest struct {
	Completed *bool `json:"completed" validate:"required"`
}
