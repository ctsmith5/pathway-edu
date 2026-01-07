package seed

import "github.com/pathway/backend/models"

func modulePatterns7() models.Module {
	return models.Module{
		ID:    "patterns-7",
		Title: "Choosing the Right Pattern",
		Content: []models.ContentBlock{
			textBlock(`## Pattern Selection Guide

Choosing the right pattern depends on your specific problem. Here's a guide to help you decide.`),
			textBlock(`## Creational Patterns

**Use Singleton when**:
- Need exactly one instance
- Global access required
- Lazy initialization needed
- ÃƒÂ¢Ã…Â¡Ã‚Â ÃƒÂ¯Ã‚Â¸Ã‚Â Use sparingly - can make testing difficult

**Use Factory when**:
- Don't know exact type at compile time
- Want to decouple creation from usage
- Need to create families of objects

**Use Builder when**:
- Object has many optional parameters
- Want to avoid telescoping constructors
- Need step-by-step construction`),
			textBlock(`## Structural Patterns

**Use Adapter when**:
- Need to make incompatible interfaces work together
- Integrating legacy code
- Wrapping third-party libraries

**Use Decorator when**:
- Need to add behavior at runtime
- Don't want to modify existing code
- Need flexible feature addition

**Use Facade when**:
- Want to simplify complex subsystem
- Need simple interface to complex system
- Want to decouple clients from subsystem`),
			textBlock(`## Behavioral Patterns

**Use Observer when**:
- One object change affects many
- Don't know how many objects need updates
- Want loose coupling

**Use Strategy when**:
- Have multiple ways to do something
- Want to switch algorithms at runtime
- Want to avoid conditional statements

**Use Command when**:
- Need to parameterize with operations
- Need undo/redo functionality
- Need to queue operations`),
			codeBlock("text", `Decision Tree:

Need to create objects?
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Need exactly one instance? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Singleton
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Don't know type? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Factory
  ÃƒÂ¢Ã¢â‚¬ÂÃ¢â‚¬ÂÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Many optional params? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Builder

Need to structure objects?
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Incompatible interfaces? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Adapter
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Add behavior dynamically? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Decorator
  ÃƒÂ¢Ã¢â‚¬ÂÃ¢â‚¬ÂÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Simplify complex system? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Facade

Need to handle interactions?
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ One-to-many notifications? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Observer
  ÃƒÂ¢Ã¢â‚¬ÂÃ…â€œÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Interchangeable algorithms? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Strategy
  ÃƒÂ¢Ã¢â‚¬ÂÃ¢â‚¬ÂÃƒÂ¢Ã¢â‚¬ÂÃ¢â€šÂ¬ Encapsulate operations? ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬â„¢ Command`),
			textBlock(`## Anti-Patterns

**Avoid these mistakes**:

ÃƒÂ¢Ã‚ÂÃ…â€™ **Using patterns everywhere**: Not every problem needs a pattern
ÃƒÂ¢Ã‚ÂÃ…â€™ **Forcing patterns**: Don't force a pattern if it doesn't fit
ÃƒÂ¢Ã‚ÂÃ…â€™ **Over-engineering**: Simple code is better than complex patterns
ÃƒÂ¢Ã‚ÂÃ…â€™ **Pattern obsession**: Patterns are tools, not goals

**Remember**:
- Patterns solve specific problems
- Start simple, refactor when needed
- Understand the problem before choosing a pattern
- Patterns should make code clearer, not more complex`),
			calloutBlock("warning", "The best pattern is often no pattern. Don't add complexity unless it solves a real problem. Start simple and refactor when patterns become necessary."),
			textBlock(`## When NOT to Use Patterns

**Skip patterns when**:
- Problem is simple and straightforward
- Pattern adds unnecessary complexity
- Team doesn't understand the pattern
- Premature optimization
- Over-engineering for future needs that may never come`),
			exerciseBlock(
				"You have a simple CRUD application with one database. Should you use the Repository pattern?",
				"It depends on your context:\n\nUse Repository if:\n- You plan to add more data sources\n- You want to test business logic easily\n- You're building a larger application\n- You want clean architecture\n\nSkip Repository if:\n- It's a simple, one-off project\n- You won't need to test business logic separately\n- You won't change data sources\n- It adds unnecessary complexity\n\nGenerally, for a simple CRUD app, you might skip it initially and add it later if needed. But if you're building something that will grow, starting with Repository is a good idea.",
				[]string{"Think about future needs", "What's the cost vs. benefit?"},
			),
		},
	}
}
