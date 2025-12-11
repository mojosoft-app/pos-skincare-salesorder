# AR Receipt API Documentation

## Overview
The AR Receipt API provides full CRUD (Create, Read, Update, Delete) operations for managing Accounts Receivable receipt records. AR receipts represent payments received from customers and can contain multiple detail line items linking to specific sales orders.

**Base URL**: `/so/api/ar-receipts`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Endpoints

### 1. Get All AR Receipts

Retrieves a list of all AR receipts with optional filtering by customer or status.

**Endpoint**: `GET /so/api/ar-receipts`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| customer_id | integer | No | Filter by customer ID |
| status_id | integer | No | Filter by status ID |
| page | integer | No | Page number (default: 1) |
| limit | integer | No | Items per page (default: 10) |

**Response Codes**:
- `200 OK` - Successfully retrieved AR receipts
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipts retrieved successfully",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "location_id": 1,
      "customer_id": 1001,
      "payment_method_id": 2,
      "doc_number": 12345,
      "doc_date": "2025-01-15",
      "posted_date": "2025-01-16",
      "total_amount": 500000.75,
      "note": "Payment for January invoices",
      "status_id": 1,
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z",
      "details": [
        {
          "id": 1,
          "ar_receipt_id": 101,
          "sales_order_id": 5001,
          "receipt_amount": 250000.50,
          "created_by": 1001,
          "updated_by": null,
          "deleted_by": null,
          "deleted_at": null,
          "created_at": "2025-01-15T10:30:00Z",
          "updated_at": "2025-01-15T10:30:00Z"
        }
      ]
    }
  ]
}
```

**Example Requests**:

Get all AR receipts:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipts" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by customer:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipts?customer_id=1001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by status:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipts?status_id=1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get AR Receipt by ID

Retrieves a single AR receipt by its unique UUID, including all associated details.

**Endpoint**: `GET /so/api/ar-receipts/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The unique UUID of the AR receipt |

**Response Codes**:
- `200 OK` - Successfully retrieved the AR receipt
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - AR receipt not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt retrieved successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "payment_method_id": 2,
    "doc_number": 12345,
    "doc_date": "2025-01-15",
    "posted_date": "2025-01-16",
    "total_amount": 500000.75,
    "note": "Payment for January invoices",
    "status_id": 1,
    "created_by": 1001,
    "updated_by": 1001,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z",
    "details": [
      {
        "id": 1,
        "ar_receipt_id": 101,
        "sales_order_id": 5001,
        "receipt_amount": 250000.50,
        "created_by": 1001,
        "updated_by": null,
        "deleted_by": null,
        "deleted_at": null,
        "created_at": "2025-01-15T10:30:00Z",
        "updated_at": "2025-01-15T10:30:00Z"
      },
      {
        "id": 2,
        "ar_receipt_id": 101,
        "sales_order_id": 5002,
        "receipt_amount": 250000.25,
        "created_by": 1001,
        "updated_by": null,
        "deleted_by": null,
        "deleted_at": null,
        "created_at": "2025-01-15T10:30:00Z",
        "updated_at": "2025-01-15T10:30:00Z"
      }
    ]
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid AR receipt ID",
  "error": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "AR receipt not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipts/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Create AR Receipt

Creates a new AR receipt with optional detail line items in a single transaction.

**Endpoint**: `POST /so/api/ar-receipts`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Request Body**:
```json
{
  "location_id": 1,
  "customer_id": 1001,
  "payment_method_id": 2,
  "doc_number": 12345,
  "doc_date": "2025-01-15",
  "posted_date": "2025-01-16",
  "total_amount": 500000.75,
  "note": "Payment for January invoices",
  "status_id": 1,
  "details": [
    {
      "sales_order_id": 5001,
      "receipt_amount": 250000.50
    },
    {
      "sales_order_id": 5002,
      "receipt_amount": 250000.25
    }
  ]
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| location_id | integer | No | The location/branch ID |
| customer_id | integer | **Yes** | The customer ID (required) |
| payment_method_id | integer | No | The payment method ID |
| doc_number | integer | No | Document number |
| doc_date | string | No | Document date (format: YYYY-MM-DD) |
| posted_date | string | No | Posted date (format: YYYY-MM-DD) |
| total_amount | number | No | Total receipt amount |
| note | string | No | Additional notes |
| status_id | integer | No | Status ID |
| details | array | No | Array of detail line items |

**Detail Object Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| sales_order_id | integer | No | Related sales order ID |
| receipt_amount | number | No | Amount allocated to this sales order |

**Response Codes**:
- `201 Created` - AR receipt created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or transaction failure

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "AR receipt created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "payment_method_id": 2,
    "doc_number": 12345,
    "doc_date": "2025-01-15",
    "posted_date": "2025-01-16",
    "total_amount": 500000.75,
    "note": "Payment for January invoices",
    "status_id": 1,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T14:30:00Z",
    "updated_at": "2025-01-15T14:30:00Z",
    "details": [
      {
        "id": 1,
        "ar_receipt_id": 101,
        "sales_order_id": 5001,
        "receipt_amount": 250000.50,
        "created_by": 1001,
        "updated_by": null,
        "deleted_by": null,
        "deleted_at": null,
        "created_at": "2025-01-15T14:30:00Z",
        "updated_at": "2025-01-15T14:30:00Z"
      }
    ]
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid request body",
  "error": "Key: 'CreateARReceiptRequest.CustomerID' Error:Field validation for 'CustomerID' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/ar-receipts" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1001,
    "payment_method_id": 2,
    "total_amount": 500000.75,
    "note": "Payment for January invoices",
    "status_id": 1,
    "details": [
      {
        "sales_order_id": 5001,
        "receipt_amount": 250000.50
      },
      {
        "sales_order_id": 5002,
        "receipt_amount": 250000.25
      }
    ]
  }'
```

---

### 4. Update AR Receipt

Updates an existing AR receipt record (does not update details - use detail endpoints separately).

**Endpoint**: `PUT /so/api/ar-receipts/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the AR receipt to update |

**Request Body**:
```json
{
  "location_id": 1,
  "customer_id": 1001,
  "payment_method_id": 2,
  "doc_number": 12345,
  "doc_date": "2025-01-15",
  "posted_date": "2025-01-16",
  "total_amount": 550000.00,
  "note": "Updated payment notes",
  "status_id": 2
}
```

**Response Codes**:
- `200 OK` - AR receipt updated successfully
- `400 Bad Request` - Invalid request body or UUID format
- `404 Not Found` - AR receipt not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt updated successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "payment_method_id": 2,
    "doc_number": 12345,
    "doc_date": "2025-01-15",
    "posted_date": "2025-01-16",
    "total_amount": 550000.00,
    "note": "Updated payment notes",
    "status_id": 2,
    "created_by": 1001,
    "updated_by": 1002,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T16:00:00Z",
    "details": []
  }
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/ar-receipts/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1001,
    "payment_method_id": 2,
    "total_amount": 550000.00,
    "note": "Updated payment notes",
    "status_id": 2
  }'
```

---

### 5. Delete AR Receipt

Soft deletes an AR receipt record (sets deleted_at timestamp).

**Endpoint**: `DELETE /so/api/ar-receipts/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the AR receipt to delete |

**Response Codes**:
- `200 OK` - AR receipt deleted successfully
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - AR receipt not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt deleted successfully",
  "data": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "AR receipt not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/ar-receipts/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### ARReceipt Object

| Field | Type | Description |
|-------|------|-------------|
| id | UUID | Unique identifier for the AR receipt (Primary Key, auto-generated) |
| location_id | integer | The location/branch ID where payment was received |
| customer_id | integer | The customer ID (required) |
| payment_method_id | integer | The payment method used (cash, transfer, etc.) |
| doc_number | integer | Document/receipt number |
| doc_date | date | Document date |
| posted_date | date | Date when payment was posted |
| total_amount | number | Total amount received |
| note | string | Additional notes or comments |
| status_id | integer | Status of the receipt (pending, posted, etc.) |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |
| details | array | Array of ARReceiptDetail objects (nested relationship) |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **UUID Primary Key**: Uses UUID for globally unique identifiers
- **Nested Creation**: Create AR receipt with details in a single transaction
- **Transaction Support**: All create operations use database transactions for data integrity
- **Soft Delete Support**: Deleted records are marked rather than physically removed
- **Multi-Tenant**: Supports tenant-specific data isolation via `X-Tenant-Code` header
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Query Filters**: Filter by customer ID or status ID
- **Auto-Preload**: Automatically loads detail relationships when retrieving receipts
- **Audit Trail**: Tracks who created, updated, and deleted each record

---

## Business Logic

### Use Cases

1. **Recording Customer Payments**: When a customer makes a payment, create an AR receipt to track it
2. **Payment Allocation**: Distribute a single payment across multiple sales orders using detail records
3. **Payment History**: Track all payments from a specific customer over time
4. **Financial Reconciliation**: Match receipts with outstanding invoices for accounting

### Relationships

- Each AR receipt belongs to one customer (`customer_id`)
- Each AR receipt can have one location (`location_id`)
- Each AR receipt uses one payment method (`payment_method_id`)
- Each AR receipt has one status (`status_id`)
- Each AR receipt can have multiple detail records (one-to-many with `ARReceiptDetail`)

### Transaction Handling

When creating an AR receipt with details:
1. Database transaction begins
2. AR receipt record is created (UUID auto-generated)
3. All detail records are created
4. Transaction commits (or rolls back on error)
5. Full receipt with details is returned

---

## Error Handling

All error responses follow a consistent format:

```json
{
  "status": "error",
  "message": "<error description>",
  "error": "<detailed error information or null>"
}
```

Common HTTP status codes:
- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid input, malformed UUID, or validation error
- `401 Unauthorized` - Missing or invalid JWT token
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error or transaction failure

---

## Usage Examples

### JavaScript (Fetch API)

```javascript
const BASE_URL = 'http://localhost:8080/so/api/ar-receipts';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all AR receipts
async function getAllReceipts() {
  const response = await fetch(BASE_URL, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Get AR receipt by UUID
async function getReceiptById(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Create new AR receipt with details
async function createReceipt(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update AR receipt
async function updateReceipt(uuid, data) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Delete AR receipt
async function deleteReceipt(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example usage
const newReceipt = await createReceipt({
  customer_id: 1001,
  payment_method_id: 2,
  total_amount: 500000.75,
  note: "Payment for January invoices",
  status_id: 1,
  details: [
    {
      sales_order_id: 5001,
      receipt_amount: 250000.50
    },
    {
      sales_order_id: 5002,
      receipt_amount: 250000.25
    }
  ]
});

console.log('Created Receipt:', newReceipt);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/ar-receipts"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all AR receipts
def get_all_receipts(customer_id=None, status_id=None):
    params = {}
    if customer_id:
        params['customer_id'] = customer_id
    if status_id:
        params['status_id'] = status_id

    response = requests.get(BASE_URL, headers=headers, params=params)
    return response.json()

# Get AR receipt by UUID
def get_receipt_by_id(receipt_uuid):
    url = f"{BASE_URL}/{receipt_uuid}"
    response = requests.get(url, headers=headers)
    return response.json()

# Create new AR receipt
def create_receipt(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update AR receipt
def update_receipt(receipt_uuid, data):
    url = f"{BASE_URL}/{receipt_uuid}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Delete AR receipt
def delete_receipt(receipt_uuid):
    url = f"{BASE_URL}/{receipt_uuid}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example usage
new_receipt_data = {
    "customer_id": 1001,
    "payment_method_id": 2,
    "total_amount": 500000.75,
    "note": "Payment for January invoices",
    "status_id": 1,
    "details": [
        {
            "sales_order_id": 5001,
            "receipt_amount": 250000.50
        },
        {
            "sales_order_id": 5002,
            "receipt_amount": 250000.25
        }
    ]
}

# Create receipt
receipt = create_receipt(new_receipt_data)
print("Created Receipt:", receipt)

# Get receipts for specific customer
customer_receipts = get_all_receipts(customer_id=1001)
print("Customer Receipts:", customer_receipts)
```

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const (
    baseURL    = "http://localhost:8080/so/api/ar-receipts"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type ARReceiptDetail struct {
    SalesOrderID  *int     `json:"sales_order_id,omitempty"`
    ReceiptAmount *float64 `json:"receipt_amount,omitempty"`
}

type ARReceiptRequest struct {
    LocationID      *int              `json:"location_id,omitempty"`
    CustomerID      *int              `json:"customer_id"`
    PaymentMethodID *int              `json:"payment_method_id,omitempty"`
    DocNumber       *int              `json:"doc_number,omitempty"`
    DocDate         *string           `json:"doc_date,omitempty"`
    PostedDate      *string           `json:"posted_date,omitempty"`
    TotalAmount     *float64          `json:"total_amount,omitempty"`
    Note            *string           `json:"note,omitempty"`
    StatusID        *int              `json:"status_id,omitempty"`
    Details         []ARReceiptDetail `json:"details,omitempty"`
}

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func makeRequest(method, url string, body interface{}) (*Response, error) {
    var reqBody io.Reader
    if body != nil {
        jsonData, err := json.Marshal(body)
        if err != nil {
            return nil, err
        }
        reqBody = bytes.NewBuffer(jsonData)
    }

    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+jwtToken)
    req.Header.Set("X-Tenant-Code", tenantCode)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result Response
    err = json.Unmarshal(respBody, &result)
    return &result, err
}

// Get all receipts
func getAllReceipts() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get receipt by UUID
func getReceiptByID(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("GET", url, nil)
}

// Create receipt
func createReceipt(receipt *ARReceiptRequest) (*Response, error) {
    return makeRequest("POST", baseURL, receipt)
}

// Update receipt
func updateReceipt(uuid string, receipt *ARReceiptRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("PUT", url, receipt)
}

// Delete receipt
func deleteReceipt(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create new receipt with details
    customerID := 1001
    paymentMethodID := 2
    totalAmount := 500000.75
    note := "Payment for January invoices"
    statusID := 1
    soID1 := 5001
    amount1 := 250000.50
    soID2 := 5002
    amount2 := 250000.25

    newReceipt := &ARReceiptRequest{
        CustomerID:      &customerID,
        PaymentMethodID: &paymentMethodID,
        TotalAmount:     &totalAmount,
        Note:            &note,
        StatusID:        &statusID,
        Details: []ARReceiptDetail{
            {
                SalesOrderID:  &soID1,
                ReceiptAmount: &amount1,
            },
            {
                SalesOrderID:  &soID2,
                ReceiptAmount: &amount2,
            },
        },
    }

    result, err := createReceipt(newReceipt)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Receipt: %+v\n", result)
}
```

---

## Best Practices

1. **Use Transactions**: Creating receipts with details uses transactions automatically - ensure proper error handling
2. **Validate Customer**: Verify customer exists before creating receipt
3. **Check Amounts**: Ensure detail amounts sum to total_amount
4. **UUID Handling**: Always validate UUID format before making API calls
5. **Update Strategy**: Update receipt header separately from details (use detail endpoints for detail changes)
6. **Soft Delete Benefits**: Maintains audit trail and data integrity
7. **Filter Large Datasets**: Use query parameters to reduce response size
8. **Handle Errors**: Always check response status and handle errors appropriately

---

## Notes

1. **Authentication**: All requests must include a valid JWT token in the Authorization header
2. **Tenant Isolation**: The `X-Tenant-Code` header is required for proper data isolation
3. **UUID Format**: AR receipt IDs are UUIDs (e.g., `550e8400-e29b-41d4-a716-446655440000`)
4. **Soft Deletes**: Records with `deleted_at` set are automatically filtered out
5. **Audit Trail**: System automatically tracks who created, updated, and deleted each record
6. **Database Table**: Data is stored in the `alana.ar_receipt` table
7. **Nested Creation**: Can create receipt with details in one API call (transactional)
8. **Update Limitation**: PUT endpoint updates receipt header only, not details

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD operations and nested creation |
