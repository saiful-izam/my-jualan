package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/auth-api/controller"
)

func NewAuthRouter(a controller.AuthController, r *gin.Engine) *gin.Engine {
	authRouter := r.Group("/api/auth")
	authRouter.POST("/login", a.Login)

	return r
}
