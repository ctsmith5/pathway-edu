package seed

import "github.com/pathway/backend/models"

func moduleScrum6() models.Module {
	return models.Module{
		ID:    "scrum-6",
		Title: "User Stories & Estimation",
		Content: []models.ContentBlock{
			textBlock(`## User Stories

User Stories are a way to express Product Backlog items from the user's perspective.

**Format**:
"As a [type of user], I want [some goal] so that [some reason]"

**Example**:
"As a customer, I want to save my payment information so that I can checkout faster next time."`),
			textBlock(`## The Three C's

User Stories have three critical aspects:

1. **Card**: Written description (the story itself)
2. **Conversation**: Discussion about details
3. **Confirmation**: Acceptance criteria (tests)`),
			calloutBlock("info", "User Stories are placeholders for conversations. The real value comes from the discussion, not just the written card."),
			textBlock(`## Acceptance Criteria

Acceptance criteria define when a story is "done":

**Example**:
- User can enter email and password
- System validates email format
- Error message shown for invalid email
- Success message shown after submission
- Email confirmation sent`),
			textBlock(`## Story Points

Story Points estimate relative effort, not time.

**Why relative?**
- Easier to compare than absolute time
- Accounts for complexity, uncertainty, and effort
- Team-specific (one team's 5 might be another's 8)

**Common scales**:
- Fibonacci: 1, 2, 3, 5, 8, 13, 21
- Powers of 2: 1, 2, 4, 8, 16
- T-shirt sizes: XS, S, M, L, XL`),
			codeBlock("text", `Estimation Example:

Story 1: "Login page" - 3 points (baseline)
Story 2: "Password reset" - 5 points (more complex)
Story 3: "User profile" - 8 points (much more complex)
Story 4: "Add comment" - 2 points (simpler than login)`),
			textBlock(`## Planning Poker

A technique for estimating:

1. Product Owner reads story
2. Team asks questions
3. Each person privately selects a card (estimate)
4. Everyone reveals simultaneously
5. Discuss differences
6. Re-estimate until consensus`),
			calloutBlock("tip", "The goal of estimation is not precision, but shared understanding. If estimates vary widely, it usually means the story needs clarification."),
			textBlock(`## Velocity

Velocity is:

- **Measure of capacity**: How many points a team completes per sprint
- **Historical data**: Based on past performance
- **Used for forecasting**: Helps predict future capacity
- **Team-specific**: Not comparable across teams`),
			textBlock(`## INVEST Criteria

Good user stories are:

- **I**ndependent: Can be developed in any order
- **N**egotiable: Details can be discussed
- **V**aluable: Delivers value to users
- **E**stimable: Team can estimate effort
- **S**mall: Can be completed in one sprint
- **T**estable: Has clear acceptance criteria`),
			exerciseBlock(
				"A user story is estimated at 13 points. The team's average velocity is 20 points per sprint. Is this story too large?",
				"Yes, this story is likely too large. At 13 points, it's 65% of the team's capacity, which makes it risky. Large stories should be broken down into smaller stories (ideally 1-8 points) to reduce risk and increase predictability. The team should split this into 2-3 smaller stories.",
				[]string{"Think about risk and predictability", "What happens if a large story isn't finished?"},
			),
		},
	}
}
