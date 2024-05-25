package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiful-izam/my-jualan/auth-api/config"
	"github.com/saiful-izam/my-jualan/auth-api/controller"
	"github.com/saiful-izam/my-jualan/auth-api/dao"
	"github.com/saiful-izam/my-jualan/auth-api/router"
	"github.com/saiful-izam/my-jualan/auth-api/service"
)

func main() {

	slog.Info("Initializing auth-api services....")

	var err error

	env, err := config.SetupEnvironment()

	if err != nil {
		slog.Error("Error while setup environment : ", err)
	}

	db, err := config.SetupDatabase(env)

	if err != nil {
		slog.Error("Error while setup Database Connecton : ", err)
	}

	// initiliaze dependency injection
	userRepo := dao.NewUserRepositoryImpl(db)
	userSvc := service.NewUserServiceImpl(userRepo)
	authController := controller.NewAuthControllerImpl(userSvc)

	r := gin.Default()
	router := router.NewAuthRouter(authController, r)

	server := &http.Server{
		Addr:    ":" + env.ServerPort,
		Handler: router,
	}

	err = server.ListenAndServe()

	if err != nil {
		slog.Error("error when initiating auth-api : ", err)
	}

	slog.Info("auth-api services running on ", "port", env.ServerPort)

}
