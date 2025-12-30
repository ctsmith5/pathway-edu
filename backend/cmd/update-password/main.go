package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/pathway/backend/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables from the backend directory
	backendPath := filepath.Join(".", "pathway", "backend")
	if err := godotenv.Load(filepath.Join(backendPath, ".env")); err != nil {
		log.Printf("Warning: No .env file found in %s or error loading: %v", backendPath, err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "pathway"
	}

	email := os.Getenv("USER_EMAIL")
	if email == "" {
		email = "test@example.com"
	}

	password := os.Getenv("USER_PASSWORD")
	if password == "" {
		password = "TestPass123"
		log.Printf("USER_PASSWORD not set, using default: %s", password)
	}

	repo, err := repository.NewMongoRepository(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer repo.Close()

	// Get user
	user, err := repo.GetUserByEmail(email)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	if user == nil {
		log.Fatalf("User %s not found", email)
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Update password in database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		log.Fatalf("Failed to parse user ID: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"password": string(hashedPassword),
		},
	}

	result, err := repo.GetDB().Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		update,
	)
	if err != nil {
		log.Fatalf("Failed to update password: %v", err)
	}

	if result.MatchedCount == 0 {
		log.Fatal("User not found for update")
	}

	log.Println("✅ Password updated successfully!")
	log.Printf("   Email: %s", email)
	log.Printf("   New Password: %s", password)
}






