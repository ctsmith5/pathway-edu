package seed

import "github.com/pathway/backend/models"

func moduleGit2() models.Module {
	return models.Module{
		ID:    "git-2",
		Title: "Git Basics: Init, Add, Commit",
		Content: []models.ContentBlock{
			textBlock(`## Setting Up Your First Repository

A Git repository (or "repo") is a folder that Git is tracking. Let's create one!

### Initializing a Repository

Navigate to your project folder and run:`),
			codeBlock("bash", `cd my-project
git init`),
			textBlock("This creates a hidden `.git` folder that stores all version control information."),
			calloutBlock("tip", "You only need to run git init once per project, at the beginning."),
			textBlock(`## The Three States of Git

Files in Git exist in three states:

1. **Modified**: You've changed the file but haven't staged it
2. **Staged**: You've marked the file to go into your next commit
3. **Committed**: The file is safely stored in your local database`),
			textBlock(`## Staging Changes

Before committing, you need to "stage" your changes. This lets you choose exactly what to include.`),
			codeBlock("bash", `# Stage a specific file
git add filename.txt

# Stage all changes
git add .

# Stage all changes in a directory
git add src/`),
			textBlock(`## Committing Changes

A commit is a snapshot of your staged changes. Always include a meaningful message!`),
			codeBlock("bash", `git commit -m "Add user authentication feature"`),
			calloutBlock("warning", "Write clear, descriptive commit messages. 'Fixed stuff' is not helpful. 'Fix login button not responding on mobile' is much better!"),
			textBlock("## Checking Status\n\nUse `git status` frequently to see what's going on:"),
			codeBlock("bash", `git status`),
			exerciseBlock(
				"Initialize a new repository, create a file called 'hello.txt', stage it, and commit it with the message 'Initial commit'",
				`git init
echo "Hello World" > hello.txt
git add hello.txt
git commit -m "Initial commit"`,
				[]string{"Use git init to start", "Create the file, then add it", "Don't forget the -m flag for your message"},
			),
			textBlock(`## Hands-On: Create Your Hello World Project

Now it's time to create your first real Git project! Follow along step by step.

### Step 1: Create Your Project Folder

Open your terminal and create a new folder for your project:`),
			codeBlock("bash", `# Create a new folder
mkdir my-first-repo

# Navigate into the folder
cd my-first-repo`),
			textBlock(`### Step 2: Initialize Git

Turn this folder into a Git repository:`),
			codeBlock("bash", `git init`),
			textBlock(`You should see: "Initialized empty Git repository in .../my-first-repo/.git/"

### Step 3: Create Your HelloWorld.txt File

Create a text file with your name and a greeting. You can use any text editor, or use the command line:`),
			codeBlock("bash", `# On Windows (PowerShell)
echo "Your Name" > HelloWorld.txt
echo "Hello World!" >> HelloWorld.txt

# On Mac/Linux
echo "Your Name" > HelloWorld.txt
echo "Hello World!" >> HelloWorld.txt`),
			calloutBlock("tip", "Replace 'Your Name' with your actual name! This file will be part of your portfolio."),
			textBlock(`Your HelloWorld.txt should look something like this:`),
			codeBlock("text", `Colin
Hello World!`),
			textBlock(`### Step 4: Check the Status

See what Git knows about your new file:`),
			codeBlock("bash", `git status`),
			textBlock(`You'll see HelloWorld.txt listed as an "untracked file" - Git sees it but isn't tracking changes yet.

### Step 5: Stage Your File`),
			codeBlock("bash", `git add HelloWorld.txt`),
			textBlock(`### Step 6: Commit Your File

Create your first commit with a meaningful message:`),
			codeBlock("bash", `git commit -m "Add HelloWorld.txt with my name and greeting"`),
			calloutBlock("info", "Congratulations! You've just created your first Git commit. This snapshot is now saved in your repository's history forever."),
			textBlock(`### Step 7: View Your Commit History`),
			codeBlock("bash", `git log`),
			textBlock(`You'll see your commit with the message, author (you!), and a unique commit hash.

**Keep this repository open** - we'll push it to GitHub in the next module!`),
			exerciseBlock(
				"You've created HelloWorld.txt and committed it. What command shows you the history of commits in your repository?",
				"git log - This shows all commits in reverse chronological order, including commit hash, author, date, and message.",
				[]string{"Think about what command shows history", "It's a short three-letter command"},
			),
		},
	}
}
