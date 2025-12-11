package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"` // Don't return password in JSON
	Role     string             `bson:"role" json:"role"` // "student", "admin"
}

type Course struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Modules     []Module           `bson:"modules" json:"modules"`
}

type Module struct {
	ID          string `bson:"id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Content     string `bson:"content" json:"content"`
	VideoURL    string `bson:"video_url" json:"video_url"`
}

type Progress struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID           primitive.ObjectID `bson:"user_id" json:"user_id"`
	CourseID         primitive.ObjectID `bson:"course_id" json:"course_id"`
	CompletedModules []string           `bson:"completed_modules" json:"completed_modules"` // List of Module IDs
	IsCompleted      bool               `bson:"is_completed" json:"is_completed"`
}


