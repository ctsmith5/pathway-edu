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
			codeBlock("typescript", `// while: when you repeat until something changes
let lives = 3;
while (lives > 0) {
  console.log("Lives:", lives);
  lives--;
}`),
			textBlock(`## For Loops: Counting Repetition

A <code>for</code> loop is the most common way to repeat code a specific number of times. It combines three things into one line:

1. **Initialization** — set up your counter variable
2. **Condition** — check if we should keep looping
3. **Increment** — update the counter after each iteration

Here's the anatomy of a for loop:`),
			imageBlock("https://storage.googleapis.com/ludicrousapps-c1ea7.firebasestorage.app/PathwayEdu/b0aca6cb-3c5c-4b30-837b-72dbc70a84a8(1).jpeg", "Diagram showing the anatomy of a JavaScript for loop with labeled parts for initialization, condition, and increment", "The three parts of a for loop: initialization (let i = 1), condition (i <= 5), and increment (i++)"),
			textBlock(`The loop above runs 5 times, with <code>i</code> taking values 1, 2, 3, 4, and 5. After the 5th iteration, <code>i</code> becomes 6, the condition <code>i <= 5</code> is no longer true, and the loop exits.

Here's the same loop in action:`),
			codeBlock("typescript", `for (let i = 1; i <= 5; i++) {
  console.log("Count:", i);
}
// Output:
// Count: 1
// Count: 2
// Count: 3
// Count: 4
// Count: 5`),
			calloutBlock("tip", "The most common bug in beginner loops is an off-by-one error (starting at 0 vs 1, or stopping too early/late)."),
			exerciseBlock(
				"Write a loop that prints the numbers 1 through 20. For each number, print whether it is even or odd.",
				"A common solution uses a for loop and the modulo operator (%). Example: if (n % 2 === 0) it's even; otherwise it's odd.",
				[]string{"Use n % 2 to check remainder", "If remainder is 0, it's even"},
			),
		},
	}
}
