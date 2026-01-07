package seed

import "github.com/pathway/backend/models"

func moduleGit4() models.Module {
	return models.Module{
		ID:    "git-4",
		Title: "Branching and Merging",
		Content: []models.ContentBlock{
			textBlock(`## What are Branches?

Branches let you diverge from the main line of development to work on features independently. Think of them as parallel universes for your code.

### Why Use Branches?

- Develop features without affecting the main code
- Experiment safely
- Work on multiple things simultaneously
- Collaborate without conflicts`),
			textBlock(`## Creating and Switching Branches`),
			codeBlock("bash", `# Create a new branch
git branch feature-login

# Switch to the branch
git checkout feature-login

# Or do both in one command
git checkout -b feature-login`),
			calloutBlock("tip", "In Git 2.23+, you can use 'git switch' instead of 'git checkout' for switching branches. It's clearer and safer!"),
			textBlock(`## Viewing Branches`),
			codeBlock("bash", `# List all local branches
git branch

# List all branches including remote
git branch -a

# See which branch you're on
git status`),
			textBlock(`## Merging Branches

When your feature is complete, merge it back into the main branch:`),
			codeBlock("bash", `# Switch to main branch
git checkout main

# Merge the feature branch
git merge feature-login`),
			textBlock(`### Types of Merges

1. **Fast-forward merge**: If main hasn't changed, Git just moves the pointer forward
2. **Three-way merge**: If both branches have new commits, Git creates a merge commit`),
			exerciseBlock(
				"Create a branch called 'add-readme', switch to it, and then switch back to main",
				`git checkout -b add-readme
# or: git branch add-readme && git checkout add-readme
git checkout main`,
				[]string{"Use -b flag to create and switch in one command", "Use checkout or switch to change branches"},
			),
			textBlock(`## Hands-On: Create Your Development Branch

Now let's practice branching with your Hello World project! You'll create a development branch and add more about yourself.

### Step 1: Make Sure You're on Main

First, ensure you're on the main branch and have the latest code:`),
			codeBlock("bash", `# Navigate to your project (if not already there)
cd my-first-repo

# Check which branch you're on
git branch

# Make sure you're on main
git checkout main`),
			textBlock(`### Step 2: Create and Switch to Development Branch`),
			codeBlock("bash", `# Create and switch to a new branch called 'development'
git checkout -b development`),
			textBlock(`You should see: "Switched to a new branch 'development'"

### Step 3: Update Your HelloWorld.txt

Open HelloWorld.txt in a text editor and add a short bio about yourself. Your file should now look something like this:`),
			codeBlock("text", `Colin
Hello World!

About Me:
- I'm learning Git and version control
- My favorite programming language is: [your answer]
- One thing I want to build: [your answer]
- Fun fact about me: [your answer]`),
			calloutBlock("tip", "Make this personal! Add your actual interests and goals. This is practice for real development where you'll document your work."),
			textBlock(`### Step 4: Stage and Commit Your Changes`),
			codeBlock("bash", `# Check what changed
git status

# Stage your changes
git add HelloWorld.txt

# Commit with a descriptive message
git commit -m "Add personal bio to HelloWorld.txt"`),
			textBlock(`### Step 5: Push Your Development Branch to GitHub`),
			codeBlock("bash", `# Push the development branch (first time)
git push -u origin development`),
			textBlock(`### Step 6: Verify on GitHub

1. Go to your repository on GitHub
2. Click the branch dropdown (it says "main")
3. You should see "development" as an option
4. Switch to it and verify your bio is there!`),
			calloutBlock("info", "You now have two branches: 'main' with your original Hello World, and 'development' with your added bio. This is exactly how professional developers work!"),
			exerciseBlock(
				"You created a development branch and pushed it to GitHub. What command creates a new branch AND switches to it in one step?",
				"git checkout -b development - The -b flag tells git to create the branch and then switch to it immediately.",
				[]string{"Think about the checkout command", "There's a flag that combines create and switch"},
			),
		},
	}
}
