package seed

import "github.com/pathway/backend/models"

func moduleSolid7() models.Module {
	return models.Module{
		ID:    "solid-7",
		Title: "Applying SOLID in Practice",
		Content: []models.ContentBlock{
			textBlock(`## Putting It All Together

SOLID principles work best when applied together. Let's see how they complement each other.`),
			textBlock(`## A Real-World Example

Consider an e-commerce order processing system:

- **SRP**: Separate classes for Order, Payment, Shipping, Notification
- **OCP**: Add new payment methods without modifying existing code
- **LSP**: Any PaymentMethod implementation works the same way
- **ISP**: Separate interfaces for Payable, Shippable, Notifiable
- **DIP**: OrderService depends on Payment interface, not concrete classes`),
			codeBlock("javascript", `// Example: SOLID e-commerce system
class Order {
  constructor(items, customer) {
    this.items = items;
    this.customer = customer;
  }
  
  total() {
    return this.items.reduce((sum, item) => sum + item.price, 0);
  }
}

// ISP: Segregated interfaces
class Payable {
  processPayment(amount) {}
}

class Shippable {
  ship(order) {}
}

// OCP: Open for extension
class CreditCardPayment extends Payable {
  processPayment(amount) {
    // Credit card logic
  }
}

class PayPalPayment extends Payable {
  processPayment(amount) {
    // PayPal logic
  }
}

// DIP: Depend on abstraction
class OrderService {
  constructor(paymentMethod, shippingService) {
    this.payment = paymentMethod; // Abstraction
    this.shipping = shippingService; // Abstraction
  }
  
  processOrder(order) {
    this.payment.processPayment(order.total());
    this.shipping.ship(order);
  }
}`),
			calloutBlock("warning", "Don't apply SOLID blindly. Start simple, and refactor when you see the need. Premature application can lead to over-engineering."),
			textBlock(`## When to Apply SOLID

**Apply SOLID when:**
- Code is becoming hard to change
- You're writing similar code multiple times
- Testing is difficult
- Adding features requires modifying many files

**Don't over-apply when:**
- The codebase is small and simple
- Requirements are still changing rapidly
- You're not sure what the future needs will be`),
			exerciseBlock(
				"Your codebase has a User class that handles authentication, sends emails, logs activity, and saves to database. How would you refactor this using SOLID principles?",
				"1. SRP: Split into User (data), AuthService, EmailService, Logger, UserRepository\n2. OCP: Make services extensible via interfaces\n3. LSP: Ensure all implementations follow their contracts\n4. ISP: Create focused interfaces (Authenticatable, Emailable, Loggable)\n5. DIP: UserService depends on these interfaces, receives them via injection",
				[]string{"Each principle addresses a different aspect", "Think about how they work together"},
			),
		},
	}
}
