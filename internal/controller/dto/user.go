package dto

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fuuk_name"`
	Email    string `json:"email"`
}
