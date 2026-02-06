package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts3() models.Module {
	return models.Module{
		ID:    "code-3",
		Title: "Repeating Actions",
		Content: []models.ContentBlock{
			textBlock("## Computers Love Repetition\n\nOne of the biggest advantages of computers is that they never get bored doing the same thing over and over. This is where loops come in!\n\n### Why Use Loops?\n\nImagine you need to print \"Hello!\" 100 times. You could write:"),
			multiCodeBlock(map[string]string{
				"csharp": `Console.WriteLine("Hello!");
Console.WriteLine("Hello!");
Console.WriteLine("Hello!");
// ... 97 more times ...`,
				"typescript": `console.log("Hello!");
console.log("Hello!");
console.log("Hello!");
// ... 97 more times ...`,
				"python": `print("Hello!")
print("Hello!")
print("Hello!")
# ... 97 more times ...`,
			}),
			textBlock("Or you could write a loop that does it for you in 3 lines!"),
			textBlock("## For Loops: Counting Repetition\n\nA for loop is perfect when you know exactly how many times you want to repeat something.\n\n### The Anatomy of a For Loop"),
			multiCodeBlock(map[string]string{
				"csharp": `for (int i = 0; i < 5; i++) {
    Console.WriteLine("Iteration: " + i);
}`,
				"typescript": `for (let i = 0; i < 5; i++) {
    console.log("Iteration: " + i);
}`,
				"python": `for i in range(5):
    print("Iteration:", i)`,
			}),
			imageBlock("https://storage.googleapis.com/ludicrousapps-c1ea7.firebasestorage.app/PathwayEdu/1f442d36-7bb3-43fc-b578-3e3d225890fa.png", "Diagram showing the anatomy of a for loop with labeled parts for initialization, condition, and increment", "The three parts of a for loop: initialization, condition, and increment"),
			textBlock("Let's break this down:\n" +
				"1. **Initialization** - Start with a counter variable\n" +
				"2. **Condition** - Keep going while this is true\n" +
				"3. **Increment** - Update the counter after each loop\n" +
				"4. **Body** - The code that runs each time\n\n" +
				"**Output:**"),
			codeBlock("text", `Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4`),
			calloutBlock("info", "Notice Python's range(5) is cleaner but works the same way - it iterates 5 times with values 0-4."),
			textBlock("## Looping Through Arrays/Lists\n\nLoops are super useful for working with collections of items."),
			multiCodeBlock(map[string]string{
				"csharp": `string[] fruits = { "apple", "banana", "cherry", "date" };

// Print each fruit
for (int i = 0; i < fruits.Length; i++) {
    Console.WriteLine(fruits[i]);
}

// Output:
// apple
// banana
// cherry
// date`,
				"typescript": `let fruits: string[] = ["apple", "banana", "cherry", "date"];

// Print each fruit
for (let i = 0; i < fruits.length; i++) {
    console.log(fruits[i]);
}

// Output:
// apple
// banana
// cherry
// date`,
				"python": `fruits = ["apple", "banana", "cherry", "date"]

# Print each fruit (Pythonic way)
for fruit in fruits:
    print(fruit)

# Or with index
for i in range(len(fruits)):
    print(fruits[i])

# Output:
# apple
# banana
# cherry
# date`,
			}),
			textBlock("**What's happening:**\n" +
				"- We access each item by its position (index)\n" +
				"- Index starts at 0\n" +
				"- The loop runs once for each item"),
			textBlock("## While Loops: Conditional Repetition\n\nA while loop keeps going as long as a condition is true. Use this when you DON'T know how many times you need to repeat.\n\n### Real-World Analogy\n" +
				"\"While the light is red, wait.\" You don't know how many seconds - you just keep checking."),
			multiCodeBlock(map[string]string{
				"csharp": `int count = 0;

while (count < 5) {
    Console.WriteLine("Count is: " + count);
    count++;
}`,
				"typescript": `let count = 0;

while (count < 5) {
    console.log("Count is: " + count);
    count++;
}`,
				"python": `count = 0

while count < 5:
    print("Count is:", count)
    count += 1`,
			}),
			textBlock("**Important:** You MUST change something inside the loop, or it will run forever!"),
			calloutBlock("warning", "An infinite loop happens when the condition never becomes false. This will crash your program! Always make sure the loop will eventually end."),
			textBlock("## Practical Example: Finding Something\n\nLet's say we have a list of students and we want to find if \"Alice\" is in it:"),
			multiCodeBlock(map[string]string{
				"csharp": `string[] students = { "Bob", "Alice", "Charlie", "Diana" };
bool found = false;

for (int i = 0; i < students.Length; i++) {
    if (students[i] == "Alice") {
        found = true;
        Console.WriteLine("Found Alice at position " + i);
        break;  // Exit the loop early!
    }
}

if (!found) {
    Console.WriteLine("Alice is not in the list.");
}`,
				"typescript": `let students: string[] = ["Bob", "Alice", "Charlie", "Diana"];
let found = false;

for (let i = 0; i < students.length; i++) {
    if (students[i] === "Alice") {
        found = true;
        console.log("Found Alice at position " + i);
        break;  // Exit the loop early!
    }
}

if (!found) {
    console.log("Alice is not in the list.");
}`,
				"python": `students = ["Bob", "Alice", "Charlie", "Diana"]
found = False

for i in range(len(students)):
    if students[i] == "Alice":
        found = True
        print("Found Alice at position", i)
        break  # Exit the loop early!

if not found:
    print("Alice is not in the list.")`,
			}),
			textBlock("**The 'break' statement** immediately exits the loop. We use it here because once we found Alice, we don't need to keep looking!"),
			exerciseBlock(
				"Write a loop that calculates the sum of all numbers from 1 to 100. Hint: Create a variable to keep track of the total, then add each number to it in the loop.",
				`C#:
int total = 0;

for (int i = 1; i <= 100; i++) {
    total = total + i;
}

Console.WriteLine("The sum is: " + total);  // Should be 5050

TypeScript:
let total = 0;

for (let i = 1; i <= 100; i++) {
    total = total + i;
}

console.log("The sum is: " + total);  // Should be 5050

Python:
total = 0

for i in range(1, 101):
    total = total + i

print("The sum is:", total)  # Should be 5050`,
				[]string{"Start total at 0", "Loop from 1 to 100", "Add i to total each time"},
			),
		},
	}
}