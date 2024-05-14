package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/server/config"
	"github.com/saiful-izam/my-jualan/server/controller"
	"github.com/saiful-izam/my-jualan/server/dao"
	"github.com/saiful-izam/my-jualan/server/env"
	"github.com/saiful-izam/my-jualan/server/router"
	"github.com/saiful-izam/my-jualan/server/service"
)

func init() {
	config.GetEnv()
}

func main() {

	slog.Info("test logging")

	env := env.GetEnvironment()

	db, err := config.ConnectDatabase(env)

	if err != nil {
		panic(err)
	}

	userRepository := dao.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserControllerImpl(userService)

	r := gin.Default()
	router := router.NewUserRouter(userController, r)

	server := &http.Server{
		Addr:    ":" + env.ServerPort,
		Handler: router,
	}

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
