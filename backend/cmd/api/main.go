package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/chat-app/backend/internal/handlers"
	"github.com/yourusername/chat-app/backend/internal/middleware"
	"github.com/yourusername/chat-app/backend/internal/repository"
)

func main() {
	// Load configuration from environment variables
	cfg := repository.DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "chat_app"),
	}

	// Initialize database
	db, err := repository.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories and handlers
	userRepo := repository.NewUserRepository(db)
	jwtSecret := getEnv("JWT_SECRET", "default-secret-key-please-change")
	authHandler := handlers.NewAuthHandler(userRepo, jwtSecret)
	profileHandler := handlers.NewProfileHandler(userRepo)

	// Initialize Gin router
	r := gin.Default()

	// Routes setup
	setupRoutes(r, authHandler, profileHandler, jwtSecret)

	// Start server
	port := getEnv("PORT", "8080")
	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func setupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, profileHandler *handlers.ProfileHandler, jwtSecret string) {
	// Public routes
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Protected routes
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		authGroup.GET("/profile", profileHandler.GetProfile)
		authGroup.PATCH("/profile", profileHandler.UpdateProfile)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
