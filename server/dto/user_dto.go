package dto

import "github.com/saiful-izam/my-jualan/server/model"

type UserDTO struct {
	Error BaseDTO
	User  []model.User
}
