package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/server/service"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (u UserControllerImpl) GetAllUsers(c *gin.Context) {
	result := u.UserService.GetListOfUsers()

	c.JSON(http.StatusOK, result)
}
