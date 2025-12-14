# Railway Environment Variables

## Your MongoDB Connection String

Add `/pathway` before the `?` to specify the database name:

```
mongodb+srv://colinsmith356_db_user:C0vfefe$22@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
```

## Railway Environment Variables to Set

Go to Railway → Your Service → Variables tab, and add these:

### 1. MONGO_URI
```
mongodb+srv://colinsmith356_db_user:C0vfefe$22@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
```

### 2. DB_NAME
```
pathway
```

### 3. JWT_SECRET
Generate a secure random string (32+ characters). You can use:
- Online generator: https://randomkeygen.com/
- Command line: `openssl rand -base64 32`

Example (generate your own!):
```
aB3xY9mN2pQ7vR5tW8uI0oP4sD6fG1hJ
```

### 4. ALLOWED_ORIGINS (Optional for now)
Leave empty or set to:
```
*
```
(You'll update this to your Vercel URL after frontend is deployed)

## Testing Locally First

Before deploying to Railway, test locally:

1. Update your local `.env` file in `pathway/backend/`:
   ```
   MONGO_URI=mongodb+srv://colinsmith356_db_user:C0vfefe$22@pathway-admin.gsx3buo.mongodb.net/pathway?retryWrites=true&w=majority
   DB_NAME=pathway
   JWT_SECRET=your-local-test-secret
   ```

2. Test connection:
   ```bash
   cd pathway/backend
   go run main.go
   ```

3. If it connects, you should see: "Connected to MongoDB!"

4. Then test the seed:
   ```bash
   go run cmd/seed/main.go
   ```

