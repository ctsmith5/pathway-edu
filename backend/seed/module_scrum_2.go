package seed

import "github.com/pathway/backend/models"

func moduleScrum2() models.Module {
	return models.Module{
		ID:    "scrum-2",
		Title: "SCRUM Roles (Product Owner, Scrum Master, Team)",
		Content: []models.ContentBlock{
			textBlock(`## The Three SCRUM Roles

SCRUM defines three distinct roles, each with specific responsibilities.`),
			textBlock(`## Product Owner

The Product Owner is responsible for:

- **Maximizing product value**: Decides what to build
- **Managing the Product Backlog**: Prioritizes and maintains items
- **Stakeholder communication**: Represents customer needs
- **Accepting work**: Decides if items meet the Definition of Done`),
			calloutBlock("tip", "The Product Owner is ONE person, not a committee. They make decisions, though they should gather input from stakeholders."),
			textBlock(`## Scrum Master

The Scrum Master serves the team by:

- **Removing impediments**: Helps team overcome obstacles
- **Facilitating events**: Ensures SCRUM events happen and are effective
- **Coaching**: Helps team understand and apply SCRUM
- **Protecting the team**: Shields team from outside interference`),
			textBlock(`## Development Team

The Development Team:

- **Self-organizes**: Decides how to do the work
- **Cross-functional**: Has all skills needed to deliver
- **Accountable**: Collectively responsible for the increment
- **Sized appropriately**: Typically 3-9 people`),
			codeBlock("text", `Key Characteristics:

Development Team:
- No titles (no "senior developer" hierarchy in SCRUM)
- No sub-teams
- Works together on all items
- Estimates work collectively`),
			calloutBlock("warning", "The Development Team is NOT managed by the Scrum Master. The team is self-organizing and decides how to accomplish the work."),
			exerciseBlock(
				"Who is responsible for deciding HOW to implement a feature in SCRUM?",
				"The Development Team is responsible for deciding HOW to implement features. The Product Owner decides WHAT to build, but the team decides the technical approach. The Scrum Master helps ensure the team can work effectively.",
				[]string{"Think about the separation of concerns", "Who has the technical expertise?"},
			),
		},
	}
}
