package seed

import "github.com/pathway/backend/models"

func moduleTypeScriptStarter() models.Module {
	return models.Module{
		ID:    "typescript-starter",
		Title: "Pull the Starter Repo",
		Content: []models.ContentBlock{
			textBlock(`## Clone Your First TypeScript Repository

Now that you have Git and VS Code set up, let's pull a real code repository and explore some fundamental programming concepts!

### What We'll Do

1. Clone a TypeScript starter repository
2. Open it in VS Code
3. Explore variables, types, and control flow
4. Run the code and see it in action`),
			calloutBlock("info", "TypeScript is JavaScript with types. It helps catch errors before your code runs and makes your code more readable."),
			textBlock(`## Step 1: Clone the Repository

Open your terminal (or VS Code's integrated terminal) and run:`),
			codeBlock("bash", `# Navigate to where you want to store projects
cd ~/projects

# Clone the starter repository
git clone https://github.com/YOUR-INSTRUCTOR/typescript-starter.git

# Navigate into the project
cd typescript-starter

# Open in VS Code
code .`),
			calloutBlock("warning", "Your instructor will provide the actual repository URL. Replace the URL above with the one they give you!"),
			textBlock(`## Step 2: Install Dependencies

Most code projects have dependencies - libraries of code others have written. Install them with:`),
			codeBlock("bash", `npm install`),
			textBlock(`This reads the package.json file and downloads everything the project needs.

## Step 3: Explore the Code

Open the src folder in VS Code. You'll see TypeScript files (.ts) containing code examples.`),
			textBlock(`### Variables and Types

In TypeScript, variables have types that tell us what kind of data they hold:`),
			codeBlock("typescript", `// String - text data
let name: string = "Alice";

// Number - numeric data
let age: number = 25;
let price: number = 19.99;

// Boolean - true or false
let isStudent: boolean = true;

// Array - list of items
let colors: string[] = ["red", "green", "blue"];

// Object - structured data
let person: { name: string; age: number } = {
  name: "Bob",
  age: 30
};`),
			calloutBlock("tip", "TypeScript can often figure out the type automatically. 'let name = \"Alice\"' works too - TypeScript knows it's a string!"),
			textBlock(`### Control Flow

Control flow determines which code runs based on conditions:`),
			codeBlock("typescript", `// If statement - make decisions
let temperature = 75;

if (temperature > 80) {
  console.log("It's hot!");
} else if (temperature > 60) {
  console.log("It's nice outside!");
} else {
  console.log("It's cold!");
}

// Loops - repeat code
for (let i = 0; i < 5; i++) {
  console.log("Count:", i);
}

// Loop through array
let fruits = ["apple", "banana", "orange"];
for (let fruit of fruits) {
  console.log("I like", fruit);
}`),
			textBlock(`## Step 4: Run the Code

Run the TypeScript code to see it in action:`),
			codeBlock("bash", `# Compile and run
npm start

# Or run in development mode (auto-reloads)
npm run dev`),
			textBlock(`Watch the terminal output. You'll see the results of the code examples!`),
			exerciseBlock(
				"Clone the starter repository, open it in VS Code, and run 'npm install' followed by 'npm start'. What output do you see in the terminal?",
				"You should see output from the TypeScript examples, including console.log statements showing variables, conditional results, and loop iterations. The exact output depends on the repository content your instructor provides.",
				[]string{"Make sure you run npm install first", "The output appears in your terminal", "Look for console.log output"},
			),
			textBlock(`## What You've Learned

- How to clone a Git repository
- How to install project dependencies with npm
- Basic TypeScript variables and types
- Control flow with if statements and loops

These concepts are the building blocks of all programming. You'll use them in every project you build!`),
		},
	}
}
