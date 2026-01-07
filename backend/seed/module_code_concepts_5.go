package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts5() models.Module {
	return models.Module{
		ID:    "code-5",
		Title: "Programming Languages 101 (C#, JavaScript, TypeScript, Go, Python)",
		Content: []models.ContentBlock{
			textBlock(`## Programming Languages 101

When people say "learn to code", what they really mean is:

- Learn to give instructions to a computer
- Learn to organize those instructions so other people can understand them
- Learn to debug when the computer does something unexpected

A programming language is just a set of rules for writing those instructions.

You will see many languages in the real world. The good news: the big ideas transfer. If you learn the fundamentals once, switching languages becomes much easier.

## Where does code run?

- Front-end (in the browser): JavaScript and TypeScript are the main languages here.
- Back-end (on a server): C#, Go, Python, and also JavaScript/TypeScript (Node.js) are common.

We teach multiple languages and concepts because real apps are built from multiple pieces that work together.

## A quick tour (plain English)

### JavaScript
- The language of the web browser
- Great for building interactive websites
- Very flexible, which is powerful but can cause confusing bugs for beginners

### TypeScript
- JavaScript with types (think: labels for data)
- Helps catch mistakes earlier and makes code easier to read on teams

### C#
- Popular for back-end services, business apps, and games (Unity)
- Strong tooling and a clear object-oriented style

### Go
- Designed for fast, reliable servers and services
- Simple syntax, great performance, strong support for concurrency (doing many tasks at once)

### Python
- Very beginner-friendly syntax
- Great for automation, scripting, and data work
- Also used for back-end web services

## Variables: mutable vs immutable (can it change?)

A variable is a name that points to a value.

- Mutable means you can change it after you create it.
- Immutable means you cannot change it (you must create a new value instead).

Here is what basic variable declarations look like in different languages.

### TypeScript (mutable with let, immutable with const)`),
			codeBlock("typescript", `let score = 10;        // mutable (can change)
score = 11;

const name = "Ava";    // immutable binding (cannot be reassigned)
// name = "Sam";       // would be an error`),
			textBlock(`### C# (mutable variables, immutable constants)`),
			codeBlock("csharp", `int score = 10;         // mutable
score = 11;

const string name = "Ava"; // constant (cannot be reassigned)
// name = "Sam";           // would be an error`),
			textBlock(`### Go (var or :=, and const)`),
			codeBlock("go", `score := 10     // mutable
score = 11

const name = "Ava" // constant`),
			textBlock(`### Python (assignment binds a name; mutability depends on the type)`),
			codeBlock("python", `score = 10
score = 11  # reassignment is allowed

name = "Ava"
# name = "Sam"  # reassignment is allowed too (strings are immutable values)`),
			calloutBlock("info", "In some languages, the variable itself is mutable/immutable. In others (like Python), the important idea is whether the underlying value (like a list) is mutable."),
			textBlock(`## Value types vs reference types (copying behavior)

This is one of the most important beginner concepts.

When you assign one variable to another, are you copying:
- the value itself (value type), or
- a reference to the same object (reference type)?

### C# example

In C#, many basic types like int are value types. Classes are reference types.`),
			codeBlock("csharp", `// Value type copy
int a = 5;
int b = a;   // b gets its own copy
b = 10;
// a is still 5

// Reference type copy (class)
class Person { public string Name { get; set; } }

var p1 = new Person { Name = "Ava" };
var p2 = p1;        // p2 points at the same object
p2.Name = "Sam";
// p1.Name is now "Sam"`),
			textBlock(`### TypeScript / JavaScript example

Primitives (number, string, boolean) behave like value copies. Objects/arrays behave like references.`),
			codeBlock("typescript", `// Value-like (primitive)
let a = 5;
let b = a;
b = 10;       // a is still 5

// Reference-like (object)
const p1 = { name: "Ava" };
const p2 = p1;
p2.name = "Sam"; // p1.name is now "Sam"`),
			textBlock(`### Go example

Go tends to copy values by default (including structs). Pointers give you reference-like behavior.`),
			codeBlock("go", `type Person struct { Name string }

// Struct assignment copies the value
p1 := Person{Name: "Ava"}
p2 := p1
p2.Name = "Sam"
// p1.Name is still "Ava"

// Pointer gives reference-like behavior
p3 := &p1
p3.Name = "Lee"
// p1.Name is now "Lee"`),
			textBlock(`### Python example

Python variables are names that refer to objects. Some objects are immutable (like ints), some are mutable (like lists).`),
			codeBlock("python", `# Immutable example
a = 5
b = a
b = 10
# a is still 5

# Mutable example (list)
xs = [1, 2]
ys = xs
ys.append(3)
# xs is now [1, 2, 3]`),
			textBlock(`## What is a class?

A class is a blueprint for creating objects.

It usually contains:
- data (fields/properties)
- behavior (methods/functions)

Here are tiny examples across languages.

### C# class (short example)`),
			codeBlock("csharp", `class Person {
  public string Name { get; }

  public Person(string name) {
    Name = name;
  }

  public void SayHi() {
    Console.WriteLine("Hi, I'm " + Name);
  }
}`),
			textBlock(`### TypeScript class (short example)`),
			codeBlock("typescript", `class Person {
  name: string;

  constructor(name: string) {
    this.name = name;
  }

  sayHi(): void {
    console.log("Hi, I'm " + this.name);
  }
}`),
			textBlock(`### Python class (short example)`),
			codeBlock("python", `class Person:
    def __init__(self, name):
        self.name = name

    def say_hi(self):
        print("Hi, I'm " + self.name)`),
			textBlock(`### Go equivalent (struct + method)`),
			codeBlock("go", `type Person struct { Name string }

func (p Person) SayHi() {
  fmt.Println("Hi, I'm " + p.Name)
}`),
			textBlock(`## Control flow (preview): if/else

Control flow is how programs make decisions.

Here is a very basic C# example:`),
			codeBlock("csharp", `int age = 17;

if (age >= 18) {
  Console.WriteLine("Adult");
} else {
  Console.WriteLine("Minor");
}`),
			textBlock(`The same idea exists in every language:`),
			codeBlock("text", `TypeScript:
if (age >= 18) { ... } else { ... }

Python:
if age >= 18:
    ...
else:
    ...

Go:
if age >= 18 {
    ...
} else {
    ...
}`),
			calloutBlock("tip", "In the next module (Control Flow), we go deeper into if/else, switch, loops, and common beginner mistakes."),
			textBlock(`## Functions and parameters (inputs to a function)

A function is a named set of steps you can reuse.

- A parameter is the name in the function definition (the input placeholder).
- An argument is the real value you pass when you call the function.

### C# example (function with a parameter)`),
			codeBlock("csharp", `static int Double(int n) {
  // n is a parameter
  return n * 2;
}

int result = Double(21); // 21 is the argument
// result is 42`),
			textBlock(`### TypeScript example (same idea)`),
			codeBlock("typescript", `function double(n: number): number {
  return n * 2;
}

const result = double(21);`),
			exerciseBlock(
				"1) In the C# if/else example above, what does it print when age is 17? 2) Write a function called Triple that takes a number and returns that number times 3.",
				"1) It prints Minor. 2) One solution: static int Triple(int n) { return n * 3; } and then call Triple(10) to get 30.",
				[]string{"If age is less than 18, the else branch runs", "A parameter is the name inside the parentheses; the argument is the value you pass in"},
			),
			exerciseBlock(
				"Write down one reason a team might choose TypeScript over JavaScript for a large project.",
				"TypeScript helps catch errors before runtime and makes code easier to understand by making data types explicit. This reduces bugs and improves teamwork in large codebases.",
				[]string{"Think about teamwork and catching mistakes early", "Types are like labels that explain what a value is"},
			),
		},
	}
}
