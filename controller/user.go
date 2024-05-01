package controller

import (
	"chat-server/dto"
	"chat-server/errors"
	"chat-server/middleware"
	"chat-server/service"
	"chat-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	Service   service.UserService
	Validator *validator.Validate
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		Service:   service,
		Validator: validator.New(),
	}
}

func (controller *UserController) RegisterRoutes(router *gin.Engine) {
	public := router.Group("/")
	public.POST("/users", controller.CreateUser)

	private := router.Group("/")
	private.Use(middleware.JwtAuth())
	private.GET("/users", controller.GetUser)
}

func (controller *UserController) CreateUser(context *gin.Context) {
	var params dto.CreateUser
	if err := context.ShouldBindJSON(&params); err != nil {
		utils.ApiErrorResponse(context, errors.InvalidRequest)
		return
	}

	err := controller.Service.CreateUser(&params)
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}

	utils.ApiSuccessResponse(context, nil)
}

func (controller *UserController) GetUser(context *gin.Context) {
	users, err := controller.Service.GetUser()
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}

	utils.ApiSuccessResponse(context, users)
}
