package router

import (
	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/server/controller"
)

func NewUserRouter(controller controller.UserController, router *gin.Engine) *gin.Engine {

	controllerRouter := router.Group("/api/user")
	controllerRouter.GET("", controller.GetAllUsers)

	return router
}
