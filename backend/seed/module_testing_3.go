package seed

import "github.com/pathway/backend/models"

func moduleTesting3() models.Module {
	return models.Module{
		ID:    "testing-3",
		Title: "Integration Testing",
		Content: []models.ContentBlock{
			textBlock(`## What is Integration Testing?

Integration testing verifies that multiple components work together correctly.

**Characteristics**:
- Tests interactions between components
- May use real dependencies
- Slower than unit tests
- More realistic scenarios`),
			textBlock(`## Why Integration Tests?

**Unit tests** verify components in isolation.
**Integration tests** verify components work together.

**Use integration tests to**:
- Verify API endpoints work
- Test database interactions
- Check component integration
- Validate data flow`),
			codeBlock("javascript", `// Integration test example

// API endpoint
app.post('/api/users', async (req, res) => {
  const user = await userService.createUser(req.body);
  res.status(201).json(user);
});

// Integration test
describe('POST /api/users', () => {
  let db;
  
  beforeAll(async () => {
    // Set up test database
    db = await setupTestDatabase();
  });
  
  afterAll(async () => {
    // Clean up
    await teardownTestDatabase(db);
  });
  
  test('creates a new user', async () => {
    // Arrange
    const userData = {
      name: 'Alice',
      email: 'alice@example.com'
    };
    
    // Act
    const response = await request(app)
      .post('/api/users')
      .send(userData)
      .expect(201);
    
    // Assert
    expect(response.body).toHaveProperty('id');
    expect(response.body.name).toBe('Alice');
    
    // Verify in database
    const user = await db.users.findOne({ email: 'alice@example.com' });
    expect(user).toBeTruthy();
  });
  
  test('returns 400 for invalid data', async () => {
    const response = await request(app)
      .post('/api/users')
      .send({ name: 'Bob' }) // Missing email
      .expect(400);
    
    expect(response.body).toHaveProperty('error');
  });
});`),
			textBlock(`## Integration Test Types

**API Integration Tests**:
- Test HTTP endpoints
- Verify request/response
- Test authentication/authorization

**Database Integration Tests**:
- Test database operations
- Verify data persistence
- Test queries and transactions

**Service Integration Tests**:
- Test services working together
- Verify data flow
- Test error handling`),
			textBlock(`## Test Database

**Use separate test database**:
- Don't use production database!
- Reset between tests
- Use transactions that rollback
- Or use in-memory database`),
			codeBlock("javascript", `// Test database setup

// Using transactions (rollback after each test)
describe('UserRepository', () => {
  let transaction;
  
  beforeEach(async () => {
    transaction = await db.beginTransaction();
  });
  
  afterEach(async () => {
    await transaction.rollback();
  });
  
  test('saves user to database', async () => {
    const repo = new UserRepository(transaction);
    const user = await repo.save({ name: 'Alice' });
    
    expect(user.id).toBeTruthy();
    const found = await repo.findById(user.id);
    expect(found.name).toBe('Alice');
  });
});`),
			calloutBlock("warning", "Never use production database for tests! Always use a separate test database or in-memory database."),
			exerciseBlock(
				"What's the difference between unit tests and integration tests?",
				"Unit tests:\n- Test individual components in isolation\n- Use mocks for dependencies\n- Fast execution\n- Many tests\n- Focus on single function/class\n\nIntegration tests:\n- Test multiple components together\n- May use real dependencies (database, APIs)\n- Slower execution\n- Fewer tests\n- Focus on interactions between components\n\nUnit tests answer 'Does this function work?' Integration tests answer 'Do these components work together?'",
				[]string{"Think about isolation vs. interaction", "What dependencies do each use?"},
			),
		},
	}
}
