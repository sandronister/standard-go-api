package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
