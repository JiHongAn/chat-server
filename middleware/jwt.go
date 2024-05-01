package middleware

import (
	"chat-server/config"
	"chat-server/errors"
	"chat-server/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			utils.ApiErrorResponse(context, errors.InvalidAccessToken)
			context.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.Jwt.AccessTokenSecret), nil
		})

		if err != nil {
			utils.ApiErrorResponse(context, errors.InvalidAccessToken)
			context.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set("id", claims["id"])
			context.Next()
		} else {
			utils.ApiErrorResponse(context, errors.InvalidAccessToken)
			context.Abort()
		}
	}
}
