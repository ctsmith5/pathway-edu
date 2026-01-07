package seed

import "github.com/pathway/backend/models"

func moduleHTTP2() models.Module {
	return models.Module{
		ID:    "http-2",
		Title: "HTTP Methods (GET, POST, PUT, DELETE, etc.)",
		Content: []models.ContentBlock{
			textBlock(`## HTTP Methods

HTTP methods (also called verbs) indicate the action to be performed on a resource.

**Common methods**:
- GET: Retrieve data
- POST: Create new resource
- PUT: Update/replace resource
- PATCH: Partial update
- DELETE: Remove resource
- HEAD: Get headers only
- OPTIONS: Get allowed methods`),
			textBlock(`## GET - Retrieve Data

**Purpose**: Fetch data from server

**Characteristics**:
- Idempotent (same request = same result)
- Safe (doesn't modify server state)
- Can be cached
- Parameters in URL query string

**Example**:
GET /api/users?page=1&limit=10`),
			textBlock(`## POST - Create Resource

**Purpose**: Submit data to create new resource

**Characteristics**:
- Not idempotent (multiple requests = multiple resources)
- Not safe (modifies server state)
- Data in request body
- Returns created resource

**Example**:
POST /api/users
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com"
}`),
			codeBlock("javascript", `// Example: Making HTTP requests

// GET request
fetch('/api/users')
  .then(response => response.json())
  .then(data => console.log(data));

// POST request
fetch('/api/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    name: 'Alice',
    email: 'alice@example.com'
  })
})
  .then(response => response.json())
  .then(data => console.log(data));`),
			textBlock(`## PUT - Replace Resource

**Purpose**: Replace entire resource

**Characteristics**:
- Idempotent (same request = same result)
- Replaces entire resource
- Must include all fields

**Example**:
PUT /api/users/123
Content-Type: application/json

{
  "name": "Alice Updated",
  "email": "alice.new@example.com"
}`),
			textBlock(`## PATCH - Partial Update

**Purpose**: Update part of a resource

**Characteristics**:
- Idempotent
- Only sends changed fields
- More efficient than PUT

**Example**:
PATCH /api/users/123
Content-Type: application/json

{
  "email": "alice.new@example.com"
}`),
			textBlock(`## DELETE - Remove Resource

**Purpose**: Delete a resource

**Characteristics**:
- Idempotent
- Removes resource
- May return 204 No Content

**Example**:
DELETE /api/users/123`),
			calloutBlock("tip", "Use the right method for the right purpose. GET for reading, POST for creating, PUT/PATCH for updating, DELETE for removing."),
			exerciseBlock(
				"What's the difference between PUT and PATCH?",
				"PUT replaces the entire resource - you must send all fields, even unchanged ones. PATCH performs a partial update - you only send the fields you want to change. PUT is idempotent (replacing with same data has same effect), and PATCH should also be idempotent.",
				[]string{"Think about what 'replace' vs 'update' means", "What data do you need to send?"},
			),
		},
	}
}
