# Fix MongoDB TLS Connection Error

## Issue: TLS Internal Error

The "tls: internal error" suggests either:
1. Network access not configured in MongoDB Atlas
2. Connection string needs TLS parameters

## Solution 1: Check MongoDB Atlas Network Access

1. Go to MongoDB Atlas Dashboard
2. Click **"Network Access"** in left sidebar
3. Check your IP whitelist:
   - For Railway, you need to allow **all IPs** (for now): `0.0.0.0/0`
   - Or add Railway's IP ranges (more secure but complex)

4. If `0.0.0.0/0` is NOT in the list:
   - Click **"Add IP Address"**
   - Click **"Allow Access from Anywhere"**
   - This adds `0.0.0.0/0`
   - Click **"Confirm"**

## Solution 2: Try Alternative Connection String Format

If network access is correct, try adding TLS parameters:

```
mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority&tls=true
```

Or try without the database name in the path:

```
mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/?retryWrites=true&w=majority&tls=true
```

(We specify the database via DB_NAME env var anyway)

## Solution 3: Verify Connection String from Atlas

1. Go to MongoDB Atlas → **Database** → **Connect**
2. Choose **"Connect your application"**
3. Driver: **Go** (or Node.js, doesn't matter)
4. Copy the connection string
5. Replace `<password>` with your URL-encoded password: `C0vfefe%2422`
6. Add `/pathway` before the `?` if not present

## Most Likely Fix:

**Check MongoDB Atlas Network Access** - Make sure `0.0.0.0/0` is whitelisted!





