# Reminded API Documentation

## Overview
The Reminded API provides read-only endpoints to retrieve reminded records from the system. These endpoints are used to access reminder information for various operations.

**Base URL**: `/so/api/reminded`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Endpoints

### 1. Get All Reminded Records

Retrieves a list of all reminded records from the database.

**Endpoint**: `GET /so/api/reminded`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**: None

**Response Codes**:
- `200 OK` - Successfully retrieved reminded records
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Reminded records retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Reminder Type 1",
      "created_by": 1001,
      "updated_by": 1001,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "name": "Reminder Type 2",
      "created_by": 1002,
      "updated_by": 1002,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-16T14:20:00Z",
      "updated_at": "2025-01-16T14:20:00Z"
    }
  ]
}
```

**Error Response** (500 Internal Server Error):
```json
{
  "status": "error",
  "message": "Database connection not found",
  "error": null
}
```

or

```json
{
  "status": "error",
  "message": "Failed to retrieve reminded records",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/reminded" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Reminded Record by ID

Retrieves a single reminded record by its unique ID.

**Endpoint**: `GET /so/api/reminded/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the reminded record |

**Response Codes**:
- `200 OK` - Successfully retrieved the reminded record
- `400 Bad Request` - Invalid reminded ID format
- `404 Not Found` - Reminded record not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Reminded record retrieved successfully",
  "data": {
    "id": 1,
    "name": "Reminder Type 1",
    "created_by": 1001,
    "updated_by": 1001,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid reminded ID",
  "error": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "Reminded record not found",
  "error": null
}
```

**Error Response** (500 Internal Server Error):
```json
{
  "status": "error",
  "message": "Database connection not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/reminded/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### Reminded Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the reminded record (Primary Key) |
| name | string | Name or description of the reminder type |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |

---

## Features

- **Soft Delete Support**: Deleted records (with `deleted_at` set) are automatically excluded from queries
- **Multi-Tenant**: Supports tenant-specific data isolation via `X-Tenant-Code` header
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Read-Only**: This API provides only GET operations (no create, update, or delete)

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
- `400 Bad Request` - Invalid input or malformed request
- `401 Unauthorized` - Missing or invalid JWT token
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error

---

## Usage Examples

### JavaScript (Fetch API)

```javascript
// Get all reminded records
async function getAllReminded() {
  const response = await fetch('http://localhost:8080/so/api/reminded', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer YOUR_JWT_TOKEN',
      'X-Tenant-Code': 'TENANT001',
      'Content-Type': 'application/json'
    }
  });

  const data = await response.json();
  return data;
}

// Get reminded by ID
async function getRemindedById(id) {
  const response = await fetch(`http://localhost:8080/so/api/reminded/${id}`, {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer YOUR_JWT_TOKEN',
      'X-Tenant-Code': 'TENANT001',
      'Content-Type': 'application/json'
    }
  });

  const data = await response.json();
  return data;
}
```

### Python (Requests)

```python
import requests

# Configuration
BASE_URL = "http://localhost:8080/so/api"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all reminded records
def get_all_reminded():
    response = requests.get(f"{BASE_URL}/reminded", headers=headers)
    return response.json()

# Get reminded by ID
def get_reminded_by_id(reminded_id):
    response = requests.get(f"{BASE_URL}/reminded/{reminded_id}", headers=headers)
    return response.json()

# Example usage
all_reminded = get_all_reminded()
print(all_reminded)

specific_reminded = get_reminded_by_id(1)
print(specific_reminded)
```

### Go

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

const (
    baseURL    = "http://localhost:8080/so/api"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type RemindedResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func getAllReminded() (*RemindedResponse, error) {
    req, err := http.NewRequest("GET", baseURL+"/reminded", nil)
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

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result RemindedResponse
    err = json.Unmarshal(body, &result)
    return &result, err
}

func getRemindedByID(id int) (*RemindedResponse, error) {
    url := fmt.Sprintf("%s/reminded/%d", baseURL, id)
    req, err := http.NewRequest("GET", url, nil)
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

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result RemindedResponse
    err = json.Unmarshal(body, &result)
    return &result, err
}
```

---

## Notes

1. **Authentication**: All requests must include a valid JWT token in the Authorization header
2. **Tenant Isolation**: The `X-Tenant-Code` header is required for proper data isolation
3. **Soft Deletes**: Records with `deleted_at` set are automatically filtered out
4. **Read-Only**: This API does not provide create, update, or delete operations
5. **Database Table**: Data is stored in the `alana.reminded` table

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with GET endpoints |
