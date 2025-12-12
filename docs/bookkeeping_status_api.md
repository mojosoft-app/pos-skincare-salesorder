# Bookkeeping Status API Documentation

## Base URL
```
/so/api/bookkeeping-status
```

## Endpoints

### 1. Get All Bookkeeping Statuses

Retrieve all bookkeeping statuses with optional filter.

**Endpoint:** `GET /so/api/bookkeeping-status`

**Query Parameters:**
- `name` (optional, string) - Filter by name (partial match, case-insensitive)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping statuses retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Draft",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 2,
      "name": "Approved",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 3,
      "name": "Completed",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    }
  ]
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping statuses",
  "data": null
}
```

**Example Request:**
```bash
# Get all bookkeeping statuses
curl -X GET "http://localhost:8080/so/api/bookkeeping-status" \
  -H "Authorization: Bearer <token>"

# Filter by name
curl -X GET "http://localhost:8080/so/api/bookkeeping-status?name=draft" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Bookkeeping Status by ID

Retrieve a single bookkeeping status by its ID.

**Endpoint:** `GET /so/api/bookkeeping-status/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping Status ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping status retrieved successfully",
  "data": {
    "id": 1,
    "name": "Draft",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null
  }
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid bookkeeping status ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping status not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping status",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/bookkeeping-status/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The tenant database is automatically selected based on the authentication context
- This is a master data endpoint - typically read-only
- The name filter uses case-insensitive partial matching (ILIKE)
- Common bookkeeping statuses include: Draft, Approved, Completed, Rejected, etc.
- This status is referenced by the Bookkeeping entity to track the state of bookkeeping records
