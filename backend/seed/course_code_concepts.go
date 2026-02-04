package seed

import "github.com/pathway/backend/models"

func courseCodeConcepts() models.Course {
	return models.Course{
		Title:       "Coding Foundations",
		Description: "Your first steps into programming! Learn the building blocks that every programmer uses, explained in plain English with lots of examples.",
		Modules: []models.Module{
			moduleCodeConcepts1(),
			moduleCodeConcepts2(),
			moduleCodeConcepts3(),
			moduleCodeConcepts4(),
			moduleCodeConcepts5(),
		},
	}
}
