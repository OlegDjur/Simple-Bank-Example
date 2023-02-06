package dto

import (
	"sbank/internal/models"
	"time"
)

type CreateUserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserResponseDTO struct {
	Username          string    `json:"username"`
	FullName          string    `json:"ful_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type LoginUserRequestDTO struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponseDTO struct {
	AccessToken string          `json:"access_token"`
	User        UserResponseDTO `json:"user"`
}

func NewUserResponse(user models.User) *UserResponseDTO {
	return &UserResponseDTO{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
