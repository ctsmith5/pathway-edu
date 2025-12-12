# Deployment Guide

## Frontend Deployment to Vercel

### Prerequisites
1. Vercel account (free tier works)
2. GitHub repository connected
3. Backend API deployed and accessible

### Steps

1. **Push code to GitHub**
   ```bash
   git add .
   git commit -m "Prepare for Vercel deployment"
   git push origin master
   ```

2. **Deploy to Vercel**
   - Go to https://vercel.com
   - Click "New Project"
   - Import your GitHub repository
   - **Root Directory**: Set to `pathway/frontend`
   - **Framework Preset**: Vite (auto-detected)
   - **Build Command**: `npm run build` (auto-detected)
   - **Output Directory**: `dist` (auto-detected)

3. **Set Environment Variables**
   In Vercel project settings → Environment Variables, add:
   ```
   VITE_API_URL=https://your-backend-url.com/api
   ```
   Replace `your-backend-url.com` with your actual backend URL.

4. **Deploy**
   - Click "Deploy"
   - Wait for build to complete
   - Your app will be live!

### Important Notes

- **Backend must be deployed separately** - Vercel only hosts the frontend
- **CORS**: Make sure your backend allows requests from your Vercel domain
- **Environment Variables**: The `VITE_` prefix is required for Vite to expose variables to the client

## Backend Deployment Options

The Go backend needs to be deployed separately. Options:

### Option 1: Railway (Recommended - Easy)
- Go to https://railway.app
- Connect GitHub repo
- Set root directory to `pathway/backend`
- Add environment variables:
  - `MONGO_URI` (MongoDB Atlas connection string)
  - `DB_NAME=pathway`
  - `PORT` (auto-set by Railway)
  - `JWT_SECRET` (generate a secure random string)

### Option 2: Render
- Similar to Railway
- Free tier available
- Auto-deploys from GitHub

### Option 3: Fly.io
- Good for Go applications
- Free tier available
- Requires `fly.toml` configuration

### Option 4: DigitalOcean App Platform
- Paid but reliable
- Good for production

## MongoDB Setup

For production, use **MongoDB Atlas** (free tier available):

1. Go to https://www.mongodb.com/cloud/atlas
2. Create a free cluster
3. Create database user
4. Whitelist IP addresses (or use 0.0.0.0/0 for development)
5. Get connection string: `mongodb+srv://username:password@cluster.mongodb.net/pathway?retryWrites=true&w=majority`
6. Use this as `MONGO_URI` in backend environment variables

## Testing Deployment

1. Frontend: Visit your Vercel URL
2. Backend: Test API endpoints
3. Register a new user
4. Verify dashboard loads
5. Test module navigation

