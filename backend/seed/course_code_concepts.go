package seed

import "github.com/pathway/backend/models"

func courseCodeConcepts() models.Course {
	return models.Course{
		Title:       "Code Concepts",
		Description: "Fundamental programming paradigms and concepts that every developer should master.",
		Modules: []models.Module{
			moduleCodeConcepts1(),
			moduleCodeConcepts2(),
			moduleCodeConcepts3(),
			moduleCodeConcepts4(),
			moduleCodeConcepts5(),
			moduleCodeConcepts6(),
			moduleTypeScriptStarter(),
			moduleReactHelloWorld(),
			moduleFizzBuzzClass(),
			moduleFizzBuzzGaltonUI(),
		},
	}
}
