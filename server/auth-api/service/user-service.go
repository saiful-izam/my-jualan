package service

import (
	"log/slog"

	"github.com/saiful-izam/my-jualan/auth-api/dao"
	"github.com/saiful-izam/my-jualan/auth-api/dto"
	"github.com/saiful-izam/my-jualan/auth-api/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUser(id uint64) (model.User, error)
	Login(req dto.LoginDTO) (dto.SessionDTO, error)
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

func (svc *UserServiceImpl) Login(req dto.LoginDTO) (dto.SessionDTO, error) {
	var session dto.SessionDTO
	user, err := svc.UserRepository.FindUserByEmail(req.Email)

	if err != nil {
		slog.Error("error while fetching user in DB : ", err)
		return dto.SessionDTO{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		slog.Error("Password are not matched for user : ", "email", user.Email)

		return dto.SessionDTO{}, err
	}

	session = dto.SessionDTO{
		Token:        "token123",
		RefreshToken: "refreshtoken1234",
	}
	return session, nil
}
