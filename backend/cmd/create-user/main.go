package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pathway/backend/models"
	"github.com/pathway/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../../.env")

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "pathway"
	}

	// User details
	email := os.Getenv("USER_EMAIL")
	if email == "" {
		email = "test@example.com"
	}

	name := os.Getenv("USER_NAME")
	if name == "" {
		name = "Test User"
	}

	password := os.Getenv("USER_PASSWORD")
	if password == "" {
		password = "test123456"
	}

	// Initialize Repository
	repo, err := repository.NewMongoRepository(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer repo.Close()

	// Check if user already exists
	existingUser, err := repo.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		log.Printf("User %s already exists (ID: %s)", email, existingUser.ID.Hex())
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Create user
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "student",
	}

	if err := repo.CreateUser(user); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	log.Printf("✅ User created successfully!")
	log.Printf("   Email: %s", user.Email)
	log.Printf("   Name: %s", user.Name)
	log.Printf("   ID: %s", user.ID.Hex())

	// Initialize progress
	if err := repo.InitializeUserProgress(user.ID.Hex()); err != nil {
		log.Printf("Warning: Failed to initialize progress: %v", err)
	} else {
		log.Println("✅ Progress initialized for all courses")
	}
}

