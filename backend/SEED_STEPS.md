# Seed Database - Step by Step

## Quick Method: Railway CLI

### Step 1: Install Railway CLI
```bash
npm i -g @railway/cli
```

### Step 2: Login to Railway
```bash
railway login
```
(This will open a browser to authenticate)

### Step 3: Link to Your Project
```bash
railway link
```
- Select your Railway project when prompted
- Or if you're in the repo directory, it should auto-detect

### Step 4: Run Seed Command
```bash
railway run go run cmd/seed/main.go
```

This will:
- Connect to your Railway service
- Use the environment variables from Railway
- Run the seed command
- Populate the database with 7 courses

### Step 5: Verify
```bash
curl https://pathway-edu-production.up.railway.app/api/courses
```

You should see 7 courses with all their modules and content blocks.

---

## Alternative: If Railway CLI Doesn't Work

### Option A: One-time Deploy Script
Temporarily modify Railway to run seed on deploy (not recommended for production)

### Option B: Manual Database Insert
Use MongoDB Atlas Compass or MongoDB shell to manually insert data

### Option C: Create a Seed Endpoint
Add a `/api/admin/seed` endpoint that you can call once (then remove it)

Let me know which method you'd like to try!







