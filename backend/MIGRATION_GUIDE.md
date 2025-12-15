# Migration Guide: Copy User Data from Dev to Production

This guide shows how to:
1. **Setup**: Copy user from production to dev database (one-time setup)
2. **Migrate**: Copy user and progress from dev to production

## Prerequisites

- Railway CLI installed (`npm i -g @railway/cli`)
- Railway project linked
- Access to MongoDB connection string (same cluster, different database names)

## Step 1: Setup Dev Database (One-Time)

First, copy your test user from production to the new dev database.

### 1.1 Set Environment Variables in Railway

1. Go to your Railway project dashboard
2. Click on your backend service
3. Go to **Variables** tab
4. Add the following environment variables for copying to dev:

```
SOURCE_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
SOURCE_DB_NAME=pathway
TARGET_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority
TARGET_DB_NAME=pathway-dev
USER_EMAIL=test@example.com
```

**Important Notes:**
- Both use the same cluster (`pathway-admin.gsx3buo.mongodb.net`)
- Source uses database `pathway` (production)
- Target uses database `pathway-dev` (new dev database)
- URL-encode special characters in passwords (e.g., `$` becomes `%24`)

### 1.2 Copy User to Dev Database

Run the copy script via Railway CLI:

```bash
cd pathway/backend
railway run go run cmd/copy-to-dev/main.go
```

This will:
- Copy the test user from `pathway` database to `pathway-dev` database
- Copy all progress records
- Map course IDs between databases

### 1.3 Seed Dev Database with Courses

Make sure dev database has all courses:

```bash
# Temporarily set MONGO_URI to dev database
railway variables set MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority"
railway variables set DB_NAME="pathway-dev"

# Seed courses
railway run go run cmd/seed/main.go

# Restore production settings
railway variables set MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority"
railway variables set DB_NAME="pathway"
```

## Step 2: Migrate from Dev to Production

Now that you have a dev database, you can migrate changes from dev to production.

### 2.1 Set Migration Environment Variables

In Railway **Variables** tab, update/add:

```
SOURCE_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority
SOURCE_DB_NAME=pathway-dev
TARGET_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
TARGET_DB_NAME=pathway
USER_EMAIL=test@example.com
```

## Step 2: Run Migration via Railway CLI

### Option A: Railway CLI (Recommended)

1. **Install Railway CLI** (if not already installed):
   ```bash
   npm i -g @railway/cli
   ```

2. **Login to Railway**:
   ```bash
   railway login
   ```

3. **Link to your project** (if not already linked):
   ```bash
   cd pathway/backend
   railway link
   ```
   Select your Railway project when prompted

4. **Run the migration**:
   ```bash
   railway run go run cmd/migrate/main.go
   ```

   This will:
   - Use the environment variables you set in Railway
   - Connect to both source (dev) and target (production) databases
   - Copy the user and all progress records
   - Map course IDs between databases

### Option B: Railway Web Terminal

1. Go to Railway Dashboard → Your Service
2. Click on your service
3. Look for **"View Logs"** or **"Terminal"** button
4. Click to open the terminal
5. Run:
   ```bash
   cd /app  # or wherever your code is
   go run cmd/migrate/main.go
   ```

## Step 3: Verify Migration

After the migration completes, verify the data:

1. **Test login** with the migrated user:
   ```bash
   curl -X POST https://your-railway-url.up.railway.app/api/auth/login \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","password":"your-password"}'
   ```

2. **Check user progress**:
   ```bash
   curl https://your-railway-url.up.railway.app/api/courses \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
   ```

3. **Verify in frontend**:
   - Log in with `test@example.com`
   - Check that all courses appear
   - Verify progress is preserved

## What the Migration Does

1. **Fetches user** from source (dev) database by email
2. **Checks if user exists** in target (production) database
   - If exists: Uses existing user
   - If not: Creates new user in target
3. **Maps course IDs** between databases (by course title)
4. **Copies all progress records** with mapped course IDs
5. **Cleans up** existing progress (idempotent - safe to run multiple times)

## Troubleshooting

**Migration fails with connection error:**
- Verify both `SOURCE_MONGO_URI` and `TARGET_MONGO_URI` are correct
- Check MongoDB Atlas network access allows Railway IPs
- Ensure passwords are URL-encoded

**User not found:**
- Verify `USER_EMAIL` matches the email in your dev database
- Check that the user exists in the source database

**Course ID mapping fails:**
- Ensure course titles match exactly between dev and production
- Run seed on production first to ensure all courses exist

**Progress not showing:**
- Check that course IDs were mapped correctly
- Verify progress records were inserted (check MongoDB directly)

## Running Locally (Alternative)

If you prefer to run locally instead of on Railway:

1. Create a `.env` file in `pathway/backend/`:
   ```
   SOURCE_MONGO_URI=...
   SOURCE_DB_NAME=pathway
   TARGET_MONGO_URI=...
   TARGET_DB_NAME=pathway
   USER_EMAIL=test@example.com
   ```

2. Run:
   ```bash
   cd pathway/backend
   go run cmd/migrate/main.go
   ```

**Note:** Running on Railway is recommended because:
- Production environment variables are already configured
- Better network access to production database
- More secure (credentials stay in Railway)

