package seed

import "github.com/pathway/backend/models"

func moduleReactHelloWorld() models.Module {
	return models.Module{
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
	}
}
