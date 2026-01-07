package seed

import "github.com/pathway/backend/models"

func courseTesting() models.Course {
	return models.Course{
		Title:       "Testing",
		Description: "Master software testing strategies including unit tests, integration tests, and test-driven development (TDD).",
		Modules: []models.Module{
			moduleTesting1(),
			moduleTesting2(),
			moduleTesting3(),
			moduleTesting4(),
			moduleTesting5(),
			moduleTesting6(),
			moduleTesting7(),
		},
	}
}
