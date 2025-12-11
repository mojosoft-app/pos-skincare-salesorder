# AR Receipt Detail API Documentation

## Overview
The AR Receipt Detail API provides full CRUD (Create, Read, Update, Delete) operations for managing AR receipt detail records. These details represent individual line items within AR receipts, linking receipts to specific sales orders.

**Base URL**: `/so/api/ar-receipt-details`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Endpoints

### 1. Get All AR Receipt Details

Retrieves a list of all AR receipt details with optional filtering.

**Endpoint**: `GET /so/api/ar-receipt-details`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| ar_receipt_id | integer | No | Filter by AR receipt ID |
| sales_order_id | integer | No | Filter by sales order ID |

**Response Codes**:
- `200 OK` - Successfully retrieved AR receipt details
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt details retrieved successfully",
  "data": [
    {
      "id": 1,
      "ar_receipt_id": 101,
      "sales_order_id": 5001,
      "receipt_amount": 150000.50,
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "ar_receipt_id": 101,
      "sales_order_id": 5002,
      "receipt_amount": 250000.00,
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T11:00:00Z",
      "updated_at": "2025-01-15T11:00:00Z"
    }
  ]
}
```

**Example Requests**:

Get all details:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipt-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by AR receipt ID:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipt-details?ar_receipt_id=101" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by sales order ID:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipt-details?sales_order_id=5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get AR Receipt Detail by ID

Retrieves a single AR receipt detail by its unique ID.

**Endpoint**: `GET /so/api/ar-receipt-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the AR receipt detail |

**Response Codes**:
- `200 OK` - Successfully retrieved the AR receipt detail
- `400 Bad Request` - Invalid AR receipt detail ID format
- `404 Not Found` - AR receipt detail not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt detail retrieved successfully",
  "data": {
    "id": 1,
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 150000.50,
    "created_by": 1001,
    "updated_by": 1001,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  }
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipt-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Get AR Receipt Details by AR Receipt ID

Retrieves all detail records for a specific AR receipt.

**Endpoint**: `GET /so/api/ar-receipt-details/by-ar-receipt/{ar_receipt_id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| ar_receipt_id | integer | Yes | The AR receipt ID to get details for |

**Response Codes**:
- `200 OK` - Successfully retrieved AR receipt details
- `400 Bad Request` - Invalid AR receipt ID format
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt details retrieved successfully",
  "data": [
    {
      "id": 1,
      "ar_receipt_id": 101,
      "sales_order_id": 5001,
      "receipt_amount": 150000.50,
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "ar_receipt_id": 101,
      "sales_order_id": 5002,
      "receipt_amount": 250000.00,
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T11:00:00Z",
      "updated_at": "2025-01-15T11:00:00Z"
    }
  ]
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/ar-receipt-details/by-ar-receipt/101" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 4. Create AR Receipt Detail

Creates a new AR receipt detail record.

**Endpoint**: `POST /so/api/ar-receipt-details`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Request Body**:
```json
{
  "ar_receipt_id": 101,
  "sales_order_id": 5001,
  "receipt_amount": 150000.50
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ar_receipt_id | integer | Yes | The AR receipt ID this detail belongs to |
| sales_order_id | integer | No | The related sales order ID |
| receipt_amount | number | No | The amount received for this line item |

**Response Codes**:
- `201 Created` - AR receipt detail created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or server error

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "AR receipt detail created successfully",
  "data": {
    "id": 3,
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 150000.50,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T14:30:00Z",
    "updated_at": "2025-01-15T14:30:00Z"
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid request body",
  "error": "Key: 'CreateARReceiptDetailRequest.ARReceiptID' Error:Field validation for 'ARReceiptID' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/ar-receipt-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 150000.50
  }'
```

---

### 5. Update AR Receipt Detail

Updates an existing AR receipt detail record.

**Endpoint**: `PUT /so/api/ar-receipt-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the AR receipt detail to update |

**Request Body**:
```json
{
  "ar_receipt_id": 101,
  "sales_order_id": 5001,
  "receipt_amount": 175000.75
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ar_receipt_id | integer | Yes | The AR receipt ID this detail belongs to |
| sales_order_id | integer | No | The related sales order ID |
| receipt_amount | number | No | The amount received for this line item |

**Response Codes**:
- `200 OK` - AR receipt detail updated successfully
- `400 Bad Request` - Invalid request body or ID format
- `404 Not Found` - AR receipt detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt detail updated successfully",
  "data": {
    "id": 1,
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 175000.75,
    "created_by": 1001,
    "updated_by": 1002,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T15:45:00Z"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "AR receipt detail not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/ar-receipt-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 175000.75
  }'
```

---

### 6. Delete AR Receipt Detail

Soft deletes an AR receipt detail record.

**Endpoint**: `DELETE /so/api/ar-receipt-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the AR receipt detail to delete |

**Response Codes**:
- `200 OK` - AR receipt detail deleted successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - AR receipt detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "AR receipt detail deleted successfully",
  "data": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "AR receipt detail not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/ar-receipt-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### ARReceiptDetail Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the AR receipt detail (Primary Key, Auto-increment) |
| ar_receipt_id | integer | Foreign key referencing the AR receipt |
| sales_order_id | integer | Foreign key referencing the sales order |
| receipt_amount | number | The amount received for this detail line item |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Soft Delete Support**: Deleted records are marked rather than physically removed
- **Multi-Tenant**: Supports tenant-specific data isolation via `X-Tenant-Code` header
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Query Filters**: Filter by AR receipt ID or sales order ID
- **Custom Lookup**: Dedicated endpoint to retrieve all details for a specific AR receipt
- **Audit Trail**: Tracks who created, updated, and deleted each record

---

## Business Logic

### Use Cases

1. **Recording Payment Details**: When a customer makes a payment that covers multiple sales orders, create separate detail records for each order
2. **Payment Allocation**: Track how a single receipt amount is distributed across different sales orders
3. **Reconciliation**: Match receipt details with outstanding sales orders for accounting purposes

### Relationships

- Each AR receipt detail belongs to one AR receipt (`ar_receipt_id`)
- Each AR receipt detail can be linked to one sales order (`sales_order_id`)
- An AR receipt can have multiple detail records

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
- `400 Bad Request` - Invalid input or malformed request
- `401 Unauthorized` - Missing or invalid JWT token
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error

---

## Usage Examples

### JavaScript (Fetch API)

```javascript
const BASE_URL = 'http://localhost:8080/so/api/ar-receipt-details';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all AR receipt details
async function getAllDetails() {
  const response = await fetch(BASE_URL, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Get details by AR receipt ID
async function getDetailsByARReceiptID(arReceiptId) {
  const response = await fetch(`${BASE_URL}/by-ar-receipt/${arReceiptId}`, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Create new detail
async function createDetail(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update detail
async function updateDetail(id, data) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Delete detail
async function deleteDetail(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example usage
const newDetail = await createDetail({
  ar_receipt_id: 101,
  sales_order_id: 5001,
  receipt_amount: 150000.50
});

console.log('Created:', newDetail);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/ar-receipt-details"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all AR receipt details
def get_all_details():
    response = requests.get(BASE_URL, headers=headers)
    return response.json()

# Get details by AR receipt ID
def get_details_by_ar_receipt_id(ar_receipt_id):
    url = f"{BASE_URL}/by-ar-receipt/{ar_receipt_id}"
    response = requests.get(url, headers=headers)
    return response.json()

# Create new detail
def create_detail(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update detail
def update_detail(detail_id, data):
    url = f"{BASE_URL}/{detail_id}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Delete detail
def delete_detail(detail_id):
    url = f"{BASE_URL}/{detail_id}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example usage
new_detail = create_detail({
    "ar_receipt_id": 101,
    "sales_order_id": 5001,
    "receipt_amount": 150000.50
})

print("Created:", new_detail)

# Get all details for a specific AR receipt
details = get_details_by_ar_receipt_id(101)
print("Details for AR Receipt 101:", details)
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
    baseURL    = "http://localhost:8080/so/api/ar-receipt-details"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type ARReceiptDetail struct {
    ID            int      `json:"id"`
    ARReceiptID   *int     `json:"ar_receipt_id"`
    SalesOrderID  *int     `json:"sales_order_id"`
    ReceiptAmount *float64 `json:"receipt_amount"`
    CreatedBy     *int64   `json:"created_by"`
    UpdatedBy     *int64   `json:"updated_by"`
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

// Get all details
func getAllDetails() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get details by AR receipt ID
func getDetailsByARReceiptID(arReceiptID int) (*Response, error) {
    url := fmt.Sprintf("%s/by-ar-receipt/%d", baseURL, arReceiptID)
    return makeRequest("GET", url, nil)
}

// Create detail
func createDetail(detail map[string]interface{}) (*Response, error) {
    return makeRequest("POST", baseURL, detail)
}

// Update detail
func updateDetail(id int, detail map[string]interface{}) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("PUT", url, detail)
}

// Delete detail
func deleteDetail(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create new detail
    newDetail := map[string]interface{}{
        "ar_receipt_id":   101,
        "sales_order_id":  5001,
        "receipt_amount": 150000.50,
    }

    result, err := createDetail(newDetail)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created: %+v\n", result)

    // Get details for AR receipt 101
    details, err := getDetailsByARReceiptID(101)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Details: %+v\n", details)
}
```

---

## Best Practices

1. **Always Validate AR Receipt ID**: Ensure the `ar_receipt_id` exists before creating details
2. **Use Transactions**: When creating/updating multiple details, use database transactions
3. **Check Totals**: Ensure the sum of detail amounts matches the AR receipt total
4. **Soft Delete Benefits**: Soft deletes maintain data integrity and audit trails
5. **Filter Efficiently**: Use query parameters to filter large datasets on the server side
6. **Handle Errors Gracefully**: Always check response status codes and handle errors appropriately

---

## Notes

1. **Authentication**: All requests must include a valid JWT token in the Authorization header
2. **Tenant Isolation**: The `X-Tenant-Code` header is required for proper data isolation
3. **Soft Deletes**: Records with `deleted_at` set are automatically filtered out from queries
4. **Audit Trail**: The system automatically tracks who created, updated, and deleted each record
5. **Database Table**: Data is stored in the `alana.ar_receipt_detail` table
6. **Foreign Keys**: Ensure referenced AR receipts and sales orders exist before creating details

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD operations |
