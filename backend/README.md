# Pathway Backend API

Go backend API for the Pathway educational platform.

## Features

- RESTful API with Gin framework
- JWT authentication
- MongoDB database integration
- User progress tracking
- Course and module management
- Structured content blocks

## Local Development

### Prerequisites

- Go 1.21 or higher
- MongoDB (local or Atlas)
- Git

### Setup

1. **Clone repository**
   ```bash
   git clone <repo-url>
   cd pathway/backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment**
   Create a `.env` file:
   ```env
   MONGO_URI=mongodb://localhost:27017
   DB_NAME=pathway
   PORT=8080
   JWT_SECRET=your-secret-key-here
   ALLOWED_ORIGINS=*
   ```

4. **Start MongoDB** (if using local)
   ```bash
   # Windows (if installed as service)
   # MongoDB should start automatically
   
   # Or start manually
   mongod
   ```

5. **Seed database** (optional)
   ```bash
   go run cmd/seed/main.go
   ```

6. **Run server**
   ```bash
   go run main.go
   ```

   Server will start on `http://localhost:8080`

## API Endpoints

### Public Endpoints

- `GET /api/health` - Health check
- `GET /api/courses` - Get all courses
- `GET /api/courses/:id` - Get single course
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user

### Protected Endpoints (require JWT)

- `GET /api/user/me` - Get current user
- `GET /api/user/progress` - Get user's course progress

## Project Structure

```
backend/
├── cmd/
│   └── seed/          # Database seeding command
├── handlers/          # HTTP request handlers
├── middleware/        # Middleware (auth, CORS)
├── models/           # Data models
├── repository/       # Database access layer
├── seed/             # Seed data
├── main.go           # Application entry point
└── go.mod            # Go dependencies
```

## Environment Variables

See `.env.example` for all required environment variables.

## Deployment

See `DEPLOYMENT.md` for detailed deployment instructions.

Quick deploy to Railway:
1. Push code to GitHub
2. Connect Railway to repo
3. Set root directory to `pathway/backend`
4. Add environment variables
5. Deploy!

## Testing

```bash
# Health check
curl http://localhost:8080/api/health

# Register user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com","password":"test123"}'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123"}'
```

## License

MIT

