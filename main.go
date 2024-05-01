package main

import (
	"chat-server/config"
	"chat-server/controller"
	"chat-server/database"
	"chat-server/errors"
	"chat-server/service"
	"chat-server/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDatabase()

	// Router
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		utils.ApiErrorResponse(c, errors.ApiNotFound)
	})

	// User
	userService := service.NewUserService(database.DB)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(router)

	// Run
	router.Run(config.AppConfig.Server.Port)
}
