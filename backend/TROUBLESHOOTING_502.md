# Troubleshooting 502 Error on Railway

## Check Railway Logs

1. Go to Railway Dashboard → Your Service
2. Click **"Deployments"** tab
3. Click on the latest deployment
4. Click **"View Logs"** or look at the logs panel

Look for errors like:
- "Failed to connect to MongoDB"
- "Port not found"
- "Environment variable missing"
- Any panic/error messages

## Common Issues:

### 1. Missing Environment Variables

Check that these are set in Railway → Variables:
- ✅ `MONGO_URI` - Your MongoDB Atlas connection string
- ✅ `DB_NAME` - `pathway`
- ✅ `JWT_SECRET` - A secure random string
- ✅ `ALLOWED_ORIGINS` - Can be empty or `*` for now

### 2. MongoDB Connection Failing

- Verify `MONGO_URI` is correct
- Check MongoDB Atlas Network Access allows Railway IPs (use `0.0.0.0/0` for testing)
- Verify database user credentials

### 3. Port Configuration

Railway sets `PORT` automatically, but verify:
- Your code reads `PORT` from environment (it does)
- No hardcoded port conflicts

### 4. Application Not Starting

Check logs for:
- Build errors
- Runtime panics
- Missing dependencies

## Quick Fix Steps:

1. **Check Logs** - See what error is shown
2. **Verify Environment Variables** - All 4 should be set
3. **Check MongoDB** - Test connection string locally first
4. **Redeploy** - After fixing env vars, trigger new deployment

## Share the Logs

Copy the error messages from Railway logs and we can fix it!

