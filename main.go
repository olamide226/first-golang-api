package main

import (
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")
	db.InitDB("./db/database.db")
	server := gin.Default()

	routes.InitRoutes(server)
	
	server.Run(":8080")
}
