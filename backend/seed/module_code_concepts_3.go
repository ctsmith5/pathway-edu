package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts3() models.Module {
	return models.Module{
		ID:    "code-3",
		Title: "Protocol Oriented Programming",
		Content: []models.ContentBlock{
			textBlock(`## Protocol-Oriented Programming

Protocol-Oriented Programming (POP) emphasizes defining behavior through protocols (interfaces) rather than through class inheritance.

### Why Protocol-Oriented?

- More flexible than class inheritance
- Allows composition over inheritance
- Better suited for value types
- Enables retroactive modeling`),
			calloutBlock("tip", "Swift popularized this paradigm, but the concepts apply to any language with interfaces or protocols."),
		},
	}
}
