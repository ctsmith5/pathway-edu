package seed

import "github.com/pathway/backend/models"

func moduleCodeConcepts1() models.Module {
	return models.Module{
		ID:    "code-1",
		Title: "Object Oriented Programming",
		Content: []models.ContentBlock{
			textBlock(`## What is Object-Oriented Programming?

Object-Oriented Programming (OOP) is a programming paradigm based on the concept of "objects" that contain data and code.

### The Four Pillars of OOP

1. **Encapsulation**: Bundling data + methods together to make things easier and safer to use
2. **Inheritance**: Creating new classes based on existing ones
3. **Polymorphism**: Different types sharing a common “plug” (interface) so they can be swapped
4. **Abstraction**: Hiding complex implementation details behind a simple way to use something

---

## 1) Encapsulation (the “capsule” idea)

Encapsulation means we **bundle**:
- the **data** (variables / state)
- and the **methods** that use that data (functions)

…into one unit (a class), so it’s **easy to use** and harder to mess up by accident.

Think of it like a **medicine capsule**: you don’t need to know what’s inside or how it was made—you just take it the correct way.

Encapsulation is useful because:
- **It improves usability**: you get a simple “public” way to use something
- **It improves safety**: the important internal data can be protected so you can’t accidentally break it
- **It makes changes easier later**: you can change the inside without forcing everyone to change how they use it

In many languages, we do this by choosing what’s **public** (anyone can use) and what’s **private** (only the class can touch). That way, you interact with objects through a safe set of methods instead of directly changing important values.

---

## 2) Inheritance (reusing what already exists)

Inheritance means you can create a **new class** that automatically gets the **properties and methods** from a **base class** (also called a parent class).

This is awesome because it lets you reuse code instead of copy/pasting it.

### Quick example (C#)

Here’s a base class and a subclass. Notice how the subclass uses properties/methods from the base class **without re-writing them**:

` + "```csharp\n" + `// Base class (parent)
class Vehicle
{
    public string Brand { get; }
    public int MaxSpeed { get; }

    public Vehicle(string brand, int maxSpeed)
    {
        Brand = brand;
        MaxSpeed = maxSpeed;
    }

    public void Honk()
    {
        Console.WriteLine(Brand + " says: Beep beep!");
    }
}

// Subclass (child)
class Car : Vehicle
{
    public int Doors { get; }

    public Car(string brand, int maxSpeed, int doors) : base(brand, maxSpeed)
    {
        Doors = doors;
    }
}

// Using it
var myCar = new Car("Toyota", 180, 4);
Console.WriteLine(myCar.Brand);     // inherited from Vehicle
Console.WriteLine(myCar.MaxSpeed);  // inherited from Vehicle
myCar.Honk();                       // inherited method
` + "```\n" + `

### Vehicle analogy (easy mental model)

Imagine a parent class called **Vehicle**. A bunch of things can be vehicles:
- **Car**
- **Truck**
- **Bicycle**
- **Moped**
- **Segway**

They all share some “vehicle stuff” like:
- has a way to move
- can have a max speed
- might have a brand

But each one can also have unique features:
- a **Truck** might have ` + "`cargoCapacity`" + `
- a **Bicycle** might have ` + "`numberOfGears`" + `
- a **Car** might have ` + "`doors`" + `

Inheritance helps us organize that shared stuff in one place.

---

## 3) Polymorphism (same “plug”, different device)

Polymorphism means different types can be used through the **same interface**, even if they’re different on the inside.

### GPU / PCIe analogy

Most graphics cards plug into a computer using **PCIe**. You don’t need a totally new computer design for every graphics card brand:
- NVIDIA GPU
- AMD GPU
- (or a different model)

They all use the same “slot” (interface), so you can **swap** them without rewriting the whole computer.

In programming, an **interface** is like that PCIe slot:
- the interface defines the “shape” of what’s allowed (what methods exist)
- different classes can implement the same interface in their own way
- your code can talk to the interface and doesn’t care which exact class is behind it

---

## 4) Abstraction (use it without knowing every detail)

Abstraction means you focus on **what something does**, not **how it does it** internally.

### Coffee maker analogy

If you use a coffee maker, you know the inputs:
- water
- coffee grounds
- heat (electricity)

And you know the output:
- coffee

But you don’t need to understand:
- the exact temperature control system
- the tubing layout
- the internal timing logic

You just press “brew” and it works.

That’s abstraction in code:
- you get a simple way to use something
- the complicated details stay hidden

---

### Why this matters

These four ideas help you write code that is:
- easier to understand
- easier to reuse
- easier to test
- easier to change later (without breaking everything)`),
			calloutBlock("info", "OOP helps manage complexity in large codebases by organizing code into reusable, modular components."),
		},
	}
}
