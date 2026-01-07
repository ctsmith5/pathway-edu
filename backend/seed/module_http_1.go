package seed

import "github.com/pathway/backend/models"

func moduleHTTP1() models.Module {
	return models.Module{
		ID:    "http-1",
		Title: "HTTP Fundamentals",
		Content: []models.ContentBlock{
			textBlock(`## What is HTTP?

HTTP (HyperText Transfer Protocol) is the foundation of data communication on the World Wide Web.

**Key characteristics**:
- **Stateless**: Each request is independent
- **Request-Response**: Client sends request, server responds
- **Application layer**: Works on top of TCP/IP
- **Text-based**: Human-readable protocol`),
			calloutBlock("info", "HTTP was created by Tim Berners-Lee in 1989. HTTP/1.1 (1997) is still widely used, but HTTP/2 (2015) and HTTP/3 (2020) offer improvements."),
			textBlock(`## How HTTP Works

1. **Client** (browser/app) sends HTTP request
2. **Server** receives and processes request
3. **Server** sends HTTP response
4. **Connection** may close (HTTP/1.1) or stay open (HTTP/2+)`),
			codeBlock("text", `Example HTTP Request:

GET /api/users HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0
Accept: application/json

Example HTTP Response:

HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 156

{
  "users": [
    {"id": 1, "name": "Alice"},
    {"id": 2, "name": "Bob"}
  ]
}`),
			textBlock(`## HTTP Components

**Request Line**:
- Method (GET, POST, etc.)
- Path (/api/users)
- HTTP Version

**Headers**:
- Metadata about request/response
- Key-value pairs
- Examples: Content-Type, Authorization

**Body** (optional):
- Data sent with request
- Used in POST, PUT, PATCH`),
			textBlock(`## URLs and URIs

**URL** (Uniform Resource Locator):
- Full address: https://example.com/api/users?id=123
- Includes protocol, domain, path, query

**URI** (Uniform Resource Identifier):
- More general term
- URL is a type of URI

**Components**:
- **Scheme**: http:// or https://
- **Host**: example.com
- **Path**: /api/users
- **Query**: ?id=123
- **Fragment**: #section`),
			exerciseBlock(
				"What does it mean that HTTP is stateless?",
				"HTTP is stateless means each request is independent - the server doesn't remember previous requests. Each request must contain all information needed to process it. This is why we use cookies, sessions, and tokens to maintain state across requests.",
				[]string{"Think about what 'state' means", "How do websites remember you're logged in?"},
			),
		},
	}
}
