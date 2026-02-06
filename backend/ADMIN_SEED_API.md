# Admin Seeding API

The backend includes a protected endpoint to re-seed the database without needing CLI access.

## Setup

1. **Generate a seed token** (any random string):
   ```bash
   openssl rand -base64 32
   ```

2. **Add to Railway environment variables:**
   - Go to Railway Dashboard → Your Service → Variables
   - Add: `ADMIN_SEED_TOKEN=<your-generated-token>`
   - Redeploy

## Usage

Once set up, you can seed the database by sending a POST request:

```bash
curl -X POST https://your-railway-url.up.railway.app/api/admin/seed \
  -H "X-Admin-Seed-Token: <your-token>"
```

**Response (success):**
```json
{
  "message": "Courses seeded successfully",
  "seeded_at": "2026-02-05T06:10:00Z"
}
```

**Response (if token not set):**
```json
{"error": "Not found"}
```

**Response (if wrong token):**
```json
{"error": "Forbidden"}
```

## Security Notes

- The endpoint is hidden (404) if `ADMIN_SEED_TOKEN` is not set
- Only accepts POST requests with valid token in `X-Admin-Seed-Token` header
- Token should be kept secret - don't commit it to git!
- For production, consider removing the endpoint entirely or using IP restrictions

## Dev vs Prod

**Development:**
- Set `ADMIN_SEED_TOKEN` to any value
- Use it whenever you update course content

**Production:**
- Consider NOT setting the token (endpoint disabled)
- Or use a very strong random token and only seed manually when necessary
- Better yet: remove the admin routes entirely before going live
