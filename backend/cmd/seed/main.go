package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pathway/backend/repository"
	"github.com/pathway/backend/seed"
)

func main() {
	// Load environment variables
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
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer repo.Close()

	log.Println("Starting database seed...")
	if err := seed.SeedCourses(repo); err != nil {
		log.Fatalf("Failed to seed courses: %v", err)
	}

	log.Println("✅ Database seeded successfully!")
}

