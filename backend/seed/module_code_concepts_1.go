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
			codeBlock("csharp", `// Creating variables
string myName = "Alex";
int myAge = 16;
bool isStudent = true;

// Using variables
Console.WriteLine(myName);     // Prints: Alex
Console.WriteLine(myAge);      // Prints: 16

// Changing a variable
myAge = 17;
Console.WriteLine(myAge);      // Prints: 17`),
			textBlock(`## Data Types: Different Kinds of Information

Computers need to know what kind of data they're working with. Here are the most common types:

### 1. Strings (Text)
Strings are text - anything in quotes. Think of it like a message in a bottle.`),
			codeBlock("csharp", `// Strings are always in quotes
string firstName = "Maria";
string favoriteFood = "Pizza";
string message = "Hello, World!";

// You can combine strings
string greeting = "Hi, " + firstName;  // "Hi, Maria"`),
			textBlock(`### 2. Numbers
Numbers can be whole numbers (integers) or decimals (floating point).`),
			codeBlock("csharp", `int age = 16;           // Integer (whole number)
double price = 19.99;   // Double (decimal)
int temperature = -5;   // Negative numbers work too

// Math operations
int sum = 10 + 5;       // 15
int product = 4 * 3;    // 12`),
			textBlock(`### 3. Booleans (True/False)
Booleans are simple: they're either true or false. Like a light switch - it's either on or off.`),
			codeBlock("csharp", `bool isLoggedIn = true;
bool hasPermission = false;
bool isRaining = true;`),
			textBlock(`### 4. Arrays (Lists)
An array is a list of items. Think of it like a shopping list or a row of lockers.`),
			codeBlock("csharp", `string[] colors = { "red", "green", "blue" };
int[] scores = { 95, 87, 92, 88 };

// Access items by position (starts at 0!)
Console.WriteLine(colors[0]);  // "red"
Console.WriteLine(colors[2]);  // "blue"`),
			calloutBlock("warning", "Arrays start counting at 0, not 1! The first item is at position 0, the second at position 1, etc. This is a common source of bugs for beginners."),
			exerciseBlock(
				"Create variables to store your name, age, and whether you like pizza (true/false). Then create an array of your three favorite movies. Print them all to the console.",
				`string myName = "Your Name";
int myAge = 16;
bool likesPizza = true;
string[] favoriteMovies = { "Movie 1", "Movie 2", "Movie 3" };

Console.WriteLine(myName);
Console.WriteLine(myAge);
Console.WriteLine(likesPizza);
foreach (var movie in favoriteMovies) {
    Console.WriteLine(movie);
}`,
				[]string{"Use the type (string, int, bool) to create variables", "Remember to put strings in quotes", "Arrays use curly braces {} in C#"},
			),
		},
	}
}