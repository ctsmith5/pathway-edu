package seed

import "github.com/pathway/backend/models"

func modulePatterns1() models.Module {
	return models.Module{
		ID:    "patterns-1",
		Title: "Introduction to Design Patterns",
		Content: []models.ContentBlock{
			textBlock(`## What are Design Patterns?

Design patterns are reusable solutions to common problems in software design. They're templates for solving problems that occur repeatedly.

**Key points**:
- Not code you can copy-paste
- General solutions to general problems
- Language-agnostic concepts
- Proven approaches used by experts`),
			calloutBlock("info", "Design patterns were popularized by the 'Gang of Four' (GoF) book 'Design Patterns: Elements of Reusable Object-Oriented Software' published in 1994."),
			textBlock(`## Why Use Design Patterns?

**Benefits**:
- **Proven solutions**: Tested approaches to common problems
- **Communication**: Shared vocabulary for developers
- **Best practices**: Encapsulate expert knowledge
- **Flexibility**: Make code more maintainable and extensible

**When to use**:
- Solving recurring problems
- Improving code structure
- Making code more maintainable
- Learning from expert solutions`),
			textBlock(`## Types of Design Patterns

**Creational Patterns**:
- Deal with object creation
- Examples: Singleton, Factory, Builder

**Structural Patterns**:
- Deal with object composition
- Examples: Adapter, Decorator, Facade

**Behavioral Patterns**:
- Deal with object interaction
- Examples: Observer, Strategy, Command`),
			calloutBlock("warning", "Don't force patterns into your code. Use them when they solve real problems. Over-engineering with patterns can make code harder to understand."),
			textBlock(`## Pattern Structure

Most patterns describe:

1. **Problem**: What problem does it solve?
2. **Solution**: How does it solve it?
3. **Consequences**: Trade-offs and benefits
4. **Implementation**: How to code it

**Pattern template**:
- Intent
- Motivation
- Applicability
- Structure
- Participants
- Collaborations
- Consequences`),
			exerciseBlock(
				"What is a design pattern, and how is it different from a library or framework?",
				"A design pattern is a general, reusable solution to a commonly occurring problem in software design. It's a template or blueprint, not code you can directly use.\n\nA library is code you can call, and a framework is code that calls you. Design patterns are conceptual solutions that you implement yourself, while libraries and frameworks are actual code you use.\n\nPatterns are language-agnostic concepts, while libraries and frameworks are specific implementations.",
				[]string{"Think about the difference between concepts and code", "What can you directly use vs. what you implement?"},
			),
		},
	}
}
