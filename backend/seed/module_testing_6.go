package seed

import "github.com/pathway/backend/models"

func moduleTesting6() models.Module {
	return models.Module{
		ID:    "testing-6",
		Title: "Testing Best Practices",
		Content: []models.ContentBlock{
			textBlock(`## Testing Best Practices

Following best practices makes your tests more valuable and maintainable.`),
			textBlock(`## Test Organization

**Structure**:
- One test file per source file
- Group related tests
- Use descriptive names
- Follow AAA pattern (Arrange, Act, Assert)`),
			codeBlock("javascript", `// Good test organization

describe('UserService', () => {
  describe('createUser', () => {
    test('creates user with valid data', () => {
      // Arrange
      const service = new UserService(mockRepo);
      const userData = { name: 'Alice', email: 'alice@example.com' };
      
      // Act
      const user = service.createUser(userData);
      
      // Assert
      expect(user.id).toBeTruthy();
      expect(user.name).toBe('Alice');
    });
    
    test('throws error when email is missing', () => {
      // Test error case
    });
  });
  
  describe('getUser', () => {
    // Tests for getUser
  });
});`),
			textBlock(`## Test Naming

**Good test names**:
- Describe what is being tested
- Include expected behavior
- Be specific and clear

**Examples**:
- ÃƒÂ¢Ã…â€œÃ¢â‚¬Â¦ should return user when valid id provided
- ÃƒÂ¢Ã…â€œÃ¢â‚¬Â¦ throws error when email is invalid
- ÃƒÂ¢Ã‚ÂÃ…â€™ test1
- ÃƒÂ¢Ã‚ÂÃ…â€™ user test`),
			textBlock(`## Keep Tests Simple

**One assertion per test** (when possible):
- Easier to understand
- Clear what failed
- Focused on one behavior

**Avoid**:
- Complex setup
- Test interdependencies
- Testing multiple things`),
			textBlock(`## Test Independence

**Tests should**:
- Run in any order
- Not depend on other tests
- Be able to run alone
- Clean up after themselves`),
			codeBlock("javascript", `// Bad - tests depend on each other
let counter = 0;

test('increments counter', () => {
  counter++;
  expect(counter).toBe(1);
});

test('counter is 2', () => {
  expect(counter).toBe(2); // Depends on previous test!
});

// Good - independent tests
test('increments counter', () => {
  let counter = 0;
  counter++;
  expect(counter).toBe(1);
});

test('counter starts at 0', () => {
  let counter = 0;
  expect(counter).toBe(0);
});`),
			textBlock(`## Test Data

**Use factories**:
- Create test data easily
- Consistent test data
- Easy to modify

**Avoid**:
- Hard-coded test data everywhere
- Shared mutable state
- Production data`),
			codeBlock("javascript", `// Test data factory

function createUser(overrides = {}) {
  return {
    id: 1,
    name: 'Alice',
    email: 'alice@example.com',
    ...overrides
  };
}

test('creates user', () => {
  const user = createUser({ name: 'Bob' });
  expect(user.name).toBe('Bob');
  expect(user.email).toBe('alice@example.com'); // Default
});`),
			calloutBlock("tip", "Good tests are like good code - readable, maintainable, and focused. Treat test code with the same care as production code."),
			exerciseBlock(
				"What makes a test maintainable and easy to understand?",
				"A maintainable test:\n1. Has a clear, descriptive name\n2. Tests one thing\n3. Uses AAA pattern (Arrange, Act, Assert)\n4. Is independent (doesn't rely on other tests)\n5. Has minimal setup\n6. Uses factories for test data\n7. Is well-organized (grouped logically)\n8. Has clear assertions\n9. Doesn't have complex logic\n10. Is readable - someone else can understand it\n\nMaintainable tests are easy to update when requirements change and easy to debug when they fail.",
				[]string{"Think about what makes code maintainable", "What happens when a test fails?"},
			),
		},
	}
}
