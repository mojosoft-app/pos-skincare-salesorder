# Bookkeeping Detail API Documentation

## Base URL
```
/so/api/bookkeeping-detail
```

## Endpoints

### 1. Get All Bookkeeping Details

Retrieve all bookkeeping details with optional filters.

**Endpoint:** `GET /so/api/bookkeeping-detail`

**Query Parameters:**
- `bookkeeping_id` (optional, integer) - Filter by bookkeeping ID
- `type_id` (optional, integer) - Filter by transaction type ID
- `category_id` (optional, integer) - Filter by category ID
- `payment_method_id` (optional, integer) - Filter by payment method ID
- `posted_date_from` (optional, string) - Filter by posted date from (YYYY-MM-DD)
- `posted_date_to` (optional, string) - Filter by posted date to (YYYY-MM-DD)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping details retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 1,
      "type_id": 1,
      "category_id": 1,
      "payment_method_id": 1,
      "posted_date": "2024-01-15T00:00:00Z",
      "doc_number": "INV-001",
      "income": 500000.00,
      "expanse": 0.00,
      "description": "Sales transaction",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "bookkeeping": {
        "id": 1,
        "book_date": "2024-01-15",
        ...
      },
      "type": {
        "id": 1,
        "name": "Sales",
        ...
      },
      "category": {
        "id": 1,
        "name": "Product Sales",
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
  "message": "Failed to retrieve bookkeeping details",
  "data": null
}
```

**Example Request:**
```bash
# Get all bookkeeping details
curl -X GET "http://localhost:8080/so/api/bookkeeping-detail" \
  -H "Authorization: Bearer <token>"

# With filters
curl -X GET "http://localhost:8080/so/api/bookkeeping-detail?bookkeeping_id=1&type_id=1&posted_date_from=2024-01-01&posted_date_to=2024-01-31" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Bookkeeping Detail by ID

Retrieve a single bookkeeping detail by its ID.

**Endpoint:** `GET /so/api/bookkeeping-detail/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping Detail ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping detail retrieved successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "type_id": 1,
    "category_id": 1,
    "payment_method_id": 1,
    "posted_date": "2024-01-15T00:00:00Z",
    "doc_number": "INV-001",
    "income": 500000.00,
    "expanse": 0.00,
    "description": "Sales transaction",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "book_date": "2024-01-15",
      ...
    },
    "type": {
      "id": 1,
      "name": "Sales",
      ...
    },
    "category": {
      "id": 1,
      "name": "Product Sales",
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
  "message": "Invalid bookkeeping detail ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping detail not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve bookkeeping detail",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/bookkeeping-detail/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Get Bookkeeping Details by Bookkeeping ID

Retrieve all details for a specific bookkeeping record.

**Endpoint:** `GET /so/api/bookkeeping-detail/by-bookkeeping/{bookkeeping_id}`

**Path Parameters:**
- `bookkeeping_id` (required, integer) - Bookkeeping ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping details retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 1,
      "type_id": 1,
      "category_id": 1,
      "payment_method_id": 1,
      "posted_date": "2024-01-15T00:00:00Z",
      "doc_number": "INV-001",
      "income": 500000.00,
      "expanse": 0.00,
      "description": "Sales transaction",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null,
      "type": {
        "id": 1,
        "name": "Sales",
        ...
      },
      "category": {
        "id": 1,
        "name": "Product Sales",
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
  "message": "Failed to retrieve bookkeeping details",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/bookkeeping-detail/by-bookkeeping/1" \
  -H "Authorization: Bearer <token>"
```

---

### 4. Create Bookkeeping Detail

Create a new bookkeeping detail.

**Endpoint:** `POST /so/api/bookkeeping-detail`

**Request Body:**
```json
{
  "bookkeeping_id": 1,
  "type_id": 1,
  "category_id": 1,
  "payment_method_id": 1,
  "posted_date": "2024-01-15T00:00:00Z",
  "doc_number": "INV-001",
  "income": 500000.00,
  "expanse": 0.00,
  "description": "Sales transaction"
}
```

**Request Body Parameters:**
- `bookkeeping_id` (optional, integer) - Bookkeeping ID reference
- `type_id` (optional, integer) - Transaction type ID reference
- `category_id` (optional, integer) - Transaction category ID reference
- `payment_method_id` (optional, integer) - Payment method ID reference
- `posted_date` (optional, datetime) - Transaction posted date (ISO 8601 format)
- `doc_number` (optional, string) - Document number (invoice, receipt, etc.)
- `income` (optional, float) - Income amount (use for income transactions)
- `expanse` (optional, float) - Expense amount (use for expense transactions)
- `description` (optional, string) - Transaction description

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Bookkeeping detail created successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "type_id": 1,
    "category_id": 1,
    "payment_method_id": 1,
    "posted_date": "2024-01-15T00:00:00Z",
    "doc_number": "INV-001",
    "income": 500000.00,
    "expanse": 0.00,
    "description": "Sales transaction",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "created_by": 1,
    "updated_by": null,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "book_date": "2024-01-15",
      ...
    },
    "type": {
      "id": 1,
      "name": "Sales",
      ...
    },
    "category": {
      "id": 1,
      "name": "Product Sales",
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
  "message": "Failed to create bookkeeping detail",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/bookkeeping-detail" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 1,
    "type_id": 1,
    "category_id": 1,
    "payment_method_id": 1,
    "posted_date": "2024-01-15T00:00:00Z",
    "doc_number": "INV-001",
    "income": 500000.00,
    "expanse": 0.00,
    "description": "Sales transaction"
  }'
```

---

### 5. Update Bookkeeping Detail

Update an existing bookkeeping detail.

**Endpoint:** `PUT /so/api/bookkeeping-detail/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping Detail ID

**Request Body:**
```json
{
  "bookkeeping_id": 1,
  "type_id": 1,
  "category_id": 1,
  "payment_method_id": 1,
  "posted_date": "2024-01-15T00:00:00Z",
  "doc_number": "INV-001-REV",
  "income": 550000.00,
  "expanse": 0.00,
  "description": "Sales transaction - updated"
}
```

**Request Body Parameters:**
- `bookkeeping_id` (optional, integer) - Bookkeeping ID reference
- `type_id` (optional, integer) - Transaction type ID reference
- `category_id` (optional, integer) - Transaction category ID reference
- `payment_method_id` (optional, integer) - Payment method ID reference
- `posted_date` (optional, datetime) - Transaction posted date (ISO 8601 format)
- `doc_number` (optional, string) - Document number (invoice, receipt, etc.)
- `income` (optional, float) - Income amount (use for income transactions)
- `expanse` (optional, float) - Expense amount (use for expense transactions)
- `description` (optional, string) - Transaction description

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping detail updated successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 1,
    "type_id": 1,
    "category_id": 1,
    "payment_method_id": 1,
    "posted_date": "2024-01-15T00:00:00Z",
    "doc_number": "INV-001-REV",
    "income": 550000.00,
    "expanse": 0.00,
    "description": "Sales transaction - updated",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:00:00Z",
    "created_by": 1,
    "updated_by": 1,
    "deleted_by": null,
    "bookkeeping": {
      "id": 1,
      "book_date": "2024-01-15",
      ...
    },
    "type": {
      "id": 1,
      "name": "Sales",
      ...
    },
    "category": {
      "id": 1,
      "name": "Product Sales",
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
  "message": "Invalid bookkeeping detail ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping detail not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update bookkeeping detail",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/bookkeeping-detail/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 1,
    "type_id": 1,
    "category_id": 1,
    "payment_method_id": 1,
    "posted_date": "2024-01-15T00:00:00Z",
    "doc_number": "INV-001-REV",
    "income": 550000.00,
    "expanse": 0.00,
    "description": "Sales transaction - updated"
  }'
```

---

### 6. Delete Bookkeeping Detail

Soft delete a bookkeeping detail.

**Endpoint:** `DELETE /so/api/bookkeeping-detail/{id}`

**Path Parameters:**
- `id` (required, integer) - Bookkeeping Detail ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Bookkeeping detail deleted successfully",
  "data": null
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid bookkeeping detail ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Bookkeeping detail not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete bookkeeping detail",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/bookkeeping-detail/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- **Relationships:**
  - `Bookkeeping` - Parent bookkeeping record
  - `Type` - Transaction type (Sales, Purchase, etc.)
  - `Category` - Transaction category
  - `PaymentMethod` - Payment method used
- **Transaction Logic:**
  - Use `income` field for income transactions (set `expanse` to 0)
  - Use `expanse` field for expense transactions (set `income` to 0)
  - The `doc_number` field stores document reference (invoice number, receipt number, etc.)
- Date filters use the format YYYY-MM-DD for query parameters
- Multiple filters can be combined for complex queries
