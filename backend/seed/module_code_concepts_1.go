package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts1() models.Module {
	return models.Module{
		ID:    "code-1",
		Title: "Object Oriented Programming",
		Content: []models.ContentBlock{
			textBlock(`## What is Object-Oriented Programming?

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects" that contain data and code.

### The Four Pillars of OOP

1. **Encapsulation**: Bundling data and methods that operate on that data
2. **Inheritance**: Creating new classes based on existing ones
3. **Polymorphism**: Objects of different types responding to the same interface
4. **Abstraction**: Hiding complex implementation details`),
			calloutBlock("info", "OOP helps manage complexity in large codebases by organizing code into reusable, modular components."),
		},
	}
}
