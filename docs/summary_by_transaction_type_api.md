# Summary By Transaction Type API Documentation

## Overview
The Summary By Transaction Type API provides full CRUD (Create, Read, Update, Delete) operations for managing summary records grouped by transaction type. These summaries aggregate bookkeeping data by transaction type, providing quick access to financial summaries and reporting data.

**Base URL**: `/so/api/summary-by-transaction-type`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **Financial Summary Management**: Track totals by transaction type within bookkeeping records
- **Flexible Filtering**: Query by bookkeeping ID or transaction type ID
- **Custom Lookup**: Dedicated endpoint to retrieve all summaries for a specific bookkeeping record
- **Relational Data**: Automatically preloads related bookkeeping and transaction type data
- **Integer IDs**: Uses auto-incrementing integer primary keys

---

## Endpoints

### 1. Get All Summaries

Retrieves a list of all summaries by transaction type with optional filtering.

**Endpoint**: `GET /so/api/summary-by-transaction-type`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| bookkeeping_id | integer | No | Filter by bookkeeping ID |
| type_id | integer | No | Filter by transaction type ID |

**Response Codes**:
- `200 OK` - Successfully retrieved summaries
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Summaries retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 100,
      "type_id": 1,
      "total": 5000000.00,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z",
      "bookkeeping": {
        "id": 100,
        "location_id": "LOC001",
        "book_date": "2025-01-15",
        "opening": 10000000.00,
        "income": 8000000.00,
        "expanse": 3000000.00,
        "balance": 15000000.00
      },
      "type": {
        "id": 1,
        "name": "Penjualan"
      }
    },
    {
      "id": 2,
      "bookkeeping_id": 100,
      "type_id": 2,
      "total": 3000000.00,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:35:00Z",
      "updated_at": "2025-01-15T10:35:00Z",
      "bookkeeping": {
        "id": 100,
        "location_id": "LOC001",
        "book_date": "2025-01-15"
      },
      "type": {
        "id": 2,
        "name": "Pembelian"
      }
    }
  ]
}
```

**Example Requests**:

Get all summaries:
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-transaction-type" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by bookkeeping:
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-transaction-type?bookkeeping_id=100" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by transaction type:
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-transaction-type?type_id=1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Summary by ID

Retrieves a single summary by its unique ID.

**Endpoint**: `GET /so/api/summary-by-transaction-type/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the summary |

**Response Codes**:
- `200 OK` - Successfully retrieved the summary
- `400 Bad Request` - Invalid summary ID format
- `404 Not Found` - Summary not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Summary retrieved successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5000000.00,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z",
    "bookkeeping": {
      "id": 100,
      "location_id": "LOC001",
      "book_date": "2025-01-15",
      "opening": 10000000.00,
      "income": 8000000.00,
      "expanse": 3000000.00,
      "balance": 15000000.00
    },
    "type": {
      "id": 1,
      "name": "Penjualan"
    }
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "Summary not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-transaction-type/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Get Summaries by Bookkeeping ID

Retrieves all summary records for a specific bookkeeping record.

**Endpoint**: `GET /so/api/summary-by-transaction-type/by-bookkeeping/{bookkeeping_id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| bookkeeping_id | integer | Yes | The bookkeeping ID to get summaries for |

**Response Codes**:
- `200 OK` - Successfully retrieved summaries
- `400 Bad Request` - Invalid bookkeeping ID format
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Summaries retrieved successfully",
  "data": [
    {
      "id": 1,
      "bookkeeping_id": 100,
      "type_id": 1,
      "total": 5000000.00,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z",
      "type": {
        "id": 1,
        "name": "Penjualan"
      }
    },
    {
      "id": 2,
      "bookkeeping_id": 100,
      "type_id": 2,
      "total": 3000000.00,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:35:00Z",
      "updated_at": "2025-01-15T10:35:00Z",
      "type": {
        "id": 2,
        "name": "Pembelian"
      }
    }
  ]
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/summary-by-transaction-type/by-bookkeeping/100" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 4. Create Summary

Creates a new summary by transaction type record.

**Endpoint**: `POST /so/api/summary-by-transaction-type`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Request Body**:
```json
{
  "bookkeeping_id": 100,
  "type_id": 1,
  "total": 5000000.00
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| bookkeeping_id | integer | No | The bookkeeping ID this summary belongs to |
| type_id | integer | No | The transaction type ID |
| total | float | No | Total amount for this transaction type |

**Response Codes**:
- `201 Created` - Summary created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or server error

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Summary created successfully",
  "data": {
    "id": 3,
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5000000.00,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T14:30:00Z",
    "updated_at": "2025-01-15T14:30:00Z",
    "bookkeeping": {
      "id": 100,
      "location_id": "LOC001"
    },
    "type": {
      "id": 1,
      "name": "Penjualan"
    }
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid request body",
  "error": "invalid JSON format"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/summary-by-transaction-type" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5000000.00
  }'
```

---

### 5. Update Summary

Updates an existing summary record.

**Endpoint**: `PUT /so/api/summary-by-transaction-type/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the summary to update |

**Request Body**:
```json
{
  "bookkeeping_id": 100,
  "type_id": 1,
  "total": 5500000.00
}
```

**Response Codes**:
- `200 OK` - Summary updated successfully
- `400 Bad Request` - Invalid request body or ID format
- `404 Not Found` - Summary not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Summary updated successfully",
  "data": {
    "id": 1,
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5500000.00,
    "created_by": 1001,
    "updated_by": 1002,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T16:00:00Z",
    "bookkeeping": {
      "id": 100,
      "location_id": "LOC001"
    },
    "type": {
      "id": 1,
      "name": "Penjualan"
    }
  }
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/summary-by-transaction-type/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5500000.00
  }'
```

---

### 6. Delete Summary

Soft deletes a summary record.

**Endpoint**: `DELETE /so/api/summary-by-transaction-type/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the summary to delete |

**Response Codes**:
- `200 OK` - Summary deleted successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Summary not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Summary deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/summary-by-transaction-type/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### SummaryByTransactionType Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the summary (Primary Key, Auto-increment) |
| bookkeeping_id | integer | Foreign key referencing the bookkeeping record |
| type_id | integer | Foreign key referencing the transaction type |
| total | float | Total amount for this transaction type |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |
| bookkeeping | object | Related bookkeeping record (preloaded) |
| type | object | Related transaction type record (preloaded) |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Flexible Querying**: Filter by bookkeeping ID or transaction type ID
- **Custom Lookup**: Dedicated endpoint for retrieving all summaries of a specific bookkeeping record
- **Relational Data**: Automatically preloads related bookkeeping and transaction type data
- **Soft Delete Support**: Deleted records are marked rather than physically removed
- **Multi-Tenant**: Supports tenant-specific data isolation
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Audit Trail**: Tracks who created, updated, and deleted each record

---

## Business Logic

### Use Cases

1. **Financial Summary Reporting**: Generate quick summaries of bookkeeping data by transaction type
2. **Transaction Type Analysis**: Analyze totals grouped by transaction type
3. **Bookkeeping Overview**: Get a quick overview of a bookkeeping record's transaction type distribution
4. **Data Aggregation**: Aggregate detailed bookkeeping transactions by type
5. **Financial Dashboard**: Provide summary data for financial dashboards and reports

### Relationships

- Each summary belongs to one bookkeeping record (`bookkeeping_id`)
- Each summary references one transaction type (`type_id`)
- A bookkeeping record can have multiple summaries (one per transaction type)
- Each summary aggregates transactions of a specific type

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
- `400 Bad Request` - Invalid input or validation error
- `401 Unauthorized` - Missing or invalid JWT token
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error

---

## Usage Examples

### JavaScript (Fetch API)

```javascript
const BASE_URL = 'http://localhost:8080/so/api/summary-by-transaction-type';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all summaries for a bookkeeping record
async function getSummariesByBookkeeping(bookkeepingId) {
  const response = await fetch(
    `${BASE_URL}/by-bookkeeping/${bookkeepingId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Get summary by ID
async function getSummaryById(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Create new summary
async function createSummary(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update summary
async function updateSummary(id, data) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Delete summary
async function deleteSummary(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example: Create summary
const newSummary = await createSummary({
  bookkeeping_id: 100,
  type_id: 1,
  total: 5000000.00
});

console.log('Created Summary:', newSummary);

// Get all summaries for a bookkeeping record
const summaries = await getSummariesByBookkeeping(100);
console.log('Bookkeeping Summaries:', summaries);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/summary-by-transaction-type"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all summaries for a bookkeeping record
def get_summaries_by_bookkeeping(bookkeeping_id):
    url = f"{BASE_URL}/by-bookkeeping/{bookkeeping_id}"
    response = requests.get(url, headers=headers)
    return response.json()

# Get summary by ID
def get_summary_by_id(summary_id):
    url = f"{BASE_URL}/{summary_id}"
    response = requests.get(url, headers=headers)
    return response.json()

# Create new summary
def create_summary(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update summary
def update_summary(summary_id, data):
    url = f"{BASE_URL}/{summary_id}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Delete summary
def delete_summary(summary_id):
    url = f"{BASE_URL}/{summary_id}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example usage
summary_data = {
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5000000.00
}

# Create summary
new_summary = create_summary(summary_data)
print("Created Summary:", new_summary)

# Get all summaries for the bookkeeping record
summaries = get_summaries_by_bookkeeping(100)
print("Bookkeeping Summaries:", summaries)

# Update summary total
update_data = {
    "bookkeeping_id": 100,
    "type_id": 1,
    "total": 5500000.00
}
updated_summary = update_summary(new_summary['data']['id'], update_data)
print("Updated Summary:", updated_summary)
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
    baseURL    = "http://localhost:8080/so/api/summary-by-transaction-type"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type SummaryByTransactionTypeRequest struct {
    BookkeepingID *int     `json:"bookkeeping_id,omitempty"`
    TypeID        *int     `json:"type_id,omitempty"`
    Total         *float64 `json:"total,omitempty"`
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

// Get all summaries
func getAllSummaries() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get summaries by bookkeeping ID
func getSummariesByBookkeepingID(bookkeepingID int) (*Response, error) {
    url := fmt.Sprintf("%s/by-bookkeeping/%d", baseURL, bookkeepingID)
    return makeRequest("GET", url, nil)
}

// Get summary by ID
func getSummaryByID(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("GET", url, nil)
}

// Create summary
func createSummary(summary *SummaryByTransactionTypeRequest) (*Response, error) {
    return makeRequest("POST", baseURL, summary)
}

// Update summary
func updateSummary(id int, summary *SummaryByTransactionTypeRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("PUT", url, summary)
}

// Delete summary
func deleteSummary(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create summary
    bookkeepingID := 100
    typeID := 1
    total := 5000000.00

    newSummary := &SummaryByTransactionTypeRequest{
        BookkeepingID: &bookkeepingID,
        TypeID:        &typeID,
        Total:         &total,
    }

    result, err := createSummary(newSummary)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Summary: %+v\n", result)

    // Get all summaries for the bookkeeping record
    summaries, err := getSummariesByBookkeepingID(100)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Bookkeeping Summaries: %+v\n", summaries)
}
```

---

## Best Practices

1. **Data Integrity**: Ensure bookkeeping and transaction type records exist before creating summaries
2. **Use Filters**: When querying many summaries, use filters to reduce response size
3. **Batch Operations**: Consider creating summaries as part of bookkeeping record creation
4. **Update Strategy**: Update summaries when underlying bookkeeping details change
5. **Foreign Keys**: Validate referenced bookkeeping and transaction type IDs
6. **Decimal Precision**: Handle monetary values with appropriate decimal precision
7. **Error Handling**: Always validate data before sending requests
8. **Aggregate Consistency**: Ensure summary totals match the sum of related bookkeeping details

---

## Notes

1. **Authentication**: All requests must include a valid JWT token
2. **Tenant Isolation**: `X-Tenant-Code` header is required
3. **Optional Fields**: All fields are optional for create/update operations
4. **Integer IDs**: Uses auto-incrementing integer IDs (not UUIDs)
5. **Soft Deletes**: Deleted records are filtered out automatically
6. **Audit Trail**: System tracks who created, updated, and deleted records
7. **Database Table**: Data stored in `alana.summary_by_transaction_type` table
8. **Foreign Keys**: Ensure referenced bookkeeping and transaction types exist
9. **Relational Preloading**: Related data automatically loaded for convenience
10. **Dedicated Lookup**: Use `/by-bookkeeping/{bookkeeping_id}` endpoint for efficient bookkeeping-specific queries

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD operations |
