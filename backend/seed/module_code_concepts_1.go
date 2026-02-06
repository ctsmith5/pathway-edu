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
			multiCodeBlock(map[string]string{
				"csharp": `// Creating variables
string myName = "Alex";
int myAge = 16;
bool isStudent = true;

// Using variables
Console.WriteLine(myName);     // Prints: Alex
Console.WriteLine(myAge);      // Prints: 16

// Changing a variable
myAge = 17;
Console.WriteLine(myAge);      // Prints: 17`,
				"typescript": `// Creating variables
let myName: string = "Alex";
let myAge: number = 16;
let isStudent: boolean = true;

// Using variables
console.log(myName);     // Prints: Alex
console.log(myAge);      // Prints: 16

// Changing a variable
myAge = 17;
console.log(myAge);      // Prints: 17`,
				"python": `# Creating variables
my_name = "Alex"
my_age = 16
is_student = True

# Using variables
print(my_name)     # Prints: Alex
print(my_age)      # Prints: 16

# Changing a variable
my_age = 17
print(my_age)      # Prints: 17`,
			}),
			textBlock(`## Data Types: Different Kinds of Information

Computers need to know what kind of data they're working with. Here are the most common types:

### 1. Strings (Text)
Strings are text - anything in quotes. Think of it like a message in a bottle.`),
			multiCodeBlock(map[string]string{
				"csharp": `// Strings are always in quotes
string firstName = "Maria";
string favoriteFood = "Pizza";
string message = "Hello, World!";

// You can combine strings
string greeting = "Hi, " + firstName;  // "Hi, Maria"`,
				"typescript": `// Strings are always in quotes
let firstName: string = "Maria";
let favoriteFood: string = "Pizza";
let message: string = "Hello, World!";

// You can combine strings
let greeting = "Hi, " + firstName;  // "Hi, Maria"`,
				"python": `# Strings are always in quotes
first_name = "Maria"
favorite_food = 'Pizza'
message = "Hello, World!"

# You can combine strings
greeting = "Hi, " + first_name  # "Hi, Maria"`,
			}),
			textBlock(`### 2. Numbers
Numbers can be whole numbers (integers) or decimals (floating point).`),
			multiCodeBlock(map[string]string{
				"csharp": `int age = 16;           // Integer (whole number)
double price = 19.99;   // Double (decimal)
int temperature = -5;   // Negative numbers work too

// Math operations
int sum = 10 + 5;       // 15
int product = 4 * 3;    // 12`,
				"typescript": `let age: number = 16;           // Integer
let price: number = 19.99;      // Decimal
let temperature: number = -5;   // Negative numbers work too

// Math operations
let sum = 10 + 5;       // 15
let product = 4 * 3;    // 12`,
				"python": `age = 16               # Integer
price = 19.99          # Float (decimal)
temperature = -5       # Negative numbers work too

# Math operations
sum = 10 + 5           # 15
product = 4 * 3        # 12`,
			}),
			textBlock(`### 3. Booleans (True/False)
Booleans are simple: they're either true or false. Like a light switch - it's either on or off.`),
			multiCodeBlock(map[string]string{
				"csharp": `bool isLoggedIn = true;
bool hasPermission = false;
bool isRaining = true;`,
				"typescript": `let isLoggedIn: boolean = true;
let hasPermission: boolean = false;
let isRaining: boolean = true;`,
				"python": `is_logged_in = True
has_permission = False
is_raining = True`,
			}),
			textBlock(`### 4. Arrays/Lists
An array is a list of items. Think of it like a shopping list or a row of lockers.`),
			multiCodeBlock(map[string]string{
				"csharp": `string[] colors = { "red", "green", "blue" };
int[] scores = { 95, 87, 92, 88 };

// Access items by position (starts at 0!)
Console.WriteLine(colors[0]);  // "red"
Console.WriteLine(colors[2]);  // "blue"`,
				"typescript": `let colors: string[] = ["red", "green", "blue"];
let scores: number[] = [95, 87, 92, 88];

// Access items by position (starts at 0!)
console.log(colors[0]);  // "red"
console.log(colors[2]);  // "blue"`,
				"python": `colors = ["red", "green", "blue"]
scores = [95, 87, 92, 88]

# Access items by position (starts at 0!)
print(colors[0])  # "red"
print(colors[2])  # "blue"`,
			}),
			calloutBlock("warning", "Arrays start counting at 0, not 1! The first item is at position 0, the second at position 1, etc. This is a common source of bugs for beginners."),
			exerciseBlock(
				"Create variables to store your name, age, and whether you like pizza (true/false). Then create an array/list of your three favorite movies. Print them all to the console.",
				`C#:
string myName = "Your Name";
int myAge = 16;
bool likesPizza = true;
string[] favoriteMovies = { "Movie 1", "Movie 2", "Movie 3" };

Console.WriteLine(myName);
Console.WriteLine(myAge);
Console.WriteLine(likesPizza);
foreach (var movie in favoriteMovies) {
    Console.WriteLine(movie);
}

TypeScript:
let myName: string = "Your Name";
let myAge: number = 16;
let likesPizza: boolean = true;
let favoriteMovies: string[] = ["Movie 1", "Movie 2", "Movie 3"];

console.log(myName);
console.log(myAge);
console.log(likesPizza);
favoriteMovies.forEach(movie => console.log(movie));

Python:
my_name = "Your Name"
my_age = 16
likes_pizza = True
favorite_movies = ["Movie 1", "Movie 2", "Movie 3"]

print(my_name)
print(my_age)
print(likes_pizza)
for movie in favorite_movies:
    print(movie)`,
				[]string{"Use the appropriate type for your language", "Remember to put strings in quotes", "Arrays use [] in TypeScript/Python, {} in C#"},
			),
		},
	}
}