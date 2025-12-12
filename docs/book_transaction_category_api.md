# Book Transaction Category API Documentation

## Base URL
```
/so/api/book-transaction-category
```

## Endpoints

### 1. Get All Book Transaction Categories

Retrieve all book transaction categories with optional filter.

**Endpoint:** `GET /so/api/book-transaction-category`

**Query Parameters:**
- `name` (optional, string) - Filter by name (partial match, case-insensitive)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction categories retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Product Sales",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 2,
      "name": "Service Revenue",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 3,
      "name": "Office Supplies",
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
  "message": "Failed to retrieve book transaction categories",
  "data": null
}
```

**Example Request:**
```bash
# Get all book transaction categories
curl -X GET "http://localhost:8080/so/api/book-transaction-category" \
  -H "Authorization: Bearer <token>"

# Filter by name
curl -X GET "http://localhost:8080/so/api/book-transaction-category?name=sales" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Book Transaction Category by ID

Retrieve a single book transaction category by its ID.

**Endpoint:** `GET /so/api/book-transaction-category/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Category ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction category retrieved successfully",
  "data": {
    "id": 1,
    "name": "Product Sales",
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
  "message": "Invalid book transaction category ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction category not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve book transaction category",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/book-transaction-category/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Create Book Transaction Category

Create a new book transaction category.

**Endpoint:** `POST /so/api/book-transaction-category`

**Request Body:**
```json
{
  "name": "Product Sales"
}
```

**Request Body Parameters:**
- `name` (required, string) - Transaction category name

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Book transaction category created successfully",
  "data": {
    "id": 1,
    "name": "Product Sales",
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
  "message": "Failed to create book transaction category",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/book-transaction-category" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Product Sales"
  }'
```

---

### 4. Update Book Transaction Category

Update an existing book transaction category.

**Endpoint:** `PUT /so/api/book-transaction-category/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Category ID

**Request Body:**
```json
{
  "name": "Product Sales - Updated"
}
```

**Request Body Parameters:**
- `name` (required, string) - Transaction category name

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction category updated successfully",
  "data": {
    "id": 1,
    "name": "Product Sales - Updated",
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
  "message": "Invalid book transaction category ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction category not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update book transaction category",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/book-transaction-category/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Product Sales - Updated"
  }'
```

---

### 5. Delete Book Transaction Category

Soft delete a book transaction category.

**Endpoint:** `DELETE /so/api/book-transaction-category/{id}`

**Path Parameters:**
- `id` (required, integer) - Book Transaction Category ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Book transaction category deleted successfully",
  "data": null
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid book transaction category ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Book transaction category not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete book transaction category",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/book-transaction-category/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- The `name` field is required when creating or updating a transaction category
- The name filter uses case-insensitive partial matching (ILIKE)
- **Common transaction categories include:**
  - **Income Categories:** Product Sales, Service Revenue, Consulting Fees, Commission Income, etc.
  - **Expense Categories:** Office Supplies, Utilities, Rent, Salaries, Marketing, etc.
  - **Other Categories:** Bank Fees, Taxes, Equipment, Maintenance, etc.
- This entity is referenced by:
  - `BookkeepingDetail` - To categorize individual transactions in bookkeeping details
- Categories provide more detailed classification than transaction types, allowing for better financial reporting and analysis
