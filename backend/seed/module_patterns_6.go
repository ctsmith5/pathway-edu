package seed

import "github.com/pathway/backend/models"

func modulePatterns6() models.Module {
	return models.Module{
		ID:    "patterns-6",
		Title: "Repository Pattern",
		Content: []models.ContentBlock{
			textBlock(`## Repository Pattern

The Repository pattern provides an abstraction layer between business logic and data access.

**Purpose**:
- Encapsulate data access logic
- Provide collection-like interface
- Decouple business logic from data source`),
			textBlock(`## Why Use Repository?

**Benefits**:
- **Abstraction**: Business logic doesn't know about database
- **Testability**: Easy to mock data access
- **Flexibility**: Can switch data sources easily
- **Separation**: Clear boundary between layers

**Use when**:
- Need to abstract data access
- Want to test business logic independently
- May need to change data source
- Want clean separation of concerns`),
			codeBlock("javascript", `// Repository Pattern

// Interface
class UserRepository {
  findById(id) {
    throw new Error('Must implement');
  }
  
  findAll() {
    throw new Error('Must implement');
  }
  
  save(user) {
    throw new Error('Must implement');
  }
  
  delete(id) {
    throw new Error('Must implement');
  }
}

// Implementation
class MySQLUserRepository extends UserRepository {
  constructor(db) {
    super();
    this.db = db;
  }
  
  findById(id) {
    return this.db.query('SELECT * FROM users WHERE id = ?', [id]);
  }
  
  findAll() {
    return this.db.query('SELECT * FROM users');
  }
  
  save(user) {
    if (user.id) {
      return this.db.query('UPDATE users SET ? WHERE id = ?', [user, user.id]);
    } else {
      return this.db.query('INSERT INTO users SET ?', [user]);
    }
  }
  
  delete(id) {
    return this.db.query('DELETE FROM users WHERE id = ?', [id]);
  }
}

// In-memory implementation for testing
class InMemoryUserRepository extends UserRepository {
  constructor() {
    super();
    this.users = [];
  }
  
  findById(id) {
    return Promise.resolve(this.users.find(u => u.id === id));
  }
  
  findAll() {
    return Promise.resolve(this.users);
  }
  
  save(user) {
    if (user.id) {
      const index = this.users.findIndex(u => u.id === user.id);
      this.users[index] = user;
    } else {
      user.id = Date.now();
      this.users.push(user);
    }
    return Promise.resolve(user);
  }
  
  delete(id) {
    this.users = this.users.filter(u => u.id !== id);
    return Promise.resolve();
  }
}`),
			textBlock(`## Using Repository

**Business Logic**:
- Uses repository interface
- Doesn't know about database
- Easy to test with mock repository`),
			codeBlock("javascript", `// Business logic using repository

class UserService {
  constructor(userRepository) {
    this.repository = userRepository;
  }
  
  async getUser(id) {
    return await this.repository.findById(id);
  }
  
  async getAllUsers() {
    return await this.repository.findAll();
  }
  
  async createUser(userData) {
    // Business logic
    if (!userData.email) {
      throw new Error('Email required');
    }
    
    return await this.repository.save(userData);
  }
}

// Usage with real database
const mysqlRepo = new MySQLUserRepository(db);
const userService = new UserService(mysqlRepo);

// Usage in tests with mock
const mockRepo = new InMemoryUserRepository();
const testService = new UserService(mockRepo);`),
			calloutBlock("tip", "Repository pattern is essential for testability. You can easily swap real database with in-memory implementation for testing."),
			textBlock(`## Repository vs DAO

**Repository**:
- Collection-like interface
- Domain-centric
- Business logic aware
- Can aggregate multiple data sources

**DAO (Data Access Object)**:
- Table-centric
- Closer to database
- One DAO per table
- Less business logic`),
			exerciseBlock(
				"What are the main benefits of the Repository pattern, and when would you use it?",
				"Main benefits:\n1. Abstraction - business logic doesn't depend on data source\n2. Testability - easy to mock for unit tests\n3. Flexibility - can switch databases without changing business logic\n4. Separation - clear boundary between data access and business logic\n\nUse it when:\n- You need to abstract data access\n- You want to test business logic independently\n- You might need to change data sources\n- You want clean architecture with clear layers",
				[]string{"Think about testability", "What problems does it solve?"},
			),
		},
	}
}
