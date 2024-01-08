package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {

	events := server.Group("/events")
	events.Use(middlewares.Authenticate)
	events.GET("/", getEvents)
	events.GET("/:id", getEvent)
	events.POST("/", postEvent)
	events.PUT("/:id", updateEvent)
	events.DELETE("/:id", deleteEvent)

	userRoutes := server.Group("/user")
	userRoutes.POST("/signup", signup)
	userRoutes.POST("/login", login)

	// routes not found
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "not found"})
	})
}
