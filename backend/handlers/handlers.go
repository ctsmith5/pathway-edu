package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pathway/backend/repository"
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
	// Mock data for now
	courses := []gin.H{
		{"id": "1", "title": "Intro to Programming", "description": "Basics of Go and Logic"},
		{"id": "2", "title": "Web Development", "description": "HTML, CSS, and React"},
	}
	c.JSON(http.StatusOK, courses)
}
