package seed

import "github.com/pathway/backend/models"

func moduleSolid1() models.Module {
	return models.Module{
		ID:    "solid-1",
		Title: "Introduction to SOLID Principles",
		Content: []models.ContentBlock{
			textBlock(`## What are SOLID Principles?

SOLID is an acronym for five object-oriented design principles that make software designs more understandable, flexible, and maintainable.

### The Five Principles

1. **S** - Single Responsibility Principle
2. **O** - Open/Closed Principle
3. **L** - Liskov Substitution Principle
4. **I** - Interface Segregation Principle
5. **D** - Dependency Inversion Principle`),
			calloutBlock("info", "SOLID principles were introduced by Robert C. Martin (Uncle Bob) in the early 2000s. They're guidelines, not strict rules - apply them thoughtfully."),
			textBlock(`## Why SOLID Matters

Following SOLID principles leads to:

- **Maintainability**: Easier to understand and modify code
- **Testability**: Components can be tested in isolation
- **Flexibility**: Changes don't cascade through the system
- **Reusability**: Components can be reused in different contexts
- **Scalability**: Codebase grows without becoming unmanageable`),
			textBlock(`## A Word of Caution

SOLID principles are tools, not goals. Over-applying them can lead to:
- Over-engineering
- Unnecessary complexity
- Premature abstraction

Use them when they solve real problems, not just because they exist.`),
			exerciseBlock(
				"What problem do SOLID principles primarily solve?",
				"SOLID principles solve the problem of code that becomes difficult to maintain, test, and extend as a project grows. They help prevent code rot and technical debt by promoting good design practices.",
				[]string{"Think about what happens to code over time", "Consider what makes code hard to change"},
			),
		},
	}
}
