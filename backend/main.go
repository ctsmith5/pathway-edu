package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pathway/backend/handlers"
	"github.com/pathway/backend/middleware"
	"github.com/pathway/backend/repository"
)

func main() {
	// Load environment variables if needed (optional for now)
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "pathway"
	}

	// Initialize Repository
	repo, err := repository.NewMongoRepository(mongoURI, dbName)
	if err != nil {
		// For development, we might want to continue even if DB fails, or panic.
		// Let's log and panic for now as DB is critical.
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer repo.Close()

	// Initialize Handlers
	h := handlers.NewHandler(repo)

	// Setup Router
	r := gin.Default()

	// CORS Configuration
	// Set ALLOWED_ORIGINS environment variable in production (e.g., "https://your-app.vercel.app")
	// Leave unset for development (defaults to "*" - allows all)
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*" // Development default
	}

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		// Set allowed origin
		if allowedOrigins == "*" {
			// Development: allow all origins
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			// Production: use specific origin from environment variable
			// For multiple origins, you can extend this to check against a list
			c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		}
		
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Routes
	api := r.Group("/api")
	{
		// Public routes
		api.GET("/health", h.HealthCheck)
		api.GET("/courses", h.GetCourses)
		api.GET("/courses/:id", h.GetCourseByID)

		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}

		// Protected routes (require authentication)
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/me", h.GetCurrentUser)
			user.GET("/progress", h.GetUserProgress)
		}
	}

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
