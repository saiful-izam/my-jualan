package dto

type LoginDTO struct {
	Email    string
	Password string
}

type RegisterDTO struct {
	Username string
	LoginDTO
}
