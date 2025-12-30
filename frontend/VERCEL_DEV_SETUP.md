# Vercel Dev Environment Setup

This guide shows how to set up a separate dev frontend deployment on Vercel.

## Option 1: Separate Vercel Project (Recommended)

### Step 1: Create Dev Project in Vercel

1. Go to https://vercel.com/dashboard
2. Click **"Add New Project"**
3. Import the same GitHub repository
4. Name it: `pathway-edu-dev` (or similar)
5. Set **Root Directory** to: `pathway/frontend`

### Step 2: Configure Dev Environment Variables

In the dev project → Settings → Environment Variables:

**For Dev Database:**
```
VITE_API_URL=https://your-railway-backend.up.railway.app/api
```

**Note:** You can use the same Railway backend, or if you set up a separate dev backend, use that URL instead.

### Step 3: Deploy from Dev Branch

1. Create a `dev` branch:
   ```bash
   git checkout -b dev
   git push origin dev
   ```

2. In Vercel dev project → Settings → Git
3. Set **Production Branch** to: `dev`
4. Or use **Preview Deployments** for automatic PR deployments

### Step 4: Update Railway CORS

Add your dev Vercel URL to Railway's `ALLOWED_ORIGINS`:

In Railway → Variables:
```
ALLOWED_ORIGINS=https://pathway-edu.vercel.app,https://pathway-edu-dev.vercel.app
```

(Comma-separated for multiple origins)

---

## Option 2: Use Vercel Environment-Specific Variables

If you want to keep one project but use different API URLs:

### Step 1: Set Environment-Specific Variables

In your Vercel project → Settings → Environment Variables:

**Variable:** `VITE_API_URL`

**Production:**
- Value: `https://your-production-backend.up.railway.app/api`
- Environment: ✅ Production only

**Preview:**
- Value: `https://your-backend.up.railway.app/api` (or dev backend if separate)
- Environment: ✅ Preview only

**Development:**
- Value: `http://localhost:8080/api`
- Environment: ✅ Development only

### Step 2: Use Preview Deployments

- Push to a feature branch → Vercel creates preview deployment
- Preview deployments use "Preview" environment variables
- Main branch uses "Production" environment variables

---

## Option 3: Separate Dev Backend + Frontend

If you want completely separate dev infrastructure:

### Dev Backend (Railway)
1. Create a new Railway service: `pathway-backend-dev`
2. Use dev database: `pathway-dev`
3. Set environment variables:
   ```
   MONGO_URI=mongodb+srv://...@pathway-admin.gsx3buo.mongodb.net/pathway-dev?...
   DB_NAME=pathway-dev
   ```

### Dev Frontend (Vercel)
1. Create new Vercel project: `pathway-edu-dev`
2. Set `VITE_API_URL` to dev backend URL
3. Deploy from `dev` branch

---

## Recommended Setup

For your current situation, I recommend **Option 1** (separate Vercel project):

**Benefits:**
- ✅ Clear separation between dev and production
- ✅ Can test changes without affecting production
- ✅ Easy to see which is which
- ✅ Can use same backend (just different database via environment)

**Setup:**
1. Create `pathway-edu-dev` project in Vercel
2. Connect to same GitHub repo
3. Set Root Directory: `pathway/frontend`
4. Set `VITE_API_URL` to your Railway backend (same as production)
5. Deploy from `dev` branch
6. Update Railway `ALLOWED_ORIGINS` to include both URLs

**Workflow:**
- Work on `dev` branch → Deploys to dev Vercel
- Merge to `master` → Deploys to production Vercel
- Both point to same Railway backend, but you can test with dev database by changing `MONGO_URI` in Railway

---

## Quick Setup Commands

```bash
# Create dev branch
git checkout -b dev
git push origin dev

# Then in Vercel:
# 1. Create new project: pathway-edu-dev
# 2. Connect to repo
# 3. Set root: pathway/frontend
# 4. Set production branch: dev
# 5. Add VITE_API_URL environment variable
# 6. Deploy
```






