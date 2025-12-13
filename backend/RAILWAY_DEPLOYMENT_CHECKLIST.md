# Railway Deployment Checklist

Use this checklist when deploying to Railway.

## Pre-Deployment

- [ ] Code pushed to GitHub
- [ ] MongoDB Atlas cluster created
- [ ] MongoDB database user created
- [ ] MongoDB network access configured (allow Railway IPs or 0.0.0.0/0 for testing)
- [ ] Connection string obtained from MongoDB Atlas
- [ ] JWT secret generated (`openssl rand -base64 32`)

## Railway Setup

- [ ] Account created at https://railway.app
- [ ] GitHub account connected
- [ ] New project created
- [ ] Repository selected
- [ ] Root directory set to: `pathway/backend`

## Environment Variables

Add these in Railway → Variables:

- [ ] `MONGO_URI` = `mongodb+srv://username:password@cluster.mongodb.net/pathway?retryWrites=true&w=majority`
- [ ] `DB_NAME` = `pathway`
- [ ] `JWT_SECRET` = `<your-generated-secret>`
- [ ] `ALLOWED_ORIGINS` = `https://your-frontend.vercel.app` (or leave empty for `*`)

## Deployment

- [ ] Railway auto-detected Go (or manually set)
- [ ] Build command: `go build -o server .` (or auto)
- [ ] Start command: `./server` (or auto)
- [ ] Deployment started
- [ ] Deployment completed successfully
- [ ] Got deployment URL (e.g., `https://pathway-backend.railway.app`)

## Post-Deployment

- [ ] Test health endpoint: `curl https://your-backend.railway.app/api/health`
- [ ] Test courses endpoint: `curl https://your-backend.railway.app/api/courses`
- [ ] Seed database (run `go run cmd/seed/main.go` via Railway CLI or web terminal)
- [ ] Test registration: `curl -X POST https://your-backend.railway.app/api/auth/register ...`
- [ ] Test login: `curl -X POST https://your-backend.railway.app/api/auth/login ...`
- [ ] Update Vercel frontend with backend URL

## Frontend Update

- [ ] Go to Vercel project settings
- [ ] Add/update environment variable: `VITE_API_URL=https://your-backend.railway.app/api`
- [ ] Redeploy frontend
- [ ] Test full flow: Register → Login → Dashboard → Module View

## Troubleshooting

If deployment fails:
- [ ] Check Railway logs for errors
- [ ] Verify Go version compatibility
- [ ] Check environment variables are set correctly
- [ ] Verify MongoDB connection string format
- [ ] Check MongoDB network access allows Railway IPs

If API calls fail:
- [ ] Verify CORS `ALLOWED_ORIGINS` matches frontend URL
- [ ] Check backend URL is correct in frontend
- [ ] Verify JWT secret is set
- [ ] Check MongoDB connection is working

## Security Review

- [ ] JWT secret is strong and random
- [ ] MongoDB password is strong
- [ ] CORS restricted to specific domain (not `*`)
- [ ] No secrets in Git repository
- [ ] MongoDB network access restricted (production)

