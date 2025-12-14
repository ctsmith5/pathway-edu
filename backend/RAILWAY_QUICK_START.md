# Railway Deployment - Quick Start Guide

## Step 1: Connect GitHub Repository

1. In Railway dashboard, click **"New Project"**
2. Select **"Deploy from GitHub repo"**
3. Authorize Railway to access your GitHub (if first time)
4. Select your repository: `pathway-edu` (or whatever you named it)
5. Click **"Deploy Now"**

## Step 2: Configure Service

Railway will auto-detect Go, but you need to set the root directory:

1. Click on your service (it will have a random name)
2. Go to **Settings** tab
3. Under **"Root Directory"**, set it to: `pathway/backend`
4. Railway will automatically use the `railway.json` config file

## Step 3: Set Environment Variables

1. Go to the **Variables** tab in your Railway service
2. Click **"New Variable"** and add each of these:

### Required Variables:

**MONGO_URI**
```
mongodb+srv://username:password@cluster.mongodb.net/pathway?retryWrites=true&w=majority
```
*(Replace with your actual MongoDB Atlas connection string)*

**DB_NAME**
```
pathway
```

**JWT_SECRET**
```
<generate-a-secure-random-string>
```
*Generate one with: `openssl rand -base64 32` (or use an online generator)*

**ALLOWED_ORIGINS**
```
https://your-frontend.vercel.app
```
*(Leave empty for now if you haven't deployed frontend yet - it will default to `*`)*

## Step 4: Deploy

1. Railway will automatically start building once you save the root directory
2. Watch the **Deployments** tab for build progress
3. Wait for deployment to complete (usually 2-3 minutes)
4. Once deployed, Railway will give you a URL like: `https://pathway-backend-production.up.railway.app`

## Step 5: Test Your Deployment

Once deployed, test the endpoints:

```bash
# Health check
curl https://your-railway-url.up.railway.app/api/health

# Should return: {"status":"ok","message":"Pathway Backend is running"}
```

## Step 6: Seed the Database

After successful deployment, you need to seed the database:

### Option A: Railway Web Terminal
1. Go to your service in Railway
2. Click on the service
3. Look for **"View Logs"** or **"Open Terminal"** option
4. Run: `go run cmd/seed/main.go`

### Option B: Railway CLI
1. Install Railway CLI: `npm i -g @railway/cli`
2. Login: `railway login`
3. Link project: `railway link`
4. Run seed: `railway run go run cmd/seed/main.go`

### Option C: One-time Deploy Script
You can temporarily modify `railway.json` to run seed on first deploy, then remove it.

## Step 7: Get Your Backend URL

1. In Railway, go to your service
2. Click **Settings** → **Networking**
3. Your public URL will be shown (or generate a custom domain)
4. Copy this URL - you'll need it for your frontend!

## Troubleshooting

**Build fails:**
- Check that Root Directory is set to `pathway/backend`
- Verify `go.mod` is in that directory
- Check Railway logs for specific errors

**Database connection fails:**
- Verify `MONGO_URI` is correct (check for typos)
- Ensure MongoDB Atlas network access allows Railway IPs (use `0.0.0.0/0` for testing)
- Check database user credentials

**Service won't start:**
- Check that `PORT` environment variable is set (Railway sets this automatically)
- Verify start command is `./server`
- Check logs for runtime errors

## Next Steps

After backend is deployed:
1. Copy your Railway backend URL
2. Deploy frontend to Vercel
3. Set `VITE_API_URL` in Vercel to your Railway URL
4. Update `ALLOWED_ORIGINS` in Railway to your Vercel frontend URL
5. Test the full stack!

