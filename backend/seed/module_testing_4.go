package seed

import "github.com/pathway/backend/models"

func moduleTesting4() models.Module {
	return models.Module{
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
	}
}
