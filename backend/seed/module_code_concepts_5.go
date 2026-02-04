package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts5() models.Module {
	return models.Module{
		ID:    "code-5",
		Title: "Common Mistakes Beginners Make",
		Content: []models.ContentBlock{
			textBlock(`## Everyone Makes Mistakes

Programming is hard, and every programmer - even professionals with 20 years of experience - makes mistakes every day. The difference is knowing how to spot and fix them!

Here are the most common mistakes beginners make, and how to avoid them.`),
			calloutBlock("info", "Mistakes are how you learn! Don't get discouraged. Every bug you fix makes you a better programmer."),
			textBlock(`## Mistake #1: Off-by-One Errors

This is the most common bug in programming. Arrays start at 0, not 1!

` + "```javascript\n" + `let colors = ["red", "green", "blue"];

// WRONG - This will give you undefined!
console.log(colors[3]);  // undefined (there is no 4th item)

// CORRECT - The last item is at position 2
console.log(colors[2]);  // "blue"

// Better yet, use length - 1
let lastIndex = colors.length - 1;
console.log(colors[lastIndex]);  // "blue"
` + "```" + `
**How to avoid it:** Always remember: first item is [0], last item is [length - 1]`),
			textBlock(`## Mistake #2: Using = Instead of ===

A single equals (=) assigns a value. Triple equals (===) compares values.

` + "```javascript\n" + `let x = 5;

// WRONG - This assigns 10 to x, then checks if 10 is truthy
if (x = 10) {
    console.log("This always runs!");
}

// CORRECT - This actually compares
if (x === 10) {
    console.log("x is 10");
} else {
    console.log("x is not 10");  // This runs
}
` + "```" + `
**How to avoid it:** Read your conditions out loud. \"If x equals 10\" should use ===.`),
			textBlock(`## Mistake #3: Infinite Loops

Forgetting to update your loop variable causes the loop to run forever!

` + "```javascript\n" + `// WRONG - i never changes, so this runs forever!
for (let i = 0; i < 5; i) {
    console.log(i);
}

// WRONG - Same problem with while loops
let count = 0;
while (count < 5) {
    console.log(count);
    // Oops! Forgot to increment count
}

// CORRECT
for (let i = 0; i < 5; i++) {
    console.log(i);
}
` + "```" + `
**How to avoid it:** Always check that your loop variable will eventually make the condition false.`),
			textBlock(`## Mistake #4: Scope Issues

Variables created inside a function (or loop) can't be used outside of it.

` + "```javascript\n" + `function calculate() {
    let result = 42;
    return result;
}

calculate();
console.log(result);  // ERROR! result is not defined here

// CORRECT - Capture the return value
let answer = calculate();
console.log(answer);  // 42
` + "```" + `
**How to avoid it:** Remember where you create your variables. If you need a value outside a function, return it and store it in a variable.`),
			textBlock(`## Mistake #5: String vs Number Confusion

JavaScript can be tricky with numbers that look like strings.

` + "```javascript\n" + `let a = "5";
let b = "3";

// WRONG - This concatenates (joins) strings!
console.log(a + b);  // "53" (not 8!)

// CORRECT - Convert to numbers first
console.log(Number(a) + Number(b));  // 8

// Or make sure they're numbers from the start
let c = 5;
let d = 3;
console.log(c + d);  // 8
` + "```" + `
**How to avoid it:** Be careful with user input - it's usually a string even if it looks like a number. Use Number() to convert.`),
			calloutBlock("tip", "When something isn't working, use console.log() to print out your variables at different points. This helps you see exactly what's happening!"),
			textBlock(`## Debugging Strategy: The Console is Your Friend

When your code isn't working:

1. **Read the error message carefully** - It usually tells you exactly what's wrong
2. **Add console.log() statements** - Print your variables to see their values
3. **Check one thing at a time** - Don't change 5 things at once
4. **Take a break** - Sometimes walking away helps you see the problem
5. **Explain it out loud** - Seriously! Explain your code to a rubber duck (or a friend). Hearing it out loud helps you spot errors.

### Example: Debugging with console.log

` + "```javascript\n" + `function calculateTotal(price, quantity) {
    console.log("Price:", price);      // Check what price is
    console.log("Quantity:", quantity); // Check what quantity is
    
    let total = price * quantity;
    console.log("Total:", total);       // Check the result
    
    return total;
}

calculateTotal("10", 5);  // Oops! Price is a string!
` + "```"),
			exerciseBlock(
				"Find and fix the bug in this code:",
				`function sumArray(numbers) {
    let sum = 0;
    for (let i = 0; i <= numbers.length; i++) {
        sum += numbers[i];
    }
    return sum;
}

console.log(sumArray([1, 2, 3]));  // Should be 6, but gets undefined`,
				`The bug is in the loop condition: i <= numbers.length should be i < numbers.length.

When i equals numbers.length (which is 3), numbers[3] is undefined.
undefined + number = NaN (Not a Number).

Fixed version:
for (let i = 0; i < numbers.length; i++) {`,
				[]string{"Check the loop condition carefully", "What happens when i equals the array length?", "Remember: last valid index is length - 1"},
			),
		},
	}
}
