package seed

import "github.com/pathway/backend/models"

func moduleScrum1() models.Module {
	return models.Module{
		ID:    "scrum-1",
		Title: "Introduction to Agile & SCRUM",
		Content: []models.ContentBlock{
			textBlock(`## What is Agile?

Agile is a mindset and set of values for software development that emphasizes:

- **Individuals and interactions** over processes and tools
- **Working software** over comprehensive documentation
- **Customer collaboration** over contract negotiation
- **Responding to change** over following a plan`),
			calloutBlock("info", "The Agile Manifesto was created in 2001 by 17 software developers. It's not about abandoning processes or documentation, but prioritizing what matters most."),
			textBlock(`## What is SCRUM?

SCRUM is a framework for implementing Agile. It provides:

- **Structure**: Roles, events, and artifacts
- **Process**: How teams work together
- **Values**: Commitment, courage, focus, openness, respect

SCRUM is lightweight, simple to understand, but difficult to master.`),
			textBlock(`## The SCRUM Framework

SCRUM consists of:

1. **Roles**: Product Owner, Scrum Master, Development Team
2. **Events**: Sprint, Sprint Planning, Daily Scrum, Sprint Review, Sprint Retrospective
3. **Artifacts**: Product Backlog, Sprint Backlog, Increment`),
			textBlock(`## Why SCRUM?

SCRUM helps teams:

- Deliver value frequently
- Adapt to changing requirements
- Improve continuously
- Increase transparency
- Reduce risk`),
			exerciseBlock(
				"What's the difference between Agile and SCRUM?",
				"Agile is a philosophy and set of values about how to approach software development. SCRUM is a specific framework that implements Agile principles. You can be Agile without using SCRUM (using Kanban, XP, etc.), but SCRUM is always Agile.",
				[]string{"Think about philosophy vs. implementation", "Is SCRUM the only way to be Agile?"},
			),
		},
	}
}
