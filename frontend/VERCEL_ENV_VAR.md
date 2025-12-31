# Set VITE_API_URL in Vercel

## Steps:

1. **Go to Vercel Dashboard**
   - https://vercel.com/dashboard

2. **Click on your project** (pathway-edu or whatever you named it)

3. **Go to Settings**
   - Click **"Settings"** tab at the top

4. **Click "Environment Variables"**
   - In the left sidebar, click **"Environment Variables"**

5. **Add New Variable**
   - Click **"Add New"** button
   - **Key:** `VITE_API_URL`
   - **Value:** `https://pathway-edu-production.up.railway.app/api`
   - **Environment:** Select all three:
     - ✅ Production
     - ✅ Preview
     - ✅ Development
   - Click **"Save"**

6. **Redeploy**
   - After saving, Vercel will ask if you want to redeploy
   - Click **"Redeploy"** or go to **"Deployments"** tab
   - Click the three dots (⋯) on the latest deployment
   - Click **"Redeploy"**

## Important Notes:

- **VITE_** prefix is required - Vite only exposes variables that start with `VITE_`
- Must redeploy after adding env vars - they're baked into the build
- The value should NOT have a trailing slash: `...railway.app/api` (not `...railway.app/api/`)

## Verify It's Set:

After redeploy, check:
1. Visit your Vercel URL
2. Open browser console (F12)
3. Check Network tab
4. Look for API calls - they should go to `pathway-edu-production.up.railway.app` not `localhost:8080`







