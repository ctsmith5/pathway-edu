package seed

import (
	"log"

	"github.com/pathway/backend/models"
	"github.com/pathway/backend/repository"
)

// Helper function to create a text block
func textBlock(markdown string) models.ContentBlock {
	return models.ContentBlock{
		Type: "text",
		Data: map[string]interface{}{
			"markdown": markdown,
		},
	}
}

// Helper function to create a code block
func codeBlock(language, code string) models.ContentBlock {
	return models.ContentBlock{
		Type: "code",
		Data: map[string]interface{}{
			"language": language,
			"code":     code,
		},
	}
}

// Helper function to create an image block
func imageBlock(url, alt, caption string) models.ContentBlock {
	return models.ContentBlock{
		Type: "image",
		Data: map[string]interface{}{
			"url":     url,
			"alt":     alt,
			"caption": caption,
		},
	}
}

// Helper function to create a callout block
func calloutBlock(variant, text string) models.ContentBlock {
	return models.ContentBlock{
		Type: "callout",
		Data: map[string]interface{}{
			"variant": variant, // "info", "tip", "warning", "danger"
			"text":    text,
		},
	}
}

// Helper function to create an exercise block
func exerciseBlock(prompt, solution string, hints []string) models.ContentBlock {
	return models.ContentBlock{
		Type: "exercise",
		Data: map[string]interface{}{
			"prompt":   prompt,
			"solution": solution,
			"hints":    hints,
		},
	}
}

// Helper function to create a video block
func videoBlock(url, title string) models.ContentBlock {
	return models.ContentBlock{
		Type: "video",
		Data: map[string]interface{}{
			"url":   url,
			"title": title,
		},
	}
}

func SeedCourses(repo repository.Repository) error {
	// Clear existing courses
	if err := repo.DeleteAllCourses(); err != nil {
		return err
	}

	courses := []models.Course{
		{
			Title:       "Git",
			Description: "Learn version control with Git - the essential tool for modern software development. Master branching, merging, and collaboration workflows.",
			Modules: []models.Module{
				{
					ID:    "git-1",
					Title: "Introduction to Version Control",
					Content: []models.ContentBlock{
						textBlock(`## What is Version Control?

Version control is a system that records changes to files over time so you can recall specific versions later. Think of it as a time machine for your code.

### Why Use Version Control?

- **Track Changes**: See exactly what changed, when, and by whom
- **Collaboration**: Work with others without overwriting each other's work
- **Backup**: Never lose your work - every version is saved
- **Experimentation**: Try new ideas without fear of breaking things`),
						calloutBlock("info", "Git was created by Linus Torvalds in 2005 for Linux kernel development. Today, it's the most widely used version control system in the world."),
						textBlock(`### Types of Version Control Systems

1. **Local Version Control**: Simple database on your computer
2. **Centralized Version Control**: Single server holds all versions (e.g., SVN)
3. **Distributed Version Control**: Every developer has a full copy (e.g., Git)

Git is a *distributed* version control system, which means you have the complete history on your local machine.`),
						exerciseBlock(
							"Why might a distributed version control system be better for open source projects?",
							"Distributed systems allow developers to work offline, have faster operations since most are local, provide better redundancy (every clone is a backup), and enable more flexible workflows like forking.",
							[]string{"Think about what happens if the central server goes down", "Consider developers in different time zones"},
						),
					},
				},
				{
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
					},
				},
				{
					ID:    "git-3",
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
					},
				},
				{
					ID:    "git-4",
					Title: "Remote Repositories",
					Content: []models.ContentBlock{
						textBlock(`## Working with Remote Repositories

Remote repositories are versions of your project hosted on the internet or network. They enable collaboration with others.

### Popular Git Hosting Services

- **GitHub**: Most popular, great for open source
- **GitLab**: Strong CI/CD features, can self-host
- **Bitbucket**: Integrates well with Atlassian tools`),
						textBlock(`## Connecting to a Remote`),
						codeBlock("bash", `# Add a remote repository
git remote add origin https://github.com/username/repo.git

# View your remotes
git remote -v`),
						calloutBlock("info", "'origin' is the conventional name for your primary remote repository. You can have multiple remotes with different names."),
						textBlock(`## Pushing Changes

Send your local commits to the remote repository:`),
						codeBlock("bash", `# Push to remote (first time, sets upstream)
git push -u origin main

# Subsequent pushes
git push`),
						textBlock(`## Pulling Changes

Get changes from the remote repository:`),
						codeBlock("bash", `# Fetch and merge changes
git pull

# Just fetch (doesn't merge)
git fetch`),
						textBlock(`## Cloning a Repository

To get a copy of an existing remote repository:`),
						codeBlock("bash", `git clone https://github.com/username/repo.git`),
						exerciseBlock(
							"What's the difference between git fetch and git pull?",
							"git fetch downloads changes from the remote but doesn't integrate them into your working branch. git pull does both - it fetches AND merges the changes. Using fetch is safer when you want to review changes before merging.",
							[]string{"Think about what happens to your local files", "Consider when you might want to see changes before applying them"},
						),
					},
				},
				{
					ID:    "git-5",
					Title: "Collaboration Workflows",
					Content: []models.ContentBlock{
						textBlock(`## Git Collaboration Workflows

When working with a team, you need a structured approach to collaboration.

### Feature Branch Workflow

The most common workflow:

1. Create a branch for each feature
2. Make commits on that branch
3. Push the branch to remote
4. Open a Pull Request
5. Review, discuss, and merge`),
						codeBlock("bash", `# Start a new feature
git checkout -b feature/user-profile

# Work and commit
git add .
git commit -m "Add user profile page"

# Push to remote
git push -u origin feature/user-profile`),
						textBlock(`## Pull Requests (PRs)

Pull Requests are a way to propose changes and request code review. They:

- Show the diff between branches
- Enable code review and discussion
- Can run automated tests
- Create a record of why changes were made`),
						calloutBlock("tip", "Write clear PR descriptions explaining WHAT you changed and WHY. Link to related issues if applicable."),
						textBlock(`## Code Review Best Practices

As a **reviewer**:
- Be constructive and kind
- Ask questions rather than make demands
- Focus on the code, not the person
- Approve when it's good enough, not perfect

As an **author**:
- Keep PRs small and focused
- Respond to feedback graciously
- Explain your reasoning when needed`),
						textBlock(`## Forking Workflow

Used in open source projects:

1. Fork the repository (create your own copy)
2. Clone your fork
3. Make changes
4. Push to your fork
5. Open a PR to the original repo`),
						exerciseBlock(
							"You want to contribute to an open source project. What are the steps?",
							"1. Fork the repository on GitHub\n2. Clone your fork: git clone https://github.com/YOUR-USERNAME/repo.git\n3. Create a feature branch: git checkout -b my-feature\n4. Make changes and commit them\n5. Push to your fork: git push origin my-feature\n6. Open a Pull Request from your fork to the original repository",
							[]string{"Start by getting your own copy of the repo", "You can't push directly to someone else's repo"},
						),
					},
				},
				{
					ID:    "git-6",
					Title: "Resolving Conflicts",
					Content: []models.ContentBlock{
						textBlock(`## What are Merge Conflicts?

Merge conflicts occur when Git can't automatically merge changes because the same lines were modified differently in both branches.

### When Do Conflicts Happen?

- Two people edit the same line of a file
- One person deletes a file another person modified
- Complex structural changes that Git can't reconcile`),
						textBlock(`## Identifying Conflicts

When a conflict occurs, Git marks the file:`),
						codeBlock("text", `<<<<<<< HEAD
your changes here
=======
their changes here
>>>>>>> feature-branch`),
						calloutBlock("warning", "Never commit files with conflict markers still in them! Always resolve conflicts completely."),
						textBlock(`## Resolving Conflicts

1. Open the conflicted file
2. Find the conflict markers
3. Decide which changes to keep (or combine them)
4. Remove the conflict markers
5. Stage and commit the resolved file`),
						codeBlock("bash", `# After manually fixing the file
git add resolved-file.txt
git commit -m "Resolve merge conflict in resolved-file.txt"`),
						textBlock(`## Prevention is Better Than Cure

To minimize conflicts:

- Pull frequently to stay up to date
- Keep branches short-lived
- Communicate with your team about who's working on what
- Break large changes into smaller PRs`),
						calloutBlock("tip", "Many editors (VS Code, IntelliJ) have built-in merge conflict resolution tools that make this process much easier!"),
						exerciseBlock(
							"You see conflict markers in your file. What are the three sections and what do they represent?",
							"The three sections are:\n1. <<<<<<< HEAD - The start of your current branch's changes\n2. ======= - The separator between the two versions\n3. >>>>>>> branch-name - The end marker showing the incoming branch's changes\n\nEverything between HEAD and ======= is your version. Everything between ======= and the branch name is the incoming version.",
							[]string{"Look at the markers: HEAD, =======, and the branch name", "HEAD refers to your current position"},
						),
					},
				},
				{
					ID:    "git-7",
					Title: "Advanced Git: Rebase, Cherry-pick, Stash",
					Content: []models.ContentBlock{
						textBlock(`## Git Rebase

Rebase is an alternative to merging that creates a linear history by moving your branch's commits on top of another branch.`),
						codeBlock("bash", `# Rebase your feature branch onto main
git checkout feature-branch
git rebase main`),
						textBlock(`### Merge vs Rebase

| Merge | Rebase |
|-------|--------|
| Preserves complete history | Creates linear history |
| Creates merge commits | No extra commits |
| Safer for shared branches | Best for local branches |
| Shows when branches converged | Looks like work happened sequentially |`),
						calloutBlock("danger", "Never rebase commits that have been pushed and shared with others! This rewrites history and will cause problems for your teammates."),
						textBlock(`## Interactive Rebase

Clean up your commit history before sharing:`),
						codeBlock("bash", `# Rebase last 3 commits interactively
git rebase -i HEAD~3`),
						textBlock(`This lets you:
- **pick**: Keep the commit
- **reword**: Change the commit message
- **squash**: Combine with previous commit
- **drop**: Remove the commit`),
						textBlock(`## Cherry-pick

Apply a specific commit from one branch to another:`),
						codeBlock("bash", `# Apply a specific commit to current branch
git cherry-pick abc1234`),
						calloutBlock("tip", "Cherry-pick is great for applying hotfixes from main to a release branch, or grabbing a specific feature commit."),
						textBlock(`## Git Stash

Temporarily save changes without committing:`),
						codeBlock("bash", `# Stash your changes
git stash

# List stashes
git stash list

# Apply most recent stash
git stash pop

# Apply a specific stash
git stash apply stash@{2}`),
						textBlock(`### When to Use Stash

- Need to switch branches but aren't ready to commit
- Want to pull changes but have local modifications
- Need to quickly context-switch to fix a bug`),
						exerciseBlock(
							"You're working on a feature but need to urgently fix a bug on main. You're not ready to commit your feature work. What do you do?",
							`git stash                    # Save your work-in-progress
git checkout main           # Switch to main
git checkout -b hotfix      # Create a hotfix branch
# ... fix the bug and commit ...
git checkout main
git merge hotfix
git checkout feature-branch # Go back to your feature
git stash pop               # Restore your work-in-progress`,
							[]string{"You need to save your current work temporarily", "Use stash to store uncommitted changes"},
						),
					},
				},
			},
		},
		{
			Title:       "Architecture - SOLID",
			Description: "Master the SOLID principles of object-oriented design to write maintainable, scalable, and robust software architectures.",
			Modules:     []models.Module{},
		},
		{
			Title:       "SCRUM",
			Description: "Learn the SCRUM framework for agile project management. Understand sprints, standups, and how to deliver value iteratively.",
			Modules:     []models.Module{},
		},
		{
			Title:       "HTTP Networking",
			Description: "Deep dive into HTTP protocols, REST APIs, request/response cycles, and modern web communication patterns.",
			Modules:     []models.Module{},
		},
		{
			Title:       "Design Patterns",
			Description: "Learn proven software design patterns to solve common programming challenges elegantly and efficiently.",
			Modules:     []models.Module{},
		},
		{
			Title:       "Testing",
			Description: "Master software testing strategies including unit tests, integration tests, and test-driven development (TDD).",
			Modules:     []models.Module{},
		},
		{
			Title:       "Code Concepts",
			Description: "Fundamental programming paradigms and concepts that every developer should master.",
			Modules: []models.Module{
				{
					ID:    "code-1",
					Title: "Object Oriented Programming",
					Content: []models.ContentBlock{
						textBlock(`## What is Object-Oriented Programming?

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects" that contain data and code.

### The Four Pillars of OOP

1. **Encapsulation**: Bundling data and methods that operate on that data
2. **Inheritance**: Creating new classes based on existing ones
3. **Polymorphism**: Objects of different types responding to the same interface
4. **Abstraction**: Hiding complex implementation details`),
						calloutBlock("info", "OOP helps manage complexity in large codebases by organizing code into reusable, modular components."),
					},
				},
				{
					ID:    "code-2",
					Title: "Functional Programming",
					Content: []models.ContentBlock{
						textBlock(`## What is Functional Programming?

Functional Programming (FP) is a paradigm that treats computation as the evaluation of mathematical functions, avoiding changing state and mutable data.

### Key Concepts

- **Pure Functions**: Same input always produces same output, no side effects
- **Immutability**: Data cannot be changed after creation
- **First-Class Functions**: Functions can be passed as arguments and returned from other functions
- **Higher-Order Functions**: Functions that operate on other functions`),
						codeBlock("javascript", `// Pure function example
const add = (a, b) => a + b;

// Higher-order function
const map = (arr, fn) => arr.map(fn);

// Immutability - create new array instead of modifying
const numbers = [1, 2, 3];
const doubled = numbers.map(n => n * 2); // [2, 4, 6]`),
					},
				},
				{
					ID:    "code-3",
					Title: "Protocol Oriented Programming",
					Content: []models.ContentBlock{
						textBlock(`## Protocol-Oriented Programming

Protocol-Oriented Programming (POP) emphasizes defining behavior through protocols (interfaces) rather than through class inheritance.

### Why Protocol-Oriented?

- More flexible than class inheritance
- Allows composition over inheritance
- Better suited for value types
- Enables retroactive modeling`),
						calloutBlock("tip", "Swift popularized this paradigm, but the concepts apply to any language with interfaces or protocols."),
					},
				},
				{
					ID:    "code-4",
					Title: "Functions & Closures",
					Content: []models.ContentBlock{
						textBlock(`## Understanding Functions and Closures

Functions are reusable blocks of code. Closures are functions that capture and remember their surrounding context.

### What is a Closure?

A closure is a function that has access to variables from its outer (enclosing) scope, even after that outer function has returned.`),
						codeBlock("javascript", `function createCounter() {
  let count = 0;  // This variable is "captured" by the closure
  
  return function() {
    count += 1;
    return count;
  };
}

const counter = createCounter();
console.log(counter()); // 1
console.log(counter()); // 2
console.log(counter()); // 3`),
						calloutBlock("info", "Closures are powerful for creating private state, callbacks, and functional patterns like currying."),
					},
				},
			},
		},
	}

	// Insert all courses
	for _, course := range courses {
		if err := repo.CreateCourse(&course); err != nil {
			log.Printf("Error creating course %s: %v", course.Title, err)
			return err
		}
		log.Printf("Created course: %s", course.Title)
	}

	log.Println("Successfully seeded all courses!")
	return nil
}
