package seed

import "github.com/pathway/backend/models"

func moduleGit6() models.Module {
	return models.Module{
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
	}
}
