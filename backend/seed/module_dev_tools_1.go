package seed

import "github.com/pathway/backend/models"

func moduleDevTools1() models.Module {
	return models.Module{
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
	}
}
