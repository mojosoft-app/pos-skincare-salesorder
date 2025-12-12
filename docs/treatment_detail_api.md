# Treatment Detail API Documentation

## Overview
The Treatment Detail API provides full CRUD (Create, Read, Update, Delete) operations for managing treatment detail records. Each detail record represents items or products used during a treatment session, including quantity tracking for inventory management.

**Base URL**: `/so/api/treatment-details`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **Item Usage Tracking**: Track items and products used during treatments
- **Flexible Filtering**: Query by treatment ID, item ID, or get all details
- **Custom Lookup**: Dedicated endpoint to retrieve all details for a specific treatment
- **Simple Structure**: Straightforward item tracking with quantity management
- **Integer IDs**: Uses auto-incrementing integer primary keys

---

## Endpoints

### 1. Get All Treatment Details

Retrieves a list of all treatment details with optional filtering.

**Endpoint**: `GET /so/api/treatment-details`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| treatment_id | integer | No | Filter by treatment ID |
| item_id | integer | No | Filter by item/product ID |

**Response Codes**:
- `200 OK` - Successfully retrieved treatment details
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment details retrieved successfully",
  "data": [
    {
      "id": 1,
      "treatment_id": 5001,
      "item_id": 101,
      "unit_id": 1,
      "quantity": 2,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "treatment_id": 5001,
      "item_id": 102,
      "unit_id": 1,
      "quantity": 1,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:35:00Z",
      "updated_at": "2025-01-15T10:35:00Z"
    }
  ]
}
```

**Example Requests**:

Get all details:
```bash
curl -X GET "http://localhost:8080/so/api/treatment-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by treatment:
```bash
curl -X GET "http://localhost:8080/so/api/treatment-details?treatment_id=5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by item:
```bash
curl -X GET "http://localhost:8080/so/api/treatment-details?item_id=101" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Treatment Detail by ID

Retrieves a single treatment detail by its unique ID.

**Endpoint**: `GET /so/api/treatment-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the treatment detail |

**Response Codes**:
- `200 OK` - Successfully retrieved the treatment detail
- `400 Bad Request` - Invalid detail ID format
- `404 Not Found` - Treatment detail not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment detail retrieved successfully",
  "data": {
    "id": 1,
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 2,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "Treatment detail not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/treatment-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Get Treatment Details by Treatment ID

Retrieves all detail records for a specific treatment.

**Endpoint**: `GET /so/api/treatment-details/by-treatment/{treatment_id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| treatment_id | integer | Yes | The treatment ID to get details for |

**Response Codes**:
- `200 OK` - Successfully retrieved treatment details
- `400 Bad Request` - Invalid treatment ID format
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment details retrieved successfully",
  "data": [
    {
      "id": 1,
      "treatment_id": 5001,
      "item_id": 101,
      "unit_id": 1,
      "quantity": 2,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "treatment_id": 5001,
      "item_id": 102,
      "unit_id": 1,
      "quantity": 1,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:35:00Z",
      "updated_at": "2025-01-15T10:35:00Z"
    }
  ]
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/treatment-details/by-treatment/5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 4. Create Treatment Detail

Creates a new treatment detail record.

**Endpoint**: `POST /so/api/treatment-details`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Request Body**:
```json
{
  "treatment_id": 5001,
  "item_id": 101,
  "unit_id": 1,
  "quantity": 2
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| treatment_id | integer | No | The treatment ID this detail belongs to |
| item_id | integer | No | The item/product ID |
| unit_id | integer | No | The unit of measurement ID |
| quantity | integer | **Yes** | Quantity of item used (required) |

**Response Codes**:
- `201 Created` - Treatment detail created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or server error

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Treatment detail created successfully",
  "data": {
    "id": 3,
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 2,
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
  "error": "Key: 'TreatmentDetailRequest.Quantity' Error:Field validation for 'Quantity' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/treatment-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 2
  }'
```

---

### 5. Update Treatment Detail

Updates an existing treatment detail record.

**Endpoint**: `PUT /so/api/treatment-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the treatment detail to update |

**Request Body**:
```json
{
  "treatment_id": 5001,
  "item_id": 101,
  "unit_id": 1,
  "quantity": 3
}
```

**Response Codes**:
- `200 OK` - Treatment detail updated successfully
- `400 Bad Request` - Invalid request body or ID format
- `404 Not Found` - Treatment detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment detail updated successfully",
  "data": {
    "id": 1,
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 3,
    "created_by": 1001,
    "updated_by": 1002,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T16:00:00Z"
  }
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/treatment-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 3
  }'
```

---

### 6. Delete Treatment Detail

Soft deletes a treatment detail record.

**Endpoint**: `DELETE /so/api/treatment-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the treatment detail to delete |

**Response Codes**:
- `200 OK` - Treatment detail deleted successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Treatment detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment detail deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/treatment-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### TreatmentDetail Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the treatment detail (Primary Key, Auto-increment) |
| treatment_id | integer | Foreign key referencing the treatment |
| item_id | integer | Foreign key referencing the item/product |
| unit_id | integer | Foreign key referencing the unit of measurement |
| quantity | integer | Quantity of item used (required) |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Flexible Querying**: Filter by treatment ID or item ID
- **Custom Lookup**: Dedicated endpoint for retrieving all details of a specific treatment
- **Simple Item Tracking**: Track items and quantities used in treatments
- **Soft Delete Support**: Deleted records are marked rather than physically removed
- **Multi-Tenant**: Supports tenant-specific data isolation
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Audit Trail**: Tracks who created, updated, and deleted each record

---

## Business Logic

### Use Cases

1. **Item Usage Recording**: Record items and products used during a treatment session
2. **Inventory Tracking**: Track consumption of items for inventory management
3. **Treatment Analysis**: Query all items used in a specific treatment
4. **Item Usage Reports**: Generate reports on item usage across treatments
5. **Quantity Adjustment**: Update quantities if recording errors occur

### Relationships

- Each treatment detail belongs to one treatment (`treatment_id`)
- Each treatment detail references one item/product (`item_id`)
- Each treatment detail has one unit of measurement (`unit_id`)
- A treatment can have multiple detail records (one-to-many)

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
const BASE_URL = 'http://localhost:8080/so/api/treatment-details';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all details for a treatment
async function getDetailsByTreatment(treatmentId) {
  const response = await fetch(
    `${BASE_URL}/by-treatment/${treatmentId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Get detail by ID
async function getDetailById(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
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

// Example: Create detail
const newDetail = await createDetail({
  treatment_id: 5001,
  item_id: 101,
  unit_id: 1,
  quantity: 2
});

console.log('Created Detail:', newDetail);

// Get all details for a treatment
const treatmentDetails = await getDetailsByTreatment(5001);
console.log('Treatment Details:', treatmentDetails);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/treatment-details"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all details for a treatment
def get_details_by_treatment(treatment_id):
    url = f"{BASE_URL}/by-treatment/{treatment_id}"
    response = requests.get(url, headers=headers)
    return response.json()

# Get detail by ID
def get_detail_by_id(detail_id):
    url = f"{BASE_URL}/{detail_id}"
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
detail_data = {
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 2
}

# Create detail
new_detail = create_detail(detail_data)
print("Created Detail:", new_detail)

# Get all details for the treatment
treatment_details = get_details_by_treatment(5001)
print("Treatment Details:", treatment_details)

# Update detail quantity
update_data = {
    "treatment_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "quantity": 3
}
updated_detail = update_detail(new_detail['data']['id'], update_data)
print("Updated Detail:", updated_detail)
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
    baseURL    = "http://localhost:8080/so/api/treatment-details"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type TreatmentDetailRequest struct {
    TreatmentID *int `json:"treatment_id,omitempty"`
    ItemID      *int `json:"item_id,omitempty"`
    UnitID      *int `json:"unit_id,omitempty"`
    Quantity    *int `json:"quantity"`
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

// Get details by treatment ID
func getDetailsByTreatmentID(treatmentID int) (*Response, error) {
    url := fmt.Sprintf("%s/by-treatment/%d", baseURL, treatmentID)
    return makeRequest("GET", url, nil)
}

// Get detail by ID
func getDetailByID(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("GET", url, nil)
}

// Create detail
func createDetail(detail *TreatmentDetailRequest) (*Response, error) {
    return makeRequest("POST", baseURL, detail)
}

// Update detail
func updateDetail(id int, detail *TreatmentDetailRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("PUT", url, detail)
}

// Delete detail
func deleteDetail(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create detail
    treatmentID := 5001
    itemID := 101
    unitID := 1
    quantity := 2

    newDetail := &TreatmentDetailRequest{
        TreatmentID: &treatmentID,
        ItemID:      &itemID,
        UnitID:      &unitID,
        Quantity:    &quantity,
    }

    result, err := createDetail(newDetail)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Detail: %+v\n", result)

    // Get all details for the treatment
    details, err := getDetailsByTreatmentID(5001)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Treatment Details: %+v\n", details)
}
```

---

## Best Practices

1. **Required Field**: Always provide `quantity` field - it's required
2. **Validate Quantity**: Ensure quantity is greater than 0
3. **Use Filters**: When querying many details, use filters to reduce response size
4. **Batch Operations**: Use the treatment endpoint for creating treatments with multiple details in one transaction
5. **Update Strategy**: Update individual details separately from the parent treatment
6. **Foreign Keys**: Ensure referenced treatments and items exist before creating details
7. **Error Handling**: Always validate required fields before sending requests
8. **Item Tracking**: Use this API to maintain accurate records of item usage for inventory management

---

## Notes

1. **Authentication**: All requests must include a valid JWT token
2. **Tenant Isolation**: `X-Tenant-Code` header is required
3. **Required Fields**: `quantity` is required for create/update operations
4. **Integer IDs**: Uses auto-incrementing integer IDs (not UUIDs)
5. **Soft Deletes**: Deleted records are filtered out automatically
6. **Audit Trail**: System tracks who created, updated, and deleted records
7. **Database Table**: Data stored in `alana.treatment_detail` table
8. **Foreign Keys**: Ensure referenced treatments, items, and units exist
9. **Simple Structure**: No automatic calculations - straightforward item tracking
10. **Dedicated Lookup**: Use `/by-treatment/{treatment_id}` endpoint for efficient treatment-specific queries

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD operations |
