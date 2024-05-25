package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/auth-api/dto"
	"github.com/saiful-izam/my-jualan/auth-api/service"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	UserService service.UserService
}

func NewAuthControllerImpl(userService service.UserService) *AuthControllerImpl {
	return &AuthControllerImpl{
		UserService: userService,
	}
}

func (auth *AuthControllerImpl) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from login api",
	})
}

func (auth *AuthControllerImpl) Register(c *gin.Context) {
	var request dto.RegisterDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request body",
		})
	}
}
