package seed

import "github.com/pathway/backend/models"

func moduleFizzBuzzGaltonUI() models.Module {
	return models.Module{
		ID:    "fizzbuzz-galton-ui",
		Title: "Coding Exercise: FizzBuzz Galton Board Animation (React + TypeScript)",
		Content: []models.ContentBlock{
			textBlock(`## What You're Building

YouÃƒÂ¢Ã¢â€šÂ¬Ã¢â€žÂ¢ll build a small animated UI that *visualizes* your FizzBuzz engine.

Imagine numbers dropping from the top like a Galton board:
- Each number "falls" down the screen
- It lands in one of four buckets:
  - Fizz
  - Buzz
  - FizzBuzz
  - Number

This teaches you:
- How to use your TypeScript class from React
- How to model data in state
- How to animate UI changes with CSS
- How to build a small system end-to-end

## Step 1: Create a component file

Create:

src/components/GaltonFizzBuzz.tsx

Paste this starter code:`),
			codeBlock("tsx", `import { useEffect, useMemo, useRef, useState } from "react";
import { FizzBuzzEngine } from "../lib/FizzBuzzEngine";
import "./GaltonFizzBuzz.css";

type Bucket = "Fizz" | "Buzz" | "FizzBuzz" | "Number";

type Ball = {
  id: string;
  n: number;
  bucket: Bucket;
  leftPx: number;
};

function bucketFor(value: string): Bucket {
  if (value === "FizzBuzz") return "FizzBuzz";
  if (value === "Fizz") return "Fizz";
  if (value === "Buzz") return "Buzz";
  return "Number";
}

export default function GaltonFizzBuzz() {
  const engine = useMemo(() => new FizzBuzzEngine(), []);

  const nextN = useRef(1);
  const [balls, setBalls] = useState<Ball[]>([]);
  const [counts, setCounts] = useState<Record<Bucket, number>>({
    Fizz: 0,
    Buzz: 0,
    FizzBuzz: 0,
    Number: 0,
  });

  useEffect(() => {
    const interval = window.setInterval(() => {
      const n = nextN.current++;
      const value = engine.valueFor(n);
      const bucket = bucketFor(value);

      // Place balls roughly above their target bucket, with a little random wobble
      const bucketCenters: Record<Bucket, number> = {
        Fizz: 110,
        Buzz: 250,
        FizzBuzz: 390,
        Number: 530,
      };

      const leftPx = bucketCenters[bucket] + (Math.random() * 40 - 20);

      const ball: Ball = {
        id: String(Date.now()) + "-" + String(Math.random()),
        n,
        bucket,
        leftPx,
      };

      setBalls((prev) => [...prev, ball]);
    }, 650);

    return () => window.clearInterval(interval);
  }, [engine]);

  const handleBallLanded = (ball: Ball) => {
    setBalls((prev) => prev.filter((b) => b.id !== ball.id));
    setCounts((prev) => ({ ...prev, [ball.bucket]: prev[ball.bucket] + 1 }));
  };

  return (
    <div className="galton">
      <h2>FizzBuzz Galton Board</h2>

      <div className="board">
        {balls.map((ball) => (
          <div
            key={ball.id}
            className={"ball ball--" + ball.bucket.toLowerCase()}
            style={{ left: ball.leftPx }}
            onAnimationEnd={() => handleBallLanded(ball)}
            title={"n=" + ball.n + " -> " + engine.valueFor(ball.n)}
          >
            {ball.n}
          </div>
        ))}

        <div className="buckets">
          <div className="bucket">
            <div className="bucket__label">Fizz</div>
            <div className="bucket__count">{counts.Fizz}</div>
          </div>
          <div className="bucket">
            <div className="bucket__label">Buzz</div>
            <div className="bucket__count">{counts.Buzz}</div>
          </div>
          <div className="bucket">
            <div className="bucket__label">FizzBuzz</div>
            <div className="bucket__count">{counts.FizzBuzz}</div>
          </div>
          <div className="bucket">
            <div className="bucket__label">Number</div>
            <div className="bucket__count">{counts.Number}</div>
          </div>
        </div>
      </div>
    </div>
  );
}`),
			textBlock(`## Step 2: Add the CSS animation

Create:

src/components/GaltonFizzBuzz.css

Paste this CSS:`),
			codeBlock("css", `.galton {
  padding: 24px;
  font-family: system-ui, Arial, sans-serif;
}

.board {
  position: relative;
  height: 520px;
  border: 1px solid rgba(255,255,255,0.15);
  border-radius: 12px;
  background: rgba(255,255,255,0.03);
  overflow: hidden;
}

.ball {
  position: absolute;
  top: -40px;
  width: 36px;
  height: 36px;
  border-radius: 999px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 700;
  color: #0f172a;
  animation: fall 1.2s cubic-bezier(.2,.9,.25,1) forwards;
  box-shadow: 0 8px 18px rgba(0,0,0,0.25);
}

.ball--fizz { background: #a7f3d0; }
.ball--buzz { background: #bfdbfe; }
.ball--fizzbuzz { background: #fecaca; }
.ball--number { background: #fde68a; }

@keyframes fall {
  0% { transform: translateY(0); opacity: 0.9; }
  80% { transform: translateY(420px); }
  100% { transform: translateY(450px); opacity: 1; }
}

.buckets {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 140px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 12px;
  padding: 12px;
  background: rgba(0,0,0,0.2);
  border-top: 1px solid rgba(255,255,255,0.12);
}

.bucket {
  border: 1px solid rgba(255,255,255,0.15);
  border-radius: 10px;
  padding: 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.bucket__label {
  font-weight: 700;
  color: rgba(255,255,255,0.9);
}

.bucket__count {
  font-weight: 800;
  font-size: 20px;
  color: rgba(255,255,255,0.9);
}`),
			textBlock(`## Step 3: Render it in App.tsx

Update src/App.tsx to render your new component:`),
			codeBlock("tsx", `import GaltonFizzBuzz from "./components/GaltonFizzBuzz";

export default function App() {
  return <GaltonFizzBuzz />;
}`),
			calloutBlock("tip", "If the balls all land in 'Number', check your FizzBuzz rules first. The UI is only as correct as the engine!"),
			exerciseBlock(
				"Add two buttons: Pause/Resume and Reset Counts. (Hint: store a boolean in state and only create the interval when running.)",
				"A typical solution uses a running boolean in state and conditionally starts/stops the interval in useEffect. Reset counts by setting the counts state back to zeros and clearing balls.",
				[]string{"Add a `running` state boolean", "In useEffect, only setInterval when running is true", "Reset counts by setting state back to the initial object"},
			),
		},
	}
}
