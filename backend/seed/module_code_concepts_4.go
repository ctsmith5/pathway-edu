package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts4() models.Module {
	return models.Module{
		ID:    "code-4",
		Title: "Functions & Closures",
		Content: []models.ContentBlock{
			textBlock(`## Understanding Functions and Closures

Functions are reusable blocks of code. Closures are functions that capture and remember their surrounding context.

### What is a Closure?

A closure is a function that has access to variables from its outer (enclosing) scope, even after that outer function has returned.`),
			codeBlock("javascript", `function createCounter() {
  let count = 0;  // This variable is "captured" by the closure
  
  return function() {
    count += 1;
    return count;
  };
}

const counter = createCounter();
console.log(counter()); // 1
console.log(counter()); // 2
console.log(counter()); // 3`),
			calloutBlock("info", "Closures are powerful for creating private state, callbacks, and functional patterns like currying."),
		},
	}
}
