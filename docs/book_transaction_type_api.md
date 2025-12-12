# Book Transaction Type API Documentation

## Base URL
```
/so/api/book-transaction-type
```

## Endpoints

### 1. Get All Book Transaction Types

Retrieve all book transaction types with optional filter.

**Endpoint:** `GET /so/api/book-transaction-type`

**Query Parameters:**
- `name` (optional, string) - Filter by name (partial match, case-insensitive)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction types retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Sales",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 2,
      "name": "Purchase",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 3,
      "name": "Expense",
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
  "message": "Failed to retrieve book transaction types",
  "data": null
}
```

**Example Request:**
```bash
# Get all book transaction types
curl -X GET "http://localhost:8080/so/api/book-transaction-type" \
  -H "Authorization: Bearer <token>"

# Filter by name
curl -X GET "http://localhost:8080/so/api/book-transaction-type?name=sales" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Book Transaction Type by ID

Retrieve a single book transaction type by its ID.

**Endpoint:** `GET /so/api/book-transaction-type/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Type ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction type retrieved successfully",
  "data": {
    "id": 1,
    "name": "Sales",
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
  "message": "Invalid book transaction type ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction type not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve book transaction type",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/book-transaction-type/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Create Book Transaction Type

Create a new book transaction type.

**Endpoint:** `POST /so/api/book-transaction-type`

**Request Body:**
```json
{
  "name": "Sales"
}
```

**Request Body Parameters:**
- `name` (required, string) - Transaction type name

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Book transaction type created successfully",
  "data": {
    "id": 1,
    "name": "Sales",
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
  "message": "Invalid request body",
  "data": "error details"
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to create book transaction type",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/book-transaction-type" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sales"
  }'
```

---

### 4. Update Book Transaction Type

Update an existing book transaction type.

**Endpoint:** `PUT /so/api/book-transaction-type/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Type ID

**Request Body:**
```json
{
  "name": "Sales - Updated"
}
```

**Request Body Parameters:**
- `name` (required, string) - Transaction type name

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction type updated successfully",
  "data": {
    "id": 1,
    "name": "Sales - Updated",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:00:00Z",
    "created_by": 1,
    "updated_by": 1,
    "deleted_by": null
  }
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid book transaction type ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction type not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update book transaction type",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/book-transaction-type/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sales - Updated"
  }'
```

---

### 5. Delete Book Transaction Type

Soft delete a book transaction type.

**Endpoint:** `DELETE /so/api/book-transaction-type/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Type ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction type deleted successfully",
  "data": null
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid book transaction type ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction type not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete book transaction type",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/book-transaction-type/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- The `name` field is required when creating or updating a transaction type
- The name filter uses case-insensitive partial matching (ILIKE)
- Common transaction types include: Sales, Purchase, Expense, Income, Transfer, Refund, etc.
- This entity is referenced by:
  - `BookkeepingDetail` - To categorize transaction types in bookkeeping details
  - `SummaryByTransactionType` - For transaction type summaries
  - `SummaryByTransactionTypeAndPaymentMethod` - For combined summaries
