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

This is the most common bug in programming. Arrays/lists start at 0, not 1!`),
			multiCodeBlock(map[string]string{
				"csharp": `string[] colors = { "red", "green", "blue" };

// WRONG - This will give you an error!
Console.WriteLine(colors[3]);  // IndexOutOfRangeException

// CORRECT - The last item is at position 2
Console.WriteLine(colors[2]);  // "blue"

// Better yet, use Length - 1
int lastIndex = colors.Length - 1;
Console.WriteLine(colors[lastIndex]);  // "blue"`,
				"typescript": `let colors: string[] = ["red", "green", "blue"];

// WRONG - This will give you undefined!
console.log(colors[3]);  // undefined

// CORRECT - The last item is at position 2
console.log(colors[2]);  // "blue"

// Better yet, use length - 1
let lastIndex = colors.length - 1;
console.log(colors[lastIndex]);  // "blue"`,
				"python": `colors = ["red", "green", "blue"]

# WRONG - This will give you an IndexError!
print(colors[3])  # IndexError

# CORRECT - The last item is at position 2
print(colors[2])  # "blue"

# Better yet, use negative indexing
print(colors[-1])  # "blue" (Python only!)`,
			}),
			textBlock(`**How to avoid it:** Always remember: first item is [0], last item is [Length - 1] (or [-1] in Python!)`),
			textBlock(`## Mistake #2: Using = Instead of ==

- A single equals (=) assigns a value
- Double equals (==) compares values`),
			multiCodeBlock(map[string]string{
				"csharp": `int x = 5;

// WRONG - This won't compile in C#
if (x = 10) {
    Console.WriteLine("This always runs!");
}

// CORRECT - Use == to compare
if (x == 10) {
    Console.WriteLine("x equals 10");
} else {
    Console.WriteLine("x is not 10");  // This runs
}

// C# catches this error at compile time!`,
				"typescript": `let x = 5;

// WRONG - This assigns 10 to x, then checks if 10 is truthy
if (x = 10) {
    console.log("This always runs!");
}

// CORRECT - Use === to compare
if (x === 10) {
    console.log("x equals 10");
} else {
    console.log("x is not 10");  // This runs
}`,
				"python": `x = 5

# WRONG - This assigns 10 to x, always evaluates to True
if x = 10:  # SyntaxError in Python!
    print("This always runs!")

# CORRECT - Use == to compare
if x == 10:
    print("x equals 10")
else:
    print("x is not 10")  # This runs

# Python catches this error!`,
			}),
			textBlock(`**How to avoid it:** C# and Python catch this error at compile/parse time. TypeScript might not, so use === for strict comparison.`),
			textBlock(`## Mistake #3: Infinite Loops

Forgetting to update your loop variable causes the loop to run forever!`),
			multiCodeBlock(map[string]string{
				"csharp": `// WRONG - i never changes, so this runs forever!
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
}`,
				"typescript": `// WRONG - i never changes, so this runs forever!
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
}`,
				"python": `# WRONG - i never changes, so this runs forever!
i = 0
while i < 5:
    print(i)
    # Oops! Forgot to increment i

# CORRECT
for i in range(5):
    print(i)

# Or with while
i = 0
while i < 5:
    print(i)
    i += 1`,
			}),
			textBlock(`**How to avoid it:** Always check that your loop variable will eventually make the condition false.`),
			textBlock(`## Mistake #4: Scope Issues

Variables created inside a function can't be used outside of it.`),
			multiCodeBlock(map[string]string{
				"csharp": `int Calculate() {
    int result = 42;
    return result;
}

Calculate();
Console.WriteLine(result);  // ERROR! result is not defined here

// CORRECT - Capture the return value
int answer = Calculate();
Console.WriteLine(answer);  // 42`,
				"typescript": `function calculate() {
    let result = 42;
    return result;
}

calculate();
console.log(result);  // ERROR! result is not defined here

// CORRECT - Capture the return value
let answer = calculate();
console.log(answer);  // 42`,
				"python": `def calculate():
    result = 42
    return result

calculate()
print(result)  # ERROR! result is not defined here

# CORRECT - Capture the return value
answer = calculate()
print(answer)  # 42`,
			}),
			textBlock(`**How to avoid it:** Remember where you create your variables. If you need a value outside a function, return it and store it in a variable.`),
			textBlock(`## Mistake #5: String vs Number Confusion`),
			multiCodeBlock(map[string]string{
				"csharp": `string a = "5";
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
Console.WriteLine(c + d);  // 8`,
				"typescript": `let a = "5";
let b = "3";

// WRONG - This concatenates (joins) strings!
console.log(a + b);  // "53" (not 8!)

// CORRECT - Parse to numbers first
let numA = parseInt(a);
let numB = parseInt(b);
console.log(numA + numB);  // 8

// Or make sure they're numbers from the start
let c = 5;
let d = 3;
console.log(c + d);  // 8`,
				"python": `a = "5"
b = "3"

# WRONG - This concatenates (joins) strings!
print(a + b)  # "53" (not 8!)

# CORRECT - Parse to numbers first
num_a = int(a)
num_b = int(b)
print(num_a + num_b)  # 8

# Or make sure they're numbers from the start
c = 5
d = 3
print(c + d)  # 8`,
			}),
			textBlock(`**How to avoid it:** Be careful with user input - it's usually a string. Use Parse/int()/parseInt() to convert.`),
			calloutBlock("tip", "When something isn't working, use Console.WriteLine/console.log/print to print out your variables at different points. This helps you see exactly what's happening!"),
			textBlock(`## Debugging Strategy: The Console is Your Friend

When your code isn't working:

1. **Read the error message carefully** - It usually tells you exactly what's wrong
2. **Add print statements** - Print your variables to see their values
3. **Check one thing at a time** - Don't change 5 things at once
4. **Take a break** - Sometimes walking away helps you see the problem
5. **Explain it out loud** - Seriously! Explain your code to a rubber duck (or a friend). Hearing it out loud helps you spot errors.

### Example: Debugging with print statements`),
			multiCodeBlock(map[string]string{
				"csharp": `int CalculateTotal(string price, int quantity) {
    Console.WriteLine("Price: " + price);
    Console.WriteLine("Quantity: " + quantity);
    
    int priceNum = int.Parse(price);
    int total = priceNum * quantity;
    Console.WriteLine("Total: " + total);
    
    return total;
}

CalculateTotal("10", 5);`,
				"typescript": `function calculateTotal(price: string, quantity: number): number {
    console.log("Price:", price);
    console.log("Quantity:", quantity);
    
    let priceNum = parseInt(price);
    let total = priceNum * quantity;
    console.log("Total:", total);
    
    return total;
}

calculateTotal("10", 5);`,
				"python": `def calculate_total(price, quantity):
    print("Price:", price)
    print("Quantity:", quantity)
    
    price_num = int(price)
    total = price_num * quantity
    print("Total:", total)
    
    return total

calculate_total("10", 5)`,
			}),
			exerciseBlock(
				"Find and fix the bug in this code:\n\n"+
				"C#:\n"+
				"int SumArray(int[] numbers) {\n"+
				"    int sum = 0;\n"+
				"    for (int i = 0; i <= numbers.Length; i++) {\n"+
				"        sum += numbers[i];\n"+
				"    }\n"+
				"    return sum;\n"+
				"}\n\n"+
				"TypeScript:\n"+
				"function sumArray(numbers: number[]): number {\n"+
				"    let sum = 0;\n"+
				"    for (let i = 0; i <= numbers.length; i++) {\n"+
				"        sum += numbers[i];\n"+
				"    }\n"+
				"    return sum;\n"+
				"}\n\n"+
				"Python:\n"+
				"def sum_array(numbers):\n"+
				"    total = 0\n"+
				"    for i in range(len(numbers) + 1):\n"+
				"        total += numbers[i]\n"+
				"    return total",
				"The bug is in the loop condition: <= should be < (or range should not add 1).\n\n"+
				"When i equals the array length, you go past the end.\n\n"+
				"Fixed versions:\n"+
				"C#: for (int i = 0; i < numbers.Length; i++)\n"+
				"TypeScript: for (let i = 0; i < numbers.length; i++)\n"+
				"Python: for i in range(len(numbers)):",
				[]string{"Check the loop condition carefully", "What happens when i equals the array length?", "Remember: last valid index is Length - 1"},
			),
		},
	}
}