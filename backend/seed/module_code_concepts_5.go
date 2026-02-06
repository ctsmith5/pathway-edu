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

This is the most common bug in programming. Arrays start at 0, not 1!`),
			codeBlock("csharp", `string[] colors = { "red", "green", "blue" };

// WRONG - This will give you an error!
Console.WriteLine(colors[3]);  // IndexOutOfRangeException (there is no 4th item)

// CORRECT - The last item is at position 2
Console.WriteLine(colors[2]);  // "blue"

// Better yet, use Length - 1
int lastIndex = colors.Length - 1;
Console.WriteLine(colors[lastIndex]);  // "blue"`),
			textBlock(`**How to avoid it:** Always remember: first item is [0], last item is [Length - 1]`),
			textBlock(`## Mistake #2: Using = Instead of ==

- A single equals (=) assigns a value
- Double equals (==) compares values`),
			codeBlock("csharp", `int x = 5;

// WRONG - This assigns 10 to x, then checks if 10 is true
if (x = 10) {  // This won't compile in C#!
    Console.WriteLine("This always runs!");
}

// CORRECT - Use == to compare
if (x == 10) {
    Console.WriteLine("x equals 10");
} else {
    Console.WriteLine("x is not 10");  // This runs
}`),
			textBlock(`**How to avoid it:** Read your conditions out loud. "If x equals 10" should use ==. In C#, the compiler catches this error, unlike JavaScript!`),
			textBlock(`## Mistake #3: Infinite Loops

Forgetting to update your loop variable causes the loop to run forever!`),
			codeBlock("csharp", `// WRONG - i never changes, so this runs forever!
for (int i = 0; i < 5; i) {
    Console.WriteLine(i);
}

// WRONG - Same problem with while loops
int count = 0;
while (count < 5) {
    Console.WriteLine(count);
    // Oops! Forgot to increment count
}

// CORRECT
for (int i = 0; i < 5; i++) {
    Console.WriteLine(i);
}`),
			textBlock(`**How to avoid it:** Always check that your loop variable will eventually make the condition false.`),
			textBlock(`## Mistake #4: Scope Issues

Variables created inside a function (or loop) can't be used outside of it.`),
			codeBlock("csharp", `int Calculate() {
    int result = 42;
    return result;
}

Calculate();
Console.WriteLine(result);  // ERROR! result is not defined here

// CORRECT - Capture the return value
int answer = Calculate();
Console.WriteLine(answer);  // 42`),
			textBlock(`**How to avoid it:** Remember where you create your variables. If you need a value outside a function, return it and store it in a variable.`),
			textBlock(`## Mistake #5: String vs Number Confusion

C# is stricter than JavaScript, but you can still make this mistake:`),
			codeBlock("csharp", `string a = "5";
string b = "3";

// WRONG - This concatenates (joins) strings!
Console.WriteLine(a + b);  // "53" (not 8!)

// CORRECT - Parse to numbers first
int numA = int.Parse(a);
int numB = int.Parse(b);
Console.WriteLine(numA + numB);  // 8

// Or make sure they're numbers from the start
int c = 5;
int d = 3;
Console.WriteLine(c + d);  // 8`),
			textBlock(`**How to avoid it:** Be careful with user input - it's usually a string. Use int.Parse(), double.Parse(), etc. to convert.`),
			calloutBlock("tip", "When something isn't working, use Console.WriteLine() to print out your variables at different points. This helps you see exactly what's happening!"),
			textBlock(`## Debugging Strategy: The Console is Your Friend

When your code isn't working:

1. **Read the error message carefully** - It usually tells you exactly what's wrong
2. **Add Console.WriteLine() statements** - Print your variables to see their values
3. **Check one thing at a time** - Don't change 5 things at once
4. **Take a break** - Sometimes walking away helps you see the problem
5. **Explain it out loud** - Seriously! Explain your code to a rubber duck (or a friend). Hearing it out loud helps you spot errors.

### Example: Debugging with Console.WriteLine`),
			codeBlock("csharp", `int CalculateTotal(string price, int quantity) {
    Console.WriteLine("Price: " + price);      // Check what price is
    Console.WriteLine("Quantity: " + quantity); // Check what quantity is
    
    int priceNum = int.Parse(price);
    int total = priceNum * quantity;
    Console.WriteLine("Total: " + total);       // Check the result
    
    return total;
}

CalculateTotal("10", 5);  // Works because we parse the string!`),
			exerciseBlock(
				"Find and fix the bug in this code:\n\n"+
				"```csharp\n"+
				"int SumArray(int[] numbers) {\n"+
				"    int sum = 0;\n"+
				"    for (int i = 0; i <= numbers.Length; i++) {\n"+
				"        sum += numbers[i];\n"+
				"    }\n"+
				"    return sum;\n"+
				"}\n\n"+
				"Console.WriteLine(SumArray(new int[] {1, 2, 3}));  // Should be 6, but throws error\n"+
				"```",
				"The bug is in the loop condition: i <= numbers.Length should be i < numbers.Length.\n\n"+
				"When i equals numbers.Length (which is 3), numbers[3] throws IndexOutOfRangeException.\n\n"+
				"Fixed version:\n"+
				"for (int i = 0; i < numbers.Length; i++) {",
				[]string{"Check the loop condition carefully", "What happens when i equals the array length?", "Remember: last valid index is Length - 1"},
			),
		},
	}
}