package model

type UserRequest struct {
	UserID *int `uri:"user_id" validate:"required,gt=0"`
}
