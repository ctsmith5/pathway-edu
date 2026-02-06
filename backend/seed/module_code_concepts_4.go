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
			multiCodeBlock(map[string]string{
				"csharp": `// ReturnType FunctionName(Parameters)
// {
//     // Code to run
//     return value;  // Optional
// }`,
				"typescript": `// function functionName(parameters): ReturnType {
//     // Code to run
//     return value;  // Optional
// }`,
				"python": `# def function_name(parameters):
#     # Code to run
#     return value  # Optional`,
			}),
			textBlock(`### Simple Example: Say Hello`),
			multiCodeBlock(map[string]string{
				"csharp": `void SayHello() {
    Console.WriteLine("Hello there!");
}

// Using (calling) the function
SayHello();  // Prints: Hello there!
SayHello();  // Prints: Hello there!`,
				"typescript": `function sayHello(): void {
    console.log("Hello there!");
}

// Using (calling) the function
sayHello();  // Prints: Hello there!
sayHello();  // Prints: Hello there!`,
				"python": `def say_hello():
    print("Hello there!")

# Using (calling) the function
say_hello()  # Prints: Hello there!
say_hello()  # Prints: Hello there!`,
			}),
			textBlock(`We defined the function once, then used it twice!`),
			textBlock(`## Functions with Parameters

Parameters let you pass information INTO the function.`),
			multiCodeBlock(map[string]string{
				"csharp": `void Greet(string name) {
    Console.WriteLine("Hello, " + name + "!");
}

Greet("Alice");   // Hello, Alice!
Greet("Bob");     // Hello, Bob!
Greet("Charlie"); // Hello, Charlie!`,
				"typescript": `function greet(name: string): void {
    console.log("Hello, " + name + "!");
}

greet("Alice");   // Hello, Alice!
greet("Bob");     // Hello, Bob!
greet("Charlie"); // Hello, Charlie!`,
				"python": `def greet(name):
    print("Hello, " + name + "!")

greet("Alice")   # Hello, Alice!
greet("Bob")     # Hello, Bob!
greet("Charlie") # Hello, Charlie!`,
			}),
			textBlock(`You can have multiple parameters:`),
			multiCodeBlock(map[string]string{
				"csharp": `void Add(int a, int b) {
    Console.WriteLine(a + b);
}

Add(3, 5);   // 8
Add(10, 20); // 30`,
				"typescript": `function add(a: number, b: number): void {
    console.log(a + b);
}

add(3, 5);   // 8
add(10, 20); // 30`,
				"python": `def add(a, b):
    print(a + b)

add(3, 5)   # 8
add(10, 20) # 30`,
			}),
			textBlock(`## Functions that Return Values

Sometimes you want the function to give you something back. Use the **return** keyword.`),
			multiCodeBlock(map[string]string{
				"csharp": `int Multiply(int a, int b) {
    return a * b;
}

int result = Multiply(4, 5);
Console.WriteLine(result);  // 20

// You can use it directly too
Console.WriteLine(Multiply(3, 7));  // 21`,
				"typescript": `function multiply(a: number, b: number): number {
    return a * b;
}

let result = multiply(4, 5);
console.log(result);  // 20

// You can use it directly too
console.log(multiply(3, 7));  // 21`,
				"python": `def multiply(a, b):
    return a * b

result = multiply(4, 5)
print(result)  # 20

# You can use it directly too
print(multiply(3, 7))  # 21`,
			}),
			textBlock(`**Important:** Once a function hits 'return', it stops immediately and gives back the value.`),
			multiCodeBlock(map[string]string{
				"csharp": `string CheckAge(int age) {
    if (age >= 18) {
        return "Adult";
    }
    return "Minor";  // This only runs if age < 18
}

Console.WriteLine(CheckAge(20));  // Adult
Console.WriteLine(CheckAge(15));  // Minor`,
				"typescript": `function checkAge(age: number): string {
    if (age >= 18) {
        return "Adult";
    }
    return "Minor";  // This only runs if age < 18
}

console.log(checkAge(20));  // Adult
console.log(checkAge(15));  // Minor`,
				"python": `def check_age(age):
    if age >= 18:
        return "Adult"
    return "Minor"  # This only runs if age < 18

print(check_age(20))  # Adult
print(check_age(15))  # Minor`,
			}),
			calloutBlock("tip", "Functions should do ONE thing well. If your function is doing many different things, break it into smaller functions!"),
			textBlock(`## Putting It Together: A Real Example

Let's calculate the area of a rectangle, but make it reusable:`),
			multiCodeBlock(map[string]string{
				"csharp": `int CalculateArea(int width, int height) {
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
Console.WriteLine("Room perimeter: " + perimeter + " feet");`,
				"typescript": `function calculateArea(width: number, height: number): number {
    return width * height;
}

function calculatePerimeter(width: number, height: number): number {
    return 2 * (width + height);
}

// Using our functions
let roomWidth = 10;
let roomHeight = 8;

let area = calculateArea(roomWidth, roomHeight);
let perimeter = calculatePerimeter(roomWidth, roomHeight);

console.log("Room area: " + area + " square feet");
console.log("Room perimeter: " + perimeter + " feet");`,
				"python": `def calculate_area(width, height):
    return width * height

def calculate_perimeter(width, height):
    return 2 * (width + height)

# Using our functions
room_width = 10
room_height = 8

area = calculate_area(room_width, room_height)
perimeter = calculate_perimeter(room_width, room_height)

print("Room area:", area, "square feet")
print("Room perimeter:", perimeter, "feet")`,
			}),
			textBlock(`Now we can calculate area and perimeter for ANY rectangle without rewriting the math!`),
			exerciseBlock(
				"Write a function called 'IsEven' that takes a number and returns true if it's even, false if it's odd. Then test it with the numbers 4, 7, and 10.",
				`C#:
bool IsEven(int number) {
    return number % 2 == 0;
}

Console.WriteLine(IsEven(4));   // true
Console.WriteLine(IsEven(7));   // false
Console.WriteLine(IsEven(10));  // true

TypeScript:
function isEven(number: number): boolean {
    return number % 2 === 0;
}

console.log(isEven(4));   // true
console.log(isEven(7));   // false
console.log(isEven(10));  // true

Python:
def is_even(number):
    return number % 2 == 0

print(is_even(4))   # True
print(is_even(7))   # False
print(is_even(10))  # True`,
				[]string{"Use the % operator (gives remainder)", "If number % 2 equals 0, it's even", "Return true/True or false/False"},
			),
		},
	}
}