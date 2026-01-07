package seed

import "github.com/pathway/backend/models"

func modulePatterns2() models.Module {
	return models.Module{
		ID:    "patterns-2",
		Title: "Creational Patterns (Singleton, Factory)",
		Content: []models.ContentBlock{
			textBlock(`## Creational Patterns

Creational patterns deal with object creation mechanisms, trying to create objects in a manner suitable to the situation.

**Common patterns**:
- Singleton
- Factory
- Builder
- Prototype
- Abstract Factory`),
			textBlock(`## Singleton Pattern

**Intent**: Ensure a class has only one instance and provide global access to it.

**Use when**:
- Exactly one instance is needed
- Instance must be accessible globally
- Lazy initialization is desired

**Example uses**:
- Database connections
- Logging
- Configuration managers`),
			codeBlock("javascript", `// Singleton Pattern

class DatabaseConnection {
  constructor() {
    if (DatabaseConnection.instance) {
      return DatabaseConnection.instance;
    }
    
    this.connection = this.connect();
    DatabaseConnection.instance = this;
  }
  
  connect() {
    // Connection logic
    return { status: 'connected' };
  }
  
  query(sql) {
    // Query logic
    return 'results';
  }
}

// Usage
const db1 = new DatabaseConnection();
const db2 = new DatabaseConnection();

console.log(db1 === db2); // true - same instance`),
			calloutBlock("warning", "Singleton can make testing difficult and create hidden dependencies. Use sparingly and consider dependency injection instead."),
			textBlock(`## Factory Pattern

**Intent**: Create objects without specifying the exact class of object that will be created.

**Use when**:
- Don't know exact types at compile time
- Want to decouple object creation from usage
- Need to create families of related objects

**Example uses**:
- UI component creation
- Database driver selection
- Payment method processing`),
			codeBlock("javascript", `// Factory Pattern

class PaymentMethod {
  processPayment(amount) {
    throw new Error('Must implement');
  }
}

class CreditCard extends PaymentMethod {
  processPayment(amount) {
    return 'Processing ' + amount + ' via credit card';
  }
}

class PayPal extends PaymentMethod {
  processPayment(amount) {
    return 'Processing ' + amount + ' via PayPal';
  }
}

class PaymentFactory {
  static create(type) {
    switch(type) {
      case 'creditcard':
        return new CreditCard();
      case 'paypal':
        return new PayPal();
      default:
        throw new Error('Unknown payment type');
    }
  }
}

// Usage
const payment = PaymentFactory.create('creditcard');
payment.processPayment(100);`),
			textBlock(`## Builder Pattern

**Intent**: Construct complex objects step by step.

**Use when**:
- Object has many optional parameters
- Want to avoid telescoping constructor
- Need different representations

**Example uses**:
- SQL query builders
- HTTP request builders
- Configuration objects`),
			codeBlock("javascript", `// Builder Pattern

class QueryBuilder {
  constructor() {
    this.query = {
      select: [],
      from: null,
      where: [],
      orderBy: null
    };
  }
  
  select(fields) {
    this.query.select = fields;
    return this;
  }
  
  from(table) {
    this.query.from = table;
    return this;
  }
  
  where(condition) {
    this.query.where.push(condition);
    return this;
  }
  
  orderBy(field) {
    this.query.orderBy = field;
    return this;
  }
  
  build() {
    return this.query;
  }
}

// Usage
const query = new QueryBuilder()
  .select(['name', 'email'])
  .from('users')
  .where('age > 18')
  .orderBy('name')
  .build();`),
			calloutBlock("tip", "Factory is great when you don't know the exact type. Builder is great when you have many optional parameters."),
			exerciseBlock(
				"When would you use Singleton vs Factory pattern?",
				"Use Singleton when you need exactly one instance of a class that should be globally accessible (like a database connection or logger).\n\nUse Factory when you need to create objects but don't know the exact type at compile time, or when you want to decouple object creation from usage (like creating different payment methods based on user choice).\n\nThey solve different problems: Singleton ensures one instance, Factory handles object creation logic.",
				[]string{"What problem does each solve?", "Think about when you'd need one instance vs. multiple types"},
			),
		},
	}
}
