package seed

import "github.com/pathway/backend/models"

func moduleSolid5() models.Module {
	return models.Module{
		ID:    "solid-5",
		Title: "Interface Segregation Principle (ISP)",
		Content: []models.ContentBlock{
			textBlock(`## Interface Segregation Principle

**Clients should not be forced to depend on interfaces they do not use.**

Create specific, focused interfaces rather than large, general-purpose ones.`),
			textBlock(`## The Problem with Fat Interfaces

When interfaces are too large:
- Classes must implement methods they don't need
- Changes to one part affect unrelated classes
- Code becomes harder to understand and maintain`),
			codeBlock("javascript", `// ❌ BAD: Fat interface
class Worker {
  work() {}
  eat() {}
  sleep() {}
}

class Human extends Worker {
  work() { /* humans work */ }
  eat() { /* humans eat */ }
  sleep() { /* humans sleep */ }
}

class Robot extends Worker {
  work() { /* robots work */ }
  eat() { /* robots don't eat! */ }
  sleep() { /* robots don't sleep! */ }
}

// ✅ GOOD: Segregated interfaces
class Workable {
  work() {}
}

class Eatable {
  eat() {}
}

class Sleepable {
  sleep() {}
}

class Human implements Workable, Eatable, Sleepable {
  work() { /* humans work */ }
  eat() { /* humans eat */ }
  sleep() { /* humans sleep */ }
}

class Robot implements Workable {
  work() { /* robots work */ }
  // No need to implement eat() or sleep()
}`),
			calloutBlock("info", "ISP is about creating lean, focused interfaces. Each interface should represent a specific role or capability."),
			textBlock(`## Benefits of ISP

- **Flexibility**: Classes only depend on what they need
- **Maintainability**: Changes to one interface don't affect unrelated classes
- **Clarity**: Interfaces clearly communicate their purpose
- **Testability**: Easier to mock and test specific behaviors`),
			exerciseBlock(
				"You have a Printer interface with print(), scan(), fax(), and copy() methods. A basic printer only needs print(). How would you refactor this?",
				"Split into separate interfaces:\n- Printable: print()\n- Scannable: scan()\n- Faxable: fax()\n- Copyable: copy()\n\nA basic printer implements only Printable. Advanced printers can implement multiple interfaces as needed.",
				[]string{"Not all printers have all capabilities", "Think about what each device actually needs"},
			),
		},
	}
}
