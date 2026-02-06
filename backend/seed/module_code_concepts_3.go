package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts3() models.Module {
	return models.Module{
		ID:    "code-3",
		Title: "Repeating Actions",
		Content: []models.ContentBlock{
			textBlock("## Computers Love Repetition\n\nOne of the biggest advantages of computers is that they never get bored doing the same thing over and over. This is where loops come in!\n\n### Why Use Loops?\n\nImagine you need to print \"Hello!\" 100 times. You could write:\n" +
				"```javascript\n" +
				"console.log(\"Hello!\");\n" +
				"console.log(\"Hello!\");\n" +
				"console.log(\"Hello!\");\n" +
				"// ... 97 more times ...\n" +
				"```\n\n" +
				"Or you could write a loop that does it for you in 3 lines!"),
			textBlock("## For Loops: Counting Repetition\n\nA for loop is perfect when you know exactly how many times you want to repeat something.\n\n### The Anatomy of a For Loop\n\n" +
				"```javascript\n" +
				"for (let i = 0; i < 5; i++) {\n" +
				"    console.log(\"Iteration: \" + i);\n" +
				"}\n" +
				"```"),
			imageBlock("https://storage.googleapis.com/ludicrousapps-c1ea7.firebasestorage.app/PathwayEdu/1f442d36-7bb3-43fc-b578-3e3d225890fa.png", "Diagram showing the anatomy of a JavaScript for loop with labeled parts for initialization, condition, and increment", "The three parts of a for loop: initialization (let i = 0), condition (i < 5), and increment (i++)"),
			textBlock("Let's break this down:\n" +
				"1. **let i = 0** - Start with i equal to 0 (initialization)\n" +
				"2. **i < 5** - Keep going while i is less than 5 (condition)\n" +
				"3. **i++** - Add 1 to i after each loop (increment)\n" +
				"4. **{ ... }** - The code that runs each time\n\n" +
				"**Output:**\n" +
				"```\n" +
				"Iteration: 0\n" +
				"Iteration: 1\n" +
				"Iteration: 2\n" +
				"Iteration: 3\n" +
				"Iteration: 4\n" +
				"```"),
			calloutBlock("info", "The variable 'i' is just a convention. It stands for 'index' or 'iterator'. You can use any name you want, like 'count' or 'step'."),
			textBlock("## Looping Through Arrays\n\nLoops are super useful for working with arrays (lists of items).\n\n" +
				"```javascript\n" +
				"let fruits = [\"apple\", \"banana\", \"cherry\", \"date\"];\n\n" +
				"// Print each fruit\n" +
				"for (let i = 0; i < fruits.length; i++) {\n" +
				"    console.log(fruits[i]);\n" +
				"}\n\n" +
				"// Output:\n" +
				"// apple\n" +
				"// banana\n" +
				"// cherry\n" +
				"// date\n" +
				"```\n\n" +
				"**What's happening:**\n" +
				"- `fruits.length` gives us 4 (there are 4 items)\n" +
				"- `fruits[i]` gets the item at position i\n" +
				"- We go from i=0 to i=3 (all 4 items)"),
			textBlock("## While Loops: Conditional Repetition\n\nA while loop keeps going as long as a condition is true. Use this when you DON'T know how many times you need to repeat.\n\n### Real-World Analogy\n" +
				"\"While the light is red, wait.\" You don't know how many seconds - you just keep checking.\n\n" +
				"```javascript\n" +
				"let count = 0;\n\n" +
				"while (count < 5) {\n" +
				"    console.log(\"Count is: \" + count);\n" +
				"    count++;\n" +
				"}\n" +
				"```\n\n" +
				"**Important:** You MUST change something inside the loop, or it will run forever! In the example above, we increment count each time."),
			calloutBlock("warning", "An infinite loop happens when the condition never becomes false. This will crash your program! Always make sure the loop will eventually end."),
			textBlock("## Practical Example: Finding Something\n\nLet's say we have a list of students and we want to find if \"Alice\" is in it:\n\n" +
				"```javascript\n" +
				"let students = [\"Bob\", \"Alice\", \"Charlie\", \"Diana\"];\n" +
				"let found = false;\n\n" +
				"for (let i = 0; i < students.length; i++) {\n" +
				"    if (students[i] === \"Alice\") {\n" +
				"        found = true;\n" +
				"        console.log(\"Found Alice at position \" + i);\n" +
				"        break;  // Exit the loop early!\n" +
				"    }\n" +
				"}\n\n" +
				"if (!found) {\n" +
				"    console.log(\"Alice is not in the list.\");\n" +
				"}\n" +
				"```\n\n" +
				"**The 'break' statement** immediately exits the loop. We use it here because once we found Alice, we don't need to keep looking!"),
			exerciseBlock(
				"Write a loop that calculates the sum of all numbers from 1 to 100. Hint: Create a variable to keep track of the total, then add each number to it in the loop.",
				"let total = 0;\n\nfor (let i = 1; i <= 100; i++) {\n    total = total + i;\n}\n\nconsole.log(\"The sum is: \" + total);  // Should be 5050",
				[]string{"Start total at 0", "Loop from 1 to 100", "Add i to total each time"},
			),
		},
	}
}