package model

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100"`
	Timezone  string `json:"timezone" validate:"required"`
}

type RegisterUserCommand struct {
	Email        string `json:"id"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Timezone     string `json:"timezone"`
}

type UserProfileCommand struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Timezone  string `json:"timezone"`
}

type UserProfileCommandRes struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Timezone  string `json:"timezone"`
	UpdatedAt string `json:"updated_at"`
}
