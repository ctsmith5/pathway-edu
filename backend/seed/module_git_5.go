package seed

import "github.com/pathway/backend/models"

func moduleGit5() models.Module {
	return models.Module{
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
	}
}
