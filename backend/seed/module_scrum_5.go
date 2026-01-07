package seed

import "github.com/pathway/backend/models"

func moduleScrum5() models.Module {
	return models.Module{
		ID:    "scrum-5",
		Title: "Sprint Planning",
		Content: []models.ContentBlock{
			textBlock(`## Sprint Planning Overview

Sprint Planning initiates the sprint by planning the work to be performed. It's time-boxed to a maximum of 8 hours for a one-month sprint.

**Participants**: Entire Scrum Team
**Facilitator**: Scrum Master`),
			textBlock(`## Part 1: What Can Be Done?

The Product Owner presents:

- **Product Backlog**: Ordered list of items
- **Product Goal**: Long-term objective
- **Forecast**: What might be achievable

The Development Team:

- Considers capacity
- Considers past performance
- Selects items they believe they can complete
- Creates a **Sprint Goal**`),
			textBlock(`## The Sprint Goal

The Sprint Goal:

- Provides focus and coherence
- Guides the team on why they're building the increment
- Can be achieved through various Product Backlog items
- If the goal becomes obsolete, the sprint may be cancelled`),
			calloutBlock("tip", "A good Sprint Goal is specific enough to provide direction but flexible enough to allow the team to adapt as they learn."),
			textBlock(`## Part 2: How Will It Be Done?

The Development Team plans:

- **How to build**: Technical approach
- **Tasks**: Breaks down items into tasks
- **Collaboration**: Who works on what
- **Initial plan**: For the first days of the sprint

The Product Owner:
- Answers questions
- Clarifies requirements
- May renegotiate scope if needed`),
			codeBlock("text", `Example Sprint Planning:

Sprint Goal: "Enable users to reset their password securely"

Selected Items:
- User story: "As a user, I want to reset my password"
- Technical task: "Set up email service"
- Bug fix: "Fix password validation"

Tasks:
- Design password reset flow
- Create reset token system
- Build email template
- Write tests
- Update documentation`),
			textBlock(`## Commitment

By the end of Sprint Planning:

- **Sprint Goal** is set
- **Sprint Backlog** is created
- Team is **committed** to the goal

Note: Commitment doesn't mean guarantee. It means the team will do their best to achieve the goal.`),
			exerciseBlock(
				"During Sprint Planning, the team realizes they can't complete all the items the Product Owner wanted. What should happen?",
				"The team should communicate this to the Product Owner. They can either:\n1. Reduce the scope to what's achievable\n2. Extend the sprint (not recommended - breaks time-boxing)\n3. Adjust the Sprint Goal to be more realistic\n\nThe Product Owner and team collaborate to find the right balance. The team should never commit to more than they believe they can deliver.",
				[]string{"Think about transparency and honesty", "What's the purpose of the Sprint Goal?"},
			),
		},
	}
}
