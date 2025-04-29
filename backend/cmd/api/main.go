package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/chat-app/backend/internal/handlers"
	"github.com/yourusername/chat-app/backend/internal/middleware"
	"github.com/yourusername/chat-app/backend/internal/repository"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authHandler := handlers.NewAuthHandler(userRepo)
	profileHandler := handlers.NewProfileHandler(userRepo)

	r := gin.Default()

	// Auth routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Protected routes
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/profile", profileHandler.GetProfile)
		authGroup.PATCH("/profile", profileHandler.UpdateProfile)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
