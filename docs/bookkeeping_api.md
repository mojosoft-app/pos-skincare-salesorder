# Bookkeeping API Documentation

## Base URL
```
/so/api/bookkeeping
```

## Endpoints

### 1. Get All Bookkeeping Records

Retrieve all bookkeeping records with optional filters.

**Endpoint:** `GET /so/api/bookkeeping`

**Query Parameters:**
- `location_id` (optional, string) - Filter by location ID
- `status_id` (optional, integer) - Filter by status ID
- `book_date_from` (optional, string) - Filter by book date from (YYYY-MM-DD)
- `book_date_to` (optional, string) - Filter by book date to (YYYY-MM-DD)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping records retrieved successfully",
  "data": [
    {
      "id": 1,
      "location_id": "LOC001",
      "book_date": "2024-01-15T00:00:00Z",
      "opening": 500000.00,
      "income": 1500000.00,
      "expanse": 800000.00,
      "balance": 1200000.00,
      "note": "Daily bookkeeping",
      "status_id": 1,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "status": {
        "id": 1,
        "name": "Draft",
        ...
      },
      "details": [
        {
          "id": 1,
          "bookkeeping_id": 1,
          ...
        }
      ]
    }
  ]
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping records",
  "data": null
}
```

**Example Request:**
```bash
# Get all bookkeeping records
curl -X GET "http://localhost:8080/so/api/bookkeeping" \
  -H "Authorization: Bearer <token>"

# With filters
curl -X GET "http://localhost:8080/so/api/bookkeeping?location_id=LOC001&status_id=1&book_date_from=2024-01-01&book_date_to=2024-01-31" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Bookkeeping by ID

Retrieve a single bookkeeping record by its ID with all relationships.

**Endpoint:** `GET /so/api/bookkeeping/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping record retrieved successfully",
  "data": {
    "id": 1,
    "location_id": "LOC001",
    "book_date": "2024-01-15T00:00:00Z",
    "opening": 500000.00,
    "income": 1500000.00,
    "expanse": 800000.00,
    "balance": 1200000.00,
    "note": "Daily bookkeeping",
    "status_id": 1,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "status": {
      "id": 1,
      "name": "Draft",
      ...
    },
    "details": [
      {
        "id": 1,
        "bookkeeping_id": 1,
        ...
      }
    ],
    "summary_by_transaction_type": [
      {
        "id": 1,
        "bookkeeping_id": 1,
        "type_id": 1,
        "total": 1500000.00,
        ...
      }
    ],
    "summary_by_payment_method": [
      {
        "id": 1,
        "bookkeeping_id": 1,
        "payment_method_id": 1,
        "total": 750000.00,
        ...
      }
    ],
    "summary_by_transaction_type_and_payment_method": [
      {
        "id": 1,
        "bookkeeping_id": 1,
        "type_id": 1,
        "payment_method_id": 1,
        "total": 750000.00,
        ...
      }
    ]
  }
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

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping record not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping record",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/bookkeeping/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Get Bookkeeping by Location ID

Retrieve all bookkeeping records for a specific location.

**Endpoint:** `GET /so/api/bookkeeping/by-location/{location_id}`

**Path Parameters:**
- `location_id` (required, string) - Location ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping records retrieved successfully",
  "data": [
    {
      "id": 1,
      "location_id": "LOC001",
      "book_date": "2024-01-15T00:00:00Z",
      "opening": 500000.00,
      "income": 1500000.00,
      "expanse": 800000.00,
      "balance": 1200000.00,
      "note": "Daily bookkeeping",
      "status_id": 1,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "status": {
        "id": 1,
        "name": "Draft",
        ...
      },
      "details": [
        {
          "id": 1,
          "bookkeeping_id": 1,
          ...
        }
      ]
    }
  ]
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping records",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/bookkeeping/by-location/LOC001" \
  -H "Authorization: Bearer <token>"
```

---

### 4. Create Bookkeeping Record

Create a new bookkeeping record.

**Endpoint:** `POST /so/api/bookkeeping`

**Request Body:**
```json
{
  "location_id": "LOC001",
  "book_date": "2024-01-15T00:00:00Z",
  "opening": 500000.00,
  "income": 1500000.00,
  "expanse": 800000.00,
  "balance": 1200000.00,
  "note": "Daily bookkeeping",
  "status_id": 1
}
```

**Request Body Parameters:**
- `location_id` (optional, string) - Location ID reference
- `book_date` (optional, datetime) - Booking date (ISO 8601 format)
- `opening` (optional, float) - Opening balance
- `income` (optional, float) - Total income
- `expanse` (optional, float) - Total expense
- `balance` (optional, float) - Closing balance (opening + income - expanse)
- `note` (optional, string) - Additional notes
- `status_id` (optional, integer) - Bookkeeping status ID reference

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Bookkeeping record created successfully",
  "data": {
    "id": 1,
    "location_id": "LOC001",
    "book_date": "2024-01-15T00:00:00Z",
    "opening": 500000.00,
    "income": 1500000.00,
    "expanse": 800000.00,
    "balance": 1200000.00,
    "note": "Daily bookkeeping",
    "status_id": 1,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "status": {
      "id": 1,
      "name": "Draft",
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
  "message": "Failed to create bookkeeping record",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/bookkeeping" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "location_id": "LOC001",
    "book_date": "2024-01-15T00:00:00Z",
    "opening": 500000.00,
    "income": 1500000.00,
    "expanse": 800000.00,
    "balance": 1200000.00,
    "note": "Daily bookkeeping",
    "status_id": 1
  }'
```

---

### 5. Update Bookkeeping Record

Update an existing bookkeeping record.

**Endpoint:** `PUT /so/api/bookkeeping/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping ID

**Request Body:**
```json
{
  "location_id": "LOC001",
  "book_date": "2024-01-15T00:00:00Z",
  "opening": 500000.00,
  "income": 1600000.00,
  "expanse": 850000.00,
  "balance": 1250000.00,
  "note": "Daily bookkeeping - updated",
  "status_id": 2
}
```

**Request Body Parameters:**
- `location_id` (optional, string) - Location ID reference
- `book_date` (optional, datetime) - Booking date (ISO 8601 format)
- `opening` (optional, float) - Opening balance
- `income` (optional, float) - Total income
- `expanse` (optional, float) - Total expense
- `balance` (optional, float) - Closing balance (opening + income - expanse)
- `note` (optional, string) - Additional notes
- `status_id` (optional, integer) - Bookkeeping status ID reference

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping record updated successfully",
  "data": {
    "id": 1,
    "location_id": "LOC001",
    "book_date": "2024-01-15T00:00:00Z",
    "opening": 500000.00,
    "income": 1600000.00,
    "expanse": 850000.00,
    "balance": 1250000.00,
    "note": "Daily bookkeeping - updated",
    "status_id": 2,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:00:00Z",
    "created_by": 1,
    "updated_by": 1,
    "deleted_by": null,
    "status": {
      "id": 2,
      "name": "Approved",
      ...
    }
  }
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

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping record not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update bookkeeping record",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/bookkeeping/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "location_id": "LOC001",
    "book_date": "2024-01-15T00:00:00Z",
    "opening": 500000.00,
    "income": 1600000.00,
    "expanse": 850000.00,
    "balance": 1250000.00,
    "note": "Daily bookkeeping - updated",
    "status_id": 2
  }'
```

---

### 6. Delete Bookkeeping Record

Soft delete a bookkeeping record.

**Endpoint:** `DELETE /so/api/bookkeeping/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping record deleted successfully",
  "data": null
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

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping record not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete bookkeeping record",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/bookkeeping/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- **Relationships:**
  - `Status` - Bookkeeping status reference
  - `Details` - Related bookkeeping detail records
  - `SummaryByTransactionType` - Summary grouped by transaction type
  - `SummaryByPaymentMethod` - Summary grouped by payment method
  - `SummaryByTransactionTypeAndPaymentMethod` - Summary grouped by both transaction type and payment method
- The GetByID endpoint loads all relationships including all summary tables for comprehensive data retrieval
- Date filters use the format YYYY-MM-DD for query parameters
- The balance should typically equal: opening + income - expanse
