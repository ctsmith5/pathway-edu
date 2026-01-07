package seed

import "github.com/pathway/backend/models"

func moduleTesting1() models.Module {
	return models.Module{
		ID:    "testing-1",
		Title: "Testing Fundamentals",
		Content: []models.ContentBlock{
			textBlock(`## Why Test Software?

Testing ensures your code works correctly and helps prevent bugs from reaching production.

**Benefits**:
- **Confidence**: Know your code works
- **Documentation**: Tests show how code should be used
- **Refactoring**: Safe to change code with tests
- **Debugging**: Tests help identify problems quickly
- **Quality**: Forces you to write better code`),
			calloutBlock("info", "Good tests are like a safety net - they give you confidence to make changes and catch problems before users do."),
			textBlock(`## Testing Pyramid

The testing pyramid shows the ideal distribution of tests:

**Unit Tests** (bottom, most):
- Test individual functions/classes
- Fast, isolated, many
- Foundation of testing

**Integration Tests** (middle):
- Test components working together
- Moderate speed, moderate number
- Verify interactions

**E2E Tests** (top, fewest):
- Test complete user flows
- Slow, fewer tests
- Verify system works end-to-end`),
			codeBlock("text", `Testing Pyramid:

        /\
       /  \     E2E Tests (few)
      /____\
     /      \    Integration Tests (some)
    /________\
   /          \  Unit Tests (many)
  /____________\`),
			textBlock(`## Types of Testing

**Unit Testing**:
- Test individual units in isolation
- Fast and focused
- Most common type

**Integration Testing**:
- Test multiple components together
- Verify they work correctly together

**End-to-End Testing**:
- Test complete user workflows
- Simulate real user scenarios

**Manual Testing**:
- Human tester performs actions
- Useful for UX and exploratory testing`),
			textBlock(`## Test Characteristics

**Good tests are**:
- **Fast**: Run quickly
- **Independent**: Don't depend on other tests
- **Repeatable**: Same results every time
- **Self-validating**: Pass or fail automatically
- **Timely**: Written before or with code`),
			calloutBlock("tip", "Follow the AAA pattern: Arrange (set up), Act (execute), Assert (verify). This makes tests clear and readable."),
			exerciseBlock(
				"Why is the testing pyramid shape (many unit tests, fewer integration tests, fewest E2E tests) recommended?",
				"The pyramid shape is recommended because:\n1. Unit tests are fast and cheap - you can have many\n2. Integration tests are slower - have fewer\n3. E2E tests are slowest and most brittle - have fewest\n\nThis maximizes test coverage and confidence while minimizing execution time. Having too many slow tests makes the test suite slow, which means developers run tests less often.",
				[]string{"Think about speed and cost", "What happens if you have too many slow tests?"},
			),
		},
	}
}
