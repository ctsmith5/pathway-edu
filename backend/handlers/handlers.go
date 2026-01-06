package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pathway/backend/repository"
	"github.com/pathway/backend/seed"
)

type Handler struct {
	Repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Pathway Backend is running",
	})
}

func (h *Handler) GetCourses(c *gin.Context) {
	courses, err := h.Repo.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch courses",
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetCourseByID retrieves a single course by its ID
func (h *Handler) GetCourseByID(c *gin.Context) {
	courseID := c.Param("id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID required"})
		return
	}

	course, err := h.Repo.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// GetUserProgress retrieves the authenticated user's progress across all courses
func (h *Handler) GetUserProgress(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Get courses with progress
	coursesWithProgress, err := h.Repo.GetUserProgressWithCourses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch progress"})
		return
	}

	// If no progress exists, initialize it
	if len(coursesWithProgress) == 0 {
		if err := h.Repo.InitializeUserProgress(userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize progress"})
			return
		}
		// Fetch again after initialization
		coursesWithProgress, err = h.Repo.GetUserProgressWithCourses(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch progress"})
			return
		}
	}

	c.JSON(http.StatusOK, coursesWithProgress)
}

// CompleteModuleRequest represents the request body for completing a module
type CompleteModuleRequest struct {
	CourseID string `json:"course_id" binding:"required"`
	ModuleID string `json:"module_id" binding:"required"`
}

// CompleteModule marks a module as complete for the authenticated user
func (h *Handler) CompleteModule(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req CompleteModuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Repo.MarkModuleComplete(userID, req.CourseID, req.ModuleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark module as complete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Module marked as complete",
		"course_id": req.CourseID,
		"module_id": req.ModuleID,
	})
}

// AdminSeedCourses reseeds the courses collection.
// This is intentionally protected by an env var + header token so it can't be triggered accidentally.
//
// To enable:
// - Set ADMIN_SEED_TOKEN on the backend service
// - Call POST /api/admin/seed with header X-Admin-Seed-Token: <token>
func (h *Handler) AdminSeedCourses(c *gin.Context) {
	seedToken := os.Getenv("ADMIN_SEED_TOKEN")
	if seedToken == "" {
		// Hide the endpoint if not configured
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	provided := c.GetHeader("X-Admin-Seed-Token")
	if provided == "" || provided != seedToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := seed.SeedCourses(h.Repo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed courses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Courses seeded successfully",
		"seeded_at": time.Now().UTC().Format(time.RFC3339),
	})
}
