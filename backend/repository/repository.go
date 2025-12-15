package repository

import (
	"context"
	"log"
	"time"

	"github.com/pathway/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Close()
	// Course methods
	GetAllCourses() ([]models.Course, error)
	GetCourseByID(id string) (*models.Course, error)
	CreateCourse(course *models.Course) error
	DeleteAllCourses() error
	// User methods
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	// Progress methods
	GetUserProgress(userID string) ([]models.Progress, error)
	InitializeUserProgress(userID string) error
	GetUserProgressWithCourses(userID string) ([]models.CourseWithProgress, error)
}

type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoRepository(uri string, dbName string) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the primary
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return &MongoRepository{
		client: client,
		db:     client.Database(dbName),
	}, nil
}

func (r *MongoRepository) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.client.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
}

// GetDB returns the underlying MongoDB database (for migration scripts)
func (r *MongoRepository) GetDB() *mongo.Database {
	return r.db
}

// ==================== Course Methods ====================

// GetAllCourses retrieves all courses from the database
func (r *MongoRepository) GetAllCourses() ([]models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection("courses").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []models.Course
	if err = cursor.All(ctx, &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

// GetCourseByID retrieves a specific course by ID
func (r *MongoRepository) GetCourseByID(id string) (*models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var course models.Course
	err = r.db.Collection("courses").FindOne(ctx, bson.M{"_id": objectID}).Decode(&course)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

// CreateCourse inserts a new course into the database
func (r *MongoRepository) CreateCourse(course *models.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("courses").InsertOne(ctx, course)
	return err
}

// DeleteAllCourses removes all courses from the database (useful for seeding)
func (r *MongoRepository) DeleteAllCourses() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("courses").DeleteMany(ctx, bson.M{})
	return err
}

// ==================== User Methods ====================

// CreateUser inserts a new user into the database
func (r *MongoRepository) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	// Set the ID on the user object
	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetUserByEmail finds a user by their email address
func (r *MongoRepository) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := r.db.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID finds a user by their ID
func (r *MongoRepository) GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.db.Collection("users").FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ==================== Progress Methods ====================

// GetUserProgress retrieves all progress records for a user
func (r *MongoRepository) GetUserProgress(userID string) ([]models.Progress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.db.Collection("progress").Find(ctx, bson.M{"user_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var progress []models.Progress
	if err = cursor.All(ctx, &progress); err != nil {
		return nil, err
	}

	return progress, nil
}

// InitializeUserProgress creates progress entries for all courses for a new user
func (r *MongoRepository) InitializeUserProgress(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Get all courses
	courses, err := r.GetAllCourses()
	if err != nil {
		return err
	}

	// Create progress entry for each course
	var progressDocs []interface{}
	for _, course := range courses {
		progress := models.Progress{
			UserID:           userObjectID,
			CourseID:         course.ID,
			CompletedModules: []string{},
			IsCompleted:      false,
		}
		progressDocs = append(progressDocs, progress)
	}

	if len(progressDocs) > 0 {
		_, err = r.db.Collection("progress").InsertMany(ctx, progressDocs)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetUserProgressWithCourses retrieves courses with user's progress data
func (r *MongoRepository) GetUserProgressWithCourses(userID string) ([]models.CourseWithProgress, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Get all courses
	courses, err := r.GetAllCourses()
	if err != nil {
		return nil, err
	}

	// Get user's progress
	cursor, err := r.db.Collection("progress").Find(ctx, bson.M{"user_id": userObjectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var progressList []models.Progress
	if err = cursor.All(ctx, &progressList); err != nil {
		return nil, err
	}

	// Create a map for quick lookup
	progressMap := make(map[string]models.Progress)
	for _, p := range progressList {
		progressMap[p.CourseID.Hex()] = p
	}

	// Combine courses with progress
	var result []models.CourseWithProgress
	for _, course := range courses {
		cwp := models.CourseWithProgress{
			Course:           course,
			CompletedModules: []string{},
			IsCompleted:      false,
		}

		if progress, exists := progressMap[course.ID.Hex()]; exists {
			cwp.CompletedModules = progress.CompletedModules
			cwp.IsCompleted = progress.IsCompleted
		}

		// Calculate progress percentage
		totalModules := len(course.Modules)
		completedCount := len(cwp.CompletedModules)
		if totalModules > 0 {
			cwp.ProgressPercent = float64(completedCount) / float64(totalModules) * 100
		}

		result = append(result, cwp)
	}

	return result, nil
}
