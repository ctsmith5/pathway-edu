# How to Check Railway Logs

## Steps:

1. **Go to Railway Dashboard**
2. **Click on your service** (pathway-edu)
3. **Click "Deployments" tab**
4. **Click on the latest deployment** (the most recent one)
5. **Look at the logs panel** - it should show real-time logs

## What to Look For:

### Common Error Messages:

**MongoDB Connection:**
- "Failed to connect to MongoDB"
- "no reachable servers"
- "authentication failed"
- "network timeout"

**Application Errors:**
- "panic:"
- "fatal error:"
- "cannot find package"
- "undefined:"

**Port/Startup:**
- "address already in use"
- "listen tcp :8080: bind: address already in use"

## Copy the Error

Look for lines in **RED** or lines that say **"error"**, **"panic"**, or **"fatal"**.

Copy the error message and we can fix it!

## Alternative: View Logs in Real-Time

1. Railway Dashboard → Your Service
2. Click **"View Logs"** button (if available)
3. This shows live logs from the running service






