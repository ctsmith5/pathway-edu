package seed

import "github.com/pathway/backend/models"

func moduleScrum4() models.Module {
	return models.Module{
		ID:    "scrum-4",
		Title: "SCRUM Artifacts (Product Backlog, Sprint Backlog, Increment)",
		Content: []models.ContentBlock{
			textBlock(`## SCRUM Artifacts

SCRUM has three artifacts that provide transparency and opportunities for inspection:

1. **Product Backlog**: Ordered list of everything needed
2. **Sprint Backlog**: Work selected for the current sprint
3. **Increment**: Sum of all completed work`),
			textBlock(`## Product Backlog

The Product Backlog is:

- **Ordered**: Most valuable items first
- **Emergent**: Never complete, always evolving
- **Detailed appropriately**: Top items are detailed, lower items are vague
- **Owned by Product Owner**: But team contributes estimates

**Contains**:
- Features
- Bugs
- Technical work
- Knowledge acquisition`),
			calloutBlock("info", "The Product Backlog is the single source of truth for what needs to be done. Nothing gets done that's not in the backlog."),
			textBlock(`## Product Backlog Items

Each item has:

- **Description**: What it is
- **Order**: Priority/position
- **Estimate**: Effort (usually Story Points)
- **Value**: Why it matters

**Refinement**:
- Continuously refined
- Items at top are "ready" (clear, feasible, testable)
- Team helps clarify and estimate`),
			textBlock(`## Sprint Backlog

The Sprint Backlog:

- **Selected during Sprint Planning**: Team commits to items
- **Owned by Development Team**: They decide how to do it
- **Visible to all**: Shows progress
- **Updated daily**: During Daily Scrum

**Contains**:
- Selected Product Backlog items
- Plan for delivering the increment
- At least one high-priority process improvement`),
			textBlock(`## The Increment

The Increment is:

- **Sum of all work**: All Product Backlog items completed in the sprint
- **Potentially releasable**: Meets Definition of Done
- **Usable**: Stakeholders can use it
- **Additive**: Each increment builds on previous ones

**Definition of Done**:
- Criteria that must be met for work to be "done"
- Includes testing, documentation, code review, etc.
- Prevents technical debt`),
			calloutBlock("warning", "An increment must be 'done' according to the Definition of Done. Partially done work doesn't count as an increment."),
			exerciseBlock(
				"What's the difference between the Product Backlog and Sprint Backlog?",
				"The Product Backlog contains ALL work for the product, ordered by value. The Sprint Backlog contains only the work selected for the current sprint, plus the plan for how to do it. The Sprint Backlog is a subset of the Product Backlog.",
				[]string{"Think about scope - what's the difference in size?", "When is each created and updated?"},
			),
		},
	}
}
