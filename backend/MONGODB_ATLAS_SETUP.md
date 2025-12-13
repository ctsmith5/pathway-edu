# MongoDB Atlas Setup Guide

## Step-by-Step Instructions

### 1. Create Account
- Go to https://www.mongodb.com/cloud/atlas/register
- Sign up with email or Google account
- Verify your email

### 2. Create Free Cluster
- Click "Build a Database"
- Choose **M0 FREE** tier (512MB storage, shared)
- Select a cloud provider (AWS recommended)
- Choose a region closest to you
- Click "Create"

### 3. Create Database User
- Go to "Database Access" in the left sidebar
- Click "Add New Database User"
- Choose "Password" authentication
- Username: Create a secure username (e.g., `pathway-admin`)
- Password: Generate a strong password (click "Autogenerate Secure Password" or create your own)
- **Save the password!** You'll need it for the connection string
- Set user privileges: "Atlas admin" (for development) or "Read and write to any database"
- Click "Add User"

### 4. Configure Network Access
- Go to "Network Access" in the left sidebar
- Click "Add IP Address"
- For development/testing: Click "Allow Access from Anywhere" (adds `0.0.0.0/0`)
  - **Note**: For production, restrict to specific IPs (Railway IPs, Vercel IPs, etc.)
- Click "Confirm"

### 5. Get Connection String
- Go to "Database" → "Connect"
- Choose "Connect your application"
- Driver: **Go** (or Node.js, doesn't matter for connection string)
- Version: Latest
- Copy the connection string
- It looks like: `mongodb+srv://<username>:<password>@cluster0.xxxxx.mongodb.net/?retryWrites=true&w=majority`
- Replace `<username>` and `<password>` with your database user credentials
- Add database name: `mongodb+srv://username:password@cluster0.xxxxx.mongodb.net/pathway?retryWrites=true&w=majority`

### 6. Test Connection Locally
Update your local `.env` file:
```
MONGO_URI=mongodb+srv://username:password@cluster0.xxxxx.mongodb.net/pathway?retryWrites=true&w=majority
DB_NAME=pathway
```

Then test:
```bash
cd pathway/backend
go run main.go
```

If it connects successfully, you'll see: "Connected to MongoDB!"

## Security Notes

- **Never commit** your connection string to Git
- Use environment variables for all secrets
- In production, restrict Network Access to specific IPs
- Use strong passwords for database users
- Enable MongoDB Atlas monitoring and alerts

## Troubleshooting

**Connection timeout:**
- Check Network Access settings (must allow your IP)
- Verify username/password are correct
- Check if cluster is running (not paused)

**Authentication failed:**
- Verify username and password in connection string
- Check user privileges in Database Access

**Database not found:**
- MongoDB Atlas creates databases automatically on first write
- The database name in connection string is just a preference
- Your seed script will create the collections

