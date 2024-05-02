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

type FriendController struct {
	Service   service.FriendService
	Validator *validator.Validate
}

func NewFriendController(service service.FriendService) *FriendController {
	return &FriendController{
		Service:   service,
		Validator: validator.New(),
	}
}

func (controller *FriendController) RegisterRoutes(router *gin.Engine) {
	private := router.Group("/")
	private.Use(middleware.JwtAuth())
	private.GET("/friends", controller.GetFriends)
	private.GET("/friends/requests", controller.GetFriendRequests)
	private.POST("/friends/requests", controller.CreateFriendRequest)
	private.POST("/friends/accepts", controller.FriendRequestAccept)
	private.DELETE("/friends/:id", controller.DeleteFriend)
}

// GetFriends 친구 목록 조회
func (controller *FriendController) GetFriends(context *gin.Context) {
	friends, err := controller.Service.GetFriends(context.GetString("id"))
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}
	utils.ApiSuccessResponse(context, friends)
}

// GetFriendRequests 친구 요청 조회
func (controller *FriendController) GetFriendRequests(context *gin.Context) {
	friendRequests, err := controller.Service.GetFriendRequests(context.GetString("id"))
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}
	utils.ApiSuccessResponse(context, friendRequests)
}

// CreateFriendRequest 친구 요청 생성
func (controller *FriendController) CreateFriendRequest(context *gin.Context) {
	var params dto.CreateFriendRequest

	if err := context.ShouldBindJSON(&params); err != nil {
		utils.ApiErrorResponse(context, errors.InvalidRequest)
		return
	}

	err := controller.Service.CreateFriendRequest(context.GetString("id"), params.FriendId)
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}
	utils.ApiSuccessResponse(context, nil)
}

// FriendRequestAccept 친구 요청 수락
func (controller *FriendController) FriendRequestAccept(context *gin.Context) {
	var params dto.CreateFriendRequest

	if err := context.ShouldBindJSON(&params); err != nil {
		utils.ApiErrorResponse(context, errors.InvalidRequest)
		return
	}

	err := controller.Service.FriendRequestAccept(context.GetString("id"), params.FriendId)
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}
	utils.ApiSuccessResponse(context, nil)
}

// DeleteFriend 친구 삭제
func (controller *FriendController) DeleteFriend(context *gin.Context) {
	friendId := context.Param("id")

	if friendId == "" {
		utils.ApiErrorResponse(context, errors.InvalidRequest)
		return
	}

	err := controller.Service.DeleteFriend(context.GetString("id"), friendId)
	if err != nil {
		utils.ApiErrorResponse(context, err)
		return
	}
	utils.ApiSuccessResponse(context, nil)
}
