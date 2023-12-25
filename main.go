package main

import (
	"diary_api/controllers"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})

}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controllers.AddEntry)
	protectedRoutes.GET("/entry", controllers.GetAllEntries)
	protectedRoutes.GET("/entry/:id", controllers.GetEntryById)
	protectedRoutes.PUT("/entry/:id", controllers.UpdateEntry)
	protectedRoutes.DELETE("/entry/:id", controllers.DeleteEntry)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
