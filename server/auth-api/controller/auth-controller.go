package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/auth-api/common"
	"github.com/saiful-izam/my-jualan/auth-api/dto"
	"github.com/saiful-izam/my-jualan/auth-api/model"
	"github.com/saiful-izam/my-jualan/auth-api/service"
	"golang.org/x/crypto/bcrypt"
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
	var request dto.LoginDTO
	var errorResponse dto.ErrorDto

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		errorResponse = dto.ErrorDto{
			Error: common.INVALID_REQUEST_BODY,
		}
	}

	if errorResponse.Error != "" {
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
}

func (auth *AuthControllerImpl) Register(c *gin.Context) {
	var request dto.RegisterDTO
	var errorResponse dto.ErrorDto

	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse = dto.ErrorDto{
			Error: "invalid request body",
		}
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		slog.Error("Error when hashing the password : ", err)
		errorResponse = dto.ErrorDto{
			Error: "error when hashing the password..",
		}
	}

	user := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPwd),
	}

	err = auth.UserService.SaveUser(&user)

	if err != nil {
		errorResponse = dto.ErrorDto{
			Error: "error when saving the user..",
		}
	}

	if errorResponse.Error != "" {
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	c.JSON(http.StatusCreated, user)
}
