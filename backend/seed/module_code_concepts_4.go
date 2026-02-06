package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts4() models.Module {
	return models.Module{
		ID:    "code-4",
		Title: "Organizing Code with Functions",
		Content: []models.ContentBlock{
			textBlock(`## What is a Function?

A function is a reusable block of code that performs a specific task. Think of it like a recipe - you write it once, then you can use it whenever you need it.

### Why Use Functions?

1. **Reusability**: Write once, use many times
2. **Organization**: Break big problems into smaller pieces
3. **Readability**: Give complex code a simple name
4. **Testing**: Test small parts individually

### Real-World Analogy
A function is like a toaster:
- **Input**: You put in bread (parameters/arguments)
- **Process**: The toaster does its thing (function body)
- **Output**: You get toast (return value)

You don't need to know HOW the toaster works - you just use it!`),
			textBlock(`## Creating a Function

Here's the basic structure:`),
			codeBlock("csharp", `// ReturnType FunctionName(Parameters)
// {
//     // Code to run
//     return value;  // Optional
// }`),
			textBlock(`### Simple Example: Say Hello`),
			codeBlock("csharp", `void SayHello() {
    Console.WriteLine("Hello there!");
}

// Using (calling) the function
SayHello();  // Prints: Hello there!
SayHello();  // Prints: Hello there!`),
			textBlock(`We defined the function once, then used it twice!`),
			textBlock(`## Functions with Parameters

Parameters let you pass information INTO the function.`),
			codeBlock("csharp", `void Greet(string name) {
    Console.WriteLine("Hello, " + name + "!");
}

Greet("Alice");   // Hello, Alice!
Greet("Bob");     // Hello, Bob!
Greet("Charlie"); // Hello, Charlie!`),
			textBlock(`You can have multiple parameters:`),
			codeBlock("csharp", `void Add(int a, int b) {
    Console.WriteLine(a + b);
}

Add(3, 5);   // 8
Add(10, 20); // 30`),
			textBlock(`## Functions that Return Values

Sometimes you want the function to give you something back. Use the **return** keyword.`),
			codeBlock("csharp", `int Multiply(int a, int b) {
    return a * b;
}

int result = Multiply(4, 5);
Console.WriteLine(result);  // 20

// You can use it directly too
Console.WriteLine(Multiply(3, 7));  // 21`),
			textBlock(`**Important:** Once a function hits 'return', it stops immediately and gives back the value.`),
			codeBlock("csharp", `string CheckAge(int age) {
    if (age >= 18) {
        return "Adult";
    }
    return "Minor";  // This only runs if age < 18
}

Console.WriteLine(CheckAge(20));  // Adult
Console.WriteLine(CheckAge(15));  // Minor`),
			calloutBlock("tip", "Functions should do ONE thing well. If your function is doing many different things, break it into smaller functions!"),
			textBlock(`## Putting It Together: A Real Example

Let's calculate the area of a rectangle, but make it reusable:`),
			codeBlock("csharp", `int CalculateArea(int width, int height) {
    return width * height;
}

int CalculatePerimeter(int width, int height) {
    return 2 * (width + height);
}

// Using our functions
int roomWidth = 10;
int roomHeight = 8;

int area = CalculateArea(roomWidth, roomHeight);
int perimeter = CalculatePerimeter(roomWidth, roomHeight);

Console.WriteLine("Room area: " + area + " square feet");
Console.WriteLine("Room perimeter: " + perimeter + " feet");`),
			textBlock(`Now we can calculate area and perimeter for ANY rectangle without rewriting the math!`),
			exerciseBlock(
				"Write a function called 'IsEven' that takes a number and returns true if it's even, false if it's odd. Then test it with the numbers 4, 7, and 10.",
				`bool IsEven(int number) {
    return number % 2 == 0;
}

Console.WriteLine(IsEven(4));   // true
Console.WriteLine(IsEven(7));   // false
Console.WriteLine(IsEven(10));  // true`,
				[]string{"Use the % operator (gives remainder)", "If number % 2 equals 0, it's even", "Return true or false"},
			),
		},
	}
}