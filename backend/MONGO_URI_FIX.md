# Fix MongoDB Connection String

## The Problem

Your password contains a `$` character which needs to be URL-encoded in the connection string.

**Original password:** `C0vfefe$22`
**URL-encoded `$`:** `%24`

## Fixed Connection String

In Railway → Variables → `MONGO_URI`, use this:

```
mongodb+srv://colinsmith356_db_user:C0vfefe%2422@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
```

**Key change:** `C0vfefe$22` → `C0vfefe%2422` (the `$` is now `%24`)

## URL Encoding Reference

If your password has other special characters:
- `$` = `%24`
- `@` = `%40`
- `#` = `%23`
- `%` = `%25`
- `&` = `%26`
- `+` = `%2B`
- `=` = `%3D`
- `?` = `%3F`

## Steps to Fix

1. Go to Railway → Your Service → Variables
2. Edit `MONGO_URI`
3. Replace with the encoded version above
4. Save
5. Railway will auto-redeploy
6. Check logs - should see "Connected to MongoDB!"




