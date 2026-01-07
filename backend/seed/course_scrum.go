package seed

import "github.com/pathway/backend/models"

func courseScrum() models.Course {
	return models.Course{
		Title:       "SCRUM",
		Description: "Learn the SCRUM framework for agile project management. Understand sprints, standups, and how to deliver value iteratively.",
		Modules: []models.Module{
			moduleScrum1(),
			moduleScrum2(),
			moduleScrum3(),
			moduleScrum4(),
			moduleScrum5(),
			moduleScrum6(),
			moduleScrum7(),
		},
	}
}
