package dto

type CreateUser struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type GetUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
