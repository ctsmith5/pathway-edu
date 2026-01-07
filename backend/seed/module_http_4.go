package seed

import "github.com/pathway/backend/models"

func moduleHTTP4() models.Module {
	return models.Module{
		ID:    "http-4",
		Title: "Headers & Cookies",
		Content: []models.ContentBlock{
			textBlock(`## HTTP Headers

Headers are key-value pairs that provide metadata about requests and responses.

**Types**:
- **Request headers**: Sent by client
- **Response headers**: Sent by server
- **General headers**: Apply to both
- **Entity headers**: Describe the body`),
			textBlock(`## Common Request Headers

**Host**: 
- Domain name of server
- Required in HTTP/1.1

**User-Agent**:
- Client application info
- Browser, version, OS

**Accept**:
- What content types client accepts
- Accept: application/json

**Authorization**:
- Credentials for authentication
- Authorization: Bearer <token>

**Content-Type**:
- Type of data in body
- Content-Type: application/json

**Cookie**:
- Cookies to send
- Cookie: sessionId=abc123`),
			textBlock(`## Common Response Headers

**Content-Type**:
- Type of data in body
- Content-Type: application/json

**Set-Cookie**:
- Instructs client to set cookie
- Set-Cookie: sessionId=abc123; Path=/; HttpOnly

**Cache-Control**:
- Caching directives
- Cache-Control: no-cache

**Location**:
- Redirect URL
- Used with 3xx status codes

**ETag**:
- Entity tag for caching
- ETag: "abc123"`),
			textBlock(`## What are Cookies?

Cookies are small pieces of data stored by the browser and sent with requests.

**Uses**:
- Session management
- Personalization
- Tracking
- Authentication tokens`),
			codeBlock("javascript", `// Setting cookies (server-side example)

// Set-Cookie header
Set-Cookie: sessionId=abc123; Path=/; HttpOnly; Secure; SameSite=Strict

// Cookie attributes:
// - Path: Where cookie is valid
// - Domain: Which domain can access
// - HttpOnly: Not accessible via JavaScript
// - Secure: Only sent over HTTPS
// - SameSite: CSRF protection
// - Max-Age/Expires: When cookie expires`),
			textBlock(`## Cookie Security

**HttpOnly**:
- Prevents JavaScript access
- Protects against XSS attacks

**Secure**:
- Only sent over HTTPS
- Protects against man-in-the-middle

**SameSite**:
- Strict: Never sent cross-site
- Lax: Sent for top-level navigation
- None: Always sent (requires Secure)

**Domain**:
- Restricts which domains can access
- Don't set for other domains!`),
			textBlock(`## Alternatives to Cookies

**Local Storage**:
- Stored in browser
- Accessible via JavaScript
- Not sent automatically
- Larger capacity

**Session Storage**:
- Like Local Storage
- Cleared when tab closes

**JWT Tokens**:
- Stored in Local Storage or memory
- Sent in Authorization header
- Stateless authentication`),
			calloutBlock("tip", "Use HttpOnly cookies for sensitive data like session IDs. Use Local Storage for non-sensitive data that needs JavaScript access."),
			exerciseBlock(
				"Why should authentication tokens be stored in HttpOnly cookies rather than Local Storage?",
				"HttpOnly cookies are not accessible to JavaScript, which protects them from XSS (Cross-Site Scripting) attacks. If an attacker injects malicious JavaScript, they can't steal tokens from HttpOnly cookies. Local Storage is accessible to JavaScript, making it vulnerable to XSS attacks.",
				[]string{"Think about XSS attacks", "What can JavaScript access?"},
			),
		},
	}
}
