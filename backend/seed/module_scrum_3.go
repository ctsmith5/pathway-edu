package seed

import "github.com/pathway/backend/models"

func moduleScrum3() models.Module {
	return models.Module{
		ID:    "scrum-3",
		Title: "SCRUM Events (Sprint, Daily Standup, Retrospective)",
		Content: []models.ContentBlock{
			textBlock(`## SCRUM Events

SCRUM has five formal events, all designed to create regularity and minimize meetings:

1. **Sprint**: The container for all other events
2. **Sprint Planning**: Plan the work for the sprint
3. **Daily Scrum**: 15-minute team sync
4. **Sprint Review**: Inspect the increment
5. **Sprint Retrospective**: Improve the process`),
			textBlock(`## The Sprint

The Sprint is:

- **Time-boxed**: 1-4 weeks (most commonly 2 weeks)
- **Fixed length**: Same duration for all sprints
- **Goal-focused**: Has a Sprint Goal
- **No changes**: Scope is locked during the sprint

During a sprint:
- No changes that endanger the Sprint Goal
- Quality doesn't decrease
- Scope may be clarified and renegotiated`),
			textBlock(`## Sprint Planning

**Time-box**: 8 hours for a 4-week sprint (proportionally less for shorter sprints)

**Two parts:**

1. **What can be done?** - Product Owner presents backlog, team selects items
2. **How will it be done?** - Team plans the work

**Output**: Sprint Goal and Sprint Backlog`),
			textBlock(`## Daily Scrum

**Time-box**: 15 minutes, same time and place every day

**Purpose**: 
- Plan work for the next 24 hours
- Inspect progress toward Sprint Goal
- Identify impediments

**Format**: Each team member answers:
- What did I do yesterday?
- What will I do today?
- Are there any impediments?`),
			calloutBlock("tip", "The Daily Scrum is NOT a status report to the Scrum Master. It's for the Development Team to synchronize and plan."),
			textBlock(`## Sprint Review

**Time-box**: 4 hours for a 4-week sprint

**Purpose**: Inspect the increment and adapt the Product Backlog

**Attendees**: Scrum Team + Stakeholders

**Activities**:
- Demo working software
- Discuss what was done
- Gather feedback
- Update Product Backlog`),
			textBlock(`## Sprint Retrospective

**Time-box**: 3 hours for a 4-week sprint

**Purpose**: Plan improvements for the next sprint

**Questions**:
- What went well?
- What could be improved?
- What will we do differently?

**Output**: Actionable improvements`),
			exerciseBlock(
				"Why are SCRUM events time-boxed?",
				"Time-boxing ensures events don't drag on unnecessarily, creates predictability, and forces focus. It also prevents over-planning and ensures the team spends most of their time on actual development work rather than meetings.",
				[]string{"Think about what happens without time limits", "What's the cost of long meetings?"},
			),
		},
	}
}
