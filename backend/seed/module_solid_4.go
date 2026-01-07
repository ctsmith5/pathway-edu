package seed

import "github.com/pathway/backend/models"

func moduleSolid4() models.Module {
	return models.Module{
		ID:    "solid-4",
		Title: "Liskov Substitution Principle (LSP)",
		Content: []models.ContentBlock{
			textBlock(`## Liskov Substitution Principle

**Objects of a superclass should be replaceable with objects of its subclasses without breaking the application.**

Named after Barbara Liskov, this principle ensures that inheritance is used correctly.`),
			textBlock(`## The Contract

Subtypes must:
- Accept the same input parameters as their base type
- Return compatible types
- Not throw exceptions that the base type doesn't throw
- Maintain the same behavioral guarantees`),
			codeBlock("javascript", `// ❌ BAD: Violates LSP
class Rectangle {
  constructor(width, height) {
    this.width = width;
    this.height = height;
  }
  
  setWidth(width) {
    this.width = width;
  }
  
  setHeight(height) {
    this.height = height;
  }
  
  area() {
    return this.width * this.height;
  }
}

class Square extends Rectangle {
  constructor(side) {
    super(side, side);
  }
  
  setWidth(width) {
    this.width = width;
    this.height = width; // Problem!
  }
  
  setHeight(height) {
    this.width = height; // Problem!
    this.height = height;
  }
}

// This breaks when substituting Square for Rectangle
function testArea(rect) {
  rect.setWidth(5);
  rect.setHeight(4);
  return rect.area(); // Expects 20, but Square gives 16!
}

// ✅ GOOD: Don't force Square to be a Rectangle
class Shape {
  area() {
    throw new Error('Must implement');
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

class Square extends Shape {
  constructor(side) {
    super();
    this.side = side;
  }
  
  area() {
    return this.side ** 2;
  }
}`),
			calloutBlock("tip", "If you find yourself checking the type of a subclass before using it, you're likely violating LSP. The code should work with the base type."),
			exerciseBlock(
				"A Bird class has a fly() method. A Penguin class extends Bird but can't fly. How does this violate LSP and how would you fix it?",
				"This violates LSP because code expecting a Bird to fly would break with a Penguin. Solutions:\n1. Don't make Penguin extend Bird if it can't fulfill Bird's contract\n2. Create a FlyingBird subclass that Bird extends, and have Penguin extend Bird directly\n3. Use composition: have a Flyable interface that only flying birds implement",
				[]string{"Think about what guarantees Bird makes", "Can all birds fly?"},
			),
		},
	}
}
