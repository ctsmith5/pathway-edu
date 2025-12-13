# Backend Deployment Guide

## Quick Start: Railway Deployment

### Prerequisites
1. MongoDB Atlas account and cluster (see `MONGODB_ATLAS_SETUP.md`)
2. GitHub repository with backend code
3. Railway account (https://railway.app)

### Step 1: Prepare MongoDB Atlas
Follow the instructions in `MONGODB_ATLAS_SETUP.md` to:
- Create a cluster
- Create a database user
- Configure network access
- Get your connection string

### Step 2: Deploy to Railway

1. **Sign up/Login**
   - Go to https://railway.app
   - Sign up with GitHub

2. **Create New Project**
   - Click "New Project"
   - Select "Deploy from GitHub repo"
   - Choose your repository

3. **Configure Service**
   - Railway will auto-detect Go
   - **Root Directory**: Set to `pathway/backend`
   - Railway will use the `railway.json` config file

4. **Set Environment Variables**
   Go to your service → Variables tab, add:
   
   ```
   MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/pathway?retryWrites=true&w=majority
   DB_NAME=pathway
   JWT_SECRET=<generate-with-openssl-rand-base64-32>
   ALLOWED_ORIGINS=https://your-frontend.vercel.app
   ```
   
   **Generate JWT Secret:**
   ```bash
   openssl rand -base64 32
   ```

5. **Deploy**
   - Railway will automatically build and deploy
   - Wait for deployment to complete
   - Get your deployment URL (e.g., `https://pathway-backend.railway.app`)

### Step 3: Seed Production Database

After deployment, you need to seed the database with course data.

**Option A: SSH into Railway and run seed**
```bash
# Railway provides a CLI or web terminal
cd /app
go run cmd/seed/main.go
```

**Option B: Create a one-time deployment script**
Railway can run a command on deploy. You can temporarily add:
```json
"deploy": {
  "startCommand": "./server && go run cmd/seed/main.go"
}
```
(Then remove the seed part after first run)

**Option C: Use Railway's one-off command**
- Go to your service in Railway
- Use the "Run Command" feature
- Run: `go run cmd/seed/main.go`

### Step 4: Test Deployment

Test your endpoints:
```bash
# Health check
curl https://your-backend.railway.app/api/health

# Get courses
curl https://your-backend.railway.app/api/courses

# Register user
curl -X POST https://your-backend.railway.app/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"test123"}'
```

### Step 5: Update Frontend

Update your Vercel frontend environment variable:
```
VITE_API_URL=https://your-backend.railway.app/api
```

## Environment Variables Reference

| Variable | Required | Description | Example |
|----------|----------|-------------|---------|
| `MONGO_URI` | Yes | MongoDB connection string | `mongodb+srv://...` |
| `DB_NAME` | No | Database name (default: `pathway`) | `pathway` |
| `PORT` | No | Server port (auto-set by Railway) | `8080` |
| `JWT_SECRET` | Yes | Secret for JWT signing | Random 32+ char string |
| `ALLOWED_ORIGINS` | No | CORS allowed origins (default: `*`) | `https://app.vercel.app` |

## Alternative Deployment Platforms

### Render
1. Create new Web Service
2. Connect GitHub repo
3. Root Directory: `pathway/backend`
4. Build Command: `go build -o server .`
5. Start Command: `./server`
6. Add environment variables
7. Deploy

### Fly.io
1. Install Fly CLI: `curl -L https://fly.io/install.sh | sh`
2. Login: `fly auth login`
3. Create app: `fly launch` (in `pathway/backend` directory)
4. Set secrets: `fly secrets set MONGO_URI=... JWT_SECRET=...`
5. Deploy: `fly deploy`

## Troubleshooting

**Build fails:**
- Check Go version in `go.mod` matches platform
- Verify all dependencies are in `go.sum`
- Check Railway logs for specific errors

**Database connection fails:**
- Verify `MONGO_URI` is correct
- Check MongoDB Atlas network access (must allow Railway IPs)
- Verify database user credentials

**CORS errors:**
- Set `ALLOWED_ORIGINS` to your frontend URL
- Check that frontend is using correct backend URL
- Verify CORS headers in response

**JWT authentication fails:**
- Ensure `JWT_SECRET` is set and consistent
- Check token expiration (default: 24 hours)
- Verify Authorization header format: `Bearer <token>`

## Security Checklist

- [ ] `JWT_SECRET` is a strong random string (32+ characters)
- [ ] MongoDB user has strong password
- [ ] MongoDB network access restricted (not `0.0.0.0/0` in production)
- [ ] `ALLOWED_ORIGINS` set to specific frontend domain(s)
- [ ] No secrets committed to Git
- [ ] HTTPS enabled (automatic on Railway)

## Monitoring

Railway provides:
- Real-time logs
- Metrics dashboard
- Deployment history
- Environment variable management

Set up alerts for:
- Deployment failures
- High error rates
- Database connection issues

