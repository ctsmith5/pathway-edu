package seed

import "github.com/pathway/backend/models"

func moduleHTTP5() models.Module {
	return models.Module{
		ID:    "http-5",
		Title: "REST API Design",
		Content: []models.ContentBlock{
			textBlock(`## What is REST?

REST (Representational State Transfer) is an architectural style for designing web services.

**Key principles**:
- Stateless
- Client-server architecture
- Uniform interface
- Resource-based`),
			textBlock(`## RESTful Design

**Resources**:
- Everything is a resource
- Identified by URLs
- Nouns, not verbs

**Examples**:
- âœ… GET /api/users
- âœ… GET /api/users/123
- âœ… POST /api/users
- âŒ GET /api/getUser
- âŒ POST /api/createUser`),
			codeBlock("text", `RESTful URL Structure:

Collection:
GET    /api/users           # List all users
POST   /api/users           # Create user

Resource:
GET    /api/users/123       # Get user 123
PUT    /api/users/123       # Replace user 123
PATCH  /api/users/123       # Update user 123
DELETE /api/users/123       # Delete user 123

Sub-resources:
GET    /api/users/123/posts  # Get posts by user 123
POST   /api/users/123/posts  # Create post for user 123`),
			textBlock(`## REST Best Practices

**Use proper HTTP methods**:
- GET for reading
- POST for creating
- PUT/PATCH for updating
- DELETE for removing

**Use proper status codes**:
- 200 for success
- 201 for created
- 204 for no content
- 400 for client errors
- 500 for server errors

**Version your API**:
- /api/v1/users
- /api/v2/users`),
			textBlock(`## Request/Response Format

**Request**:
- Headers for metadata
- Body for data (JSON)
- Query params for filtering

**Response**:
- Consistent structure
- Error format
- Pagination metadata`),
			codeBlock("json", `// Example REST API Response

// Success response
{
  "data": {
    "id": 123,
    "name": "Alice",
    "email": "alice@example.com"
  }
}

// List response with pagination
{
  "data": [
    {"id": 1, "name": "Alice"},
    {"id": 2, "name": "Bob"}
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50
  }
}

// Error response
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Email is required",
    "details": {
      "email": "This field is required"
    }
  }
}`),
			textBlock(`## Filtering, Sorting, Pagination

**Filtering**:
GET /api/users?status=active&role=admin

**Sorting**:
GET /api/users?sort=name&order=asc

**Pagination**:
GET /api/users?page=1&limit=10

**Search**:
GET /api/users?search=alice`),
			calloutBlock("tip", "Keep URLs clean and intuitive. Use query parameters for optional things like filtering and pagination, not for required resource identification."),
			exerciseBlock(
				"Design RESTful endpoints for a blog API with posts and comments. Include endpoints for listing, creating, updating, and deleting.",
				"Posts:\n- GET /api/posts (list)\n- POST /api/posts (create)\n- GET /api/posts/123 (get one)\n- PUT /api/posts/123 (replace)\n- PATCH /api/posts/123 (update)\n- DELETE /api/posts/123 (delete)\n\nComments:\n- GET /api/posts/123/comments (list comments for post)\n- POST /api/posts/123/comments (create comment)\n- GET /api/comments/456 (get comment)\n- PATCH /api/comments/456 (update)\n- DELETE /api/comments/456 (delete)",
				[]string{"Think about resource hierarchy", "What are the nouns?"},
			),
		},
	}
}
