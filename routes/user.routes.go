package routes

import (
	"log"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing user", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		log.Println("error saving user", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error saving user", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})

}

func login(ctx *gin.Context) {
	var user models.UserLogin
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing user", "error": err.Error()})
		return
	}

	err = user.Authenticate()
	if err != nil {
		log.Println("error authenticating user", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error authenticating user"})
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println("error generating token", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user authenticated", "token": token})

}
