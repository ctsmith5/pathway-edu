package seed

import "github.com/pathway/backend/models"

func modulePatterns3() models.Module {
	return models.Module{
		ID:    "patterns-3",
		Title: "Structural Patterns (Adapter, Decorator)",
		Content: []models.ContentBlock{
			textBlock(`## Structural Patterns

Structural patterns deal with object composition - how classes and objects are composed to form larger structures.

**Common patterns**:
- Adapter
- Decorator
- Facade
- Proxy
- Bridge`),
			textBlock(`## Adapter Pattern

**Intent**: Allow incompatible interfaces to work together.

**Use when**:
- Want to use existing class with incompatible interface
- Need to integrate third-party libraries
- Want to create reusable class for unrelated classes

**Example uses**:
- Integrating legacy code
- Wrapping third-party APIs
- Making incompatible interfaces compatible`),
			codeBlock("javascript", `// Adapter Pattern

// Old interface (incompatible)
class OldPaymentSystem {
  makePayment(amount, currency) {
    return 'Paid ' + amount + ' ' + currency;
  }
}

// New interface (what we want)
class PaymentProcessor {
  pay(amount) {
    throw new Error('Must implement');
  }
}

// Adapter
class PaymentAdapter extends PaymentProcessor {
  constructor(oldSystem) {
    super();
    this.oldSystem = oldSystem;
  }
  
  pay(amount) {
    // Adapt old interface to new
    return this.oldSystem.makePayment(amount, 'USD');
  }
}

// Usage
const oldSystem = new OldPaymentSystem();
const adapter = new PaymentAdapter(oldSystem);
adapter.pay(100); // Works with new interface`),
			calloutBlock("info", "Adapter is like a translator - it makes two incompatible interfaces work together without changing either one."),
			textBlock(`## Decorator Pattern

**Intent**: Attach additional responsibilities to objects dynamically.

**Use when**:
- Want to add behavior to objects at runtime
- Don't want to modify existing code
- Need to add features flexibly

**Example uses**:
- Adding logging to methods
- Adding caching to functions
- Adding validation to inputs
- UI component enhancements`),
			codeBlock("javascript", `// Decorator Pattern

class Coffee {
  cost() {
    return 5;
  }
  
  description() {
    return 'Coffee';
  }
}

// Decorators
class MilkDecorator {
  constructor(coffee) {
    this.coffee = coffee;
  }
  
  cost() {
    return this.coffee.cost() + 2;
  }
  
  description() {
    return this.coffee.description() + ', Milk';
  }
}

class SugarDecorator {
  constructor(coffee) {
    this.coffee = coffee;
  }
  
  cost() {
    return this.coffee.cost() + 1;
  }
  
  description() {
    return this.coffee.description() + ', Sugar';
  }
}

// Usage
let coffee = new Coffee();
coffee = new MilkDecorator(coffee);
coffee = new SugarDecorator(coffee);

console.log(coffee.description()); // "Coffee, Milk, Sugar"
console.log(coffee.cost()); // 8`),
			textBlock(`## Facade Pattern

**Intent**: Provide a simplified interface to a complex subsystem.

**Use when**:
- Want to hide complexity
- Need simple interface to complex system
- Want to decouple clients from subsystem

**Example uses**:
- API wrappers
- Library interfaces
- System abstractions`),
			codeBlock("javascript", `// Facade Pattern

// Complex subsystem
class CPU {
  freeze() { return 'CPU frozen'; }
  jump(position) { return 'Jump to ' + position; }
  execute() { return 'CPU executing'; }
}

class Memory {
  load(position, data) { return 'Load ' + data + ' at ' + position; }
}

class HardDrive {
  read(lba, size) { return 'Read ' + size + ' bytes from ' + lba; }
}

// Facade
class Computer {
  constructor() {
    this.cpu = new CPU();
    this.memory = new Memory();
    this.hardDrive = new HardDrive();
  }
  
  start() {
    // Simple interface to complex system
    this.cpu.freeze();
    this.memory.load('0x0', this.hardDrive.read('0', '1024'));
    this.cpu.jump('0x0');
    this.cpu.execute();
    return 'Computer started';
  }
}

// Usage - simple interface
const computer = new Computer();
computer.start();`),
			calloutBlock("tip", "Facade simplifies complex systems. Adapter makes incompatible interfaces work together. Decorator adds behavior dynamically."),
			exerciseBlock(
				"What's the difference between Adapter and Facade patterns?",
				"Adapter makes incompatible interfaces work together - it translates between two different interfaces so they can collaborate.\n\nFacade provides a simplified interface to a complex subsystem - it hides complexity behind a simple interface.\n\nAdapter is about compatibility (making A work with B), while Facade is about simplicity (hiding complexity behind a simple interface).",
				[]string{"What problem does each solve?", "Think about compatibility vs. simplicity"},
			),
		},
	}
}
