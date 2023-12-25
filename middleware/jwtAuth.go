package middleware

import (
	"diary_api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := helpers.ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required!"})
			context.Abort()
			return
		}
		context.Next()
	}
}
