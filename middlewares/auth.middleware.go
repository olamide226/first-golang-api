package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	token = utils.SplitToken(token)
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized, not token provided"})
		return
	}
	claims, err := utils.ValidateToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": err.Error()})
		return
	}
	ctx.Set("claims", claims)
	ctx.Next()
}