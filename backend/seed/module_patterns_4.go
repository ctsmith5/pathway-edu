package seed

import "github.com/pathway/backend/models"

func modulePatterns4() models.Module {
	return models.Module{
		ID:    "patterns-4",
		Title: "Behavioral Patterns (Observer, Strategy)",
		Content: []models.ContentBlock{
			textBlock(`## Behavioral Patterns

Behavioral patterns deal with object interaction - how objects communicate and distribute responsibility.

**Common patterns**:
- Observer
- Strategy
- Command
- Chain of Responsibility
- State`),
			textBlock(`## Observer Pattern

**Intent**: Define a one-to-many dependency between objects so when one changes, all dependents are notified.

**Use when**:
- Change to one object requires changing others
- Don't know how many objects need to be updated
- Want loose coupling between objects

**Example uses**:
- Event systems
- Model-View architecture
- Publish-subscribe systems`),
			codeBlock("javascript", `// Observer Pattern

class Subject {
  constructor() {
    this.observers = [];
  }
  
  subscribe(observer) {
    this.observers.push(observer);
  }
  
  unsubscribe(observer) {
    this.observers = this.observers.filter(o => o !== observer);
  }
  
  notify(data) {
    this.observers.forEach(observer => observer.update(data));
  }
}

class Observer {
  update(data) {
    throw new Error('Must implement');
  }
}

class EmailObserver extends Observer {
  update(data) {
    console.log('Sending email: ' + data);
  }
}

class SMSObserver extends Observer {
  update(data) {
    console.log('Sending SMS: ' + data);
  }
}

// Usage
const subject = new Subject();
subject.subscribe(new EmailObserver());
subject.subscribe(new SMSObserver());

subject.notify('User logged in'); // Both observers notified`),
			calloutBlock("info", "Observer pattern is the foundation of many event systems. React's state management uses similar concepts."),
			textBlock(`## Strategy Pattern

**Intent**: Define a family of algorithms, encapsulate each one, and make them interchangeable.

**Use when**:
- Have multiple ways to perform a task
- Want to switch algorithms at runtime
- Want to avoid conditional statements

**Example uses**:
- Sorting algorithms
- Payment methods
- Compression algorithms
- Validation strategies`),
			codeBlock("javascript", `// Strategy Pattern

class Sorter {
  constructor(strategy) {
    this.strategy = strategy;
  }
  
  setStrategy(strategy) {
    this.strategy = strategy;
  }
  
  sort(data) {
    return this.strategy.sort(data);
  }
}

class QuickSort {
  sort(data) {
    return [...data].sort((a, b) => a - b);
  }
}

class BubbleSort {
  sort(data) {
    const arr = [...data];
    for (let i = 0; i < arr.length; i++) {
      for (let j = 0; j < arr.length - i - 1; j++) {
        if (arr[j] > arr[j + 1]) {
          [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        }
      }
    }
    return arr;
  }
}

// Usage
const sorter = new Sorter(new QuickSort());
sorter.sort([3, 1, 4, 2]); // Uses QuickSort

sorter.setStrategy(new BubbleSort());
sorter.sort([3, 1, 4, 2]); // Uses BubbleSort`),
			textBlock(`## Command Pattern

**Intent**: Encapsulate a request as an object, allowing parameterization and queuing of requests.

**Use when**:
- Need to parameterize objects with operations
- Need to queue operations
- Need to support undo/redo

**Example uses**:
- Undo/redo functionality
- Macro recording
- Job queues
- Transaction systems`),
			codeBlock("javascript", `// Command Pattern

class Command {
  execute() {
    throw new Error('Must implement');
  }
  
  undo() {
    throw new Error('Must implement');
  }
}

class Light {
  on() { return 'Light is ON'; }
  off() { return 'Light is OFF'; }
}

class LightOnCommand extends Command {
  constructor(light) {
    super();
    this.light = light;
  }
  
  execute() {
    return this.light.on();
  }
  
  undo() {
    return this.light.off();
  }
}

class Invoker {
  constructor() {
    this.history = [];
  }
  
  execute(command) {
    this.history.push(command);
    return command.execute();
  }
  
  undo() {
    if (this.history.length > 0) {
      const command = this.history.pop();
      return command.undo();
    }
  }
}

// Usage
const light = new Light();
const command = new LightOnCommand(light);
const invoker = new Invoker();

invoker.execute(command); // Light on
invoker.undo(); // Light off`),
			calloutBlock("tip", "Observer is about notifications. Strategy is about interchangeable algorithms. Command is about encapsulating operations."),
			exerciseBlock(
				"When would you use Observer vs Strategy pattern?",
				"Use Observer when you need one object to notify multiple other objects about changes (like an event system, where one event triggers multiple handlers).\n\nUse Strategy when you have multiple ways to accomplish the same task and want to switch between them (like different sorting algorithms, payment methods, or validation rules).\n\nObserver is about notifications and dependencies. Strategy is about interchangeable algorithms.",
				[]string{"What problem does each solve?", "Think about notifications vs. algorithms"},
			),
		},
	}
}
