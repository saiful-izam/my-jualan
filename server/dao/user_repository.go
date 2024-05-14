package dao

import (
	"github.com/saiful-izam/my-jualan/server/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() []model.User
	FindById(id int) model.User
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&model.User{})

	return &UserRepositoryImpl{
		Db: db,
	}
}

func (u UserRepositoryImpl) GetAll() []model.User {
	var users []model.User

	result := u.Db.Find(&users)

	if result.Error != nil {
		panic(result.Error)
	}

	return users
}

func (u UserRepositoryImpl) FindById(id int) model.User {
	var user model.User

	return user
}
