# Production Deployment Checklist

## ✅ Code Pushed to Master
- All course content added (7 courses, 35+ modules)
- Migration scripts created
- Code pushed to GitHub master branch

## Next Steps

### 1. Railway Auto-Deploy (or Manual Redeploy)

Railway should automatically deploy from master branch. If not:

1. Go to Railway Dashboard → Your Service
2. Click "Deployments" tab
3. Click "Redeploy" on latest deployment

### 2. Seed Production Database

**Option A: Railway CLI**
```bash
cd pathway/backend
railway login  # (opens browser - run in your own terminal)
railway link
railway run go run cmd/seed/main.go
```

**Option B: Railway Web Terminal**
1. Railway Dashboard → Your Service
2. Click "View Logs" or "Terminal"
3. Run: `go run cmd/seed/main.go`

**Option C: Set Environment and Run Locally**
```powershell
cd pathway/backend
$env:MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority"
$env:DB_NAME="pathway"
go run cmd/seed/main.go
```

### 3. Verify Vercel Deployment

1. Go to Vercel Dashboard
2. Check that master branch deployed successfully
3. Verify `VITE_API_URL` is set:
   - Should be: `https://your-railway-backend.up.railway.app/api`
4. If needed, trigger redeploy

### 4. Test Production

1. Visit your Vercel URL
2. Register/Login:
   - Email: `test@example.com`
   - Password: `test123456`
3. Verify:
   - ✅ All 7 courses appear
   - ✅ Courses have modules with content
   - ✅ Can navigate through modules
   - ✅ Progress tracking works

## What's Deployed

**Backend (Railway):**
- ✅ All course content (7 courses, 35+ modules)
- ✅ Migration scripts
- ✅ User creation scripts
- ✅ All API endpoints

**Frontend (Vercel):**
- ✅ All UI components
- ✅ Course rendering with ContentRenderer
- ✅ Progress tracking
- ✅ Authentication

**Database (MongoDB Atlas - pathway database):**
- ⏳ Needs seeding: Run `go run cmd/seed/main.go`

## Quick Commands Reference

```bash
# Seed production database (from Railway)
railway run go run cmd/seed/main.go

# Or locally (if you have MONGO_URI set)
cd pathway/backend
go run cmd/seed/main.go

# Verify courses endpoint
curl https://your-railway-url.up.railway.app/api/courses
```





