# Setup Dev Database - Quick Guide

## Option 1: Run Locally (Easiest for Setup)

### Step 1: Create `.env` file in `pathway/backend/`

Create a file called `.env` with:

```
SOURCE_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
SOURCE_DB_NAME=pathway
TARGET_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority
TARGET_DB_NAME=pathway-dev
USER_EMAIL=test@example.com
```

### Step 2: Copy User to Dev Database

```powershell
cd pathway/backend
go run cmd/copy-to-dev/main.go
```

### Step 3: Seed Dev Database with Courses

Temporarily update your `.env`:

```
MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority
DB_NAME=pathway-dev
```

Then run:
```powershell
go run cmd/seed/main.go
```

Then restore your original `MONGO_URI` and `DB_NAME` in `.env`.

---

## Option 2: Run on Railway (After Login)

### Step 1: Login to Railway

Open a new terminal and run:
```powershell
railway login
```
(This will open a browser for authentication)

### Step 2: Link Project

```powershell
cd pathway/backend
railway link
```
(Select your Railway project)

### Step 3: Set Environment Variables in Railway

Go to Railway Dashboard → Your Service → Variables tab, add:

```
SOURCE_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
SOURCE_DB_NAME=pathway
TARGET_MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority
TARGET_DB_NAME=pathway-dev
USER_EMAIL=test@example.com
```

### Step 4: Copy User to Dev

```powershell
railway run go run cmd/copy-to-dev/main.go
```

### Step 5: Seed Dev Database

Temporarily update Railway variables:
```powershell
railway variables set MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway-dev?retryWrites=true&w=majority"
railway variables set DB_NAME="pathway-dev"
railway run go run cmd/seed/main.go
```

Then restore:
```powershell
railway variables set MONGO_URI="mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority"
railway variables set DB_NAME="pathway"
```






