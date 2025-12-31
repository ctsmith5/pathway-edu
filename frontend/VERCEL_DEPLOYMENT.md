# Deploy Frontend to Vercel

## Prerequisites
- Code pushed to GitHub
- Vercel account (free tier works)
- Backend URL: `https://pathway-edu-production.up.railway.app`

## Step-by-Step Deployment

### Step 1: Push Code to GitHub (if not already)
```bash
cd pathway  # (the inner pathway folder with .git)
git add .
git commit -m "Prepare for Vercel deployment"
git push origin master
```

### Step 2: Create Vercel Project
1. Go to https://vercel.com
2. Sign up/Login (use GitHub to connect)
3. Click **"Add New Project"**
4. Import your GitHub repository: `pathway-edu` (or your repo name)
5. Click **"Import"**

### Step 3: Configure Project Settings

**Root Directory:**
- Click **"Edit"** next to Root Directory
- Set to: `pathway/frontend`
- Click **"Continue"**

**Framework Preset:**
- Should auto-detect: **Vite**
- If not, select **"Vite"**

**Build Settings:**
- Build Command: `npm run build` (auto-detected)
- Output Directory: `dist` (auto-detected)
- Install Command: `npm install` (auto-detected)

### Step 4: Set Environment Variables

**Before deploying, click "Environment Variables" and add:**

**Variable Name:** `VITE_API_URL`

**Variable Value:** 
```
https://pathway-edu-production.up.railway.app/api
```

**Important:** Make sure to add this BEFORE clicking "Deploy"!

### Step 5: Deploy

1. Click **"Deploy"**
2. Wait for build to complete (usually 1-2 minutes)
3. Vercel will give you a URL like: `https://pathway-edu.vercel.app`

### Step 6: Update Railway CORS

After Vercel deployment:

1. Go to Railway → Your Service → Variables
2. Edit `ALLOWED_ORIGINS`
3. Set to your Vercel URL: `https://pathway-edu.vercel.app`
4. Save (Railway will auto-redeploy)

### Step 7: Test Full Stack

1. Visit your Vercel URL
2. Click "Login" → Register a new account
3. You should be redirected to Dashboard
4. See all your courses (with "coming soon" for empty ones)

## Troubleshooting

**Build fails:**
- Check that Root Directory is `pathway/frontend`
- Verify all dependencies are in `package.json`
- Check Vercel build logs

**API calls fail:**
- Verify `VITE_API_URL` is set correctly
- Check Railway CORS allows your Vercel domain
- Check browser console for CORS errors

**404 on page refresh:**
- The `vercel.json` rewrites should handle this
- If not, check Vercel routing settings

## Your URLs

- **Frontend:** `https://pathway-edu.vercel.app` (or your custom domain)
- **Backend:** `https://pathway-edu-production.up.railway.app`







