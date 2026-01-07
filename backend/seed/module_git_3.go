package seed

import "github.com/pathway/backend/models"

func moduleGit3() models.Module {
	return models.Module{
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
	}
}
