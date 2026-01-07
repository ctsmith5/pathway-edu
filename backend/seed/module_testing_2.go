package seed

import "github.com/pathway/backend/models"

func moduleTesting2() models.Module {
	return models.Module{
		ID:    "testing-2",
		Title: "Unit Testing",
		Content: []models.ContentBlock{
			textBlock(`## What is Unit Testing?

Unit testing tests individual units of code (functions, methods, classes) in isolation.

**Characteristics**:
- Tests one thing at a time
- Fast execution
- No external dependencies
- Isolated from other code`),
			textBlock(`## Writing Unit Tests

**Structure** (AAA Pattern):

1. **Arrange**: Set up test data and conditions
2. **Act**: Execute the code being tested
3. **Assert**: Verify the results`),
			codeBlock("javascript", `// Example: Unit test

function add(a, b) {
  return a + b;
}

// Test
describe('add function', () => {
  test('adds two positive numbers', () => {
    // Arrange
    const a = 2;
    const b = 3;
    
    // Act
    const result = add(a, b);
    
    // Assert
    expect(result).toBe(5);
  });
  
  test('adds negative numbers', () => {
    // Arrange
    const a = -2;
    const b = -3;
    
    // Act
    const result = add(a, b);
    
    // Assert
    expect(result).toBe(-5);
  });
});`),
			textBlock(`## Test Cases

**Good test cases cover**:
- **Happy path**: Normal, expected behavior
- **Edge cases**: Boundary conditions
- **Error cases**: Invalid inputs
- **Null/empty**: Missing or empty data`),
			codeBlock("javascript", `// Comprehensive unit test

function divide(a, b) {
  if (b === 0) {
    throw new Error('Division by zero');
  }
  return a / b;
}

describe('divide function', () => {
  test('divides positive numbers', () => {
    expect(divide(10, 2)).toBe(5);
  });
  
  test('divides negative numbers', () => {
    expect(divide(-10, -2)).toBe(5);
  });
  
  test('handles decimal results', () => {
    expect(divide(1, 3)).toBeCloseTo(0.333, 3);
  });
  
  test('throws error on division by zero', () => {
    expect(() => divide(10, 0)).toThrow('Division by zero');
  });
  
  test('handles zero numerator', () => {
    expect(divide(0, 5)).toBe(0);
  });
});`),
			textBlock(`## Mocking and Stubbing

**Mocks**: Replace dependencies with fake implementations

**Why mock?**:
- Isolate unit under test
- Control dependencies
- Speed up tests
- Test error conditions`),
			codeBlock("javascript", `// Example with mocking

class UserService {
  constructor(userRepository) {
    this.repository = userRepository;
  }
  
  async getUser(id) {
    if (!id) {
      throw new Error('ID required');
    }
    return await this.repository.findById(id);
  }
}

// Test with mock
describe('UserService', () => {
  test('gets user by id', async () => {
    // Arrange - create mock
    const mockRepo = {
      findById: jest.fn().mockResolvedValue({ id: 1, name: 'Alice' })
    };
    const service = new UserService(mockRepo);
    
    // Act
    const user = await service.getUser(1);
    
    // Assert
    expect(user).toEqual({ id: 1, name: 'Alice' });
    expect(mockRepo.findById).toHaveBeenCalledWith(1);
  });
  
  test('throws error when id is missing', async () => {
    const mockRepo = { findById: jest.fn() };
    const service = new UserService(mockRepo);
    
    await expect(service.getUser(null)).rejects.toThrow('ID required');
    expect(mockRepo.findById).not.toHaveBeenCalled();
  });
});`),
			calloutBlock("tip", "Test one thing per test. If a test fails, you should immediately know what's wrong."),
			exerciseBlock(
				"What makes a good unit test?",
				"A good unit test:\n1. Tests one thing (single responsibility)\n2. Is fast (runs in milliseconds)\n3. Is independent (doesn't rely on other tests)\n4. Is repeatable (same result every time)\n5. Is readable (clear what it's testing)\n6. Uses AAA pattern (Arrange, Act, Assert)\n7. Has a descriptive name\n8. Tests behavior, not implementation\n9. Uses mocks for external dependencies\n10. Covers happy path, edge cases, and errors",
				[]string{"Think about what makes tests useful", "What happens when a test fails?"},
			),
		},
	}
}
