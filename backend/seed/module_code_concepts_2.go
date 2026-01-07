package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts2() models.Module {
	return models.Module{
		ID:    "code-2",
		Title: "Functional Programming",
		Content: []models.ContentBlock{
			textBlock(`## What is Functional Programming?

Functional Programming (FP) is a paradigm that treats computation as the evaluation of mathematical functions, avoiding changing state and mutable data.

### Key Concepts

- **Pure Functions**: Same input always produces same output, no side effects
- **Immutability**: Data cannot be changed after creation
- **First-Class Functions**: Functions can be passed as arguments and returned from other functions
- **Higher-Order Functions**: Functions that operate on other functions`),
			codeBlock("javascript", `// Pure function example
const add = (a, b) => a + b;

// Higher-order function
const map = (arr, fn) => arr.map(fn);

// Immutability - create new array instead of modifying
const numbers = [1, 2, 3];
const doubled = numbers.map(n => n * 2); // [2, 4, 6]`),
		},
	}
}
