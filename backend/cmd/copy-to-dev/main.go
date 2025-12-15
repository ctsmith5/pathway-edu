package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pathway/backend/models"
	"github.com/pathway/backend/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Load environment variables
	// Try loading from current directory and from backend directory
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../../.env")

	// Source (production) database
	sourceURI := os.Getenv("SOURCE_MONGO_URI")
	if sourceURI == "" {
		log.Fatal("SOURCE_MONGO_URI environment variable is required")
	}
	sourceDBName := os.Getenv("SOURCE_DB_NAME")
	if sourceDBName == "" {
		sourceDBName = "pathway"
	}

	// Target (dev) database - same cluster, different database name
	targetURI := os.Getenv("TARGET_MONGO_URI")
	if targetURI == "" {
		// If not set, use same URI but will change database name
		targetURI = sourceURI
	}
	targetDBName := os.Getenv("TARGET_DB_NAME")
	if targetDBName == "" {
		targetDBName = "pathway-dev"
	}

	// User email to copy
	userEmail := os.Getenv("USER_EMAIL")
	if userEmail == "" {
		userEmail = "test@example.com"
		log.Printf("USER_EMAIL not set, defaulting to: %s", userEmail)
	}

	// Connect to source database
	log.Println("Connecting to source (production) database...")
	sourceRepo, err := repository.NewMongoRepository(sourceURI, sourceDBName)
	if err != nil {
		log.Fatalf("Failed to connect to source MongoDB: %v", err)
	}
	defer sourceRepo.Close()

	// Connect to target database (same cluster, different DB)
	log.Println("Connecting to target (dev) database...")
	targetRepo, err := repository.NewMongoRepository(targetURI, targetDBName)
	if err != nil {
		log.Fatalf("Failed to connect to target MongoDB: %v", err)
	}
	defer targetRepo.Close()

	// Copy user and progress
	if err := copyUserToDev(sourceRepo, targetRepo, userEmail); err != nil {
		log.Fatalf("Copy failed: %v", err)
	}

	log.Println("✅ User copied to dev database successfully!")
}

func copyUserToDev(sourceRepo, targetRepo repository.Repository, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get target repository as MongoRepository to access db directly
	targetMongo, ok := targetRepo.(*repository.MongoRepository)
	if !ok {
		return fmt.Errorf("target repository is not a MongoRepository")
	}

	// 1. Get user from source (production)
	log.Printf("Fetching user: %s from production database...", email)
	sourceUser, err := sourceRepo.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to get user from source: %v", err)
	}
	log.Printf("Found user: %s (ID: %s)", sourceUser.Email, sourceUser.ID.Hex())

	// 2. Check if user already exists in target (dev)
	targetUser, err := targetRepo.GetUserByEmail(email)
	var targetUserID primitive.ObjectID
	if err == nil && targetUser != nil {
		log.Printf("User %s already exists in dev database (ID: %s)", email, targetUser.ID.Hex())
		log.Println("Skipping user creation, using existing user...")
		targetUserID = targetUser.ID
	} else {
		// 3. Copy user to target (dev) - without ID to generate new one
		log.Println("Copying user to dev database...")
		newUser := *sourceUser
		newUser.ID = primitive.NilObjectID // Let MongoDB generate new ID
		
		if err := targetRepo.CreateUser(&newUser); err != nil {
			return fmt.Errorf("failed to create user in dev: %v", err)
		}
		
		// Get the newly created user to get its ID
		createdUser, err := targetRepo.GetUserByEmail(email)
		if err != nil {
			return fmt.Errorf("failed to get created user: %v", err)
		}
		targetUserID = createdUser.ID
		log.Printf("User created in dev database (ID: %s)", targetUserID.Hex())
	}

	// 4. Get progress from source (production)
	log.Println("Fetching progress from production database...")
	sourceProgress, err := sourceRepo.GetUserProgress(sourceUser.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed to get progress from source: %v", err)
	}
	log.Printf("Found %d progress records in production", len(sourceProgress))

	// 5. Get all courses from both databases to map course IDs
	sourceCourses, err := sourceRepo.GetAllCourses()
	if err != nil {
		return fmt.Errorf("failed to get source courses: %v", err)
	}
	targetCourses, err := targetRepo.GetAllCourses()
	if err != nil {
		return fmt.Errorf("failed to get target courses: %v", err)
	}

	// Create course ID mapping (source course ID -> target course ID) by title
	courseIDMap := make(map[primitive.ObjectID]primitive.ObjectID)
	for _, sourceCourse := range sourceCourses {
		for _, targetCourse := range targetCourses {
			if sourceCourse.Title == targetCourse.Title {
				courseIDMap[sourceCourse.ID] = targetCourse.ID
				log.Printf("Mapped course: %s -> %s", sourceCourse.ID.Hex(), targetCourse.ID.Hex())
				break
			}
		}
	}

	// 6. Copy progress records
	progressCollection := targetMongo.GetDB().Collection("progress")

	// Delete existing progress for this user in target (idempotent)
	log.Println("Cleaning up existing progress in dev database...")
	_, err = progressCollection.DeleteMany(ctx, bson.M{"user_id": targetUserID})
	if err != nil {
		return fmt.Errorf("failed to delete existing progress: %v", err)
	}

	// Insert progress records with mapped course IDs
	var progressDocs []interface{}
	for _, prog := range sourceProgress {
		targetCourseID, exists := courseIDMap[prog.CourseID]
		if !exists {
			log.Printf("Warning: Course ID %s not found in dev, skipping progress record", prog.CourseID.Hex())
			continue
		}

		newProgress := models.Progress{
			UserID:           targetUserID, // Use dev user ID
			CourseID:         targetCourseID,
			CompletedModules: prog.CompletedModules,
			IsCompleted:      prog.IsCompleted,
		}
		progressDocs = append(progressDocs, newProgress)
	}

	if len(progressDocs) > 0 {
		log.Printf("Inserting %d progress records into dev database...", len(progressDocs))
		_, err = progressCollection.InsertMany(ctx, progressDocs)
		if err != nil {
			return fmt.Errorf("failed to insert progress: %v", err)
		}
		log.Printf("Successfully copied %d progress records", len(progressDocs))
	} else {
		log.Println("No progress records to copy")
	}

	return nil
}

