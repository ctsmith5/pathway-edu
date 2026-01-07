package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts6() models.Module {
	return models.Module{
		ID:    "code-6",
		Title: "Control Flow (How Programs Make Decisions)",
		Content: []models.ContentBlock{
			textBlock(`## Quick note before we begin

In the previous module (Programming Languages 101), you saw a quick preview of if/else and how different languages express the same ideas.

This module goes deeper and gives you the tools you will use in every program:
- decisions (if/else and switch)
- repetition (loops)
- the mental model for reading code like a story`),
			textBlock(`## Control Flow: The "Choose Your Own Adventure" of Code

Control flow is how a program decides:
- **What to do next**
- **When to repeat**
- **When to stop**

If programming is like writing instructions for a robot, control flow is how you say:
- "If the door is closed, open it"
- "Do this 10 times"
- "Keep doing this until the button is pressed"

## The core tools (TypeScript examples)

### If / Else (decisions)`),
			codeBlock("typescript", `const score = 87;

if (score >= 90) {
  console.log("A");
} else if (score >= 80) {
  console.log("B");
} else {
  console.log("Keep going!");
}`),
			textBlock(`### Switch (many options)`),
			codeBlock("typescript", `const day = "Monday";

switch (day) {
  case "Monday":
    console.log("Start strong");
    break;
  case "Friday":
    console.log("Almost there");
    break;
  default:
    console.log("Another day, another step");
}`),
			textBlock(`### Loops (repetition)`),
			codeBlock("typescript", `// for: when you know how many times
for (let i = 1; i <= 5; i++) {
  console.log("Count:", i);
}

// while: when you repeat until something changes
let lives = 3;
while (lives > 0) {
  console.log("Lives:", lives);
  lives--;
}`),
			calloutBlock("tip", "The most common bug in beginner loops is an off-by-one error (starting at 0 vs 1, or stopping too early/late)."),
			exerciseBlock(
				"Write a loop that prints the numbers 1 through 20. For each number, print whether it is even or odd.",
				"A common solution uses a for loop and the modulo operator (%). Example: if (n % 2 === 0) it's even; otherwise it's odd.",
				[]string{"Use n % 2 to check remainder", "If remainder is 0, it's even"},
			),
		},
	}
}
