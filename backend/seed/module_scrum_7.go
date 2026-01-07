package seed

import "github.com/pathway/backend/models"

func moduleScrum7() models.Module {
	return models.Module{
		ID:    "scrum-7",
		Title: "Running Effective Sprints",
		Content: []models.ContentBlock{
			textBlock(`## Sprint Execution

During a sprint, the Development Team:

- Works on Sprint Backlog items
- Collaborates daily
- Adapts the plan as needed
- Maintains quality standards
- Works toward the Sprint Goal`),
			textBlock(`## Daily Practices

**Daily Scrum**:
- Same time, same place
- 15 minutes maximum
- Focus on coordination, not status reporting
- Identify impediments

**Continuous Integration**:
- Integrate code frequently
- Run automated tests
- Keep the build green

**Pair Programming**:
- Share knowledge
- Improve code quality
- Reduce bottlenecks`),
			calloutBlock("tip", "The best sprints are boring - predictable, steady progress. Drama and heroics are signs of problems."),
			textBlock(`## Handling Impediments

Impediments block progress:

- **Identify quickly**: Raise in Daily Scrum
- **Escalate appropriately**: Scrum Master helps remove
- **Track visibly**: Use an impediment board
- **Resolve promptly**: Don't let them linger`),
			textBlock(`## Adapting During Sprint

The Sprint Backlog is:

- **Emergent**: Can be updated as team learns
- **Owned by team**: They decide how to adapt
- **Goal-focused**: Changes should support Sprint Goal

**Allowed changes**:
- Clarifying requirements
- Breaking down items further
- Adjusting tasks

**Not allowed**:
- Adding new Product Backlog items (unless Sprint Goal changes)
- Changing Sprint Goal without cancelling sprint`),
			textBlock(`## Definition of Done

The Definition of Done ensures quality:

**Example**:
- Code written and reviewed
- Unit tests written and passing
- Integration tests passing
- Documentation updated
- Deployed to staging
- Product Owner acceptance

**Benefits**:
- Prevents technical debt
- Ensures consistency
- Makes "done" clear to all`),
			textBlock(`## Sprint Cancellation

A sprint can be cancelled if:

- Sprint Goal becomes obsolete
- Market conditions change dramatically
- Technology becomes unavailable

**Process**:
- Only Product Owner can cancel
- Completed items are reviewed
- Incomplete items return to Product Backlog
- Team reflects on what happened`),
			textBlock(`## Common Anti-Patterns

**Avoid these**:

- ❌ Extending sprints when behind
- ❌ Adding work mid-sprint
- ❌ Skipping retrospectives
- ❌ Daily Scrum as status report
- ❌ Unclear Definition of Done
- ❌ Product Owner dictating how to work
- ❌ Scrum Master as team manager`),
			calloutBlock("warning", "SCRUM is simple but not easy. It requires discipline, transparency, and a commitment to continuous improvement."),
			exerciseBlock(
				"Halfway through a sprint, a critical bug is discovered in production. The Product Owner wants the team to drop everything and fix it. How should this be handled?",
				"This depends on the severity:\n\n1. If it's truly critical (system down, data loss), the Product Owner can cancel the sprint and the team addresses it immediately.\n\n2. If it's important but not critical, the Product Owner and team discuss:\n   - Can it wait until next sprint?\n   - Can we swap it for a lower-priority item?\n   - Does it change the Sprint Goal?\n\nThe key is transparency and collaboration, not just dictating changes.",
				[]string{"Think about the Sprint Goal", "What's the impact of mid-sprint changes?"},
			),
		},
	}
}
