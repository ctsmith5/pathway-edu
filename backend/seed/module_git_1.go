package seed

import "github.com/pathway/backend/models"

func moduleGit1() models.Module {
	return models.Module{
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

1. Go to GitHub â†’ Click your profile picture â†’ **Settings**
2. Scroll down to **Developer settings** (bottom of left sidebar)
3. Click **Personal access tokens** â†’ **Tokens (classic)**
4. Click **Generate new token** â†’ **Generate new token (classic)**
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
	}
}
