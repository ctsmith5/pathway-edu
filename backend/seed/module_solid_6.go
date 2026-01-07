package seed

import "github.com/pathway/backend/models"

func moduleSolid6() models.Module {
	return models.Module{
		ID:    "solid-6",
		Title: "Dependency Inversion Principle (DIP)",
		Content: []models.ContentBlock{
			textBlock(`## Dependency Inversion Principle

**High-level modules should not depend on low-level modules. Both should depend on abstractions.**

Depend on interfaces, not concrete implementations.`),
			textBlock(`## Understanding the Layers

- **High-level modules**: Business logic, application flow
- **Low-level modules**: Database access, file I/O, external APIs
- **Abstractions**: Interfaces, abstract classes, protocols`),
			codeBlock("javascript", `// ❌ BAD: High-level depends on low-level
class MySQLDatabase {
  save(data) {
    // MySQL-specific code
  }
}

class UserService {
  constructor() {
    this.db = new MySQLDatabase(); // Direct dependency!
  }
  
  createUser(user) {
    this.db.save(user);
  }
}

// ✅ GOOD: Both depend on abstraction
class Database {
  save(data) {
    throw new Error('Must implement');
  }
}

class MySQLDatabase extends Database {
  save(data) {
    // MySQL-specific code
  }
}

class PostgreSQLDatabase extends Database {
  save(data) {
    // PostgreSQL-specific code
  }
}

class UserService {
  constructor(database) { // Depends on abstraction
    this.db = database;
  }
  
  createUser(user) {
    this.db.save(user);
  }
}

// Can swap implementations easily
const userService = new UserService(new MySQLDatabase());
// or
const userService = new UserService(new PostgreSQLDatabase());`),
			calloutBlock("tip", "Dependency Injection is the technique used to implement DIP. Pass dependencies in rather than creating them internally."),
			textBlock(`## Benefits of DIP

- **Flexibility**: Easy to swap implementations
- **Testability**: Can inject mocks for testing
- **Maintainability**: Changes to low-level modules don't affect high-level ones
- **Reusability**: High-level modules can work with any implementation`),
			exerciseBlock(
				"A NotificationService directly creates instances of EmailSender and SMSSender. How would you refactor this to follow DIP?",
				"Create a Notifier interface with a send() method. EmailSender and SMSSender implement this interface. NotificationService depends on the Notifier interface and receives it via dependency injection. This allows easy swapping of notification methods and better testing.",
				[]string{"What abstraction could represent both email and SMS?", "How can you inject the dependency?"},
			),
		},
	}
}
