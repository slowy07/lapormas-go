package model

import "github.com/google/uuid"

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Name      string    `json:"name,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	UpdatedAt int64     `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=100,min=5"`
	Password string `json:"password" validate:"required,max=100,min=5"`
	Name     string `json:"name" validate:"required,max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100,min=5"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	Token string `validate:"required,max=100"`
}

type GetUserRequest struct {
	Token string `validate:"required,max=100"`
}
