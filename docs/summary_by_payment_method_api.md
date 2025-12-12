# Summary By Payment Method API Documentation

## Base URL
```
/so/api/summary-by-payment-method
```

## Endpoints

### 1. Get All Summaries by Payment Method

Retrieve all summaries with optional filters.

**Endpoint:** `GET /so/api/summary-by-payment-method`

**Query Parameters:**
- `bookkeeping_id` (optional, integer) - Filter by bookkeeping ID
- `payment_method_id` (optional, integer) - Filter by payment method ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Summaries retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 1,
      "payment_method_id": 1,
      "total": 150000.00,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "bookkeeping": {
        "id": 1,
        "date": "2024-01-15",
        ...
      },
      "payment_method": {
        "id": 1,
        "name": "Cash",
        ...
      }
    }
  ]
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve summaries",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-payment-method?bookkeeping_id=1" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Summary by ID

Retrieve a single summary by its ID.

**Endpoint:** `GET /so/api/summary-by-payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Summary ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Summary retrieved successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "payment_method_id": 1,
    "total": 150000.00,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "date": "2024-01-15",
      ...
    },
    "payment_method": {
      "id": 1,
      "name": "Cash",
      ...
    }
  }
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid summary ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Summary not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve summary",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-payment-method/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Get Summaries by Bookkeeping ID

Retrieve all summaries for a specific bookkeeping record.

**Endpoint:** `GET /so/api/summary-by-payment-method/by-bookkeeping/{bookkeeping_id}`

**Path Parameters:**
- `bookkeeping_id` (required, integer) - Bookkeeping ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Summaries retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 1,
      "payment_method_id": 1,
      "total": 150000.00,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "payment_method": {
        "id": 1,
        "name": "Cash",
        ...
      }
    },
    {
      "id": 2,
      "bookkeeping_id": 1,
      "payment_method_id": 2,
      "total": 250000.00,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "payment_method": {
        "id": 2,
        "name": "Credit Card",
        ...
      }
    }
  ]
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid bookkeeping ID",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve summaries",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-payment-method/by-bookkeeping/1" \
  -H "Authorization: Bearer <token>"
```

---

### 4. Create Summary by Payment Method

Create a new summary by payment method.

**Endpoint:** `POST /so/api/summary-by-payment-method`

**Request Body:**
```json
{
  "bookkeeping_id": 1,
  "payment_method_id": 1,
  "total": 150000.00
}
```

**Request Body Parameters:**
- `bookkeeping_id` (optional, integer) - Bookkeeping ID reference
- `payment_method_id` (optional, integer) - Payment method ID reference
- `total` (optional, float) - Total amount for this payment method

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Summary created successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "payment_method_id": 1,
    "total": 150000.00,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "date": "2024-01-15",
      ...
    },
    "payment_method": {
      "id": 1,
      "name": "Cash",
      ...
    }
  }
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid request body",
  "data": "error details"
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to create summary",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/summary-by-payment-method" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 1,
    "payment_method_id": 1,
    "total": 150000.00
  }'
```

---

### 5. Update Summary by Payment Method

Update an existing summary by payment method.

**Endpoint:** `PUT /so/api/summary-by-payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Summary ID

**Request Body:**
```json
{
  "bookkeeping_id": 1,
  "payment_method_id": 2,
  "total": 200000.00
}
```

**Request Body Parameters:**
- `bookkeeping_id` (optional, integer) - Bookkeeping ID reference
- `payment_method_id` (optional, integer) - Payment method ID reference
- `total` (optional, float) - Total amount for this payment method

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Summary updated successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "payment_method_id": 2,
    "total": 200000.00,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:00:00Z",
    "created_by": 1,
    "updated_by": 1,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "date": "2024-01-15",
      ...
    },
    "payment_method": {
      "id": 2,
      "name": "Credit Card",
      ...
    }
  }
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid summary ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Summary not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update summary",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/summary-by-payment-method/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 1,
    "payment_method_id": 2,
    "total": 200000.00
  }'
```

---

### 6. Delete Summary by Payment Method

Soft delete a summary by payment method.

**Endpoint:** `DELETE /so/api/summary-by-payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Summary ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Summary deleted successfully",
  "data": null
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid summary ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Summary not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete summary",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/summary-by-payment-method/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- Relationships (Bookkeeping, PaymentMethod) are automatically preloaded in responses where applicable
