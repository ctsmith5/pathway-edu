package seed

import "github.com/pathway/backend/models"

func courseDevelopmentTools() models.Course {
	return models.Course{
		Title:       "Development Tools",
		Description: "Set up your development environment with the essential tools every developer needs.",
		Modules: []models.Module{
			moduleDevTools1(),
		},
	}
}
