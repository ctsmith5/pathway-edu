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
						textBlock(`## Setting Up Your GitHub Account

Before we start using Git, let's set up GitHub - the most popular platform for hosting Git repositories and collaborating with others.

### Step 1: Create Your GitHub Account

1. Go to [github.com](https://github.com)
2. Click **Sign up** in the top right corner
3. Enter your email address
4. Create a strong password
5. Choose a username (this will be visible to others, so pick something professional!)
6. Complete the verification and click **Create account**
7. Verify your email address by clicking the link GitHub sends you

### Step 2: Configure Git on Your Computer

After installing Git, you need to tell it who you are. Open your terminal (Command Prompt, PowerShell, or Terminal) and run these commands:`),
						codeBlock("bash", `# Set your name (use your real name)
git config --global user.name "Your Name"

# Set your email (use the same email as your GitHub account)
git config --global user.email "your.email@example.com"

# Verify your settings
git config --list`),
						calloutBlock("tip", "The --global flag means these settings apply to all your Git projects. You only need to do this once per computer!"),
						textBlock(`### Step 3: Set Up Authentication

To push code to GitHub, you need to authenticate. The recommended method is using a Personal Access Token (PAT):

1. Go to GitHub → Click your profile picture → **Settings**
2. Scroll down to **Developer settings** (bottom of left sidebar)
3. Click **Personal access tokens** → **Tokens (classic)**
4. Click **Generate new token** → **Generate new token (classic)**
5. Give it a name like "My Laptop"
6. Set expiration (90 days is a good start)
7. Check the **repo** scope (this allows full access to repositories)
8. Click **Generate token**
9. **Copy the token immediately** - you won't be able to see it again!`),
						calloutBlock("warning", "Treat your Personal Access Token like a password. Never share it or commit it to a repository!"),
						textBlock(`When Git asks for your password (during push/pull), use this token instead of your GitHub password.

You're now ready to start using Git and GitHub together!`),
						exerciseBlock(
							"Configure Git with your name and email. What commands did you use?",
							`git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"`,
							[]string{"Use git config with the --global flag", "You need to set both user.name and user.email"},
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
				},
				{
					ID:    "git-3",
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
						textBlock(`## Hands-On: Push Your Hello World to GitHub

Now let's push the repository you created in the previous module to GitHub!

### Step 1: Create a Repository on GitHub

1. Go to [github.com](https://github.com) and sign in
2. Click the **+** icon in the top right corner
3. Select **New repository**
4. Name it **my-first-repo** (same as your local folder)
5. Leave it as **Public** (so your instructor can see it)
6. **Do NOT** initialize with a README, .gitignore, or license (your local repo already has content)
7. Click **Create repository**`),
						calloutBlock("warning", "Make sure you DON'T check any of the initialization options! We already have a local repository with commits."),
						textBlock(`### Step 2: Connect Your Local Repository to GitHub

GitHub will show you commands - use the "push an existing repository" option. In your terminal (make sure you're in the my-first-repo folder):`),
						codeBlock("bash", `# Add GitHub as your remote (replace YOUR-USERNAME with your GitHub username)
git remote add origin https://github.com/YOUR-USERNAME/my-first-repo.git

# Verify the remote was added
git remote -v`),
						textBlock(`### Step 3: Push Your Code to GitHub`),
						codeBlock("bash", `# Push your main branch to GitHub
git push -u origin main`),
						calloutBlock("tip", "The -u flag sets up tracking, so future pushes only need 'git push'. If asked for credentials, use your GitHub username and your Personal Access Token (not your password!)."),
						textBlock(`### Step 4: Verify on GitHub

1. Go to your repository on GitHub: https://github.com/YOUR-USERNAME/my-first-repo
2. You should see your HelloWorld.txt file!
3. Click on it to see the contents - your name and "Hello World!"

**Congratulations!** Your code is now on GitHub for the world to see!`),
						exerciseBlock(
							"You've pushed your first repository! What command did you use to connect your local repo to GitHub?",
							"git remote add origin https://github.com/YOUR-USERNAME/my-first-repo.git - This adds a remote called 'origin' pointing to your GitHub repository.",
							[]string{"Think about adding a remote", "The remote is conventionally called 'origin'"},
						),
					},
				},
				{
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
						textBlock(`## Hands-On: Your First Pull Request

Now it's time to complete your first real collaboration workflow! You'll create a Pull Request to merge your development branch into main and request a code review.

### Step 1: Go to Your Repository on GitHub

1. Navigate to https://github.com/YOUR-USERNAME/my-first-repo
2. You should see a banner saying "development had recent pushes" with a **Compare & pull request** button
3. If you don't see the banner, click the **Pull requests** tab, then **New pull request**`),
						calloutBlock("tip", "GitHub often shows a helpful banner when you've recently pushed a branch. This is the fastest way to create a PR!"),
						textBlock(`### Step 2: Create the Pull Request

1. Make sure the base branch is **main** and the compare branch is **development**
2. You'll see the changes you made (your bio additions)

Fill in the PR form:

**Title:** Add personal bio to HelloWorld.txt

**Description:**`),
						codeBlock("markdown", `## Summary
- Added a personal bio section to HelloWorld.txt
- Includes my background, interests, and goals

## Changes Made
- Updated HelloWorld.txt with "About Me" section
- Added favorite programming language
- Added project goals
- Added a fun fact

## Testing
- Verified the file displays correctly on GitHub
- Confirmed all information is accurate`),
						calloutBlock("info", "Writing good PR descriptions is a professional skill. Get in the habit now - your future teammates will thank you!"),
						textBlock(`### Step 3: Request a Code Review

On the right side of the PR page:

1. Find the **Reviewers** section
2. Click the gear icon
3. Search for your instructor's GitHub username
4. Select them to request a review`),
						calloutBlock("warning", "Ask your instructor for their GitHub username so you can request them as a reviewer. They'll review your work and provide feedback!"),
						textBlock(`### Step 4: Submit the Pull Request

1. Review everything looks correct
2. Click the green **Create pull request** button

**Congratulations!** You've just created your first Pull Request!

### What Happens Next?

1. **Reviewer gets notified**: Your instructor will receive an email/notification
2. **Review process**: They may:
   - Approve the PR (you can then merge it)
   - Request changes (you'll need to update your branch)
   - Leave comments (you can discuss and respond)
3. **Merge**: Once approved, you (or the reviewer) can merge the changes into main

### Step 5: After Approval - Merge Your PR

Once your instructor approves:

1. Click the green **Merge pull request** button
2. Click **Confirm merge**
3. Optionally, delete the development branch (GitHub will offer this)

Your bio is now part of the main branch!`),
						textBlock(`### Step 6: Update Your Local Repository

After merging on GitHub, update your local main branch:`),
						codeBlock("bash", `# Switch to main
git checkout main

# Pull the merged changes
git pull origin main

# Verify your bio is now in main
cat HelloWorld.txt`),
						calloutBlock("tip", "Always pull after merging on GitHub to keep your local repository in sync!"),
						textBlock(`## Summary: What You've Learned

In this hands-on project, you've completed a full Git workflow:

1. **Set up GitHub** - Created an account and configured Git
2. **Created a repository** - Both locally and on GitHub
3. **Made your first commit** - HelloWorld.txt with your name
4. **Pushed to GitHub** - Shared your code with the world
5. **Created a branch** - Separated development from main
6. **Made changes on a branch** - Added your bio
7. **Created a Pull Request** - Proposed merging your changes
8. **Requested a code review** - Got feedback from your instructor
9. **Merged your changes** - Completed the workflow

This is exactly how professional developers work every day!`),
						exerciseBlock(
							"You've completed your first Pull Request! What are the three possible outcomes when a reviewer looks at your PR?",
							"1. Approve - The reviewer is happy with your changes and you can merge\n2. Request changes - The reviewer wants you to modify something before merging\n3. Comment - The reviewer has questions or suggestions but hasn't made a decision yet\n\nAll three are normal parts of the code review process!",
							[]string{"Think about what a reviewer might do after looking at your code", "There are three main actions a reviewer can take"},
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
			Modules: []models.Module{
				{
					ID:    "solid-1",
					Title: "Introduction to SOLID Principles",
					Content: []models.ContentBlock{
						textBlock(`## What are SOLID Principles?

SOLID is an acronym for five object-oriented design principles that make software designs more understandable, flexible, and maintainable.

### The Five Principles

1. **S** - Single Responsibility Principle
2. **O** - Open/Closed Principle
3. **L** - Liskov Substitution Principle
4. **I** - Interface Segregation Principle
5. **D** - Dependency Inversion Principle`),
						calloutBlock("info", "SOLID principles were introduced by Robert C. Martin (Uncle Bob) in the early 2000s. They're guidelines, not strict rules - apply them thoughtfully."),
						textBlock(`## Why SOLID Matters

Following SOLID principles leads to:

- **Maintainability**: Easier to understand and modify code
- **Testability**: Components can be tested in isolation
- **Flexibility**: Changes don't cascade through the system
- **Reusability**: Components can be reused in different contexts
- **Scalability**: Codebase grows without becoming unmanageable`),
						textBlock(`## A Word of Caution

SOLID principles are tools, not goals. Over-applying them can lead to:
- Over-engineering
- Unnecessary complexity
- Premature abstraction

Use them when they solve real problems, not just because they exist.`),
						exerciseBlock(
							"What problem do SOLID principles primarily solve?",
							"SOLID principles solve the problem of code that becomes difficult to maintain, test, and extend as a project grows. They help prevent code rot and technical debt by promoting good design practices.",
							[]string{"Think about what happens to code over time", "Consider what makes code hard to change"},
						),
					},
				},
				{
					ID:    "solid-2",
					Title: "Single Responsibility Principle (SRP)",
					Content: []models.ContentBlock{
						textBlock(`## Single Responsibility Principle

**A class should have only one reason to change.**

Every class should have a single, well-defined responsibility. If a class has multiple reasons to change, it's doing too much.`),
						textBlock(`## Understanding "One Reason to Change"

A "reason to change" means a class should be responsible to only one stakeholder or concern:

- **User management** - one class
- **Email sending** - another class
- **Data persistence** - yet another class`),
						codeBlock("javascript", `// ❌ BAD: Multiple responsibilities
class User {
  constructor(name, email) {
    this.name = name;
    this.email = email;
  }
  
  saveToDatabase() {
    // Database logic
  }
  
  sendEmail() {
    // Email logic
  }
  
  validateEmail() {
    // Validation logic
  }
}

// ✅ GOOD: Single responsibility
class User {
  constructor(name, email) {
    this.name = name;
    this.email = email;
  }
}

class UserRepository {
  save(user) {
    // Database logic
  }
}

class EmailService {
  send(user, message) {
    // Email logic
  }
}

class EmailValidator {
  validate(email) {
    // Validation logic
  }
}`),
						calloutBlock("tip", "Ask yourself: 'If the requirements for X change, would I need to modify this class?' If the answer is yes for multiple X's, you're violating SRP."),
						exerciseBlock(
							"A class handles user authentication, logs all login attempts, and sends welcome emails. How would you refactor this to follow SRP?",
							"Split into three classes:\n1. AuthService - handles authentication logic\n2. LoginLogger - logs login attempts\n3. EmailService - sends welcome emails\n\nThe User class would coordinate these services but not contain their logic.",
							[]string{"Each responsibility should be its own class", "Think about what changes independently"},
						),
					},
				},
				{
					ID:    "solid-3",
					Title: "Open/Closed Principle (OCP)",
					Content: []models.ContentBlock{
						textBlock(`## Open/Closed Principle

**Software entities should be open for extension but closed for modification.**

You should be able to add new functionality without changing existing code.`),
						textBlock(`## What Does This Mean?

- **Open for Extension**: You can add new behavior
- **Closed for Modification**: You don't need to change existing code

This is achieved through:
- Inheritance
- Polymorphism
- Interfaces/Abstract classes`),
						codeBlock("javascript", `// ❌ BAD: Must modify existing code
class AreaCalculator {
  calculate(shape) {
    if (shape.type === 'circle') {
      return Math.PI * shape.radius ** 2;
    } else if (shape.type === 'rectangle') {
      return shape.width * shape.height;
    }
    // Adding a new shape requires modifying this method!
  }
}

// ✅ GOOD: Open for extension
class Shape {
  area() {
    throw new Error('Must implement area()');
  }
}

class Circle extends Shape {
  constructor(radius) {
    super();
    this.radius = radius;
  }
  
  area() {
    return Math.PI * this.radius ** 2;
  }
}

class Rectangle extends Shape {
  constructor(width, height) {
    super();
    this.width = width;
    this.height = height;
  }
  
  area() {
    return this.width * this.height;
  }
}

// Can add new shapes without modifying existing code!
class Triangle extends Shape {
  constructor(base, height) {
    super();
    this.base = base;
    this.height = height;
  }
  
  area() {
    return (this.base * this.height) / 2;
  }
}`),
						calloutBlock("warning", "OCP doesn't mean you can never modify code. It means you shouldn't have to modify stable, tested code to add new features."),
						exerciseBlock(
							"You have a PaymentProcessor class with if/else statements for different payment methods (credit card, PayPal). How would you refactor this to follow OCP?",
							"Create a PaymentMethod interface/abstract class with a process() method. Each payment method (CreditCard, PayPal, etc.) implements this interface. The PaymentProcessor can then work with any PaymentMethod without knowing the specific type, and new payment methods can be added without modifying existing code.",
							[]string{"Think about polymorphism", "How can you make it extensible without changing existing code?"},
						),
					},
				},
				{
					ID:    "solid-4",
					Title: "Liskov Substitution Principle (LSP)",
					Content: []models.ContentBlock{
						textBlock(`## Liskov Substitution Principle

**Objects of a superclass should be replaceable with objects of its subclasses without breaking the application.**

Named after Barbara Liskov, this principle ensures that inheritance is used correctly.`),
						textBlock(`## The Contract

Subtypes must:
- Accept the same input parameters as their base type
- Return compatible types
- Not throw exceptions that the base type doesn't throw
- Maintain the same behavioral guarantees`),
						codeBlock("javascript", `// ❌ BAD: Violates LSP
class Rectangle {
  constructor(width, height) {
    this.width = width;
    this.height = height;
  }
  
  setWidth(width) {
    this.width = width;
  }
  
  setHeight(height) {
    this.height = height;
  }
  
  area() {
    return this.width * this.height;
  }
}

class Square extends Rectangle {
  constructor(side) {
    super(side, side);
  }
  
  setWidth(width) {
    this.width = width;
    this.height = width; // Problem!
  }
  
  setHeight(height) {
    this.width = height; // Problem!
    this.height = height;
  }
}

// This breaks when substituting Square for Rectangle
function testArea(rect) {
  rect.setWidth(5);
  rect.setHeight(4);
  return rect.area(); // Expects 20, but Square gives 16!
}

// ✅ GOOD: Don't force Square to be a Rectangle
class Shape {
  area() {
    throw new Error('Must implement');
  }
}

class Rectangle extends Shape {
  constructor(width, height) {
    super();
    this.width = width;
    this.height = height;
  }
  
  area() {
    return this.width * this.height;
  }
}

class Square extends Shape {
  constructor(side) {
    super();
    this.side = side;
  }
  
  area() {
    return this.side ** 2;
  }
}`),
						calloutBlock("tip", "If you find yourself checking the type of a subclass before using it, you're likely violating LSP. The code should work with the base type."),
						exerciseBlock(
							"A Bird class has a fly() method. A Penguin class extends Bird but can't fly. How does this violate LSP and how would you fix it?",
							"This violates LSP because code expecting a Bird to fly would break with a Penguin. Solutions:\n1. Don't make Penguin extend Bird if it can't fulfill Bird's contract\n2. Create a FlyingBird subclass that Bird extends, and have Penguin extend Bird directly\n3. Use composition: have a Flyable interface that only flying birds implement",
							[]string{"Think about what guarantees Bird makes", "Can all birds fly?"},
						),
					},
				},
				{
					ID:    "solid-5",
					Title: "Interface Segregation Principle (ISP)",
					Content: []models.ContentBlock{
						textBlock(`## Interface Segregation Principle

**Clients should not be forced to depend on interfaces they do not use.**

Create specific, focused interfaces rather than large, general-purpose ones.`),
						textBlock(`## The Problem with Fat Interfaces

When interfaces are too large:
- Classes must implement methods they don't need
- Changes to one part affect unrelated classes
- Code becomes harder to understand and maintain`),
						codeBlock("javascript", `// ❌ BAD: Fat interface
class Worker {
  work() {}
  eat() {}
  sleep() {}
}

class Human extends Worker {
  work() { /* humans work */ }
  eat() { /* humans eat */ }
  sleep() { /* humans sleep */ }
}

class Robot extends Worker {
  work() { /* robots work */ }
  eat() { /* robots don't eat! */ }
  sleep() { /* robots don't sleep! */ }
}

// ✅ GOOD: Segregated interfaces
class Workable {
  work() {}
}

class Eatable {
  eat() {}
}

class Sleepable {
  sleep() {}
}

class Human implements Workable, Eatable, Sleepable {
  work() { /* humans work */ }
  eat() { /* humans eat */ }
  sleep() { /* humans sleep */ }
}

class Robot implements Workable {
  work() { /* robots work */ }
  // No need to implement eat() or sleep()
}`),
						calloutBlock("info", "ISP is about creating lean, focused interfaces. Each interface should represent a specific role or capability."),
						textBlock(`## Benefits of ISP

- **Flexibility**: Classes only depend on what they need
- **Maintainability**: Changes to one interface don't affect unrelated classes
- **Clarity**: Interfaces clearly communicate their purpose
- **Testability**: Easier to mock and test specific behaviors`),
						exerciseBlock(
							"You have a Printer interface with print(), scan(), fax(), and copy() methods. A basic printer only needs print(). How would you refactor this?",
							"Split into separate interfaces:\n- Printable: print()\n- Scannable: scan()\n- Faxable: fax()\n- Copyable: copy()\n\nA basic printer implements only Printable. Advanced printers can implement multiple interfaces as needed.",
							[]string{"Not all printers have all capabilities", "Think about what each device actually needs"},
						),
					},
				},
				{
					ID:    "solid-6",
					Title: "Dependency Inversion Principle (DIP)",
					Content: []models.ContentBlock{
						textBlock(`## Dependency Inversion Principle

**High-level modules should not depend on low-level modules. Both should depend on abstractions.**

Depend on interfaces, not concrete implementations.`),
						textBlock(`## Understanding the Layers

- **High-level modules**: Business logic, application flow
- **Low-level modules**: Database access, file I/O, external APIs
- **Abstractions**: Interfaces, abstract classes, protocols`),
						codeBlock("javascript", `// ❌ BAD: High-level depends on low-level
class MySQLDatabase {
  save(data) {
    // MySQL-specific code
  }
}

class UserService {
  constructor() {
    this.db = new MySQLDatabase(); // Direct dependency!
  }
  
  createUser(user) {
    this.db.save(user);
  }
}

// ✅ GOOD: Both depend on abstraction
class Database {
  save(data) {
    throw new Error('Must implement');
  }
}

class MySQLDatabase extends Database {
  save(data) {
    // MySQL-specific code
  }
}

class PostgreSQLDatabase extends Database {
  save(data) {
    // PostgreSQL-specific code
  }
}

class UserService {
  constructor(database) { // Depends on abstraction
    this.db = database;
  }
  
  createUser(user) {
    this.db.save(user);
  }
}

// Can swap implementations easily
const userService = new UserService(new MySQLDatabase());
// or
const userService = new UserService(new PostgreSQLDatabase());`),
						calloutBlock("tip", "Dependency Injection is the technique used to implement DIP. Pass dependencies in rather than creating them internally."),
						textBlock(`## Benefits of DIP

- **Flexibility**: Easy to swap implementations
- **Testability**: Can inject mocks for testing
- **Maintainability**: Changes to low-level modules don't affect high-level ones
- **Reusability**: High-level modules can work with any implementation`),
						exerciseBlock(
							"A NotificationService directly creates instances of EmailSender and SMSSender. How would you refactor this to follow DIP?",
							"Create a Notifier interface with a send() method. EmailSender and SMSSender implement this interface. NotificationService depends on the Notifier interface and receives it via dependency injection. This allows easy swapping of notification methods and better testing.",
							[]string{"What abstraction could represent both email and SMS?", "How can you inject the dependency?"},
						),
					},
				},
				{
					ID:    "solid-7",
					Title: "Applying SOLID in Practice",
					Content: []models.ContentBlock{
						textBlock(`## Putting It All Together

SOLID principles work best when applied together. Let's see how they complement each other.`),
						textBlock(`## A Real-World Example

Consider an e-commerce order processing system:

- **SRP**: Separate classes for Order, Payment, Shipping, Notification
- **OCP**: Add new payment methods without modifying existing code
- **LSP**: Any PaymentMethod implementation works the same way
- **ISP**: Separate interfaces for Payable, Shippable, Notifiable
- **DIP**: OrderService depends on Payment interface, not concrete classes`),
						codeBlock("javascript", `// Example: SOLID e-commerce system
class Order {
  constructor(items, customer) {
    this.items = items;
    this.customer = customer;
  }
  
  total() {
    return this.items.reduce((sum, item) => sum + item.price, 0);
  }
}

// ISP: Segregated interfaces
class Payable {
  processPayment(amount) {}
}

class Shippable {
  ship(order) {}
}

// OCP: Open for extension
class CreditCardPayment extends Payable {
  processPayment(amount) {
    // Credit card logic
  }
}

class PayPalPayment extends Payable {
  processPayment(amount) {
    // PayPal logic
  }
}

// DIP: Depend on abstraction
class OrderService {
  constructor(paymentMethod, shippingService) {
    this.payment = paymentMethod; // Abstraction
    this.shipping = shippingService; // Abstraction
  }
  
  processOrder(order) {
    this.payment.processPayment(order.total());
    this.shipping.ship(order);
  }
}`),
						calloutBlock("warning", "Don't apply SOLID blindly. Start simple, and refactor when you see the need. Premature application can lead to over-engineering."),
						textBlock(`## When to Apply SOLID

**Apply SOLID when:**
- Code is becoming hard to change
- You're writing similar code multiple times
- Testing is difficult
- Adding features requires modifying many files

**Don't over-apply when:**
- The codebase is small and simple
- Requirements are still changing rapidly
- You're not sure what the future needs will be`),
						exerciseBlock(
							"Your codebase has a User class that handles authentication, sends emails, logs activity, and saves to database. How would you refactor this using SOLID principles?",
							"1. SRP: Split into User (data), AuthService, EmailService, Logger, UserRepository\n2. OCP: Make services extensible via interfaces\n3. LSP: Ensure all implementations follow their contracts\n4. ISP: Create focused interfaces (Authenticatable, Emailable, Loggable)\n5. DIP: UserService depends on these interfaces, receives them via injection",
							[]string{"Each principle addresses a different aspect", "Think about how they work together"},
						),
					},
				},
			},
		},
		{
			Title:       "SCRUM",
			Description: "Learn the SCRUM framework for agile project management. Understand sprints, standups, and how to deliver value iteratively.",
			Modules: []models.Module{
				{
					ID:    "scrum-1",
					Title: "Introduction to Agile & SCRUM",
					Content: []models.ContentBlock{
						textBlock(`## What is Agile?

Agile is a mindset and set of values for software development that emphasizes:

- **Individuals and interactions** over processes and tools
- **Working software** over comprehensive documentation
- **Customer collaboration** over contract negotiation
- **Responding to change** over following a plan`),
						calloutBlock("info", "The Agile Manifesto was created in 2001 by 17 software developers. It's not about abandoning processes or documentation, but prioritizing what matters most."),
						textBlock(`## What is SCRUM?

SCRUM is a framework for implementing Agile. It provides:

- **Structure**: Roles, events, and artifacts
- **Process**: How teams work together
- **Values**: Commitment, courage, focus, openness, respect

SCRUM is lightweight, simple to understand, but difficult to master.`),
						textBlock(`## The SCRUM Framework

SCRUM consists of:

1. **Roles**: Product Owner, Scrum Master, Development Team
2. **Events**: Sprint, Sprint Planning, Daily Scrum, Sprint Review, Sprint Retrospective
3. **Artifacts**: Product Backlog, Sprint Backlog, Increment`),
						textBlock(`## Why SCRUM?

SCRUM helps teams:

- Deliver value frequently
- Adapt to changing requirements
- Improve continuously
- Increase transparency
- Reduce risk`),
						exerciseBlock(
							"What's the difference between Agile and SCRUM?",
							"Agile is a philosophy and set of values about how to approach software development. SCRUM is a specific framework that implements Agile principles. You can be Agile without using SCRUM (using Kanban, XP, etc.), but SCRUM is always Agile.",
							[]string{"Think about philosophy vs. implementation", "Is SCRUM the only way to be Agile?"},
						),
					},
				},
				{
					ID:    "scrum-2",
					Title: "SCRUM Roles (Product Owner, Scrum Master, Team)",
					Content: []models.ContentBlock{
						textBlock(`## The Three SCRUM Roles

SCRUM defines three distinct roles, each with specific responsibilities.`),
						textBlock(`## Product Owner

The Product Owner is responsible for:

- **Maximizing product value**: Decides what to build
- **Managing the Product Backlog**: Prioritizes and maintains items
- **Stakeholder communication**: Represents customer needs
- **Accepting work**: Decides if items meet the Definition of Done`),
						calloutBlock("tip", "The Product Owner is ONE person, not a committee. They make decisions, though they should gather input from stakeholders."),
						textBlock(`## Scrum Master

The Scrum Master serves the team by:

- **Removing impediments**: Helps team overcome obstacles
- **Facilitating events**: Ensures SCRUM events happen and are effective
- **Coaching**: Helps team understand and apply SCRUM
- **Protecting the team**: Shields team from outside interference`),
						textBlock(`## Development Team

The Development Team:

- **Self-organizes**: Decides how to do the work
- **Cross-functional**: Has all skills needed to deliver
- **Accountable**: Collectively responsible for the increment
- **Sized appropriately**: Typically 3-9 people`),
						codeBlock("text", `Key Characteristics:

Development Team:
- No titles (no "senior developer" hierarchy in SCRUM)
- No sub-teams
- Works together on all items
- Estimates work collectively`),
						calloutBlock("warning", "The Development Team is NOT managed by the Scrum Master. The team is self-organizing and decides how to accomplish the work."),
						exerciseBlock(
							"Who is responsible for deciding HOW to implement a feature in SCRUM?",
							"The Development Team is responsible for deciding HOW to implement features. The Product Owner decides WHAT to build, but the team decides the technical approach. The Scrum Master helps ensure the team can work effectively.",
							[]string{"Think about the separation of concerns", "Who has the technical expertise?"},
						),
					},
				},
				{
					ID:    "scrum-3",
					Title: "SCRUM Events (Sprint, Daily Standup, Retrospective)",
					Content: []models.ContentBlock{
						textBlock(`## SCRUM Events

SCRUM has five formal events, all designed to create regularity and minimize meetings:

1. **Sprint**: The container for all other events
2. **Sprint Planning**: Plan the work for the sprint
3. **Daily Scrum**: 15-minute team sync
4. **Sprint Review**: Inspect the increment
5. **Sprint Retrospective**: Improve the process`),
						textBlock(`## The Sprint

The Sprint is:

- **Time-boxed**: 1-4 weeks (most commonly 2 weeks)
- **Fixed length**: Same duration for all sprints
- **Goal-focused**: Has a Sprint Goal
- **No changes**: Scope is locked during the sprint

During a sprint:
- No changes that endanger the Sprint Goal
- Quality doesn't decrease
- Scope may be clarified and renegotiated`),
						textBlock(`## Sprint Planning

**Time-box**: 8 hours for a 4-week sprint (proportionally less for shorter sprints)

**Two parts:**

1. **What can be done?** - Product Owner presents backlog, team selects items
2. **How will it be done?** - Team plans the work

**Output**: Sprint Goal and Sprint Backlog`),
						textBlock(`## Daily Scrum

**Time-box**: 15 minutes, same time and place every day

**Purpose**: 
- Plan work for the next 24 hours
- Inspect progress toward Sprint Goal
- Identify impediments

**Format**: Each team member answers:
- What did I do yesterday?
- What will I do today?
- Are there any impediments?`),
						calloutBlock("tip", "The Daily Scrum is NOT a status report to the Scrum Master. It's for the Development Team to synchronize and plan."),
						textBlock(`## Sprint Review

**Time-box**: 4 hours for a 4-week sprint

**Purpose**: Inspect the increment and adapt the Product Backlog

**Attendees**: Scrum Team + Stakeholders

**Activities**:
- Demo working software
- Discuss what was done
- Gather feedback
- Update Product Backlog`),
						textBlock(`## Sprint Retrospective

**Time-box**: 3 hours for a 4-week sprint

**Purpose**: Plan improvements for the next sprint

**Questions**:
- What went well?
- What could be improved?
- What will we do differently?

**Output**: Actionable improvements`),
						exerciseBlock(
							"Why are SCRUM events time-boxed?",
							"Time-boxing ensures events don't drag on unnecessarily, creates predictability, and forces focus. It also prevents over-planning and ensures the team spends most of their time on actual development work rather than meetings.",
							[]string{"Think about what happens without time limits", "What's the cost of long meetings?"},
						),
					},
				},
				{
					ID:    "scrum-4",
					Title: "SCRUM Artifacts (Product Backlog, Sprint Backlog, Increment)",
					Content: []models.ContentBlock{
						textBlock(`## SCRUM Artifacts

SCRUM has three artifacts that provide transparency and opportunities for inspection:

1. **Product Backlog**: Ordered list of everything needed
2. **Sprint Backlog**: Work selected for the current sprint
3. **Increment**: Sum of all completed work`),
						textBlock(`## Product Backlog

The Product Backlog is:

- **Ordered**: Most valuable items first
- **Emergent**: Never complete, always evolving
- **Detailed appropriately**: Top items are detailed, lower items are vague
- **Owned by Product Owner**: But team contributes estimates

**Contains**:
- Features
- Bugs
- Technical work
- Knowledge acquisition`),
						calloutBlock("info", "The Product Backlog is the single source of truth for what needs to be done. Nothing gets done that's not in the backlog."),
						textBlock(`## Product Backlog Items

Each item has:

- **Description**: What it is
- **Order**: Priority/position
- **Estimate**: Effort (usually Story Points)
- **Value**: Why it matters

**Refinement**:
- Continuously refined
- Items at top are "ready" (clear, feasible, testable)
- Team helps clarify and estimate`),
						textBlock(`## Sprint Backlog

The Sprint Backlog:

- **Selected during Sprint Planning**: Team commits to items
- **Owned by Development Team**: They decide how to do it
- **Visible to all**: Shows progress
- **Updated daily**: During Daily Scrum

**Contains**:
- Selected Product Backlog items
- Plan for delivering the increment
- At least one high-priority process improvement`),
						textBlock(`## The Increment

The Increment is:

- **Sum of all work**: All Product Backlog items completed in the sprint
- **Potentially releasable**: Meets Definition of Done
- **Usable**: Stakeholders can use it
- **Additive**: Each increment builds on previous ones

**Definition of Done**:
- Criteria that must be met for work to be "done"
- Includes testing, documentation, code review, etc.
- Prevents technical debt`),
						calloutBlock("warning", "An increment must be 'done' according to the Definition of Done. Partially done work doesn't count as an increment."),
						exerciseBlock(
							"What's the difference between the Product Backlog and Sprint Backlog?",
							"The Product Backlog contains ALL work for the product, ordered by value. The Sprint Backlog contains only the work selected for the current sprint, plus the plan for how to do it. The Sprint Backlog is a subset of the Product Backlog.",
							[]string{"Think about scope - what's the difference in size?", "When is each created and updated?"},
						),
					},
				},
				{
					ID:    "scrum-5",
					Title: "Sprint Planning",
					Content: []models.ContentBlock{
						textBlock(`## Sprint Planning Overview

Sprint Planning initiates the sprint by planning the work to be performed. It's time-boxed to a maximum of 8 hours for a one-month sprint.

**Participants**: Entire Scrum Team
**Facilitator**: Scrum Master`),
						textBlock(`## Part 1: What Can Be Done?

The Product Owner presents:

- **Product Backlog**: Ordered list of items
- **Product Goal**: Long-term objective
- **Forecast**: What might be achievable

The Development Team:

- Considers capacity
- Considers past performance
- Selects items they believe they can complete
- Creates a **Sprint Goal**`),
						textBlock(`## The Sprint Goal

The Sprint Goal:

- Provides focus and coherence
- Guides the team on why they're building the increment
- Can be achieved through various Product Backlog items
- If the goal becomes obsolete, the sprint may be cancelled`),
						calloutBlock("tip", "A good Sprint Goal is specific enough to provide direction but flexible enough to allow the team to adapt as they learn."),
						textBlock(`## Part 2: How Will It Be Done?

The Development Team plans:

- **How to build**: Technical approach
- **Tasks**: Breaks down items into tasks
- **Collaboration**: Who works on what
- **Initial plan**: For the first days of the sprint

The Product Owner:
- Answers questions
- Clarifies requirements
- May renegotiate scope if needed`),
						codeBlock("text", `Example Sprint Planning:

Sprint Goal: "Enable users to reset their password securely"

Selected Items:
- User story: "As a user, I want to reset my password"
- Technical task: "Set up email service"
- Bug fix: "Fix password validation"

Tasks:
- Design password reset flow
- Create reset token system
- Build email template
- Write tests
- Update documentation`),
						textBlock(`## Commitment

By the end of Sprint Planning:

- **Sprint Goal** is set
- **Sprint Backlog** is created
- Team is **committed** to the goal

Note: Commitment doesn't mean guarantee. It means the team will do their best to achieve the goal.`),
						exerciseBlock(
							"During Sprint Planning, the team realizes they can't complete all the items the Product Owner wanted. What should happen?",
							"The team should communicate this to the Product Owner. They can either:\n1. Reduce the scope to what's achievable\n2. Extend the sprint (not recommended - breaks time-boxing)\n3. Adjust the Sprint Goal to be more realistic\n\nThe Product Owner and team collaborate to find the right balance. The team should never commit to more than they believe they can deliver.",
							[]string{"Think about transparency and honesty", "What's the purpose of the Sprint Goal?"},
						),
					},
				},
				{
					ID:    "scrum-6",
					Title: "User Stories & Estimation",
					Content: []models.ContentBlock{
						textBlock(`## User Stories

User Stories are a way to express Product Backlog items from the user's perspective.

**Format**:
"As a [type of user], I want [some goal] so that [some reason]"

**Example**:
"As a customer, I want to save my payment information so that I can checkout faster next time."`),
						textBlock(`## The Three C's

User Stories have three critical aspects:

1. **Card**: Written description (the story itself)
2. **Conversation**: Discussion about details
3. **Confirmation**: Acceptance criteria (tests)`),
						calloutBlock("info", "User Stories are placeholders for conversations. The real value comes from the discussion, not just the written card."),
						textBlock(`## Acceptance Criteria

Acceptance criteria define when a story is "done":

**Example**:
- User can enter email and password
- System validates email format
- Error message shown for invalid email
- Success message shown after submission
- Email confirmation sent`),
						textBlock(`## Story Points

Story Points estimate relative effort, not time.

**Why relative?**
- Easier to compare than absolute time
- Accounts for complexity, uncertainty, and effort
- Team-specific (one team's 5 might be another's 8)

**Common scales**:
- Fibonacci: 1, 2, 3, 5, 8, 13, 21
- Powers of 2: 1, 2, 4, 8, 16
- T-shirt sizes: XS, S, M, L, XL`),
						codeBlock("text", `Estimation Example:

Story 1: "Login page" - 3 points (baseline)
Story 2: "Password reset" - 5 points (more complex)
Story 3: "User profile" - 8 points (much more complex)
Story 4: "Add comment" - 2 points (simpler than login)`),
						textBlock(`## Planning Poker

A technique for estimating:

1. Product Owner reads story
2. Team asks questions
3. Each person privately selects a card (estimate)
4. Everyone reveals simultaneously
5. Discuss differences
6. Re-estimate until consensus`),
						calloutBlock("tip", "The goal of estimation is not precision, but shared understanding. If estimates vary widely, it usually means the story needs clarification."),
						textBlock(`## Velocity

Velocity is:

- **Measure of capacity**: How many points a team completes per sprint
- **Historical data**: Based on past performance
- **Used for forecasting**: Helps predict future capacity
- **Team-specific**: Not comparable across teams`),
						textBlock(`## INVEST Criteria

Good user stories are:

- **I**ndependent: Can be developed in any order
- **N**egotiable: Details can be discussed
- **V**aluable: Delivers value to users
- **E**stimable: Team can estimate effort
- **S**mall: Can be completed in one sprint
- **T**estable: Has clear acceptance criteria`),
						exerciseBlock(
							"A user story is estimated at 13 points. The team's average velocity is 20 points per sprint. Is this story too large?",
							"Yes, this story is likely too large. At 13 points, it's 65% of the team's capacity, which makes it risky. Large stories should be broken down into smaller stories (ideally 1-8 points) to reduce risk and increase predictability. The team should split this into 2-3 smaller stories.",
							[]string{"Think about risk and predictability", "What happens if a large story isn't finished?"},
						),
					},
				},
				{
					ID:    "scrum-7",
					Title: "Running Effective Sprints",
					Content: []models.ContentBlock{
						textBlock(`## Sprint Execution

During a sprint, the Development Team:

- Works on Sprint Backlog items
- Collaborates daily
- Adapts the plan as needed
- Maintains quality standards
- Works toward the Sprint Goal`),
						textBlock(`## Daily Practices

**Daily Scrum**:
- Same time, same place
- 15 minutes maximum
- Focus on coordination, not status reporting
- Identify impediments

**Continuous Integration**:
- Integrate code frequently
- Run automated tests
- Keep the build green

**Pair Programming**:
- Share knowledge
- Improve code quality
- Reduce bottlenecks`),
						calloutBlock("tip", "The best sprints are boring - predictable, steady progress. Drama and heroics are signs of problems."),
						textBlock(`## Handling Impediments

Impediments block progress:

- **Identify quickly**: Raise in Daily Scrum
- **Escalate appropriately**: Scrum Master helps remove
- **Track visibly**: Use an impediment board
- **Resolve promptly**: Don't let them linger`),
						textBlock(`## Adapting During Sprint

The Sprint Backlog is:

- **Emergent**: Can be updated as team learns
- **Owned by team**: They decide how to adapt
- **Goal-focused**: Changes should support Sprint Goal

**Allowed changes**:
- Clarifying requirements
- Breaking down items further
- Adjusting tasks

**Not allowed**:
- Adding new Product Backlog items (unless Sprint Goal changes)
- Changing Sprint Goal without cancelling sprint`),
						textBlock(`## Definition of Done

The Definition of Done ensures quality:

**Example**:
- Code written and reviewed
- Unit tests written and passing
- Integration tests passing
- Documentation updated
- Deployed to staging
- Product Owner acceptance

**Benefits**:
- Prevents technical debt
- Ensures consistency
- Makes "done" clear to all`),
						textBlock(`## Sprint Cancellation

A sprint can be cancelled if:

- Sprint Goal becomes obsolete
- Market conditions change dramatically
- Technology becomes unavailable

**Process**:
- Only Product Owner can cancel
- Completed items are reviewed
- Incomplete items return to Product Backlog
- Team reflects on what happened`),
						textBlock(`## Common Anti-Patterns

**Avoid these**:

- ❌ Extending sprints when behind
- ❌ Adding work mid-sprint
- ❌ Skipping retrospectives
- ❌ Daily Scrum as status report
- ❌ Unclear Definition of Done
- ❌ Product Owner dictating how to work
- ❌ Scrum Master as team manager`),
						calloutBlock("warning", "SCRUM is simple but not easy. It requires discipline, transparency, and a commitment to continuous improvement."),
						exerciseBlock(
							"Halfway through a sprint, a critical bug is discovered in production. The Product Owner wants the team to drop everything and fix it. How should this be handled?",
							"This depends on the severity:\n\n1. If it's truly critical (system down, data loss), the Product Owner can cancel the sprint and the team addresses it immediately.\n\n2. If it's important but not critical, the Product Owner and team discuss:\n   - Can it wait until next sprint?\n   - Can we swap it for a lower-priority item?\n   - Does it change the Sprint Goal?\n\nThe key is transparency and collaboration, not just dictating changes.",
							[]string{"Think about the Sprint Goal", "What's the impact of mid-sprint changes?"},
						),
					},
				},
			},
		},
		{
			Title:       "HTTP Networking",
			Description: "Deep dive into HTTP protocols, REST APIs, request/response cycles, and modern web communication patterns.",
			Modules: []models.Module{
				{
					ID:    "http-1",
					Title: "HTTP Fundamentals",
					Content: []models.ContentBlock{
						textBlock(`## What is HTTP?

HTTP (HyperText Transfer Protocol) is the foundation of data communication on the World Wide Web.

**Key characteristics**:
- **Stateless**: Each request is independent
- **Request-Response**: Client sends request, server responds
- **Application layer**: Works on top of TCP/IP
- **Text-based**: Human-readable protocol`),
						calloutBlock("info", "HTTP was created by Tim Berners-Lee in 1989. HTTP/1.1 (1997) is still widely used, but HTTP/2 (2015) and HTTP/3 (2020) offer improvements."),
						textBlock(`## How HTTP Works

1. **Client** (browser/app) sends HTTP request
2. **Server** receives and processes request
3. **Server** sends HTTP response
4. **Connection** may close (HTTP/1.1) or stay open (HTTP/2+)`),
						codeBlock("text", `Example HTTP Request:

GET /api/users HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0
Accept: application/json

Example HTTP Response:

HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 156

{
  "users": [
    {"id": 1, "name": "Alice"},
    {"id": 2, "name": "Bob"}
  ]
}`),
						textBlock(`## HTTP Components

**Request Line**:
- Method (GET, POST, etc.)
- Path (/api/users)
- HTTP Version

**Headers**:
- Metadata about request/response
- Key-value pairs
- Examples: Content-Type, Authorization

**Body** (optional):
- Data sent with request
- Used in POST, PUT, PATCH`),
						textBlock(`## URLs and URIs

**URL** (Uniform Resource Locator):
- Full address: https://example.com/api/users?id=123
- Includes protocol, domain, path, query

**URI** (Uniform Resource Identifier):
- More general term
- URL is a type of URI

**Components**:
- **Scheme**: http:// or https://
- **Host**: example.com
- **Path**: /api/users
- **Query**: ?id=123
- **Fragment**: #section`),
						exerciseBlock(
							"What does it mean that HTTP is stateless?",
							"HTTP is stateless means each request is independent - the server doesn't remember previous requests. Each request must contain all information needed to process it. This is why we use cookies, sessions, and tokens to maintain state across requests.",
							[]string{"Think about what 'state' means", "How do websites remember you're logged in?"},
						),
					},
				},
				{
					ID:    "http-2",
					Title: "HTTP Methods (GET, POST, PUT, DELETE, etc.)",
					Content: []models.ContentBlock{
						textBlock(`## HTTP Methods

HTTP methods (also called verbs) indicate the action to be performed on a resource.

**Common methods**:
- GET: Retrieve data
- POST: Create new resource
- PUT: Update/replace resource
- PATCH: Partial update
- DELETE: Remove resource
- HEAD: Get headers only
- OPTIONS: Get allowed methods`),
						textBlock(`## GET - Retrieve Data

**Purpose**: Fetch data from server

**Characteristics**:
- Idempotent (same request = same result)
- Safe (doesn't modify server state)
- Can be cached
- Parameters in URL query string

**Example**:
GET /api/users?page=1&limit=10`),
						textBlock(`## POST - Create Resource

**Purpose**: Submit data to create new resource

**Characteristics**:
- Not idempotent (multiple requests = multiple resources)
- Not safe (modifies server state)
- Data in request body
- Returns created resource

**Example**:
POST /api/users
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com"
}`),
						codeBlock("javascript", `// Example: Making HTTP requests

// GET request
fetch('/api/users')
  .then(response => response.json())
  .then(data => console.log(data));

// POST request
fetch('/api/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'Alice',
    email: 'alice@example.com'
  })
})
  .then(response => response.json())
  .then(data => console.log(data));`),
						textBlock(`## PUT - Replace Resource

**Purpose**: Replace entire resource

**Characteristics**:
- Idempotent (same request = same result)
- Replaces entire resource
- Must include all fields

**Example**:
PUT /api/users/123
Content-Type: application/json

{
  "name": "Alice Updated",
  "email": "alice.new@example.com"
}`),
						textBlock(`## PATCH - Partial Update

**Purpose**: Update part of a resource

**Characteristics**:
- Idempotent
- Only sends changed fields
- More efficient than PUT

**Example**:
PATCH /api/users/123
Content-Type: application/json

{
  "email": "alice.new@example.com"
}`),
						textBlock(`## DELETE - Remove Resource

**Purpose**: Delete a resource

**Characteristics**:
- Idempotent
- Removes resource
- May return 204 No Content

**Example**:
DELETE /api/users/123`),
						calloutBlock("tip", "Use the right method for the right purpose. GET for reading, POST for creating, PUT/PATCH for updating, DELETE for removing."),
						exerciseBlock(
							"What's the difference between PUT and PATCH?",
							"PUT replaces the entire resource - you must send all fields, even unchanged ones. PATCH performs a partial update - you only send the fields you want to change. PUT is idempotent (replacing with same data has same effect), and PATCH should also be idempotent.",
							[]string{"Think about what 'replace' vs 'update' means", "What data do you need to send?"},
						),
					},
				},
				{
					ID:    "http-3",
					Title: "HTTP Status Codes",
					Content: []models.ContentBlock{
						textBlock(`## HTTP Status Codes

Status codes are three-digit numbers that indicate the result of an HTTP request.

**Format**: XXX
- First digit: Category
- Last two digits: Specific code`),
						textBlock(`## Status Code Categories

**1xx - Informational**:
- Request received, continuing process
- Rarely used

**2xx - Success**:
- Request successful
- Most common: 200 OK

**3xx - Redirection**:
- Further action needed
- Common: 301, 302, 304

**4xx - Client Error**:
- Request has error
- Common: 400, 401, 403, 404

**5xx - Server Error**:
- Server failed
- Common: 500, 502, 503`),
						textBlock(`## Common 2xx Codes

**200 OK**:
- Request successful
- Most common success response

**201 Created**:
- Resource created successfully
- Used with POST
- Should include Location header

**204 No Content**:
- Success but no content to return
- Used with DELETE or PUT

**206 Partial Content**:
- Partial content returned
- Used for range requests`),
						textBlock(`## Common 3xx Codes

**301 Moved Permanently**:
- Resource moved permanently
- Browser caches redirect

**302 Found (Temporary Redirect)**:
- Resource temporarily at different location
- Browser doesn't cache

**304 Not Modified**:
- Resource hasn't changed
- Used for caching
- Client can use cached version`),
						textBlock(`## Common 4xx Codes

**400 Bad Request**:
- Request malformed
- Client error in request format

**401 Unauthorized**:
- Authentication required
- Missing or invalid credentials

**403 Forbidden**:
- Authenticated but not authorized
- Valid credentials but insufficient permissions

**404 Not Found**:
- Resource doesn't exist
- Most famous status code!

**409 Conflict**:
- Request conflicts with current state
- Used for duplicate resources

**422 Unprocessable Entity**:
- Request well-formed but semantically invalid
- Validation errors`),
						textBlock(`## Common 5xx Codes

**500 Internal Server Error**:
- Generic server error
- Something went wrong on server

**502 Bad Gateway**:
- Server acting as gateway got invalid response
- Upstream server issue

**503 Service Unavailable**:
- Server temporarily unavailable
- Overloaded or maintenance

**504 Gateway Timeout**:
- Gateway didn't receive timely response
- Upstream server timeout`),
						codeBlock("javascript", `// Handling status codes

fetch('/api/users')
  .then(response => {
    if (response.ok) {
      // 200-299 range
      return response.json();
    } else if (response.status === 404) {
      throw new Error('User not found');
    } else if (response.status === 401) {
      throw new Error('Unauthorized');
    } else {
      throw new Error('Request failed');
    }
  })
  .then(data => console.log(data))
  .catch(error => console.error(error));`),
						calloutBlock("warning", "Use status codes correctly. 401 for authentication issues, 403 for authorization issues. 400 for client errors, 500 for server errors."),
						exerciseBlock(
							"A user tries to access a resource they don't have permission for. They're logged in. What status code should you return?",
							"Return 403 Forbidden. The user is authenticated (logged in), so it's not 401. They simply don't have permission to access this resource, which is exactly what 403 means.",
							[]string{"What's the difference between 401 and 403?", "Is the user authenticated?"},
						),
					},
				},
				{
					ID:    "http-4",
					Title: "Headers & Cookies",
					Content: []models.ContentBlock{
						textBlock(`## HTTP Headers

Headers are key-value pairs that provide metadata about requests and responses.

**Types**:
- **Request headers**: Sent by client
- **Response headers**: Sent by server
- **General headers**: Apply to both
- **Entity headers**: Describe the body`),
						textBlock(`## Common Request Headers

**Host**: 
- Domain name of server
- Required in HTTP/1.1

**User-Agent**:
- Client application info
- Browser, version, OS

**Accept**:
- What content types client accepts
- Accept: application/json

**Authorization**:
- Credentials for authentication
- Authorization: Bearer <token>

**Content-Type**:
- Type of data in body
- Content-Type: application/json

**Cookie**:
- Cookies to send
- Cookie: sessionId=abc123`),
						textBlock(`## Common Response Headers

**Content-Type**:
- Type of data in body
- Content-Type: application/json

**Set-Cookie**:
- Instructs client to set cookie
- Set-Cookie: sessionId=abc123; Path=/; HttpOnly

**Cache-Control**:
- Caching directives
- Cache-Control: no-cache

**Location**:
- Redirect URL
- Used with 3xx status codes

**ETag**:
- Entity tag for caching
- ETag: "abc123"`),
						textBlock(`## What are Cookies?

Cookies are small pieces of data stored by the browser and sent with requests.

**Uses**:
- Session management
- Personalization
- Tracking
- Authentication tokens`),
						codeBlock("javascript", `// Setting cookies (server-side example)

// Set-Cookie header
Set-Cookie: sessionId=abc123; Path=/; HttpOnly; Secure; SameSite=Strict

// Cookie attributes:
// - Path: Where cookie is valid
// - Domain: Which domain can access
// - HttpOnly: Not accessible via JavaScript
// - Secure: Only sent over HTTPS
// - SameSite: CSRF protection
// - Max-Age/Expires: When cookie expires`),
						textBlock(`## Cookie Security

**HttpOnly**:
- Prevents JavaScript access
- Protects against XSS attacks

**Secure**:
- Only sent over HTTPS
- Protects against man-in-the-middle

**SameSite**:
- Strict: Never sent cross-site
- Lax: Sent for top-level navigation
- None: Always sent (requires Secure)

**Domain**:
- Restricts which domains can access
- Don't set for other domains!`),
						textBlock(`## Alternatives to Cookies

**Local Storage**:
- Stored in browser
- Accessible via JavaScript
- Not sent automatically
- Larger capacity

**Session Storage**:
- Like Local Storage
- Cleared when tab closes

**JWT Tokens**:
- Stored in Local Storage or memory
- Sent in Authorization header
- Stateless authentication`),
						calloutBlock("tip", "Use HttpOnly cookies for sensitive data like session IDs. Use Local Storage for non-sensitive data that needs JavaScript access."),
						exerciseBlock(
							"Why should authentication tokens be stored in HttpOnly cookies rather than Local Storage?",
							"HttpOnly cookies are not accessible to JavaScript, which protects them from XSS (Cross-Site Scripting) attacks. If an attacker injects malicious JavaScript, they can't steal tokens from HttpOnly cookies. Local Storage is accessible to JavaScript, making it vulnerable to XSS attacks.",
							[]string{"Think about XSS attacks", "What can JavaScript access?"},
						),
					},
				},
				{
					ID:    "http-5",
					Title: "REST API Design",
					Content: []models.ContentBlock{
						textBlock(`## What is REST?

REST (Representational State Transfer) is an architectural style for designing web services.

**Key principles**:
- Stateless
- Client-server architecture
- Uniform interface
- Resource-based`),
						textBlock(`## RESTful Design

**Resources**:
- Everything is a resource
- Identified by URLs
- Nouns, not verbs

**Examples**:
- ✅ GET /api/users
- ✅ GET /api/users/123
- ✅ POST /api/users
- ❌ GET /api/getUser
- ❌ POST /api/createUser`),
						codeBlock("text", `RESTful URL Structure:

Collection:
GET    /api/users           # List all users
POST   /api/users           # Create user

Resource:
GET    /api/users/123       # Get user 123
PUT    /api/users/123       # Replace user 123
PATCH  /api/users/123       # Update user 123
DELETE /api/users/123       # Delete user 123

Sub-resources:
GET    /api/users/123/posts  # Get posts by user 123
POST   /api/users/123/posts  # Create post for user 123`),
						textBlock(`## REST Best Practices

**Use proper HTTP methods**:
- GET for reading
- POST for creating
- PUT/PATCH for updating
- DELETE for removing

**Use proper status codes**:
- 200 for success
- 201 for created
- 204 for no content
- 400 for client errors
- 500 for server errors

**Version your API**:
- /api/v1/users
- /api/v2/users`),
						textBlock(`## Request/Response Format

**Request**:
- Headers for metadata
- Body for data (JSON)
- Query params for filtering

**Response**:
- Consistent structure
- Error format
- Pagination metadata`),
						codeBlock("json", `// Example REST API Response

// Success response
{
  "data": {
    "id": 123,
    "name": "Alice",
    "email": "alice@example.com"
  }
}

// List response with pagination
{
  "data": [
    {"id": 1, "name": "Alice"},
    {"id": 2, "name": "Bob"}
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50
  }
}

// Error response
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Email is required",
    "details": {
      "email": "This field is required"
    }
  }
}`),
						textBlock(`## Filtering, Sorting, Pagination

**Filtering**:
GET /api/users?status=active&role=admin

**Sorting**:
GET /api/users?sort=name&order=asc

**Pagination**:
GET /api/users?page=1&limit=10

**Search**:
GET /api/users?search=alice`),
						calloutBlock("tip", "Keep URLs clean and intuitive. Use query parameters for optional things like filtering and pagination, not for required resource identification."),
						exerciseBlock(
							"Design RESTful endpoints for a blog API with posts and comments. Include endpoints for listing, creating, updating, and deleting.",
							"Posts:\n- GET /api/posts (list)\n- POST /api/posts (create)\n- GET /api/posts/123 (get one)\n- PUT /api/posts/123 (replace)\n- PATCH /api/posts/123 (update)\n- DELETE /api/posts/123 (delete)\n\nComments:\n- GET /api/posts/123/comments (list comments for post)\n- POST /api/posts/123/comments (create comment)\n- GET /api/comments/456 (get comment)\n- PATCH /api/comments/456 (update)\n- DELETE /api/comments/456 (delete)",
							[]string{"Think about resource hierarchy", "What are the nouns?"},
						),
					},
				},
				{
					ID:    "http-6",
					Title: "Request/Response Cycle",
					Content: []models.ContentBlock{
						textBlock(`## The HTTP Request/Response Cycle

Understanding the complete flow of an HTTP request:

1. **Client** initiates request
2. **DNS** resolves domain to IP
3. **TCP** connection established
4. **TLS** handshake (if HTTPS)
5. **HTTP** request sent
6. **Server** processes request
7. **HTTP** response sent
8. **Connection** may close or stay open`),
						textBlock(`## Step-by-Step Flow

**1. Client Initiates Request**:
- User clicks link or JavaScript makes request
- Browser/app prepares HTTP request

**2. DNS Lookup**:
- Domain name → IP address
- example.com → 93.184.216.34

**3. TCP Connection**:
- Three-way handshake
- SYN → SYN-ACK → ACK
- Connection established`),
						codeBlock("text", `TCP Three-Way Handshake:

Client                Server
  |                     |
  |-------- SYN -------->|
  |                     |
  |<----- SYN-ACK -------|
  |                     |
  |-------- ACK -------->|
  |                     |
  |   Connection Open    |`),
						textBlock(`## HTTPS (TLS Handshake)

If using HTTPS:

**TLS Handshake**:
1. Client sends supported cipher suites
2. Server chooses cipher and sends certificate
3. Client verifies certificate
4. Shared secret established
5. Encrypted communication begins`),
						textBlock(`## HTTP Request Sent

**Request includes**:
- Request line (method, path, version)
- Headers
- Body (if applicable)

**Server receives**:
- Parses request
- Routes to handler
- Processes request`),
						textBlock(`## Server Processing

**Server**:
1. Receives request
2. Routes to appropriate handler
3. Validates request
4. Executes business logic
5. Queries database (if needed)
6. Generates response
7. Sends response`),
						textBlock(`## HTTP Response Sent

**Response includes**:
- Status line (version, status code, message)
- Headers
- Body (if applicable)

**Client receives**:
- Parses response
- Handles based on status code
- Renders/processes data`),
						codeBlock("javascript", `// Complete request/response example

// Client side
async function fetchUser(userId) {
  try {
    const response = await fetch('/api/users/' + userId, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Authorization': 'Bearer token123'
      }
    });
    
    if (!response.ok) {
      throw new Error('HTTP error! status: ' + response.status);
    }
    
    const user = await response.json();
    return user;
  } catch (error) {
    console.error('Request failed:', error);
    throw error;
  }
}

// Server side (pseudo-code)
function handleGetUser(request) {
  // 1. Validate request
  const userId = request.params.id;
  
  // 2. Authenticate
  const user = authenticate(request.headers.authorization);
  
  // 3. Authorize
  if (!canAccess(user, userId)) {
    return { status: 403, body: { error: 'Forbidden' } };
  }
  
  // 4. Query database
  const userData = database.getUser(userId);
  
  // 5. Generate response
  return {
    status: 200,
    headers: { 'Content-Type': 'application/json' },
    body: userData
  };
}`),
						calloutBlock("info", "Understanding the request/response cycle helps debug issues. Is it a DNS problem? Network issue? Server error? Client error?"),
						exerciseBlock(
							"What happens if a server takes 30 seconds to process a request, but the client times out after 10 seconds?",
							"The client will close the connection after 10 seconds, even though the server is still processing. The server may complete processing, but the response won't reach the client. This is why it's important to:\n1. Set appropriate timeout values\n2. Optimize server processing time\n3. Use async processing for long operations\n4. Provide progress updates for long-running tasks",
							[]string{"Think about what happens to the connection", "Can the server still send a response?"},
						),
					},
				},
				{
					ID:    "http-7",
					Title: "HTTPS & Security",
					Content: []models.ContentBlock{
						textBlock(`## What is HTTPS?

HTTPS (HTTP Secure) is HTTP over TLS/SSL encryption.

**Benefits**:
- **Encryption**: Data is encrypted in transit
- **Authentication**: Verifies server identity
- **Integrity**: Detects tampering
- **Privacy**: Prevents eavesdropping`),
						textBlock(`## Why HTTPS Matters

**Without HTTPS**:
- Data sent in plain text
- Anyone can intercept
- Man-in-the-middle attacks possible
- No server verification

**With HTTPS**:
- Data encrypted
- Only server can decrypt
- Certificate verifies identity
- Secure communication`),
						calloutBlock("warning", "Never send sensitive data (passwords, credit cards, personal info) over HTTP. Always use HTTPS!"),
						textBlock(`## TLS/SSL Handshake

**Process**:
1. Client initiates connection
2. Server sends certificate
3. Client verifies certificate
4. Shared secret established
5. Encrypted communication

**Certificate contains**:
- Domain name
- Public key
- Issuer (Certificate Authority)
- Expiration date`),
						textBlock(`## Certificate Authorities (CAs)

**Role**:
- Issue digital certificates
- Verify domain ownership
- Trusted by browsers

**Examples**:
- Let's Encrypt (free)
- DigiCert
- GlobalSign

**Self-signed certificates**:
- Not trusted by browsers
- OK for development
- Not for production`),
						textBlock(`## Common Security Headers

**Strict-Transport-Security (HSTS)**:
- Forces HTTPS
- Prevents downgrade attacks
- Strict-Transport-Security: max-age=31536000

**Content-Security-Policy (CSP)**:
- Prevents XSS attacks
- Controls resource loading
- Content-Security-Policy: default-src 'self'

**X-Frame-Options**:
- Prevents clickjacking
- X-Frame-Options: DENY

**X-Content-Type-Options**:
- Prevents MIME sniffing
- X-Content-Type-Options: nosniff`),
						textBlock(`## Common Vulnerabilities

**XSS (Cross-Site Scripting)**:
- Malicious scripts injected
- Prevent with: Input validation, output encoding, CSP

**CSRF (Cross-Site Request Forgery)**:
- Unauthorized actions from user's browser
- Prevent with: CSRF tokens, SameSite cookies

**SQL Injection**:
- Malicious SQL in inputs
- Prevent with: Parameterized queries, ORMs

**Man-in-the-Middle**:
- Intercepting communication
- Prevent with: HTTPS, certificate pinning`),
						codeBlock("javascript", `// Security best practices

// 1. Always use HTTPS in production
const API_URL = process.env.NODE_ENV === 'production' 
  ? 'https://api.example.com'
  : 'http://localhost:3000';

// 2. Validate and sanitize input
function sanitizeInput(input) {
  return input
    .replace(/<script>/gi, '')
    .trim()
    .escape();
}

// 3. Use parameterized queries (never concatenate!)
// ❌ BAD
const query = 'SELECT * FROM users WHERE id = ' + userId;

// ✅ GOOD
const query = 'SELECT * FROM users WHERE id = ?';
db.query(query, [userId]);

// 4. Set secure headers
app.use((req, res, next) => {
  res.setHeader('Strict-Transport-Security', 'max-age=31536000');
  res.setHeader('X-Content-Type-Options', 'nosniff');
  res.setHeader('X-Frame-Options', 'DENY');
  next();
});`),
						calloutBlock("tip", "Security is not optional. Always use HTTPS in production, validate all inputs, use parameterized queries, and set security headers."),
						exerciseBlock(
							"What's the difference between HTTP and HTTPS, and why should you always use HTTPS?",
							"HTTP sends data in plain text, while HTTPS encrypts data using TLS/SSL. HTTPS provides:\n1. Encryption - data can't be read if intercepted\n2. Authentication - verifies you're talking to the real server\n3. Integrity - detects if data was tampered with\n\nYou should always use HTTPS because:\n- Protects user data (passwords, personal info)\n- Prevents man-in-the-middle attacks\n- Required for modern web features\n- Builds user trust\n- Better SEO rankings",
							[]string{"Think about what happens to data in transit", "What can attackers do with unencrypted data?"},
						),
					},
				},
			},
		},
		{
			Title:       "Design Patterns",
			Description: "Learn proven software design patterns to solve common programming challenges elegantly and efficiently.",
			Modules: []models.Module{
				{
					ID:    "patterns-1",
					Title: "Introduction to Design Patterns",
					Content: []models.ContentBlock{
						textBlock(`## What are Design Patterns?

Design patterns are reusable solutions to common problems in software design. They're templates for solving problems that occur repeatedly.

**Key points**:
- Not code you can copy-paste
- General solutions to general problems
- Language-agnostic concepts
- Proven approaches used by experts`),
						calloutBlock("info", "Design patterns were popularized by the 'Gang of Four' (GoF) book 'Design Patterns: Elements of Reusable Object-Oriented Software' published in 1994."),
						textBlock(`## Why Use Design Patterns?

**Benefits**:
- **Proven solutions**: Tested approaches to common problems
- **Communication**: Shared vocabulary for developers
- **Best practices**: Encapsulate expert knowledge
- **Flexibility**: Make code more maintainable and extensible

**When to use**:
- Solving recurring problems
- Improving code structure
- Making code more maintainable
- Learning from expert solutions`),
						textBlock(`## Types of Design Patterns

**Creational Patterns**:
- Deal with object creation
- Examples: Singleton, Factory, Builder

**Structural Patterns**:
- Deal with object composition
- Examples: Adapter, Decorator, Facade

**Behavioral Patterns**:
- Deal with object interaction
- Examples: Observer, Strategy, Command`),
						calloutBlock("warning", "Don't force patterns into your code. Use them when they solve real problems. Over-engineering with patterns can make code harder to understand."),
						textBlock(`## Pattern Structure

Most patterns describe:

1. **Problem**: What problem does it solve?
2. **Solution**: How does it solve it?
3. **Consequences**: Trade-offs and benefits
4. **Implementation**: How to code it

**Pattern template**:
- Intent
- Motivation
- Applicability
- Structure
- Participants
- Collaborations
- Consequences`),
						exerciseBlock(
							"What is a design pattern, and how is it different from a library or framework?",
							"A design pattern is a general, reusable solution to a commonly occurring problem in software design. It's a template or blueprint, not code you can directly use.\n\nA library is code you can call, and a framework is code that calls you. Design patterns are conceptual solutions that you implement yourself, while libraries and frameworks are actual code you use.\n\nPatterns are language-agnostic concepts, while libraries and frameworks are specific implementations.",
							[]string{"Think about the difference between concepts and code", "What can you directly use vs. what you implement?"},
						),
					},
				},
				{
					ID:    "patterns-2",
					Title: "Creational Patterns (Singleton, Factory)",
					Content: []models.ContentBlock{
						textBlock(`## Creational Patterns

Creational patterns deal with object creation mechanisms, trying to create objects in a manner suitable to the situation.

**Common patterns**:
- Singleton
- Factory
- Builder
- Prototype
- Abstract Factory`),
						textBlock(`## Singleton Pattern

**Intent**: Ensure a class has only one instance and provide global access to it.

**Use when**:
- Exactly one instance is needed
- Instance must be accessible globally
- Lazy initialization is desired

**Example uses**:
- Database connections
- Logging
- Configuration managers`),
						codeBlock("javascript", `// Singleton Pattern

class DatabaseConnection {
  constructor() {
    if (DatabaseConnection.instance) {
      return DatabaseConnection.instance;
    }
    
    this.connection = this.connect();
    DatabaseConnection.instance = this;
  }
  
  connect() {
    // Connection logic
    return { status: 'connected' };
  }
  
  query(sql) {
    // Query logic
    return 'results';
  }
}

// Usage
const db1 = new DatabaseConnection();
const db2 = new DatabaseConnection();

console.log(db1 === db2); // true - same instance`),
						calloutBlock("warning", "Singleton can make testing difficult and create hidden dependencies. Use sparingly and consider dependency injection instead."),
						textBlock(`## Factory Pattern

**Intent**: Create objects without specifying the exact class of object that will be created.

**Use when**:
- Don't know exact types at compile time
- Want to decouple object creation from usage
- Need to create families of related objects

**Example uses**:
- UI component creation
- Database driver selection
- Payment method processing`),
						codeBlock("javascript", `// Factory Pattern

class PaymentMethod {
  processPayment(amount) {
    throw new Error('Must implement');
  }
}

class CreditCard extends PaymentMethod {
  processPayment(amount) {
    return 'Processing ' + amount + ' via credit card';
  }
}

class PayPal extends PaymentMethod {
  processPayment(amount) {
    return 'Processing ' + amount + ' via PayPal';
  }
}

class PaymentFactory {
  static create(type) {
    switch(type) {
      case 'creditcard':
        return new CreditCard();
      case 'paypal':
        return new PayPal();
      default:
        throw new Error('Unknown payment type');
    }
  }
}

// Usage
const payment = PaymentFactory.create('creditcard');
payment.processPayment(100);`),
						textBlock(`## Builder Pattern

**Intent**: Construct complex objects step by step.

**Use when**:
- Object has many optional parameters
- Want to avoid telescoping constructor
- Need different representations

**Example uses**:
- SQL query builders
- HTTP request builders
- Configuration objects`),
						codeBlock("javascript", `// Builder Pattern

class QueryBuilder {
  constructor() {
    this.query = {
      select: [],
      from: null,
      where: [],
      orderBy: null
    };
  }
  
  select(fields) {
    this.query.select = fields;
    return this;
  }
  
  from(table) {
    this.query.from = table;
    return this;
  }
  
  where(condition) {
    this.query.where.push(condition);
    return this;
  }
  
  orderBy(field) {
    this.query.orderBy = field;
    return this;
  }
  
  build() {
    return this.query;
  }
}

// Usage
const query = new QueryBuilder()
  .select(['name', 'email'])
  .from('users')
  .where('age > 18')
  .orderBy('name')
  .build();`),
						calloutBlock("tip", "Factory is great when you don't know the exact type. Builder is great when you have many optional parameters."),
						exerciseBlock(
							"When would you use Singleton vs Factory pattern?",
							"Use Singleton when you need exactly one instance of a class that should be globally accessible (like a database connection or logger).\n\nUse Factory when you need to create objects but don't know the exact type at compile time, or when you want to decouple object creation from usage (like creating different payment methods based on user choice).\n\nThey solve different problems: Singleton ensures one instance, Factory handles object creation logic.",
							[]string{"What problem does each solve?", "Think about when you'd need one instance vs. multiple types"},
						),
					},
				},
				{
					ID:    "patterns-3",
					Title: "Structural Patterns (Adapter, Decorator)",
					Content: []models.ContentBlock{
						textBlock(`## Structural Patterns

Structural patterns deal with object composition - how classes and objects are composed to form larger structures.

**Common patterns**:
- Adapter
- Decorator
- Facade
- Proxy
- Bridge`),
						textBlock(`## Adapter Pattern

**Intent**: Allow incompatible interfaces to work together.

**Use when**:
- Want to use existing class with incompatible interface
- Need to integrate third-party libraries
- Want to create reusable class for unrelated classes

**Example uses**:
- Integrating legacy code
- Wrapping third-party APIs
- Making incompatible interfaces compatible`),
						codeBlock("javascript", `// Adapter Pattern

// Old interface (incompatible)
class OldPaymentSystem {
  makePayment(amount, currency) {
    return 'Paid ' + amount + ' ' + currency;
  }
}

// New interface (what we want)
class PaymentProcessor {
  pay(amount) {
    throw new Error('Must implement');
  }
}

// Adapter
class PaymentAdapter extends PaymentProcessor {
  constructor(oldSystem) {
    super();
    this.oldSystem = oldSystem;
  }
  
  pay(amount) {
    // Adapt old interface to new
    return this.oldSystem.makePayment(amount, 'USD');
  }
}

// Usage
const oldSystem = new OldPaymentSystem();
const adapter = new PaymentAdapter(oldSystem);
adapter.pay(100); // Works with new interface`),
						calloutBlock("info", "Adapter is like a translator - it makes two incompatible interfaces work together without changing either one."),
						textBlock(`## Decorator Pattern

**Intent**: Attach additional responsibilities to objects dynamically.

**Use when**:
- Want to add behavior to objects at runtime
- Don't want to modify existing code
- Need to add features flexibly

**Example uses**:
- Adding logging to methods
- Adding caching to functions
- Adding validation to inputs
- UI component enhancements`),
						codeBlock("javascript", `// Decorator Pattern

class Coffee {
  cost() {
    return 5;
  }
  
  description() {
    return 'Coffee';
  }
}

// Decorators
class MilkDecorator {
  constructor(coffee) {
    this.coffee = coffee;
  }
  
  cost() {
    return this.coffee.cost() + 2;
  }
  
  description() {
    return this.coffee.description() + ', Milk';
  }
}

class SugarDecorator {
  constructor(coffee) {
    this.coffee = coffee;
  }
  
  cost() {
    return this.coffee.cost() + 1;
  }
  
  description() {
    return this.coffee.description() + ', Sugar';
  }
}

// Usage
let coffee = new Coffee();
coffee = new MilkDecorator(coffee);
coffee = new SugarDecorator(coffee);

console.log(coffee.description()); // "Coffee, Milk, Sugar"
console.log(coffee.cost()); // 8`),
						textBlock(`## Facade Pattern

**Intent**: Provide a simplified interface to a complex subsystem.

**Use when**:
- Want to hide complexity
- Need simple interface to complex system
- Want to decouple clients from subsystem

**Example uses**:
- API wrappers
- Library interfaces
- System abstractions`),
						codeBlock("javascript", `// Facade Pattern

// Complex subsystem
class CPU {
  freeze() { return 'CPU frozen'; }
  jump(position) { return 'Jump to ' + position; }
  execute() { return 'CPU executing'; }
}

class Memory {
  load(position, data) { return 'Load ' + data + ' at ' + position; }
}

class HardDrive {
  read(lba, size) { return 'Read ' + size + ' bytes from ' + lba; }
}

// Facade
class Computer {
  constructor() {
    this.cpu = new CPU();
    this.memory = new Memory();
    this.hardDrive = new HardDrive();
  }
  
  start() {
    // Simple interface to complex system
    this.cpu.freeze();
    this.memory.load('0x0', this.hardDrive.read('0', '1024'));
    this.cpu.jump('0x0');
    this.cpu.execute();
    return 'Computer started';
  }
}

// Usage - simple interface
const computer = new Computer();
computer.start();`),
						calloutBlock("tip", "Facade simplifies complex systems. Adapter makes incompatible interfaces work together. Decorator adds behavior dynamically."),
						exerciseBlock(
							"What's the difference between Adapter and Facade patterns?",
							"Adapter makes incompatible interfaces work together - it translates between two different interfaces so they can collaborate.\n\nFacade provides a simplified interface to a complex subsystem - it hides complexity behind a simple interface.\n\nAdapter is about compatibility (making A work with B), while Facade is about simplicity (hiding complexity behind a simple interface).",
							[]string{"What problem does each solve?", "Think about compatibility vs. simplicity"},
						),
					},
				},
				{
					ID:    "patterns-4",
					Title: "Behavioral Patterns (Observer, Strategy)",
					Content: []models.ContentBlock{
						textBlock(`## Behavioral Patterns

Behavioral patterns deal with object interaction - how objects communicate and distribute responsibility.

**Common patterns**:
- Observer
- Strategy
- Command
- Chain of Responsibility
- State`),
						textBlock(`## Observer Pattern

**Intent**: Define a one-to-many dependency between objects so when one changes, all dependents are notified.

**Use when**:
- Change to one object requires changing others
- Don't know how many objects need to be updated
- Want loose coupling between objects

**Example uses**:
- Event systems
- Model-View architecture
- Publish-subscribe systems`),
						codeBlock("javascript", `// Observer Pattern

class Subject {
  constructor() {
    this.observers = [];
  }
  
  subscribe(observer) {
    this.observers.push(observer);
  }
  
  unsubscribe(observer) {
    this.observers = this.observers.filter(o => o !== observer);
  }
  
  notify(data) {
    this.observers.forEach(observer => observer.update(data));
  }
}

class Observer {
  update(data) {
    throw new Error('Must implement');
  }
}

class EmailObserver extends Observer {
  update(data) {
    console.log('Sending email: ' + data);
  }
}

class SMSObserver extends Observer {
  update(data) {
    console.log('Sending SMS: ' + data);
  }
}

// Usage
const subject = new Subject();
subject.subscribe(new EmailObserver());
subject.subscribe(new SMSObserver());

subject.notify('User logged in'); // Both observers notified`),
						calloutBlock("info", "Observer pattern is the foundation of many event systems. React's state management uses similar concepts."),
						textBlock(`## Strategy Pattern

**Intent**: Define a family of algorithms, encapsulate each one, and make them interchangeable.

**Use when**:
- Have multiple ways to perform a task
- Want to switch algorithms at runtime
- Want to avoid conditional statements

**Example uses**:
- Sorting algorithms
- Payment methods
- Compression algorithms
- Validation strategies`),
						codeBlock("javascript", `// Strategy Pattern

class Sorter {
  constructor(strategy) {
    this.strategy = strategy;
  }
  
  setStrategy(strategy) {
    this.strategy = strategy;
  }
  
  sort(data) {
    return this.strategy.sort(data);
  }
}

class QuickSort {
  sort(data) {
    return [...data].sort((a, b) => a - b);
  }
}

class BubbleSort {
  sort(data) {
    const arr = [...data];
    for (let i = 0; i < arr.length; i++) {
      for (let j = 0; j < arr.length - i - 1; j++) {
        if (arr[j] > arr[j + 1]) {
          [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        }
      }
    }
    return arr;
  }
}

// Usage
const sorter = new Sorter(new QuickSort());
sorter.sort([3, 1, 4, 2]); // Uses QuickSort

sorter.setStrategy(new BubbleSort());
sorter.sort([3, 1, 4, 2]); // Uses BubbleSort`),
						textBlock(`## Command Pattern

**Intent**: Encapsulate a request as an object, allowing parameterization and queuing of requests.

**Use when**:
- Need to parameterize objects with operations
- Need to queue operations
- Need to support undo/redo

**Example uses**:
- Undo/redo functionality
- Macro recording
- Job queues
- Transaction systems`),
						codeBlock("javascript", `// Command Pattern

class Command {
  execute() {
    throw new Error('Must implement');
  }
  
  undo() {
    throw new Error('Must implement');
  }
}

class Light {
  on() { return 'Light is ON'; }
  off() { return 'Light is OFF'; }
}

class LightOnCommand extends Command {
  constructor(light) {
    super();
    this.light = light;
  }
  
  execute() {
    return this.light.on();
  }
  
  undo() {
    return this.light.off();
  }
}

class Invoker {
  constructor() {
    this.history = [];
  }
  
  execute(command) {
    this.history.push(command);
    return command.execute();
  }
  
  undo() {
    if (this.history.length > 0) {
      const command = this.history.pop();
      return command.undo();
    }
  }
}

// Usage
const light = new Light();
const command = new LightOnCommand(light);
const invoker = new Invoker();

invoker.execute(command); // Light on
invoker.undo(); // Light off`),
						calloutBlock("tip", "Observer is about notifications. Strategy is about interchangeable algorithms. Command is about encapsulating operations."),
						exerciseBlock(
							"When would you use Observer vs Strategy pattern?",
							"Use Observer when you need one object to notify multiple other objects about changes (like an event system, where one event triggers multiple handlers).\n\nUse Strategy when you have multiple ways to accomplish the same task and want to switch between them (like different sorting algorithms, payment methods, or validation rules).\n\nObserver is about notifications and dependencies. Strategy is about interchangeable algorithms.",
							[]string{"What problem does each solve?", "Think about notifications vs. algorithms"},
						),
					},
				},
				{
					ID:    "patterns-5",
					Title: "MVC Pattern",
					Content: []models.ContentBlock{
						textBlock(`## Model-View-Controller (MVC)

MVC is an architectural pattern that separates an application into three interconnected components.

**Components**:
- **Model**: Data and business logic
- **View**: User interface
- **Controller**: Handles input and coordinates Model and View`),
						textBlock(`## The Three Components

**Model**:
- Represents data and business logic
- Independent of UI
- Notifies observers of changes
- Handles data validation

**View**:
- Displays data to user
- Observes Model for changes
- Sends user input to Controller
- No business logic

**Controller**:
- Handles user input
- Updates Model
- Selects View to display
- Coordinates Model and View`),
						codeBlock("javascript", `// MVC Pattern Example

// Model
class UserModel {
  constructor() {
    this.users = [];
    this.observers = [];
  }
  
  subscribe(observer) {
    this.observers.push(observer);
  }
  
  notify() {
    this.observers.forEach(obs => obs.update(this.users));
  }
  
  addUser(user) {
    this.users.push(user);
    this.notify();
  }
  
  getUsers() {
    return this.users;
  }
}

// View
class UserView {
  constructor(controller) {
    this.controller = controller;
  }
  
  render(users) {
    console.log('Users:', users);
    // Render UI
  }
  
  update(users) {
    this.render(users);
  }
  
  handleAddUser(name) {
    this.controller.addUser(name);
  }
}

// Controller
class UserController {
  constructor(model) {
    this.model = model;
  }
  
  addUser(name) {
    this.model.addUser({ name, id: Date.now() });
  }
  
  setView(view) {
    this.view = view;
    this.model.subscribe(view);
  }
}

// Usage
const model = new UserModel();
const controller = new UserController(model);
const view = new UserView(controller);
controller.setView(view);

view.handleAddUser('Alice');`),
						textBlock(`## MVC Flow

1. **User interacts with View** (clicks button, fills form)
2. **View sends input to Controller**
3. **Controller updates Model**
4. **Model notifies observers** (including View)
5. **View updates display** based on new Model state`),
						calloutBlock("info", "MVC is one of the most widely used patterns. Many frameworks implement it: Ruby on Rails, Django, Angular, etc."),
						textBlock(`## Benefits of MVC

**Separation of Concerns**:
- Each component has clear responsibility
- Changes to one don't affect others

**Reusability**:
- Models can be used with different Views
- Views can display different Models

**Testability**:
- Each component can be tested independently
- Business logic separated from UI

**Maintainability**:
- Clear structure
- Easy to understand and modify`),
						textBlock(`## Variations

**MVP (Model-View-Presenter)**:
- Presenter handles all UI logic
- View is passive

**MVVM (Model-View-ViewModel)**:
- ViewModel binds View and Model
- Two-way data binding

**MVC vs MVVM**:
- MVC: Controller handles input
- MVVM: ViewModel provides data binding`),
						exerciseBlock(
							"In MVC, where should validation logic go - Model, View, or Controller?",
							"Validation logic should primarily go in the Model, because:\n1. Business rules belong in the Model\n2. Validation should be reusable across different Views/Controllers\n3. Ensures data integrity regardless of how it's entered\n\nHowever, the View can have basic input validation (format checking) for better UX, and the Controller can coordinate validation, but the Model should have the authoritative validation rules.",
							[]string{"Where does business logic belong?", "What needs to be reusable?"},
						),
					},
				},
				{
					ID:    "patterns-6",
					Title: "Repository Pattern",
					Content: []models.ContentBlock{
						textBlock(`## Repository Pattern

The Repository pattern provides an abstraction layer between business logic and data access.

**Purpose**:
- Encapsulate data access logic
- Provide collection-like interface
- Decouple business logic from data source`),
						textBlock(`## Why Use Repository?

**Benefits**:
- **Abstraction**: Business logic doesn't know about database
- **Testability**: Easy to mock data access
- **Flexibility**: Can switch data sources easily
- **Separation**: Clear boundary between layers

**Use when**:
- Need to abstract data access
- Want to test business logic independently
- May need to change data source
- Want clean separation of concerns`),
						codeBlock("javascript", `// Repository Pattern

// Interface
class UserRepository {
  findById(id) {
    throw new Error('Must implement');
  }
  
  findAll() {
    throw new Error('Must implement');
  }
  
  save(user) {
    throw new Error('Must implement');
  }
  
  delete(id) {
    throw new Error('Must implement');
  }
}

// Implementation
class MySQLUserRepository extends UserRepository {
  constructor(db) {
    super();
    this.db = db;
  }
  
  findById(id) {
    return this.db.query('SELECT * FROM users WHERE id = ?', [id]);
  }
  
  findAll() {
    return this.db.query('SELECT * FROM users');
  }
  
  save(user) {
    if (user.id) {
      return this.db.query('UPDATE users SET ? WHERE id = ?', [user, user.id]);
    } else {
      return this.db.query('INSERT INTO users SET ?', [user]);
    }
  }
  
  delete(id) {
    return this.db.query('DELETE FROM users WHERE id = ?', [id]);
  }
}

// In-memory implementation for testing
class InMemoryUserRepository extends UserRepository {
  constructor() {
    super();
    this.users = [];
  }
  
  findById(id) {
    return Promise.resolve(this.users.find(u => u.id === id));
  }
  
  findAll() {
    return Promise.resolve(this.users);
  }
  
  save(user) {
    if (user.id) {
      const index = this.users.findIndex(u => u.id === user.id);
      this.users[index] = user;
    } else {
      user.id = Date.now();
      this.users.push(user);
    }
    return Promise.resolve(user);
  }
  
  delete(id) {
    this.users = this.users.filter(u => u.id !== id);
    return Promise.resolve();
  }
}`),
						textBlock(`## Using Repository

**Business Logic**:
- Uses repository interface
- Doesn't know about database
- Easy to test with mock repository`),
						codeBlock("javascript", `// Business logic using repository

class UserService {
  constructor(userRepository) {
    this.repository = userRepository;
  }
  
  async getUser(id) {
    return await this.repository.findById(id);
  }
  
  async getAllUsers() {
    return await this.repository.findAll();
  }
  
  async createUser(userData) {
    // Business logic
    if (!userData.email) {
      throw new Error('Email required');
    }
    
    return await this.repository.save(userData);
  }
}

// Usage with real database
const mysqlRepo = new MySQLUserRepository(db);
const userService = new UserService(mysqlRepo);

// Usage in tests with mock
const mockRepo = new InMemoryUserRepository();
const testService = new UserService(mockRepo);`),
						calloutBlock("tip", "Repository pattern is essential for testability. You can easily swap real database with in-memory implementation for testing."),
						textBlock(`## Repository vs DAO

**Repository**:
- Collection-like interface
- Domain-centric
- Business logic aware
- Can aggregate multiple data sources

**DAO (Data Access Object)**:
- Table-centric
- Closer to database
- One DAO per table
- Less business logic`),
						exerciseBlock(
							"What are the main benefits of the Repository pattern, and when would you use it?",
							"Main benefits:\n1. Abstraction - business logic doesn't depend on data source\n2. Testability - easy to mock for unit tests\n3. Flexibility - can switch databases without changing business logic\n4. Separation - clear boundary between data access and business logic\n\nUse it when:\n- You need to abstract data access\n- You want to test business logic independently\n- You might need to change data sources\n- You want clean architecture with clear layers",
							[]string{"Think about testability", "What problems does it solve?"},
						),
					},
				},
				{
					ID:    "patterns-7",
					Title: "Choosing the Right Pattern",
					Content: []models.ContentBlock{
						textBlock(`## Pattern Selection Guide

Choosing the right pattern depends on your specific problem. Here's a guide to help you decide.`),
						textBlock(`## Creational Patterns

**Use Singleton when**:
- Need exactly one instance
- Global access required
- Lazy initialization needed
- ⚠️ Use sparingly - can make testing difficult

**Use Factory when**:
- Don't know exact type at compile time
- Want to decouple creation from usage
- Need to create families of objects

**Use Builder when**:
- Object has many optional parameters
- Want to avoid telescoping constructors
- Need step-by-step construction`),
						textBlock(`## Structural Patterns

**Use Adapter when**:
- Need to make incompatible interfaces work together
- Integrating legacy code
- Wrapping third-party libraries

**Use Decorator when**:
- Need to add behavior at runtime
- Don't want to modify existing code
- Need flexible feature addition

**Use Facade when**:
- Want to simplify complex subsystem
- Need simple interface to complex system
- Want to decouple clients from subsystem`),
						textBlock(`## Behavioral Patterns

**Use Observer when**:
- One object change affects many
- Don't know how many objects need updates
- Want loose coupling

**Use Strategy when**:
- Have multiple ways to do something
- Want to switch algorithms at runtime
- Want to avoid conditional statements

**Use Command when**:
- Need to parameterize with operations
- Need undo/redo functionality
- Need to queue operations`),
						codeBlock("text", `Decision Tree:

Need to create objects?
  ├─ Need exactly one instance? → Singleton
  ├─ Don't know type? → Factory
  └─ Many optional params? → Builder

Need to structure objects?
  ├─ Incompatible interfaces? → Adapter
  ├─ Add behavior dynamically? → Decorator
  └─ Simplify complex system? → Facade

Need to handle interactions?
  ├─ One-to-many notifications? → Observer
  ├─ Interchangeable algorithms? → Strategy
  └─ Encapsulate operations? → Command`),
						textBlock(`## Anti-Patterns

**Avoid these mistakes**:

❌ **Using patterns everywhere**: Not every problem needs a pattern
❌ **Forcing patterns**: Don't force a pattern if it doesn't fit
❌ **Over-engineering**: Simple code is better than complex patterns
❌ **Pattern obsession**: Patterns are tools, not goals

**Remember**:
- Patterns solve specific problems
- Start simple, refactor when needed
- Understand the problem before choosing a pattern
- Patterns should make code clearer, not more complex`),
						calloutBlock("warning", "The best pattern is often no pattern. Don't add complexity unless it solves a real problem. Start simple and refactor when patterns become necessary."),
						textBlock(`## When NOT to Use Patterns

**Skip patterns when**:
- Problem is simple and straightforward
- Pattern adds unnecessary complexity
- Team doesn't understand the pattern
- Premature optimization
- Over-engineering for future needs that may never come`),
						exerciseBlock(
							"You have a simple CRUD application with one database. Should you use the Repository pattern?",
							"It depends on your context:\n\nUse Repository if:\n- You plan to add more data sources\n- You want to test business logic easily\n- You're building a larger application\n- You want clean architecture\n\nSkip Repository if:\n- It's a simple, one-off project\n- You won't need to test business logic separately\n- You won't change data sources\n- It adds unnecessary complexity\n\nGenerally, for a simple CRUD app, you might skip it initially and add it later if needed. But if you're building something that will grow, starting with Repository is a good idea.",
							[]string{"Think about future needs", "What's the cost vs. benefit?"},
						),
					},
				},
			},
		},
		{
			Title:       "Testing",
			Description: "Master software testing strategies including unit tests, integration tests, and test-driven development (TDD).",
			Modules: []models.Module{
				{
					ID:    "testing-1",
					Title: "Testing Fundamentals",
					Content: []models.ContentBlock{
						textBlock(`## Why Test Software?

Testing ensures your code works correctly and helps prevent bugs from reaching production.

**Benefits**:
- **Confidence**: Know your code works
- **Documentation**: Tests show how code should be used
- **Refactoring**: Safe to change code with tests
- **Debugging**: Tests help identify problems quickly
- **Quality**: Forces you to write better code`),
						calloutBlock("info", "Good tests are like a safety net - they give you confidence to make changes and catch problems before users do."),
						textBlock(`## Testing Pyramid

The testing pyramid shows the ideal distribution of tests:

**Unit Tests** (bottom, most):
- Test individual functions/classes
- Fast, isolated, many
- Foundation of testing

**Integration Tests** (middle):
- Test components working together
- Moderate speed, moderate number
- Verify interactions

**E2E Tests** (top, fewest):
- Test complete user flows
- Slow, fewer tests
- Verify system works end-to-end`),
						codeBlock("text", `Testing Pyramid:

        /\
       /  \     E2E Tests (few)
      /____\
     /      \    Integration Tests (some)
    /________\
   /          \  Unit Tests (many)
  /____________\`),
						textBlock(`## Types of Testing

**Unit Testing**:
- Test individual units in isolation
- Fast and focused
- Most common type

**Integration Testing**:
- Test multiple components together
- Verify they work correctly together

**End-to-End Testing**:
- Test complete user workflows
- Simulate real user scenarios

**Manual Testing**:
- Human tester performs actions
- Useful for UX and exploratory testing`),
						textBlock(`## Test Characteristics

**Good tests are**:
- **Fast**: Run quickly
- **Independent**: Don't depend on other tests
- **Repeatable**: Same results every time
- **Self-validating**: Pass or fail automatically
- **Timely**: Written before or with code`),
						calloutBlock("tip", "Follow the AAA pattern: Arrange (set up), Act (execute), Assert (verify). This makes tests clear and readable."),
						exerciseBlock(
							"Why is the testing pyramid shape (many unit tests, fewer integration tests, fewest E2E tests) recommended?",
							"The pyramid shape is recommended because:\n1. Unit tests are fast and cheap - you can have many\n2. Integration tests are slower - have fewer\n3. E2E tests are slowest and most brittle - have fewest\n\nThis maximizes test coverage and confidence while minimizing execution time. Having too many slow tests makes the test suite slow, which means developers run tests less often.",
							[]string{"Think about speed and cost", "What happens if you have too many slow tests?"},
						),
					},
				},
				{
					ID:    "testing-2",
					Title: "Unit Testing",
					Content: []models.ContentBlock{
						textBlock(`## What is Unit Testing?

Unit testing tests individual units of code (functions, methods, classes) in isolation.

**Characteristics**:
- Tests one thing at a time
- Fast execution
- No external dependencies
- Isolated from other code`),
						textBlock(`## Writing Unit Tests

**Structure** (AAA Pattern):

1. **Arrange**: Set up test data and conditions
2. **Act**: Execute the code being tested
3. **Assert**: Verify the results`),
						codeBlock("javascript", `// Example: Unit test

function add(a, b) {
  return a + b;
}

// Test
describe('add function', () => {
  test('adds two positive numbers', () => {
    // Arrange
    const a = 2;
    const b = 3;
    
    // Act
    const result = add(a, b);
    
    // Assert
    expect(result).toBe(5);
  });
  
  test('adds negative numbers', () => {
    // Arrange
    const a = -2;
    const b = -3;
    
    // Act
    const result = add(a, b);
    
    // Assert
    expect(result).toBe(-5);
  });
});`),
						textBlock(`## Test Cases

**Good test cases cover**:
- **Happy path**: Normal, expected behavior
- **Edge cases**: Boundary conditions
- **Error cases**: Invalid inputs
- **Null/empty**: Missing or empty data`),
						codeBlock("javascript", `// Comprehensive unit test

function divide(a, b) {
  if (b === 0) {
    throw new Error('Division by zero');
  }
  return a / b;
}

describe('divide function', () => {
  test('divides positive numbers', () => {
    expect(divide(10, 2)).toBe(5);
  });
  
  test('divides negative numbers', () => {
    expect(divide(-10, -2)).toBe(5);
  });
  
  test('handles decimal results', () => {
    expect(divide(1, 3)).toBeCloseTo(0.333, 3);
  });
  
  test('throws error on division by zero', () => {
    expect(() => divide(10, 0)).toThrow('Division by zero');
  });
  
  test('handles zero numerator', () => {
    expect(divide(0, 5)).toBe(0);
  });
});`),
						textBlock(`## Mocking and Stubbing

**Mocks**: Replace dependencies with fake implementations

**Why mock?**:
- Isolate unit under test
- Control dependencies
- Speed up tests
- Test error conditions`),
						codeBlock("javascript", `// Example with mocking

class UserService {
  constructor(userRepository) {
    this.repository = userRepository;
  }
  
  async getUser(id) {
    if (!id) {
      throw new Error('ID required');
    }
    return await this.repository.findById(id);
  }
}

// Test with mock
describe('UserService', () => {
  test('gets user by id', async () => {
    // Arrange - create mock
    const mockRepo = {
      findById: jest.fn().mockResolvedValue({ id: 1, name: 'Alice' })
    };
    const service = new UserService(mockRepo);
    
    // Act
    const user = await service.getUser(1);
    
    // Assert
    expect(user).toEqual({ id: 1, name: 'Alice' });
    expect(mockRepo.findById).toHaveBeenCalledWith(1);
  });
  
  test('throws error when id is missing', async () => {
    const mockRepo = { findById: jest.fn() };
    const service = new UserService(mockRepo);
    
    await expect(service.getUser(null)).rejects.toThrow('ID required');
    expect(mockRepo.findById).not.toHaveBeenCalled();
  });
});`),
						calloutBlock("tip", "Test one thing per test. If a test fails, you should immediately know what's wrong."),
						exerciseBlock(
							"What makes a good unit test?",
							"A good unit test:\n1. Tests one thing (single responsibility)\n2. Is fast (runs in milliseconds)\n3. Is independent (doesn't rely on other tests)\n4. Is repeatable (same result every time)\n5. Is readable (clear what it's testing)\n6. Uses AAA pattern (Arrange, Act, Assert)\n7. Has a descriptive name\n8. Tests behavior, not implementation\n9. Uses mocks for external dependencies\n10. Covers happy path, edge cases, and errors",
							[]string{"Think about what makes tests useful", "What happens when a test fails?"},
						),
					},
				},
				{
					ID:    "testing-3",
					Title: "Integration Testing",
					Content: []models.ContentBlock{
						textBlock(`## What is Integration Testing?

Integration testing verifies that multiple components work together correctly.

**Characteristics**:
- Tests interactions between components
- May use real dependencies
- Slower than unit tests
- More realistic scenarios`),
						textBlock(`## Why Integration Tests?

**Unit tests** verify components in isolation.
**Integration tests** verify components work together.

**Use integration tests to**:
- Verify API endpoints work
- Test database interactions
- Check component integration
- Validate data flow`),
						codeBlock("javascript", `// Integration test example

// API endpoint
app.post('/api/users', async (req, res) => {
  const user = await userService.createUser(req.body);
  res.status(201).json(user);
});

// Integration test
describe('POST /api/users', () => {
  let db;
  
  beforeAll(async () => {
    // Set up test database
    db = await setupTestDatabase();
  });
  
  afterAll(async () => {
    // Clean up
    await teardownTestDatabase(db);
  });
  
  test('creates a new user', async () => {
    // Arrange
    const userData = {
      name: 'Alice',
      email: 'alice@example.com'
    };
    
    // Act
    const response = await request(app)
      .post('/api/users')
      .send(userData)
      .expect(201);
    
    // Assert
    expect(response.body).toHaveProperty('id');
    expect(response.body.name).toBe('Alice');
    
    // Verify in database
    const user = await db.users.findOne({ email: 'alice@example.com' });
    expect(user).toBeTruthy();
  });
  
  test('returns 400 for invalid data', async () => {
    const response = await request(app)
      .post('/api/users')
      .send({ name: 'Bob' }) // Missing email
      .expect(400);
    
    expect(response.body).toHaveProperty('error');
  });
});`),
						textBlock(`## Integration Test Types

**API Integration Tests**:
- Test HTTP endpoints
- Verify request/response
- Test authentication/authorization

**Database Integration Tests**:
- Test database operations
- Verify data persistence
- Test queries and transactions

**Service Integration Tests**:
- Test services working together
- Verify data flow
- Test error handling`),
						textBlock(`## Test Database

**Use separate test database**:
- Don't use production database!
- Reset between tests
- Use transactions that rollback
- Or use in-memory database`),
						codeBlock("javascript", `// Test database setup

// Using transactions (rollback after each test)
describe('UserRepository', () => {
  let transaction;
  
  beforeEach(async () => {
    transaction = await db.beginTransaction();
  });
  
  afterEach(async () => {
    await transaction.rollback();
  });
  
  test('saves user to database', async () => {
    const repo = new UserRepository(transaction);
    const user = await repo.save({ name: 'Alice' });
    
    expect(user.id).toBeTruthy();
    const found = await repo.findById(user.id);
    expect(found.name).toBe('Alice');
  });
});`),
						calloutBlock("warning", "Never use production database for tests! Always use a separate test database or in-memory database."),
						exerciseBlock(
							"What's the difference between unit tests and integration tests?",
							"Unit tests:\n- Test individual components in isolation\n- Use mocks for dependencies\n- Fast execution\n- Many tests\n- Focus on single function/class\n\nIntegration tests:\n- Test multiple components together\n- May use real dependencies (database, APIs)\n- Slower execution\n- Fewer tests\n- Focus on interactions between components\n\nUnit tests answer 'Does this function work?' Integration tests answer 'Do these components work together?'",
							[]string{"Think about isolation vs. interaction", "What dependencies do each use?"},
						),
					},
				},
				{
					ID:    "testing-4",
					Title: "Test-Driven Development (TDD)",
					Content: []models.ContentBlock{
						textBlock(`## What is TDD?

Test-Driven Development (TDD) is a development approach where you write tests before writing code.

**TDD Cycle** (Red-Green-Refactor):

1. **Red**: Write a failing test
2. **Green**: Write minimal code to pass
3. **Refactor**: Improve code while keeping tests green
4. **Repeat**`),
						calloutBlock("info", "TDD was popularized by Kent Beck in the late 1990s. It's part of Extreme Programming (XP) methodology."),
						textBlock(`## TDD Workflow

**Step 1: Red - Write Failing Test**
- Write test for feature you want
- Test should fail (code doesn't exist yet)
- This defines what you're building

**Step 2: Green - Make It Pass**
- Write minimal code to pass test
- Don't worry about perfect code yet
- Just make it work

**Step 3: Refactor - Improve**
- Improve code quality
- Keep tests passing
- Remove duplication
- Improve design`),
						codeBlock("javascript", `// TDD Example

// Step 1: RED - Write failing test
describe('Calculator', () => {
  test('adds two numbers', () => {
    const calc = new Calculator();
    expect(calc.add(2, 3)).toBe(5);
  });
});

// Test fails - Calculator doesn't exist

// Step 2: GREEN - Write minimal code
class Calculator {
  add(a, b) {
    return a + b; // Minimal implementation
  }
}

// Test passes!

// Step 3: REFACTOR - Improve (if needed)
// Code is already simple, no refactoring needed

// Repeat for next feature
test('subtracts two numbers', () => {
  const calc = new Calculator();
  expect(calc.subtract(5, 2)).toBe(3);
});`),
						textBlock(`## Benefits of TDD

**Better Design**:
- Forces you to think about interface first
- Results in better APIs
- Encourages small, focused functions

**Confidence**:
- Know your code works
- Safe to refactor
- Tests serve as documentation

**Faster Development**:
- Less debugging
- Catch bugs early
- Less time fixing issues later`),
						textBlock(`## TDD Best Practices

**Write small tests**:
- One assertion per test (when possible)
- Test one behavior at a time
- Keep tests focused

**Write tests first**:
- Before implementation
- Defines requirements
- Guides design

**Keep tests passing**:
- Don't move to next feature until tests pass
- Refactor only when green
- Maintain test suite`),
						calloutBlock("tip", "TDD is a discipline. It feels slower at first, but saves time in the long run by preventing bugs and improving design."),
						textBlock(`## When to Use TDD

**Good for**:
- Well-defined requirements
- Complex logic
- Critical features
- Learning new codebase

**Less useful for**:
- Exploratory coding
- UI prototyping
- One-off scripts
- When requirements are unclear`),
						exerciseBlock(
							"What are the three steps of TDD, and why is the order important?",
							"The three steps are:\n1. Red - Write a failing test\n2. Green - Write minimal code to pass\n3. Refactor - Improve the code\n\nThe order is important because:\n- Writing the test first (Red) defines what you're building before you build it\n- Making it pass (Green) ensures you have working code\n- Refactoring (Refactor) improves code quality while tests ensure you don't break anything\n\nThis cycle ensures you always have tests, your code works, and you can safely improve it. Skipping steps (like writing code before tests) defeats the purpose of TDD.",
							[]string{"Why write test first?", "What happens if you skip steps?"},
						),
					},
				},
				{
					ID:    "testing-5",
					Title: "Mocking & Stubbing",
					Content: []models.ContentBlock{
						textBlock(`## Mocking and Stubbing

**Mocking** and **Stubbing** are techniques to replace real dependencies with fake ones in tests.

**Why?**:
- Isolate code under test
- Control dependencies
- Speed up tests
- Test error conditions`),
						textBlock(`## Types of Test Doubles

**Dummy**: Placeholder, never used
**Fake**: Working implementation, simplified
**Stub**: Returns predefined responses
**Mock**: Verifies interactions
**Spy**: Records interactions`),
						codeBlock("javascript", `// Stub example

class EmailService {
  send(email, message) {
    // Sends real email
  }
}

// Stub - returns predefined value
const emailStub = {
  send: jest.fn().mockResolvedValue({ success: true })
};

// Use stub in test
test('sends welcome email', async () => {
  const userService = new UserService(emailStub);
  await userService.createUser({ email: 'alice@example.com' });
  
  expect(emailStub.send).toHaveBeenCalled();
});`),
						textBlock(`## Mocking External Services

**Mock HTTP requests**:
- Don't make real API calls in tests
- Control responses
- Test error scenarios`),
						codeBlock("javascript", `// Mock HTTP request

// Mock fetch
global.fetch = jest.fn();

test('fetches user data', async () => {
  // Arrange - set up mock response
  fetch.mockResolvedValue({
    ok: true,
    json: async () => ({ id: 1, name: 'Alice' })
  });
  
  // Act
  const user = await fetchUser(1);
  
  // Assert
  expect(user.name).toBe('Alice');
  expect(fetch).toHaveBeenCalledWith('/api/users/1');
});`),
						textBlock(`## Mocking Database

**Options**:
- In-memory database
- Mock repository
- Transaction rollback`),
						codeBlock("javascript", `// Mock database repository

class UserRepository {
  async findById(id) {
    // Real database call
  }
}

// Mock for testing
const mockRepo = {
  findById: jest.fn().mockResolvedValue({ id: 1, name: 'Alice' }),
  save: jest.fn().mockResolvedValue({ id: 1, name: 'Bob' })
};

test('gets user', async () => {
  const service = new UserService(mockRepo);
  const user = await service.getUser(1);
  
  expect(user.name).toBe('Alice');
  expect(mockRepo.findById).toHaveBeenCalledWith(1);
});`),
						textBlock(`## When to Mock

**Mock when**:
- Dependency is slow (database, API)
- Dependency is unreliable
- Testing error conditions
- Dependency doesn't exist yet

**Don't mock when**:
- Dependency is fast and reliable
- Testing the integration itself
- Mock adds more complexity than value`),
						calloutBlock("warning", "Don't over-mock! Mocking everything makes tests less valuable. Sometimes you want to test the real integration."),
						exerciseBlock(
							"When should you mock a dependency in tests?",
							"You should mock a dependency when:\n1. It's slow (database, network calls) - speeds up tests\n2. It's unreliable (external APIs) - makes tests deterministic\n3. You're testing error conditions - can simulate failures\n4. It doesn't exist yet - allows TDD\n5. It has side effects (sends emails, charges credit cards) - prevents real actions\n6. You want to isolate the unit under test\n\nYou shouldn't mock when:\n- The dependency is fast and simple\n- You're writing integration tests\n- The mock adds more complexity than value\n- You want to test the real integration",
							[]string{"Think about test speed and reliability", "What are you actually testing?"},
						),
					},
				},
				{
					ID:    "testing-6",
					Title: "Testing Best Practices",
					Content: []models.ContentBlock{
						textBlock(`## Testing Best Practices

Following best practices makes your tests more valuable and maintainable.`),
						textBlock(`## Test Organization

**Structure**:
- One test file per source file
- Group related tests
- Use descriptive names
- Follow AAA pattern (Arrange, Act, Assert)`),
						codeBlock("javascript", `// Good test organization

describe('UserService', () => {
  describe('createUser', () => {
    test('creates user with valid data', () => {
      // Arrange
      const service = new UserService(mockRepo);
      const userData = { name: 'Alice', email: 'alice@example.com' };
      
      // Act
      const user = service.createUser(userData);
      
      // Assert
      expect(user.id).toBeTruthy();
      expect(user.name).toBe('Alice');
    });
    
    test('throws error when email is missing', () => {
      // Test error case
    });
  });
  
  describe('getUser', () => {
    // Tests for getUser
  });
});`),
						textBlock(`## Test Naming

**Good test names**:
- Describe what is being tested
- Include expected behavior
- Be specific and clear

**Examples**:
- ✅ should return user when valid id provided
- ✅ throws error when email is invalid
- ❌ test1
- ❌ user test`),
						textBlock(`## Keep Tests Simple

**One assertion per test** (when possible):
- Easier to understand
- Clear what failed
- Focused on one behavior

**Avoid**:
- Complex setup
- Test interdependencies
- Testing multiple things`),
						textBlock(`## Test Independence

**Tests should**:
- Run in any order
- Not depend on other tests
- Be able to run alone
- Clean up after themselves`),
						codeBlock("javascript", `// Bad - tests depend on each other
let counter = 0;

test('increments counter', () => {
  counter++;
  expect(counter).toBe(1);
});

test('counter is 2', () => {
  expect(counter).toBe(2); // Depends on previous test!
});

// Good - independent tests
test('increments counter', () => {
  let counter = 0;
  counter++;
  expect(counter).toBe(1);
});

test('counter starts at 0', () => {
  let counter = 0;
  expect(counter).toBe(0);
});`),
						textBlock(`## Test Data

**Use factories**:
- Create test data easily
- Consistent test data
- Easy to modify

**Avoid**:
- Hard-coded test data everywhere
- Shared mutable state
- Production data`),
						codeBlock("javascript", `// Test data factory

function createUser(overrides = {}) {
  return {
    id: 1,
    name: 'Alice',
    email: 'alice@example.com',
    ...overrides
  };
}

test('creates user', () => {
  const user = createUser({ name: 'Bob' });
  expect(user.name).toBe('Bob');
  expect(user.email).toBe('alice@example.com'); // Default
});`),
						calloutBlock("tip", "Good tests are like good code - readable, maintainable, and focused. Treat test code with the same care as production code."),
						exerciseBlock(
							"What makes a test maintainable and easy to understand?",
							"A maintainable test:\n1. Has a clear, descriptive name\n2. Tests one thing\n3. Uses AAA pattern (Arrange, Act, Assert)\n4. Is independent (doesn't rely on other tests)\n5. Has minimal setup\n6. Uses factories for test data\n7. Is well-organized (grouped logically)\n8. Has clear assertions\n9. Doesn't have complex logic\n10. Is readable - someone else can understand it\n\nMaintainable tests are easy to update when requirements change and easy to debug when they fail.",
							[]string{"Think about what makes code maintainable", "What happens when a test fails?"},
						),
					},
				},
				{
					ID:    "testing-7",
					Title: "Continuous Testing",
					Content: []models.ContentBlock{
						textBlock(`## Continuous Testing

Continuous testing means running tests automatically and frequently throughout development.

**Benefits**:
- Catch bugs immediately
- Prevent regressions
- Enable confident refactoring
- Faster feedback loop`),
						textBlock(`## Test Automation

**Run tests**:
- Before committing code
- On every push
- In CI/CD pipeline
- On schedule

**Tools**:
- Jest, Mocha, JUnit (test runners)
- GitHub Actions, Jenkins (CI/CD)
- Pre-commit hooks`),
						codeBlock("yaml", `# GitHub Actions example

name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Setup Node.js
        uses: actions/setup-node@v2
        with:
          node-version: '18'
      - name: Install dependencies
        run: npm install
      - name: Run tests
        run: npm test`),
						textBlock(`## Test Coverage

**Coverage metrics**:
- **Line coverage**: % of lines executed
- **Branch coverage**: % of branches tested
- **Function coverage**: % of functions called

**Aim for**:
- 80%+ coverage (good target)
- 100% coverage (ideal, but not always practical)
- Focus on critical paths`),
						calloutBlock("info", "High coverage doesn't mean good tests. It's possible to have 100% coverage with useless tests. Focus on testing the right things."),
						textBlock(`## Test Maintenance

**Keep tests updated**:
- Update tests when code changes
- Remove obsolete tests
- Refactor test code
- Keep tests fast`),
						textBlock(`## Testing in CI/CD

**Continuous Integration**:
- Run tests on every commit
- Catch issues early
- Prevent broken code in main branch

**Continuous Deployment**:
- Deploy only if tests pass
- Automated testing gates
- Confidence in releases`),
						codeBlock("text", `CI/CD Pipeline:

Code Commit
    ↓
Run Tests (Unit)
    ↓
Run Tests (Integration)
    ↓
Build Application
    ↓
Deploy to Staging
    ↓
Run E2E Tests
    ↓
Deploy to Production`),
						textBlock(`## Test Strategy

**Develop**:
- Write tests as you code
- Run tests frequently
- Fix failures immediately

**Before Commit**:
- Run full test suite
- Fix all failures
- Ensure good coverage

**In CI/CD**:
- Automated test execution
- Block deployment on failures
- Generate coverage reports`),
						calloutBlock("tip", "Make tests part of your workflow. Run them frequently, fix failures immediately, and keep them fast."),
						exerciseBlock(
							"Why is continuous testing important, and how does it help development?",
							"Continuous testing is important because:\n1. Immediate feedback - catch bugs as soon as you write them\n2. Prevents regressions - know immediately if you broke something\n3. Enables refactoring - safe to improve code with test safety net\n4. Faster development - less time debugging later\n5. Better code quality - tests force better design\n6. Confidence - know your code works\n\nIt helps development by:\n- Reducing debugging time\n- Making code changes safer\n- Providing documentation\n- Enabling faster iteration\n- Catching integration issues early",
							[]string{"Think about the feedback loop", "What happens without continuous testing?"},
						),
					},
				},
			},
		},
		{
			Title:       "Development Tools",
			Description: "Set up your development environment with the essential tools every developer needs.",
			Modules: []models.Module{
				{
					ID:    "vscode-setup",
					Title: "VS Code Setup",
					Content: []models.ContentBlock{
						textBlock(`## Setting Up Visual Studio Code

Visual Studio Code (VS Code) is a free, powerful code editor that you'll use throughout your development journey. Let's get it set up!

### Why VS Code?

- **Free and Open Source**: Completely free to use
- **Lightweight but Powerful**: Fast startup, rich features
- **Extensive Extensions**: Thousands of plugins for any language or tool
- **Built-in Git**: Version control right in your editor
- **Integrated Terminal**: Run commands without leaving the editor`),
						textBlock(`## Step 1: Download VS Code

1. Go to [code.visualstudio.com](https://code.visualstudio.com/)
2. Click the big **Download** button (it auto-detects your OS)
3. Run the installer

### Windows Installation Tips

- Check **"Add to PATH"** during installation (important!)
- Check **"Add 'Open with Code' action"** for right-click menu
- Check **"Register Code as an editor for supported file types"**`),
						calloutBlock("tip", "Adding VS Code to PATH lets you open it from the terminal by typing 'code .' in any folder!"),
						textBlock(`## Step 2: First Launch

When you first open VS Code:

1. **Choose your theme**: Dark or Light (you can change later)
2. **Skip the welcome tour** for now (or take it if you want!)
3. You'll see the **Welcome** tab - this is your starting point`),
						textBlock(`## Step 3: Essential Extensions

Click the **Extensions** icon in the left sidebar (or press Ctrl+Shift+X).

Search for and install these extensions:`),
						codeBlock("text", `Essential Extensions:

1. ESLint - JavaScript/TypeScript linting
2. Prettier - Code formatter
3. GitLens - Enhanced Git features
4. Auto Rename Tag - HTML/JSX tag renaming
5. Bracket Pair Colorizer - Colored matching brackets`),
						calloutBlock("info", "Extensions supercharge VS Code. Start with these essentials, and add more as you discover what you need!"),
						textBlock(`## Step 4: Configure Settings

Open Settings: Ctrl+, (or Cmd+, on Mac)

### Recommended Settings

Click the **{}** icon in the top right to open settings.json and add:`),
						codeBlock("json", `{
  "editor.fontSize": 14,
  "editor.tabSize": 2,
  "editor.wordWrap": "on",
  "editor.formatOnSave": true,
  "editor.minimap.enabled": false,
  "files.autoSave": "afterDelay",
  "terminal.integrated.fontSize": 13
}`),
						textBlock(`## Step 5: Learn the Key Shortcuts

These will save you hours:`),
						codeBlock("text", "Essential Shortcuts:\n\nCtrl+P          - Quick file open\nCtrl+Shift+P    - Command palette (access ANY feature)\nCtrl+B          - Toggle sidebar\nCtrl+`          - Toggle terminal\nCtrl+/          - Comment/uncomment line\nCtrl+D          - Select next occurrence\nCtrl+Shift+K    - Delete line\nAlt+Up/Down     - Move line up/down\nCtrl+Space      - Trigger suggestions"),
						calloutBlock("tip", "The Command Palette (Ctrl+Shift+P) is your best friend. If you forget how to do something, search for it there!"),
						textBlock(`## Step 6: Open a Folder

VS Code works best when you open an entire folder (project):

1. Go to **File** > **Open Folder**
2. Navigate to your my-first-repo folder
3. Click **Select Folder**

Now you can see all your project files in the sidebar!`),
						textBlock("## Step 7: Using the Integrated Terminal\n\nOpen the terminal: Ctrl+` (backtick key)\n\nThis terminal is just like your regular command prompt/terminal, but built right into VS Code!\n\nTry these commands:"),
						codeBlock("bash", `# Check where you are
pwd

# List files
ls

# Check Git status
git status`),
						exerciseBlock(
							"Open VS Code, install the ESLint and Prettier extensions, and open your my-first-repo folder. What do you see in the sidebar?",
							"You should see the Explorer sidebar showing your project files, including HelloWorld.txt and the .git folder (if hidden files are shown). The file tree shows your entire project structure.",
							[]string{"Look at the left sidebar", "You should see your HelloWorld.txt file", "The Explorer view shows your folder contents"},
						),
						textBlock(`## You're Ready!

Your VS Code is now configured and ready for coding. In the upcoming exercises, you'll use VS Code to:

- Edit code files
- Run terminal commands
- Manage Git operations
- Debug your applications

**Keep VS Code open** - you'll need it for the next modules!`),
					},
				},
			},
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
				{
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
				},
				{
					ID:    "react-hello-world",
					Title: "Your First React Page",
					Content: []models.ContentBlock{
						textBlock(`## Build Your First React Web Page

Time to create something you can see in your browser! We'll use React (a popular UI library) and Vite (a fast build tool) to create a personal info page.

### What We'll Build

A simple web page about yourself that displays:
- Your name
- A greeting message
- Some facts about you
- Dynamic content that changes based on code`),
						calloutBlock("info", "React is used by Facebook, Instagram, Netflix, and thousands of other companies. Learning React opens many doors!"),
						textBlock(`## Step 1: Create a New React Project

Open your terminal and run:`),
						codeBlock("bash", `# Navigate to your projects folder
cd ~/projects

# Create a new Vite + React project
npm create vite@latest my-first-page -- --template react

# Navigate into the project
cd my-first-page

# Install dependencies
npm install`),
						textBlock(`## Step 2: Start the Development Server`),
						codeBlock("bash", `npm run dev`),
						textBlock(`You should see output showing that Vite is ready and running on http://localhost:5173/

**Open your browser to http://localhost:5173** - you'll see the default Vite + React page!`),
						calloutBlock("tip", "The development server auto-reloads when you save changes. Keep it running while you code!"),
						textBlock(`## Step 3: Edit Your First Component

Open src/App.jsx in VS Code. Replace ALL the content with:`),
						codeBlock("jsx", `function App() {
  // Your personal info - customize these!
  const myName = "Your Name";
  const greeting = "Hello, World!";
  const favoriteColor = "blue";
  const hobbies = ["coding", "reading", "gaming"];
  const yearsLearning = 0;

  return (
    <div style={{ 
      padding: "40px", 
      fontFamily: "Arial, sans-serif",
      maxWidth: "600px",
      margin: "0 auto"
    }}>
      <h1 style={{ color: favoriteColor }}>
        {greeting}
      </h1>
      
      <h2>I'm {myName}</h2>
      
      <p>
        I've been learning to code for {yearsLearning} years.
        {yearsLearning === 0 && " Just getting started!"}
      </p>

      <h3>My Hobbies:</h3>
      <ul>
        {hobbies.map((hobby, index) => (
          <li key={index}>{hobby}</li>
        ))}
      </ul>

      <button 
        onClick={() => alert("You clicked the button!")}
        style={{
          padding: "10px 20px",
          fontSize: "16px",
          cursor: "pointer"
        }}
      >
        Click Me!
      </button>
    </div>
  );
}

export default App;`),
						textBlock(`**Save the file** (Ctrl+S) and watch your browser update automatically!`),
						calloutBlock("info", "Notice how JavaScript variables like 'myName' appear inside the HTML using {curly braces}. This is JSX - HTML mixed with JavaScript!"),
						textBlock(`## Step 4: Customize Your Page

Now make it yours! Edit the variables at the top of the component:

1. Change myName to your actual name
2. Update favoriteColor to your favorite color
3. Replace the hobbies array with your real hobbies
4. Add more content below!`),
						textBlock(`## Step 5: Add More Content

Try adding these below the button:`),
						codeBlock("jsx", `      <hr style={{ margin: "30px 0" }} />
      
      <h3>About This Page</h3>
      <p>
        This is my first React page! I built it while learning 
        web development. The content above is generated from 
        JavaScript variables, not hardcoded HTML.
      </p>

      <h3>What I'm Learning:</h3>
      <ul>
        <li>Git and GitHub</li>
        <li>VS Code</li>
        <li>TypeScript basics</li>
        <li>React and Vite</li>
      </ul>`),
						exerciseBlock(
							"Customize your React page with your real name, hobbies, and favorite color. Add at least one new section about yourself. What happens when you save the file?",
							"When you save the file, Vite's hot module replacement (HMR) automatically updates your browser without a full page refresh. Your changes appear instantly! The page should now show your personalized content.",
							[]string{"Edit the variables at the top of App.jsx", "Save with Ctrl+S", "Watch your browser update automatically"},
						),
						textBlock(`## Understanding What You Built

Let's break down what's happening:

- **Variables**: const myName = "..." stores data you can use anywhere
- **JSX**: HTML-like syntax that lets you use JavaScript with {curly braces}
- **Props and State**: The foundation of React (we'll learn more later!)
- **Event Handling**: onClick runs code when you click
- **Array.map()**: Converts an array into a list of elements`),
						textBlock(`## Congratulations!

You just built your first web page with React! This is the same technology used to build:

- Facebook and Instagram
- Netflix's interface
- Airbnb's website
- Discord's app

You're on your way to becoming a web developer!`),
						calloutBlock("tip", "Keep this project! You can expand it as you learn more. Try adding images, more sections, or even multiple pages!"),
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
