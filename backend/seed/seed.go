package seed

import (
	"log"

	"github.com/pathway/backend/models"
	"github.com/pathway/backend/repository"
)

// Helper function to create a text block
func textBlock(markdown string) models.ContentBlock {
	return models.ContentBlock{
		Type: "text",
		Data: map[string]interface{}{
			"markdown": markdown,
		},
	}
}

// Helper function to create a code block
func codeBlock(language, code string) models.ContentBlock {
	return models.ContentBlock{
		Type: "code",
		Data: map[string]interface{}{
			"language": language,
			"code":     code,
		},
	}
}

// Helper function to create an image block
func imageBlock(url, alt, caption string) models.ContentBlock {
	return models.ContentBlock{
		Type: "image",
		Data: map[string]interface{}{
			"url":     url,
			"alt":     alt,
			"caption": caption,
		},
	}
}

// Helper function to create a callout block
func calloutBlock(variant, text string) models.ContentBlock {
	return models.ContentBlock{
		Type: "callout",
		Data: map[string]interface{}{
			"variant": variant, // "info", "tip", "warning", "danger"
			"text":    text,
		},
	}
}

// Helper function to create an exercise block
func exerciseBlock(prompt, solution string, hints []string) models.ContentBlock {
	return models.ContentBlock{
		Type: "exercise",
		Data: map[string]interface{}{
			"prompt":   prompt,
			"solution": solution,
			"hints":    hints,
		},
	}
}

// Helper function to create a video block
func videoBlock(url, title string) models.ContentBlock {
	return models.ContentBlock{
		Type: "video",
		Data: map[string]interface{}{
			"url":   url,
			"title": title,
		},
	}
}

func SeedCourses(repo repository.Repository) error {
	// Clear existing courses
	if err := repo.DeleteAllCourses(); err != nil {
		return err
	}

	courses := []models.Course{
		courseGit(),
		courseSolid(),
		courseScrum(),
		courseHTTP(),
		courseDesignPatterns(),
		courseTesting(),
		courseDevelopmentTools(),
		courseCodeConcepts(),
	}

	// Insert all courses
	for _, course := range courses {
		if err := repo.CreateCourse(&course); err != nil {
			log.Printf("Error creating course %s: %v", course.Title, err)
			return err
		}
		log.Printf("Created course: %s", course.Title)
	}

	log.Println("Successfully seeded all courses!")
	return nil
}
