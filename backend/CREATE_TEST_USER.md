# Create Test User

## Option 1: Via API (if backend is running)

If your backend is running locally or on Railway:

```powershell
# Register the test user
curl -X POST http://localhost:8080/api/auth/register `
  -H "Content-Type: application/json" `
  -d '{\"name\":\"Test User\",\"email\":\"test@example.com\",\"password\":\"test123456\"}'
```

Or if on Railway:
```powershell
curl -X POST https://your-railway-url.up.railway.app/api/auth/register `
  -H "Content-Type: application/json" `
  -d '{\"name\":\"Test User\",\"email\":\"test@example.com\",\"password\":\"test123456\"}'
```

## Option 2: Create User Script

Run this to create the user directly in the database:

```powershell
cd pathway/backend
$env:MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority"
$env:DB_NAME="pathway"
go run cmd/create-user/main.go
```







