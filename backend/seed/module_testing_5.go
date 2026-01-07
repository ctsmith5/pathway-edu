package seed

import "github.com/pathway/backend/models"

func moduleTesting5() models.Module {
	return models.Module{
		ID:    "testing-5",
		Title: "Mocking & Stubbing",
		Content: []models.ContentBlock{
			textBlock(`## Mocking and Stubbing

**Mocking** and **Stubbing** are techniques to replace real dependencies with fake ones in tests.

**Why?**:
- Isolate code under test
- Control dependencies
- Speed up tests
- Test error conditions`),
			textBlock(`## Types of Test Doubles

**Dummy**: Placeholder, never used
**Fake**: Working implementation, simplified
**Stub**: Returns predefined responses
**Mock**: Verifies interactions
**Spy**: Records interactions`),
			codeBlock("javascript", `// Stub example

class EmailService {
  send(email, message) {
    // Sends real email
  }
}

// Stub - returns predefined value
const emailStub = {
  send: jest.fn().mockResolvedValue({ success: true })
};

// Use stub in test
test('sends welcome email', async () => {
  const userService = new UserService(emailStub);
  await userService.createUser({ email: 'alice@example.com' });
  
  expect(emailStub.send).toHaveBeenCalled();
});`),
			textBlock(`## Mocking External Services

**Mock HTTP requests**:
- Don't make real API calls in tests
- Control responses
- Test error scenarios`),
			codeBlock("javascript", `// Mock HTTP request

// Mock fetch
global.fetch = jest.fn();

test('fetches user data', async () => {
  // Arrange - set up mock response
  fetch.mockResolvedValue({
    ok: true,
    json: async () => ({ id: 1, name: 'Alice' })
  });
  
  // Act
  const user = await fetchUser(1);
  
  // Assert
  expect(user.name).toBe('Alice');
  expect(fetch).toHaveBeenCalledWith('/api/users/1');
});`),
			textBlock(`## Mocking Database

**Options**:
- In-memory database
- Mock repository
- Transaction rollback`),
			codeBlock("javascript", `// Mock database repository

class UserRepository {
  async findById(id) {
    // Real database call
  }
}

// Mock for testing
const mockRepo = {
  findById: jest.fn().mockResolvedValue({ id: 1, name: 'Alice' }),
  save: jest.fn().mockResolvedValue({ id: 1, name: 'Bob' })
};

test('gets user', async () => {
  const service = new UserService(mockRepo);
  const user = await service.getUser(1);
  
  expect(user.name).toBe('Alice');
  expect(mockRepo.findById).toHaveBeenCalledWith(1);
});`),
			textBlock(`## When to Mock

**Mock when**:
- Dependency is slow (database, API)
- Dependency is unreliable
- Testing error conditions
- Dependency doesn't exist yet

**Don't mock when**:
- Dependency is fast and reliable
- Testing the integration itself
- Mock adds more complexity than value`),
			calloutBlock("warning", "Don't over-mock! Mocking everything makes tests less valuable. Sometimes you want to test the real integration."),
			exerciseBlock(
				"When should you mock a dependency in tests?",
				"You should mock a dependency when:\n1. It's slow (database, network calls) - speeds up tests\n2. It's unreliable (external APIs) - makes tests deterministic\n3. You're testing error conditions - can simulate failures\n4. It doesn't exist yet - allows TDD\n5. It has side effects (sends emails, charges credit cards) - prevents real actions\n6. You want to isolate the unit under test\n\nYou shouldn't mock when:\n- The dependency is fast and simple\n- You're writing integration tests\n- The mock adds more complexity than value\n- You want to test the real integration",
				[]string{"Think about test speed and reliability", "What are you actually testing?"},
			),
		},
	}
}
