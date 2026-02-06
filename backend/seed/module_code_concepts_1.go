package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts1() models.Module {
	return models.Module{
		ID:    "code-1",
		Title: "What is Code?",
		Content: []models.ContentBlock{
			textBlock(`## Welcome to Programming!

Programming is just giving instructions to a computer. Think of it like writing a recipe - you're telling the computer exactly what to do, step by step.

### Why Learn to Code?

- **Create things**: Build websites, apps, games, and tools
- **Solve problems**: Automate boring tasks, analyze data
- **Express yourself**: Code is a creative medium, like art or music
- **Career opportunities**: Programming skills are in high demand

But before we can give instructions, we need to understand the basic building blocks.`),
			calloutBlock("tip", "Don't worry if things don't make sense immediately. Programming is a skill you build over time, like learning to play an instrument."),
			textBlock(`## Variables: Storing Information

A **variable** is like a labeled box where you can store information. You can put something in the box, look at what's inside, or replace it with something else.

Think of it like this:
- A variable name is the label on the box (like "myAge")
- The value is what's inside the box (like the number 16)
- The data type is what kind of thing can go in the box (numbers, text, etc.)`),
			codeBlock("javascript", `// Creating variables
let myName = "Alex";
let myAge = 16;
let isStudent = true;

// Using variables
console.log(myName);     // Prints: Alex
console.log(myAge);      // Prints: 16

// Changing a variable
myAge = 17;
console.log(myAge);      // Prints: 17`),
			textBlock(`## Data Types: Different Kinds of Information

Computers need to know what kind of data they're working with. Here are the most common types:

### 1. Strings (Text)
Strings are text - anything in quotes. Think of it like a message in a bottle.`),
			codeBlock("javascript", `// Strings are always in quotes
let firstName = "Maria";
let favoriteFood = 'Pizza';
let message = "Hello, World!";

// You can combine strings
let greeting = "Hi, " + firstName;  // "Hi, Maria"`),
			textBlock(`### 2. Numbers
Numbers can be whole numbers (integers) or decimals (floating point).`),
			codeBlock("javascript", `let age = 16;           // Integer (whole number)
let price = 19.99;      // Float (decimal)
let temperature = -5;   // Negative numbers work too

// Math operations
let sum = 10 + 5;       // 15
let product = 4 * 3;    // 12`),
			textBlock(`### 3. Booleans (True/False)
Booleans are simple: they're either true or false. Like a light switch - it's either on or off.`),
			codeBlock("javascript", `let isLoggedIn = true;
let hasPermission = false;
let isRaining = true;`),
			textBlock(`### 4. Arrays (Lists)
An array is a list of items. Think of it like a shopping list or a row of lockers.`),
			codeBlock("javascript", `let colors = ["red", "green", "blue"];
let scores = [95, 87, 92, 88];

// Access items by position (starts at 0!)
console.log(colors[0]);  // "red"
console.log(colors[2]);  // "blue"`),
			calloutBlock("warning", "Arrays start counting at 0, not 1! The first item is at position 0, the second at position 1, etc. This is a common source of bugs for beginners."),
			exerciseBlock(
				"Create variables to store your name, age, and whether you like pizza (true/false). Then create an array of your three favorite movies. Print them all to the console.",
				`let myName = "Your Name";
let myAge = 16;
let likesPizza = true;
let favoriteMovies = ["Movie 1", "Movie 2", "Movie 3"];

console.log(myName);
console.log(myAge);
console.log(likesPizza);
console.log(favoriteMovies);`,
				[]string{"Use let to create variables", "Remember to put strings in quotes", "Arrays go in square brackets []"},
			),
		},
	}
}