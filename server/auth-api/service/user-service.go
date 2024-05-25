package service

import (
	"log/slog"

	"github.com/saiful-izam/my-jualan/auth-api/dao"
	"github.com/saiful-izam/my-jualan/auth-api/model"
)

type UserService interface {
	GetUser(id uint64) (model.User, error)
	SaveUser(user *model.User) error
}

type UserServiceImpl struct {
	UserRepository dao.UserRepository
}

func NewUserServiceImpl(repo dao.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repo,
	}
}

func (svc *UserServiceImpl) GetUser(id uint64) (model.User, error) {
	user, err := svc.UserRepository.FindUserById(id)

	if err != nil {
		slog.Error("Error while fetching user from DB : ", err)
		return model.User{}, err
	}

	return user, nil
}

func (svc *UserServiceImpl) SaveUser(user *model.User) error {
	err := svc.UserRepository.Save(user)

	if err != nil {
		slog.Error("Error while saving user into DB : ", err)
		return err
	}

	return nil
}
