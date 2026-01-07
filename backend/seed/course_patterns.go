package seed

import "github.com/pathway/backend/models"

func courseDesignPatterns() models.Course {
	return models.Course{
		Title:       "Design Patterns",
		Description: "Learn proven software design patterns to solve common programming challenges elegantly and efficiently.",
		Modules: []models.Module{
			modulePatterns1(),
			modulePatterns2(),
			modulePatterns3(),
			modulePatterns4(),
			modulePatterns5(),
			modulePatterns6(),
			modulePatterns7(),
		},
	}
}
