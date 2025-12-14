# Test Vercel Deployment

## Quick Test Checklist

### 1. Visit Your Vercel URL
- Open your Vercel deployment URL in a browser
- Should see the landing page with Pathway logo

### 2. Test Registration/Login
- Click "Login" button
- Try registering a new account
- Should redirect to dashboard after successful registration

### 3. Check Browser Console
- Open Developer Tools (F12)
- Go to Console tab
- Look for any errors, especially:
  - CORS errors
  - API connection errors
  - 404 errors

### 4. Test API Connection
- After login, dashboard should load
- Check Network tab in DevTools
- Look for calls to `/api/user/progress`
- Should return 200 status

### 5. Update Railway CORS
If you see CORS errors:
- Go to Railway → Variables
- Set `ALLOWED_ORIGINS` to your Vercel URL
- Save and wait for redeploy

## Common Issues

**CORS Error:**
- Update Railway `ALLOWED_ORIGINS` to your Vercel URL
- Format: `https://your-app.vercel.app` (no trailing slash)

**API Not Found:**
- Check `VITE_API_URL` is set in Vercel
- Should be: `https://pathway-edu-production.up.railway.app/api`

**404 on Refresh:**
- Check `vercel.json` rewrites are working
- Should redirect all routes to `index.html`

