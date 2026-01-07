package seed

import "github.com/pathway/backend/models"

func moduleHTTP7() models.Module {
	return models.Module{
		ID:    "http-7",
		Title: "HTTPS & Security",
		Content: []models.ContentBlock{
			textBlock(`## What is HTTPS?

HTTPS (HTTP Secure) is HTTP over TLS/SSL encryption.

**Benefits**:
- **Encryption**: Data is encrypted in transit
- **Authentication**: Verifies server identity
- **Integrity**: Detects tampering
- **Privacy**: Prevents eavesdropping`),
			textBlock(`## Why HTTPS Matters

**Without HTTPS**:
- Data sent in plain text
- Anyone can intercept
- Man-in-the-middle attacks possible
- No server verification

**With HTTPS**:
- Data encrypted
- Only server can decrypt
- Certificate verifies identity
- Secure communication`),
			calloutBlock("warning", "Never send sensitive data (passwords, credit cards, personal info) over HTTP. Always use HTTPS!"),
			textBlock(`## TLS/SSL Handshake

**Process**:
1. Client initiates connection
2. Server sends certificate
3. Client verifies certificate
4. Shared secret established
5. Encrypted communication

**Certificate contains**:
- Domain name
- Public key
- Issuer (Certificate Authority)
- Expiration date`),
			textBlock(`## Certificate Authorities (CAs)

**Role**:
- Issue digital certificates
- Verify domain ownership
- Trusted by browsers

**Examples**:
- Let's Encrypt (free)
- DigiCert
- GlobalSign

**Self-signed certificates**:
- Not trusted by browsers
- OK for development
- Not for production`),
			textBlock(`## Common Security Headers

**Strict-Transport-Security (HSTS)**:
- Forces HTTPS
- Prevents downgrade attacks
- Strict-Transport-Security: max-age=31536000

**Content-Security-Policy (CSP)**:
- Prevents XSS attacks
- Controls resource loading
- Content-Security-Policy: default-src 'self'

**X-Frame-Options**:
- Prevents clickjacking
- X-Frame-Options: DENY

**X-Content-Type-Options**:
- Prevents MIME sniffing
- X-Content-Type-Options: nosniff`),
			textBlock(`## Common Vulnerabilities

**XSS (Cross-Site Scripting)**:
- Malicious scripts injected
- Prevent with: Input validation, output encoding, CSP

**CSRF (Cross-Site Request Forgery)**:
- Unauthorized actions from user's browser
- Prevent with: CSRF tokens, SameSite cookies

**SQL Injection**:
- Malicious SQL in inputs
- Prevent with: Parameterized queries, ORMs

**Man-in-the-Middle**:
- Intercepting communication
- Prevent with: HTTPS, certificate pinning`),
			codeBlock("javascript", `// Security best practices

// 1. Always use HTTPS in production
const API_URL = process.env.NODE_ENV === 'production' 
  ? 'https://api.example.com'
  : 'http://localhost:3000';

// 2. Validate and sanitize input
function sanitizeInput(input) {
  return input
    .replace(/<script>/gi, '')
    .trim()
    .escape();
}

// 3. Use parameterized queries (never concatenate!)
// âŒ BAD
const query = 'SELECT * FROM users WHERE id = ' + userId;

// âœ… GOOD
const query = 'SELECT * FROM users WHERE id = ?';
db.query(query, [userId]);

// 4. Set secure headers
app.use((req, res, next) => {
  res.setHeader('Strict-Transport-Security', 'max-age=31536000');
  res.setHeader('X-Content-Type-Options', 'nosniff');
  res.setHeader('X-Frame-Options', 'DENY');
  next();
});`),
			calloutBlock("tip", "Security is not optional. Always use HTTPS in production, validate all inputs, use parameterized queries, and set security headers."),
			exerciseBlock(
				"What's the difference between HTTP and HTTPS, and why should you always use HTTPS?",
				"HTTP sends data in plain text, while HTTPS encrypts data using TLS/SSL. HTTPS provides:\n1. Encryption - data can't be read if intercepted\n2. Authentication - verifies you're talking to the real server\n3. Integrity - detects if data was tampered with\n\nYou should always use HTTPS because:\n- Protects user data (passwords, personal info)\n- Prevents man-in-the-middle attacks\n- Required for modern web features\n- Builds user trust\n- Better SEO rankings",
				[]string{"Think about what happens to data in transit", "What can attackers do with unencrypted data?"},
			),
		},
	}
}
