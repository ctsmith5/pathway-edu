package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts3() models.Module {
	return models.Module{
		ID:    "code-3",
		Title: "Repeating Actions",
		Content: []models.ContentBlock{
			textBlock(`## Computers Love Repetition

One of the biggest advantages of computers is that they never get bored doing the same thing over and over. This is where loops come in!

### Why Use Loops?

Imagine you need to print "Hello!" 100 times. You could write:
` + "```javascript\n" + `console.log("Hello!");
console.log("Hello!");
console.log("Hello!");
// ... 97 more times ...
` + "```" + `

Or you could write a loop that does it for you in 3 lines!`),
			textBlock(`## For Loops: Counting Repetition

A for loop is perfect when you know exactly how many times you want to repeat something.

### The Anatomy of a For Loop

` + "```javascript\n" + `for (let i = 0; i < 5; i++) {
    console.log("Iteration: " + i);
}
` + "```" + `

Let's break this down:
1. **let i = 0** - Start with i equal to 0 (initialization)
2. **i < 5** - Keep going while i is less than 5 (condition)
3. **i++** - Add 1 to i after each loop (increment)
4. **{ ... }** - The code that runs each time

**Output:**
` + "```\n" + `Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
` + "```"),
			calloutBlock("info", "The variable 'i' is just a convention. It stands for 'index' or 'iterator'. You can use any name you want, like 'count' or 'step'."),
			textBlock(`## Looping Through Arrays

Loops are super useful for working with arrays (lists of items).

` + "```javascript\n" + `let fruits = ["apple", "banana", "cherry", "date"];

// Print each fruit
for (let i = 0; i < fruits.length; i++) {
    console.log(fruits[i]);
}

// Output:
// apple
// banana
// cherry
// date
` + "```" + `

**What's happening:**
- `fruits.length` gives us 4 (there are 4 items)
- `fruits[i]` gets the item at position i
- We go from i=0 to i=3 (all 4 items)`),
			textBlock(`## While Loops: Conditional Repetition

A while loop keeps going as long as a condition is true. Use this when you DON'T know how many times you need to repeat.

### Real-World Analogy
"While the light is red, wait." You don't know how many seconds - you just keep checking.

` + "```javascript\n" + `let count = 0;

while (count < 5) {
    console.log("Count is: " + count);
    count++;
}
` + "```" + `

**Important:** You MUST change something inside the loop, or it will run forever! In the example above, we increment count each time.`),
			calloutBlock("warning", "An infinite loop happens when the condition never becomes false. This will crash your program! Always make sure the loop will eventually end."),
			textBlock(`## Practical Example: Finding Something

Let's say we have a list of students and we want to find if "Alice" is in it:

` + "```javascript\n" + `let students = ["Bob", "Alice", "Charlie", "Diana"];
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
}
` + "```" + `

**The 'break' statement** immediately exits the loop. We use it here because once we found Alice, we don't need to keep looking!`),
			exerciseBlock(
				"Write a loop that calculates the sum of all numbers from 1 to 100. Hint: Create a variable to keep track of the total, then add each number to it in the loop.",
				`let total = 0;

for (let i = 1; i <= 100; i++) {
    total = total + i;
}

console.log("The sum is: " + total);  // Should be 5050`,
				[]string{"Start total at 0", "Loop from 1 to 100", "Add i to total each time"},
			),
		},
	}
}
