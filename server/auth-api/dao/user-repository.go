package dao

import (
	"github.com/saiful-izam/my-jualan/auth-api/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserById(id uint64) (model.User, error)
	Save(u *model.User) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&model.User{})

	return &UserRepositoryImpl{
		DB: db,
	}
}

func (u *UserRepositoryImpl) FindUserById(id uint64) (model.User, error) {
	var user model.User

	result := u.DB.Find(&user, id)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (u *UserRepositoryImpl) Save(user *model.User) error {
	result := u.DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
