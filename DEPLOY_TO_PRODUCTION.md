# Deploy to Production - Complete Guide

## Step 1: Merge Development to Master

```bash
cd pathway  # (the inner pathway folder)
git checkout master
git merge development
git push origin master
```

## Step 2: Seed Production Database on Railway

After Railway auto-deploys (or manually redeploy), seed the database:

### Option A: Railway CLI (After Login)

```bash
cd pathway/backend
railway login  # (opens browser)
railway link  # (select your project)
railway run go run cmd/seed/main.go
```

### Option B: Railway Web Terminal

1. Go to Railway Dashboard → Your Service
2. Click "View Logs" or "Terminal"
3. Run:
   ```bash
   go run cmd/seed/main.go
   ```

## Step 3: Verify Vercel Deployment

1. Vercel should auto-deploy from master branch
2. Check Vercel dashboard for deployment status
3. Verify `VITE_API_URL` is set correctly in Vercel environment variables

## Step 4: Test Production

1. Visit your Vercel URL
2. Register/Login with test@example.com / test123456
3. Verify all 7 courses appear with content
4. Test navigating through modules

## What Gets Deployed

**Backend (Railway):**
- All course content (7 courses, 35+ modules)
- Migration scripts
- User creation scripts

**Frontend (Vercel):**
- All UI components
- Course rendering
- Progress tracking

**Database (MongoDB Atlas):**
- Needs to be seeded with: `go run cmd/seed/main.go`






