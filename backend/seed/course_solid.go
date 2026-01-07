package seed

import "github.com/pathway/backend/models"

func courseSolid() models.Course {
	return models.Course{
		Title:       "Architecture - SOLID",
		Description: "Master the SOLID principles of object-oriented design to write maintainable, scalable, and robust software architectures.",
		Modules: []models.Module{
			moduleSolid1(),
			moduleSolid2(),
			moduleSolid3(),
			moduleSolid4(),
			moduleSolid5(),
			moduleSolid6(),
			moduleSolid7(),
		},
	}
}
