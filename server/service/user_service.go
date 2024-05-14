package service

import (
	"github.com/saiful-izam/my-jualan/server/dao"
	"github.com/saiful-izam/my-jualan/server/dto"
)

type UserService interface {
	GetListOfUsers() dto.UserDTO
}

type UserServiceImpl struct {
	UserRepositoryImpl dao.UserRepository
}

func NewUserServiceImpl(repo dao.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepositoryImpl: repo,
	}
}

func (u UserServiceImpl) GetListOfUsers() dto.UserDTO {
	var result dto.UserDTO

	users := u.UserRepositoryImpl.GetAll()

	result.Error.ErrorCode = 000
	result.Error.ErrorMessage = "SUCCESS"
	result.User = append(result.User, users...)

	return result
}
