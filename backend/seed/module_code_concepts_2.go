package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts2() models.Module {
	return models.Module{
		ID:    "code-2",
		Title: "Making Decisions",
		Content: []models.ContentBlock{
			textBlock(`## Programs Need to Make Choices

Real life is full of decisions: "If it's raining, take an umbrella." Programs work the same way! They check conditions and decide what to do.

### Comparison Operators

Before we can make decisions, we need to compare things. These operators let us do that:

| Operator | Meaning | Example | Result |
|----------|---------|---------|--------|
| == | Equal to | 5 == 5 | true |
| != | Not equal to | 5 != 3 | true |
| > | Greater than | 7 > 3 | true |
| < | Less than | 2 < 5 | true |
| >= | Greater than or equal | 5 >= 5 | true |
| <= | Less than or equal | 3 <= 5 | true |`),
			codeBlock("javascript", `let age = 16;
let hasLicense = true;

console.log(age == 16);        // true
console.log(age > 18);         // false
console.log(hasLicense == true); // true`),
			calloutBlock("tip", "Use === (triple equals) instead of == when possible. It checks both value AND type. 5 === \"5\" is false, but 5 == \"5\" is true. This prevents weird bugs!"),
			textBlock(`## If Statements: The Basic Decision

An if statement checks a condition and runs code only if that condition is true.

### Real-World Analogy
Think of an if statement like a bouncer at a club:
- Check ID (condition)
- If age >= 18, let them in (run code)
- Otherwise, don't let them in (skip code)`),
			codeBlock("javascript", `let age = 16;

if (age >= 18) {
    console.log("You can vote!");
}

if (age < 18) {
    console.log("You cannot vote yet.");
}`),
			textBlock(`## If-Else: Either This or That

What if you want to do one thing OR another? Use if-else!`),
			codeBlock("javascript", `let temperature = 75;

if (temperature > 80) {
    console.log("It's hot! Wear shorts.");
} else {
    console.log("It's not too hot. Pants are fine.");
}`),
			textBlock(`## If-Else If-Else: Multiple Choices

Sometimes you have more than two options. You can chain conditions together.`),
			codeBlock("javascript", `let score = 85;

if (score >= 90) {
    console.log("Grade: A");
} else if (score >= 80) {
    console.log("Grade: B");
} else if (score >= 70) {
    console.log("Grade: C");
} else if (score >= 60) {
    console.log("Grade: D");
} else {
    console.log("Grade: F");
}`),
			calloutBlock("info", "Only ONE block will run! Once a condition is true, JavaScript stops checking. In the example above, a score of 85 prints \"Grade: B\" and skips the rest."),
			textBlock(`## Logical Operators: Combining Conditions

Sometimes you need to check multiple things at once.

### && (AND) - Both must be true
Like saying "You can drive if you have a license AND you're not tired."

### || (OR) - At least one must be true
Like saying "You can enter if you're a member OR you have a ticket."

### ! (NOT) - Flips the value
Like saying "If it's NOT raining, we'll have the picnic."`),
			codeBlock("javascript", `let hasTicket = true;
let isVIP = false;
let age = 20;

// AND example
if (hasTicket && age >= 18) {
    console.log("You can enter the concert!");
}

// OR example  
if (hasTicket || isVIP) {
    console.log("Welcome!");
}

// NOT example
let isRaining = true;
if (!isRaining) {
    console.log("Let's go outside!");
}`),
			exerciseBlock(
				"Write a program that checks if someone can ride a roller coaster. They need to be at least 48 inches tall AND either have a ticket OR be a season pass holder.",
				`let height = 50;
let hasTicket = false;
let hasSeasonPass = true;

if (height >= 48 && (hasTicket || hasSeasonPass)) {
    console.log("You can ride!");
} else {
    console.log("Sorry, you cannot ride.");
}`,
				[]string{"Check height first", "Use || for the ticket/pass condition", "Use && to combine them"},
			),
		},
	}
}