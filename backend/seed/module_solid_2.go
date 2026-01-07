package seed

import "github.com/pathway/backend/models"

func moduleSolid2() models.Module {
	return models.Module{
		ID:    "solid-2",
		Title: "Single Responsibility Principle (SRP)",
		Content: []models.ContentBlock{
			textBlock(`## Single Responsibility Principle

**A class should have only one reason to change.**

Every class should have a single, well-defined responsibility. If a class has multiple reasons to change, it's doing too much.`),
			textBlock(`## Understanding "One Reason to Change"

A "reason to change" means a class should be responsible to only one stakeholder or concern:

- **User management** - one class
- **Email sending** - another class
- **Data persistence** - yet another class`),
			codeBlock("javascript", `// ❌ BAD: Multiple responsibilities
class User {
  constructor(name, email) {
    this.name = name;
    this.email = email;
  }
  
  saveToDatabase() {
    // Database logic
  }
  
  sendEmail() {
    // Email logic
  }
  
  validateEmail() {
    // Validation logic
  }
}

// ✅ GOOD: Single responsibility
class User {
  constructor(name, email) {
    this.name = name;
    this.email = email;
  }
}

class UserRepository {
  save(user) {
    // Database logic
  }
}

class EmailService {
  send(user, message) {
    // Email logic
  }
}

class EmailValidator {
  validate(email) {
    // Validation logic
  }
}`),
			calloutBlock("tip", "Ask yourself: 'If the requirements for X change, would I need to modify this class?' If the answer is yes for multiple X's, you're violating SRP."),
			exerciseBlock(
				"A class handles user authentication, logs all login attempts, and sends welcome emails. How would you refactor this to follow SRP?",
				"Split into three classes:\n1. AuthService - handles authentication logic\n2. LoginLogger - logs login attempts\n3. EmailService - sends welcome emails\n\nThe User class would coordinate these services but not contain their logic.",
				[]string{"Each responsibility should be its own class", "Think about what changes independently"},
			),
		},
	}
}
