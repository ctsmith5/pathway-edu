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
	_ = godotenv.Load()

	// Source (dev) database
	sourceURI := os.Getenv("SOURCE_MONGO_URI")
	if sourceURI == "" {
		log.Fatal("SOURCE_MONGO_URI environment variable is required")
	}
	sourceDBName := os.Getenv("SOURCE_DB_NAME")
	if sourceDBName == "" {
		sourceDBName = "pathway"
	}

	// Target (production) database
	targetURI := os.Getenv("TARGET_MONGO_URI")
	if targetURI == "" {
		log.Fatal("TARGET_MONGO_URI environment variable is required")
	}
	targetDBName := os.Getenv("TARGET_DB_NAME")
	if targetDBName == "" {
		targetDBName = "pathway"
	}

	// User email to migrate
	userEmail := os.Getenv("USER_EMAIL")
	if userEmail == "" {
		userEmail = "test@example.com"
		log.Printf("USER_EMAIL not set, defaulting to: %s", userEmail)
	}

	// Connect to source database
	log.Println("Connecting to source database...")
	sourceRepo, err := repository.NewMongoRepository(sourceURI, sourceDBName)
	if err != nil {
		log.Fatalf("Failed to connect to source MongoDB: %v", err)
	}
	defer sourceRepo.Close()

	// Connect to target database
	log.Println("Connecting to target database...")
	targetRepo, err := repository.NewMongoRepository(targetURI, targetDBName)
	if err != nil {
		log.Fatalf("Failed to connect to target MongoDB: %v", err)
	}
	defer targetRepo.Close()

	// Migrate user and progress
	if err := migrateUserAndProgress(sourceRepo, targetRepo, userEmail); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully!")
}

func migrateUserAndProgress(sourceRepo, targetRepo repository.Repository, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get target repository as MongoRepository to access db directly
	targetMongo, ok := targetRepo.(*repository.MongoRepository)
	if !ok {
		return fmt.Errorf("target repository is not a MongoRepository")
	}

	// 1. Get user from source
	log.Printf("Fetching user: %s from source database...", email)
	sourceUser, err := sourceRepo.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to get user from source: %v", err)
	}
	log.Printf("Found user: %s (ID: %s)", sourceUser.Email, sourceUser.ID.Hex())

	// 2. Check if user already exists in target
	targetUser, err := targetRepo.GetUserByEmail(email)
	var targetUserID primitive.ObjectID
	if err == nil && targetUser != nil {
		log.Printf("User %s already exists in target database (ID: %s)", email, targetUser.ID.Hex())
		log.Println("Skipping user creation, using existing user...")
		targetUserID = targetUser.ID
	} else {
		// 3. Copy user to target (without ID to generate new one)
		log.Println("Copying user to target database...")
		newUser := *sourceUser
		newUser.ID = primitive.NilObjectID // Let MongoDB generate new ID
		
		if err := targetRepo.CreateUser(&newUser); err != nil {
			return fmt.Errorf("failed to create user in target: %v", err)
		}
		
		// Get the newly created user to get its ID
		createdUser, err := targetRepo.GetUserByEmail(email)
		if err != nil {
			return fmt.Errorf("failed to get created user: %v", err)
		}
		targetUserID = createdUser.ID
		log.Printf("User created in target database (ID: %s)", targetUserID.Hex())
	}

	// 4. Get progress from source (using source user ID)
	log.Println("Fetching progress from source database...")
	sourceProgress, err := sourceRepo.GetUserProgress(sourceUser.ID.Hex())
	if err != nil {
		return fmt.Errorf("failed to get progress from source: %v", err)
	}
	log.Printf("Found %d progress records in source", len(sourceProgress))

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

	// Delete existing progress for this user in target (idempotent - safe to run multiple times)
	log.Println("Cleaning up existing progress in target database...")
	_, err = progressCollection.DeleteMany(ctx, bson.M{"user_id": targetUserID})
	if err != nil {
		return fmt.Errorf("failed to delete existing progress: %v", err)
	}

	// Insert progress records with mapped course IDs
	var progressDocs []interface{}
	for _, prog := range sourceProgress {
		targetCourseID, exists := courseIDMap[prog.CourseID]
		if !exists {
			log.Printf("Warning: Course ID %s not found in target, skipping progress record", prog.CourseID.Hex())
			continue
		}

		newProgress := models.Progress{
			UserID:           targetUserID, // Use target user ID
			CourseID:         targetCourseID,
			CompletedModules: prog.CompletedModules,
			IsCompleted:      prog.IsCompleted,
		}
		progressDocs = append(progressDocs, newProgress)
	}

	if len(progressDocs) > 0 {
		log.Printf("Inserting %d progress records into target database...", len(progressDocs))
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

