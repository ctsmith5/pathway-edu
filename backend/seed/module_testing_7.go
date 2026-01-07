package seed

import "github.com/pathway/backend/models"

func moduleTesting7() models.Module {
	return models.Module{
		ID:    "testing-7",
		Title: "Continuous Testing",
		Content: []models.ContentBlock{
			textBlock(`## Continuous Testing

Continuous testing means running tests automatically and frequently throughout development.

**Benefits**:
- Catch bugs immediately
- Prevent regressions
- Enable confident refactoring
- Faster feedback loop`),
			textBlock(`## Test Automation

**Run tests**:
- Before committing code
- On every push
- In CI/CD pipeline
- On schedule

**Tools**:
- Jest, Mocha, JUnit (test runners)
- GitHub Actions, Jenkins (CI/CD)
- Pre-commit hooks`),
			codeBlock("yaml", `# GitHub Actions example

name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Setup Node.js
        uses: actions/setup-node@v2
        with:
          node-version: '18'
      - name: Install dependencies
        run: npm install
      - name: Run tests
        run: npm test`),
			textBlock(`## Test Coverage

**Coverage metrics**:
- **Line coverage**: % of lines executed
- **Branch coverage**: % of branches tested
- **Function coverage**: % of functions called

**Aim for**:
- 80%+ coverage (good target)
- 100% coverage (ideal, but not always practical)
- Focus on critical paths`),
			calloutBlock("info", "High coverage doesn't mean good tests. It's possible to have 100% coverage with useless tests. Focus on testing the right things."),
			textBlock(`## Test Maintenance

**Keep tests updated**:
- Update tests when code changes
- Remove obsolete tests
- Refactor test code
- Keep tests fast`),
			textBlock(`## Testing in CI/CD

**Continuous Integration**:
- Run tests on every commit
- Catch issues early
- Prevent broken code in main branch

**Continuous Deployment**:
- Deploy only if tests pass
- Automated testing gates
- Confidence in releases`),
			codeBlock("text", `CI/CD Pipeline:

Code Commit
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Run Tests (Unit)
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Run Tests (Integration)
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Build Application
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Deploy to Staging
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Run E2E Tests
    ÃƒÂ¢Ã¢â‚¬Â Ã¢â‚¬Å“
Deploy to Production`),
			textBlock(`## Test Strategy

**Develop**:
- Write tests as you code
- Run tests frequently
- Fix failures immediately

**Before Commit**:
- Run full test suite
- Fix all failures
- Ensure good coverage

**In CI/CD**:
- Automated test execution
- Block deployment on failures
- Generate coverage reports`),
			calloutBlock("tip", "Make tests part of your workflow. Run them frequently, fix failures immediately, and keep them fast."),
			exerciseBlock(
				"Why is continuous testing important, and how does it help development?",
				"Continuous testing is important because:\n1. Immediate feedback - catch bugs as soon as you write them\n2. Prevents regressions - know immediately if you broke something\n3. Enables refactoring - safe to improve code with test safety net\n4. Faster development - less time debugging later\n5. Better code quality - tests force better design\n6. Confidence - know your code works\n\nIt helps development by:\n- Reducing debugging time\n- Making code changes safer\n- Providing documentation\n- Enabling faster iteration\n- Catching integration issues early",
				[]string{"Think about the feedback loop", "What happens without continuous testing?"},
			),
		},
	}
}
