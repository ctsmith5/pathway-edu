package seed

import "github.com/pathway/backend/models"

func moduleFizzBuzzClass() models.Module {
	return models.Module{
		ID:    "fizzbuzz-class",
		Title: "Build a TypeScript Class (FizzBuzz Engine)",
		Content: []models.ContentBlock{
			textBlock(`## The Goal

You're going to build a tiny "engine" in TypeScript that solves the classic FizzBuzz challenge.

FizzBuzz rules:
- If a number is divisible by 3 ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ "Fizz"
- If a number is divisible by 5 ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ "Buzz"
- If divisible by both ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ "FizzBuzz"
- Otherwise ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ the number itself

## Step 1: Create a TypeScript React project

If your current React project is JavaScript, itÃƒÂ¢Ã¢â€šÂ¬Ã¢â€žÂ¢s totally fine to make a new one for this exercise:
`),
			codeBlock("bash", `cd ~/projects
npm create vite@latest fizzbuzz-lab -- --template react-ts
cd fizzbuzz-lab
npm install
npm run dev`),
			textBlock(`## Step 2: Add a class file

Create a new file:

src/lib/FizzBuzzEngine.ts

Paste this code:`),
			codeBlock("typescript", `export type FizzBuzzResult = "Fizz" | "Buzz" | "FizzBuzz" | string;

export class FizzBuzzEngine {
  public valueFor(n: number): FizzBuzzResult {
    if (!Number.isInteger(n)) return "Not a whole number";

    const divisibleBy3 = n % 3 === 0;
    const divisibleBy5 = n % 5 === 0;

    if (divisibleBy3 && divisibleBy5) return "FizzBuzz";
    if (divisibleBy3) return "Fizz";
    if (divisibleBy5) return "Buzz";
    return String(n);
  }
}`),
			textBlock(`## Step 3: Try it quickly in the UI

Open src/App.tsx and temporarily replace the default content with:`),
			codeBlock("tsx", `import { useMemo } from "react";
import { FizzBuzzEngine } from "./lib/FizzBuzzEngine";

export default function App() {
  const engine = useMemo(() => new FizzBuzzEngine(), []);

  const results = [];
  for (let n = 1; n <= 30; n++) {
    results.push({ n, value: engine.valueFor(n) });
  }

  return (
    <div style={{ padding: 24, fontFamily: "system-ui, Arial" }}>
      <h1>FizzBuzz (Engine Test)</h1>
      <ol>
        {results.map((r) => (
          <li key={r.n}>
            {r.n} ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ <strong>{r.value}</strong>
          </li>
        ))}
      </ol>
    </div>
  );
}`),
			calloutBlock("info", "Notice the separation: the class does the logic; the React component just displays results. This is a professional habit."),
			exerciseBlock(
				"Change the engine so that it also returns 'Bazz' for numbers divisible by 7. How would you add this without breaking the original rules?",
				"You can add another divisibility check (n % 7 === 0) and decide how to combine labels. A clean approach is to build a string result and append parts (e.g., result += 'Fizz').",
				[]string{"Try building the output string step-by-step", "Think about combining multiple rules"},
			),
		},
	}
}
