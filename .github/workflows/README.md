# GitHub Actions Setup

This repository uses GitHub Actions to automatically deploy and seed the backend on pushes to development and master branches.

## Required GitHub Secrets

You need to set up the following secrets in your GitHub repository settings (Settings → Secrets and variables → Actions):

### Development Environment

| Secret Name | Description | How to Get It |
|------------|-------------|---------------|
| `RAILWAY_DEV_TOKEN` | Railway API token for dev environment | Run `railway login` locally, then `railway token` |
| `RAILWAY_DEV_PROJECT_ID` | Your dev Railway project ID | From Railway dashboard URL or `railway status` |
| `RAILWAY_DEV_SERVICE_ID` | Your dev Railway service ID | From Railway dashboard or `railway status` |

### Production Environment

| Secret Name | Description | How to Get It |
|------------|-------------|---------------|
| `RAILWAY_PROD_TOKEN` | Railway API token for prod environment | Generate a new token from Railway dashboard |
| `RAILWAY_PROD_PROJECT_ID` | Your prod Railway project ID | From Railway dashboard URL |
| `RAILWAY_PROD_SERVICE_ID` | Your prod Railway service ID | From Railway dashboard |

## Setting Up Secrets

### 1. Get Railway Token

```bash
# Install Railway CLI if you haven't
npm install -g @railway/cli

# Login
railway login

# Generate token
railway token
```

### 2. Get Project and Service IDs

```bash
# Link to your dev project
railway link

# Get IDs
railway status
```

Or get them from the Railway dashboard URLs:
- Project ID: `https://railway.app/project/[PROJECT-ID]`
- Service ID: Found in the service settings page

### 3. Add to GitHub

1. Go to your GitHub repo → Settings → Secrets and variables → Actions
2. Click "New repository secret"
3. Add each secret from the table above

## What Happens on Push

### Development Branch
1. Code is checked out
2. Go is set up
3. Railway CLI is installed
4. Code is deployed to Railway dev environment
5. Database is seeded with `go run cmd/seed/main.go`

### Master Branch
Same process but deploys to production environment.

## Testing the Workflow

After setting up secrets, push a small change to development:

```bash
git checkout development
git commit --allow-empty -m "Test deployment workflow"
git push origin development
```

Check the Actions tab in GitHub to see the workflow run.

## Troubleshooting

### "Project not found" error
- Verify `RAILWAY_DEV_TOKEN` is correct and not expired
- Check that project ID matches your Railway dashboard

### "Service not found" error  
- Verify the service ID is correct
- Make sure the service exists in the project

### Seeding fails
- Check that the backend code compiles
- Verify MongoDB is accessible from Railway
- Check Railway logs for connection errors

## Manual Deployment

If you need to deploy manually:

```bash
cd backend

# Deploy
railway up

# Seed
railway run go run cmd/seed/main.go
```
