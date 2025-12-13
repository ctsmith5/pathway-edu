# Railway Settings Configuration

## Important: Railway Dashboard Settings

Since Railway is having trouble with the Root Directory, use these settings:

### In Railway Dashboard → Your Service → Settings:

1. **Root Directory:** 
   - Leave this **EMPTY** (or set to `/`)
   - Railway will use the repository root

2. **Builder:**
   - Select **"Dockerfile"**

3. **Dockerfile Path:**
   - Should auto-detect `Dockerfile` at root
   - Or manually set to: `Dockerfile`

4. **Start Command:**
   - Leave empty (Dockerfile CMD will handle it)
   - Or set to: `./server`

### Environment Variables:

Still add these in the Variables tab:
- `MONGO_URI=mongodb+srv://...`
- `DB_NAME=pathway`
- `JWT_SECRET=<your-secret>`
- `ALLOWED_ORIGINS=*` (or your Vercel URL)

### Why This Works:

The root `Dockerfile` at `pathway/Dockerfile` will:
- Copy files from `backend/go.mod` (relative to repo root)
- Copy all backend source code
- Build the Go application
- Run the server

This approach avoids Railway's Root Directory issues by building from the repo root.

