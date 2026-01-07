package seed

import "github.com/pathway/backend/models"

func moduleHTTP3() models.Module {
	return models.Module{
		ID:    "http-3",
		Title: "HTTP Status Codes",
		Content: []models.ContentBlock{
			textBlock(`## HTTP Status Codes

Status codes are three-digit numbers that indicate the result of an HTTP request.

**Format**: XXX
- First digit: Category
- Last two digits: Specific code`),
			textBlock(`## Status Code Categories

**1xx - Informational**:
- Request received, continuing process
- Rarely used

**2xx - Success**:
- Request successful
- Most common: 200 OK

**3xx - Redirection**:
- Further action needed
- Common: 301, 302, 304

**4xx - Client Error**:
- Request has error
- Common: 400, 401, 403, 404

**5xx - Server Error**:
- Server failed
- Common: 500, 502, 503`),
			textBlock(`## Common 2xx Codes

**200 OK**:
- Request successful
- Most common success response

**201 Created**:
- Resource created successfully
- Used with POST
- Should include Location header

**204 No Content**:
- Success but no content to return
- Used with DELETE or PUT

**206 Partial Content**:
- Partial content returned
- Used for range requests`),
			textBlock(`## Common 3xx Codes

**301 Moved Permanently**:
- Resource moved permanently
- Browser caches redirect

**302 Found (Temporary Redirect)**:
- Resource temporarily at different location
- Browser doesn't cache

**304 Not Modified**:
- Resource hasn't changed
- Used for caching
- Client can use cached version`),
			textBlock(`## Common 4xx Codes

**400 Bad Request**:
- Request malformed
- Client error in request format

**401 Unauthorized**:
- Authentication required
- Missing or invalid credentials

**403 Forbidden**:
- Authenticated but not authorized
- Valid credentials but insufficient permissions

**404 Not Found**:
- Resource doesn't exist
- Most famous status code!

**409 Conflict**:
- Request conflicts with current state
- Used for duplicate resources

**422 Unprocessable Entity**:
- Request well-formed but semantically invalid
- Validation errors`),
			textBlock(`## Common 5xx Codes

**500 Internal Server Error**:
- Generic server error
- Something went wrong on server

**502 Bad Gateway**:
- Server acting as gateway got invalid response
- Upstream server issue

**503 Service Unavailable**:
- Server temporarily unavailable
- Overloaded or maintenance

**504 Gateway Timeout**:
- Gateway didn't receive timely response
- Upstream server timeout`),
			codeBlock("javascript", `// Handling status codes

fetch('/api/users')
  .then(response => {
    if (response.ok) {
      // 200-299 range
      return response.json();
    } else if (response.status === 404) {
      throw new Error('User not found');
    } else if (response.status === 401) {
      throw new Error('Unauthorized');
    } else {
      throw new Error('Request failed');
    }
  })
  .then(data => console.log(data))
  .catch(error => console.error(error));`),
			calloutBlock("warning", "Use status codes correctly. 401 for authentication issues, 403 for authorization issues. 400 for client errors, 500 for server errors."),
			exerciseBlock(
				"A user tries to access a resource they don't have permission for. They're logged in. What status code should you return?",
				"Return 403 Forbidden. The user is authenticated (logged in), so it's not 401. They simply don't have permission to access this resource, which is exactly what 403 means.",
				[]string{"What's the difference between 401 and 403?", "Is the user authenticated?"},
			),
		},
	}
}
