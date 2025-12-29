# Seed Database on Railway

Your backend is working! Now seed the database with course data.

## Option 1: Railway Web Terminal (Easiest)

1. Go to Railway Dashboard → Your Service
2. Click on your service
3. Look for **"View Logs"** or **"Terminal"** button
4. If you see a terminal option, click it
5. Run:
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
   (Select your Railway project when prompted)

4. Run seed:
   ```bash
   railway run go run cmd/seed/main.go
   ```

## Option 3: SSH into Container

If Railway provides SSH access:
1. Get SSH command from Railway
2. Connect
3. Navigate to `/app` or wherever the code is
4. Run: `go run cmd/seed/main.go`

## Verify After Seeding

Test the courses endpoint:
```bash
curl https://pathway-edu-production.up.railway.app/api/courses
```

You should see 7 courses with their modules and content blocks.




