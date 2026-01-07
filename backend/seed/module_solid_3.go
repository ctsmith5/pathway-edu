package seed

import "github.com/pathway/backend/models"

func moduleSolid3() models.Module {
	return models.Module{
		ID:    "solid-3",
		Title: "Open/Closed Principle (OCP)",
		Content: []models.ContentBlock{
			textBlock(`## Open/Closed Principle

**Software entities should be open for extension but closed for modification.**

You should be able to add new functionality without changing existing code.`),
			textBlock(`## What Does This Mean?

- **Open for Extension**: You can add new behavior
- **Closed for Modification**: You don't need to change existing code

This is achieved through:
- Inheritance
- Polymorphism
- Interfaces/Abstract classes`),
			codeBlock("javascript", `// ❌ BAD: Must modify existing code
class AreaCalculator {
  calculate(shape) {
    if (shape.type === 'circle') {
      return Math.PI * shape.radius ** 2;
    } else if (shape.type === 'rectangle') {
      return shape.width * shape.height;
    }
    // Adding a new shape requires modifying this method!
  }
}

// ✅ GOOD: Open for extension
class Shape {
  area() {
    throw new Error('Must implement area()');
  }
}

class Circle extends Shape {
  constructor(radius) {
    super();
    this.radius = radius;
  }
  
  area() {
    return Math.PI * this.radius ** 2;
  }
}

class Rectangle extends Shape {
  constructor(width, height) {
    super();
    this.width = width;
    this.height = height;
  }
  
  area() {
    return this.width * this.height;
  }
}

// Can add new shapes without modifying existing code!
class Triangle extends Shape {
  constructor(base, height) {
    super();
    this.base = base;
    this.height = height;
  }
  
  area() {
    return (this.base * this.height) / 2;
  }
}`),
			calloutBlock("warning", "OCP doesn't mean you can never modify code. It means you shouldn't have to modify stable, tested code to add new features."),
			exerciseBlock(
				"You have a PaymentProcessor class with if/else statements for different payment methods (credit card, PayPal). How would you refactor this to follow OCP?",
				"Create a PaymentMethod interface/abstract class with a process() method. Each payment method (CreditCard, PayPal, etc.) implements this interface. The PaymentProcessor can then work with any PaymentMethod without knowing the specific type, and new payment methods can be added without modifying existing code.",
				[]string{"Think about polymorphism", "How can you make it extensible without changing existing code?"},
			),
		},
	}
}
