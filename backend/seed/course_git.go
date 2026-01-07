package seed

import "github.com/pathway/backend/models"

func courseGit() models.Course {
	return models.Course{
		Title:       "Git",
		Description: "Learn version control with Git - the essential tool for modern software development. Master branching, merging, and collaboration workflows.",
		Modules: []models.Module{
			moduleGit1(),
			moduleGit2(),
			moduleGit3(),
			moduleGit4(),
			moduleGit5(),
			moduleGit6(),
			moduleGit7(),
		},
	}
}
