package seed

import "github.com/pathway/backend/models"

func moduleHTTP6() models.Module {
	return models.Module{
		ID:    "http-6",
		Title: "Request/Response Cycle",
		Content: []models.ContentBlock{
			textBlock(`## The HTTP Request/Response Cycle

Understanding the complete flow of an HTTP request:

1. **Client** initiates request
2. **DNS** resolves domain to IP
3. **TCP** connection established
4. **TLS** handshake (if HTTPS)
5. **HTTP** request sent
6. **Server** processes request
7. **HTTP** response sent
8. **Connection** may close or stay open`),
			textBlock(`## Step-by-Step Flow

**1. Client Initiates Request**:
- User clicks link or JavaScript makes request
- Browser/app prepares HTTP request

**2. DNS Lookup**:
- Domain name â†’ IP address
- example.com â†’ 93.184.216.34

**3. TCP Connection**:
- Three-way handshake
- SYN â†’ SYN-ACK â†’ ACK
- Connection established`),
			codeBlock("text", `TCP Three-Way Handshake:

Client                Server
  |                     |
  |-------- SYN -------->|
  |                     |
  |<----- SYN-ACK -------|
  |                     |
  |-------- ACK -------->|
  |                     |
  |   Connection Open    |`),
			textBlock(`## HTTPS (TLS Handshake)

If using HTTPS:

**TLS Handshake**:
1. Client sends supported cipher suites
2. Server chooses cipher and sends certificate
3. Client verifies certificate
4. Shared secret established
5. Encrypted communication begins`),
			textBlock(`## HTTP Request Sent

**Request includes**:
- Request line (method, path, version)
- Headers
- Body (if applicable)

**Server receives**:
- Parses request
- Routes to handler
- Processes request`),
			textBlock(`## Server Processing

**Server**:
1. Receives request
2. Routes to appropriate handler
3. Validates request
4. Executes business logic
5. Queries database (if needed)
6. Generates response
7. Sends response`),
			textBlock(`## HTTP Response Sent

**Response includes**:
- Status line (version, status code, message)
- Headers
- Body (if applicable)

**Client receives**:
- Parses response
- Handles based on status code
- Renders/processes data`),
			codeBlock("javascript", `// Complete request/response example

// Client side
async function fetchUser(userId) {
  try {
    const response = await fetch('/api/users/' + userId, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Authorization': 'Bearer token123'
      }
    });
    
    if (!response.ok) {
      throw new Error('HTTP error! status: ' + response.status);
    }
    
    const user = await response.json();
    return user;
  } catch (error) {
    console.error('Request failed:', error);
    throw error;
  }
}

// Server side (pseudo-code)
function handleGetUser(request) {
  // 1. Validate request
  const userId = request.params.id;
  
  // 2. Authenticate
  const user = authenticate(request.headers.authorization);
  
  // 3. Authorize
  if (!canAccess(user, userId)) {
    return { status: 403, body: { error: 'Forbidden' } };
  }
  
  // 4. Query database
  const userData = database.getUser(userId);
  
  // 5. Generate response
  return {
    status: 200,
    headers: { 'Content-Type': 'application/json' },
    body: userData
  };
}`),
			calloutBlock("info", "Understanding the request/response cycle helps debug issues. Is it a DNS problem? Network issue? Server error? Client error?"),
			exerciseBlock(
				"What happens if a server takes 30 seconds to process a request, but the client times out after 10 seconds?",
				"The client will close the connection after 10 seconds, even though the server is still processing. The server may complete processing, but the response won't reach the client. This is why it's important to:\n1. Set appropriate timeout values\n2. Optimize server processing time\n3. Use async processing for long operations\n4. Provide progress updates for long-running tasks",
				[]string{"Think about what happens to the connection", "Can the server still send a response?"},
			),
		},
	}
}
