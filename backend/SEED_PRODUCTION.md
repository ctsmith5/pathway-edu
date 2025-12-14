# Seed Production Database

After Railway deployment succeeds, you need to seed the database with course data.

## Option 1: Railway Web Terminal (Easiest)

1. Go to Railway Dashboard → Your Service
2. Click on your service
3. Go to **"Deployments"** tab
4. Click on the latest deployment
5. Look for **"View Logs"** or **"Open Terminal"** option
6. Run:
   ```bash
   go run cmd/seed/main.go
   ```

## Option 2: Railway CLI

1. Install Railway CLI:
   ```bash
   npm i -g @railway/cli
   ```

2. Login:
   ```bash
   railway login
   ```

3. Link to your project:
   ```bash
   railway link
   ```
   (Select your Railway project)

4. Run seed command:
   ```bash
   railway run go run cmd/seed/main.go
   ```

## Option 3: One-time Deploy Script

You can temporarily add this to your service settings:
- In Railway → Service → Settings
- Add a custom command that runs seed on deploy
- (Not recommended, better to run manually)

## Verify Seed Worked

After running seed, test:
```bash
curl https://your-railway-url.up.railway.app/api/courses
```

You should see 7 courses with their modules and content blocks.

