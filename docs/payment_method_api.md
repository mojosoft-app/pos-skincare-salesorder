# Payment Method API Documentation

## Base URL
```
/so/api/payment-method
```

## Endpoints

### 1. Get All Payment Methods

Retrieve all payment methods with optional filter.

**Endpoint:** `GET /so/api/payment-method`

**Query Parameters:**
- `name` (optional, string) - Filter by name (partial match, case-insensitive)

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Payment methods retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Cash",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 2,
      "name": "Credit Card",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "created_by": 1,
      "updated_by": null,
      "deleted_by": null
    },
    {
      "id": 3,
      "name": "Bank Transfer",
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
  "message": "Failed to retrieve payment methods",
  "data": null
}
```

**Example Request:**
```bash
# Get all payment methods
curl -X GET "http://localhost:8080/so/api/payment-method" \
  -H "Authorization: Bearer <token>"

# Filter by name
curl -X GET "http://localhost:8080/so/api/payment-method?name=cash" \
  -H "Authorization: Bearer <token>"
```

---

### 2. Get Payment Method by ID

Retrieve a single payment method by its ID.

**Endpoint:** `GET /so/api/payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Payment Method ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Payment method retrieved successfully",
  "data": {
    "id": 1,
    "name": "Cash",
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
  "message": "Invalid payment method ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Payment method not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to retrieve payment method",
  "data": null
}
```

**Example Request:**
```bash
curl -X GET "http://localhost:8080/so/api/payment-method/1" \
  -H "Authorization: Bearer <token>"
```

---

### 3. Create Payment Method

Create a new payment method.

**Endpoint:** `POST /so/api/payment-method`

**Request Body:**
```json
{
  "name": "Cash"
}
```

**Request Body Parameters:**
- `name` (required, string) - Payment method name

**Response Success (201 Created):**
```json
{
  "status": "success",
  "message": "Payment method created successfully",
  "data": {
    "id": 1,
    "name": "Cash",
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
  "message": "Failed to create payment method",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X POST "http://localhost:8080/so/api/payment-method" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cash"
  }'
```

---

### 4. Update Payment Method

Update an existing payment method.

**Endpoint:** `PUT /so/api/payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Payment Method ID

**Request Body:**
```json
{
  "name": "Cash - Updated"
}
```

**Request Body Parameters:**
- `name` (required, string) - Payment method name

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Payment method updated successfully",
  "data": {
    "id": 1,
    "name": "Cash - Updated",
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
  "message": "Invalid payment method ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Payment method not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to update payment method",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X PUT "http://localhost:8080/so/api/payment-method/1" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cash - Updated"
  }'
```

---

### 5. Delete Payment Method

Soft delete a payment method.

**Endpoint:** `DELETE /so/api/payment-method/{id}`

**Path Parameters:**
- `id` (required, integer) - Payment Method ID

**Response Success (200 OK):**
```json
{
  "status": "success",
  "message": "Payment method deleted successfully",
  "data": null
}
```

**Response Error (400 Bad Request):**
```json
{
  "status": "error",
  "message": "Invalid payment method ID",
  "data": null
}
```

**Response Error (404 Not Found):**
```json
{
  "status": "error",
  "message": "Payment method not found",
  "data": null
}
```

**Response Error (500 Internal Server Error):**
```json
{
  "status": "error",
  "message": "Failed to delete payment method",
  "data": "error details"
}
```

**Example Request:**
```bash
curl -X DELETE "http://localhost:8080/so/api/payment-method/1" \
  -H "Authorization: Bearer <token>"
```

---

## Notes

- All endpoints require authentication via Bearer token
- The API uses soft delete, so deleted records are marked as deleted but not removed from the database
- The `created_by`, `updated_by`, and `deleted_by` fields are automatically populated from the authenticated user context
- The tenant database is automatically selected based on the authentication context
- The `name` field is required when creating or updating a payment method
- The name filter uses case-insensitive partial matching (ILIKE)
- **Common payment methods include:**
  - **Cash-based:** Cash, Petty Cash
  - **Card-based:** Credit Card, Debit Card, VISA, Mastercard
  - **Bank Transfer:** Bank Transfer, Wire Transfer, ACH
  - **Digital Wallets:** PayPal, Stripe, GoPay, OVO, Dana, ShopeePay
  - **Other:** Check, Money Order, Mobile Payment
- This entity is referenced by:
  - `BookkeepingDetail` - To track payment method used in individual transactions
  - `SummaryByPaymentMethod` - For payment method summaries
  - `SummaryByTransactionTypeAndPaymentMethod` - For combined summaries
- Payment methods are essential for financial tracking and reconciliation
