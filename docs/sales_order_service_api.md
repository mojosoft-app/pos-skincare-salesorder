# Sales Order Service API Documentation

## Overview
The Sales Order Service API provides full CRUD (Create, Read, Update, Delete) operations for managing service records associated with sales orders. Services represent treatments, appointments, or service-based activities linked to order items. This API includes a special endpoint for marking services as treated/completed.

**Base URL**: `/so/api/sales-order-services`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Treatment Tracking**: Track whether services have been completed (treated)
- **Service Scheduling**: Schedule services with date/time information
- **Flexible Filtering**: Query by sales order ID, treated status, or service ID
- **Status Management**: Special endpoint to mark services as treated
- **Multi-Relationship**: Link services to sales orders, details, treatments, and reminders

---

## Endpoints

### 1. Get All Sales Order Services

Retrieves a list of all sales order services with optional filtering.

**Endpoint**: `GET /so/api/sales-order-services`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| sales_order_id | integer | No | Filter by sales order ID |
| treated | boolean | No | Filter by treatment status (true/false) |
| service_id | integer | No | Filter by service type ID |

**Response Codes**:
- `200 OK` - Successfully retrieved sales order services
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order services retrieved successfully",
  "data": [
    {
      "id": 1,
      "sales_order_id": 5001,
      "sales_order_detail_id": 1,
      "service_id": 50,
      "treatment_id": 10,
      "message_log_detail_id": "msg-12345",
      "reminded_id": 1,
      "service_name": "Hair Treatment",
      "treated": false,
      "schedule": "2025-01-20T14:00:00Z",
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "sales_order_id": 5001,
      "sales_order_detail_id": 2,
      "service_id": 51,
      "treatment_id": 11,
      "message_log_detail_id": null,
      "reminded_id": 2,
      "service_name": "Facial Treatment",
      "treated": true,
      "schedule": "2025-01-18T10:00:00Z",
      "created_by": 1001,
      "updated_by": 1002,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T11:00:00Z",
      "updated_at": "2025-01-18T10:30:00Z"
    }
  ]
}
```

**Example Requests**:

Get all services:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-services" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by sales order:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-services?sales_order_id=5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by treatment status (completed services):
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-services?treated=true" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by service type:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-services?service_id=50" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Sales Order Service by ID

Retrieves a single sales order service by its unique ID.

**Endpoint**: `GET /so/api/sales-order-services/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the sales order service |

**Response Codes**:
- `200 OK` - Successfully retrieved the sales order service
- `400 Bad Request` - Invalid service ID format
- `404 Not Found` - Sales order service not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order service retrieved successfully",
  "data": {
    "id": 1,
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "treatment_id": 10,
    "message_log_detail_id": "msg-12345",
    "reminded_id": 1,
    "service_name": "Hair Treatment",
    "treated": false,
    "schedule": "2025-01-20T14:00:00Z",
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
  "message": "Sales order service not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-services/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Create Sales Order Service

Creates a new sales order service record.

**Endpoint**: `POST /so/api/sales-order-services`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Request Body**:
```json
{
  "sales_order_id": 5001,
  "sales_order_detail_id": 1,
  "service_id": 50,
  "treatment_id": 10,
  "message_log_detail_id": "msg-12345",
  "reminded_id": 1,
  "service_name": "Hair Treatment",
  "treated": false,
  "schedule": "2025-01-20T14:00:00Z"
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| sales_order_id | integer | No | The sales order ID |
| sales_order_detail_id | integer | No | The related order detail ID |
| service_id | integer | No | The service type ID |
| treatment_id | integer | No | The treatment type ID |
| message_log_detail_id | string | No | Message log detail reference ID |
| reminded_id | integer | No | The reminder type ID |
| service_name | string | No | Service name or description |
| treated | boolean | No | Treatment completion status (default: false) |
| schedule | string | No | Scheduled date/time (ISO 8601 format) |

**Response Codes**:
- `201 Created` - Sales order service created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or server error

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Sales order service created successfully",
  "data": {
    "id": 3,
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "treatment_id": 10,
    "message_log_detail_id": "msg-12345",
    "reminded_id": 1,
    "service_name": "Hair Treatment",
    "treated": false,
    "schedule": "2025-01-20T14:00:00Z",
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
  "error": "detailed validation error message"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/sales-order-services" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "service_name": "Hair Treatment",
    "treated": false,
    "schedule": "2025-01-20T14:00:00Z"
  }'
```

---

### 4. Update Sales Order Service

Updates an existing sales order service record.

**Endpoint**: `PUT /so/api/sales-order-services/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the sales order service to update |

**Request Body**:
```json
{
  "sales_order_id": 5001,
  "sales_order_detail_id": 1,
  "service_id": 50,
  "treatment_id": 10,
  "message_log_detail_id": "msg-12345-updated",
  "reminded_id": 1,
  "service_name": "Hair Treatment - Updated",
  "treated": true,
  "schedule": "2025-01-21T15:00:00Z"
}
```

**Response Codes**:
- `200 OK` - Sales order service updated successfully
- `400 Bad Request` - Invalid request body or ID format
- `404 Not Found` - Sales order service not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order service updated successfully",
  "data": {
    "id": 1,
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "treatment_id": 10,
    "message_log_detail_id": "msg-12345-updated",
    "reminded_id": 1,
    "service_name": "Hair Treatment - Updated",
    "treated": true,
    "schedule": "2025-01-21T15:00:00Z",
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
curl -X PUT "http://localhost:8080/so/api/sales-order-services/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "sales_order_id": 5001,
    "service_id": 50,
    "service_name": "Hair Treatment - Updated",
    "treated": true,
    "schedule": "2025-01-21T15:00:00Z"
  }'
```

---

### 5. Delete Sales Order Service

Soft deletes a sales order service record.

**Endpoint**: `DELETE /so/api/sales-order-services/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the sales order service to delete |

**Response Codes**:
- `200 OK` - Sales order service deleted successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Sales order service not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order service deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/sales-order-services/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 6. Mark Service as Treated

A convenient endpoint to mark a service as treated/completed without sending the full update payload.

**Endpoint**: `PATCH /so/api/sales-order-services/{id}/mark-treated`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the sales order service to mark as treated |

**Request Body**: None required

**Response Codes**:
- `200 OK` - Service marked as treated successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Sales order service not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Service marked as treated successfully",
  "data": {
    "id": 1,
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "treatment_id": 10,
    "service_name": "Hair Treatment",
    "treated": true,
    "schedule": "2025-01-20T14:00:00Z",
    "created_by": 1001,
    "updated_by": 1002,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-20T14:30:00Z"
  }
}
```

**Example Request**:
```bash
curl -X PATCH "http://localhost:8080/so/api/sales-order-services/1/mark-treated" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### SalesOrderService Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier (Primary Key, Auto-increment) |
| sales_order_id | integer | Foreign key to sales order |
| sales_order_detail_id | integer | Foreign key to sales order detail |
| service_id | integer | Foreign key to service type |
| treatment_id | integer | Foreign key to treatment type |
| message_log_detail_id | string | Reference to message log detail |
| reminded_id | integer | Foreign key to reminder type |
| service_name | string | Service name or description |
| treated | boolean | Whether service is completed (default: false) |
| schedule | string | Scheduled date/time (ISO 8601 format) |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when soft deleted (null if active) |
| created_at | timestamp | Timestamp when created |
| updated_at | timestamp | Timestamp when last updated |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Treatment Status Tracking**: Track completion of services with `treated` flag
- **Convenient Status Update**: Dedicated endpoint for marking services as treated
- **Service Scheduling**: Schedule services with date/time information
- **Flexible Filtering**: Filter by sales order, treated status, or service type
- **Multi-Relationship Support**: Link to orders, details, treatments, reminders, and messages
- **Soft Delete Support**: Maintains data integrity and audit trails
- **Multi-Tenant**: Tenant-specific data isolation
- **JWT Authentication**: Secure access control
- **Audit Trail**: Tracks who created, updated, and deleted records

---

## Business Logic

### Use Cases

1. **Service Scheduling**: Schedule treatments or appointments for customers
2. **Treatment Tracking**: Monitor which services have been completed
3. **Appointment Management**: Link services to specific order items
4. **Reminder System**: Connect services to reminder notifications
5. **Service Completion**: Mark services as treated when appointments are fulfilled
6. **Service Reporting**: Query services by status, type, or order for reporting

### Relationships

- Each service can belong to one sales order (`sales_order_id`)
- Each service can belong to one sales order detail (`sales_order_detail_id`)
- Each service references one service type (`service_id`)
- Each service references one treatment type (`treatment_id`)
- Each service can have one reminder type (`reminded_id`)
- Each service can reference one message log detail (`message_log_detail_id`)

### Treatment Status Workflow

```
1. Service Created → treated = false (default)
2. Service Scheduled → schedule date/time set
3. Appointment Occurs → Use PATCH endpoint to mark as treated
4. Service Completed → treated = true
```

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
const BASE_URL = 'http://localhost:8080/so/api/sales-order-services';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all services for a sales order
async function getServicesBySalesOrder(salesOrderId) {
  const response = await fetch(
    `${BASE_URL}?sales_order_id=${salesOrderId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Get untreated services
async function getUntreatedServices() {
  const response = await fetch(
    `${BASE_URL}?treated=false`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Create new service
async function createService(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update service
async function updateService(id, data) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Mark service as treated (convenient method)
async function markAsTreated(id) {
  const response = await fetch(`${BASE_URL}/${id}/mark-treated`, {
    method: 'PATCH',
    headers: headers
  });
  return await response.json();
}

// Delete service
async function deleteService(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example: Create and later mark as treated
const newService = await createService({
  sales_order_id: 5001,
  sales_order_detail_id: 1,
  service_id: 50,
  service_name: "Hair Treatment",
  treated: false,
  schedule: "2025-01-20T14:00:00Z"
});

console.log('Created Service:', newService);

// Later, when appointment is completed
const treated = await markAsTreated(newService.data.id);
console.log('Service Marked as Treated:', treated);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/sales-order-services"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all services for a sales order
def get_services_by_sales_order(sales_order_id):
    params = {"sales_order_id": sales_order_id}
    response = requests.get(BASE_URL, headers=headers, params=params)
    return response.json()

# Get services by treatment status
def get_services_by_status(treated):
    params = {"treated": str(treated).lower()}
    response = requests.get(BASE_URL, headers=headers, params=params)
    return response.json()

# Create new service
def create_service(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update service
def update_service(service_id, data):
    url = f"{BASE_URL}/{service_id}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Mark service as treated
def mark_as_treated(service_id):
    url = f"{BASE_URL}/{service_id}/mark-treated"
    response = requests.patch(url, headers=headers)
    return response.json()

# Delete service
def delete_service(service_id):
    url = f"{BASE_URL}/{service_id}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example: Schedule service and mark as treated later
service_data = {
    "sales_order_id": 5001,
    "sales_order_detail_id": 1,
    "service_id": 50,
    "service_name": "Hair Treatment",
    "treated": False,
    "schedule": "2025-01-20T14:00:00Z"
}

# Create service
new_service = create_service(service_data)
print("Created Service:", new_service)

# Get all untreated services
untreated = get_services_by_status(False)
print("Untreated Services:", untreated)

# Mark as treated when appointment is done
if new_service['status'] == 'success':
    service_id = new_service['data']['id']
    treated = mark_as_treated(service_id)
    print("Service Marked as Treated:", treated)
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
    baseURL    = "http://localhost:8080/so/api/sales-order-services"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type SalesOrderServiceRequest struct {
    SalesOrderID       *int    `json:"sales_order_id,omitempty"`
    SalesOrderDetailID *int    `json:"sales_order_detail_id,omitempty"`
    ServiceID          *int    `json:"service_id,omitempty"`
    TreatmentID        *int    `json:"treatment_id,omitempty"`
    MessageLogDetailID *string `json:"message_log_detail_id,omitempty"`
    RemindedID         *int    `json:"reminded_id,omitempty"`
    ServiceName        *string `json:"service_name,omitempty"`
    Treated            *bool   `json:"treated,omitempty"`
    Schedule           *string `json:"schedule,omitempty"`
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

// Get all services
func getAllServices() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get service by ID
func getServiceByID(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("GET", url, nil)
}

// Create service
func createService(service *SalesOrderServiceRequest) (*Response, error) {
    return makeRequest("POST", baseURL, service)
}

// Update service
func updateService(id int, service *SalesOrderServiceRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("PUT", url, service)
}

// Mark as treated
func markAsTreated(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d/mark-treated", baseURL, id)
    return makeRequest("PATCH", url, nil)
}

// Delete service
func deleteService(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create service and mark as treated
    salesOrderID := 5001
    detailID := 1
    serviceID := 50
    serviceName := "Hair Treatment"
    treated := false
    schedule := "2025-01-20T14:00:00Z"

    newService := &SalesOrderServiceRequest{
        SalesOrderID:       &salesOrderID,
        SalesOrderDetailID: &detailID,
        ServiceID:          &serviceID,
        ServiceName:        &serviceName,
        Treated:            &treated,
        Schedule:           &schedule,
    }

    result, err := createService(newService)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Service: %+v\n", result)

    // Mark as treated
    if result.Status == "success" {
        // Extract service ID from result
        // Note: You'll need to type assert the data to extract the ID
        treated, err := markAsTreated(1) // Use actual ID from result
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Printf("Service Marked as Treated: %+v\n", treated)
    }
}
```

---

## Best Practices

1. **Use Mark-Treated Endpoint**: Use the PATCH endpoint for simple status updates instead of full PUT
2. **Schedule Format**: Always use ISO 8601 format for schedule dates (e.g., "2025-01-20T14:00:00Z")
3. **Filter Efficiently**: Use query parameters to reduce response sizes
4. **Track Status**: Regularly query untreated services for appointment management
5. **Link Relationships**: Properly link services to orders, details, and treatments
6. **Handle Reminders**: Use reminded_id to trigger appropriate notifications
7. **Error Handling**: Implement proper error handling for 404 and 500 errors
8. **Batch Operations**: Consider querying by sales_order_id to get all related services
9. **Status Updates**: Update `treated` status promptly after appointments
10. **Audit Trail**: Leverage created_by/updated_by for accountability

---

## Notes

1. **Authentication**: JWT token required in Authorization header
2. **Tenant Isolation**: `X-Tenant-Code` header required
3. **Soft Deletes**: Deleted records filtered out automatically
4. **Default Status**: Services are created with `treated = false` by default
5. **Audit Trail**: System tracks creation, updates, and deletions
6. **Database Table**: Data stored in `alana.sales_order_service` table
7. **Schedule Format**: Use ISO 8601 format for schedule field
8. **Status Tracking**: `treated` field is boolean (true/false)
9. **Convenient Endpoint**: Use PATCH `/mark-treated` for quick status updates
10. **Multi-Relationship**: Services can link to multiple related entities

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD and mark-treated endpoint |
