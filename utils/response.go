package utils

import (
	"chat-server/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}

func ApiErrorResponse(c *gin.Context, err error) {
	if appErr, ok := err.(errors.AppError); ok {
		response := ApiResponse{
			Success: false,
			Code:    appErr.Code,
			Message: appErr.Message,
		}
		c.JSON(appErr.Status, response)
	} else {
		c.JSON(500, ApiResponse{
			Success: false,
			Code:    -1,
			Message: "Internal Server Error",
		})
	}
}

func ApiSuccessResponse(c *gin.Context, data interface{}) {
	response := ApiResponse{
		Success: true,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}
