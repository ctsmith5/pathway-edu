package seed

import "github.com/pathway/backend/models"

func moduleGit7() models.Module {
	return models.Module{
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
	}
}
