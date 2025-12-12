# Treatment API Documentation

## Overview
The Treatment API provides full CRUD (Create, Read, Update, Delete) operations for managing treatment records. Treatments represent medical/beauty treatment sessions linked to patients, doctors, nurses, beauticians, and sales orders. Each treatment can contain multiple detail line items for tracking items/products used during the treatment.

**Base URL**: `/so/api/treatments`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **UUID Primary Key**: Uses globally unique identifiers for treatments
- **Nested Creation**: Create treatment with details in one API call
- **Transaction Support**: All nested creations use database transactions for data integrity
- **Auto-Preload**: Automatically loads details relationships
- **Flexible Filtering**: Query by status ID, patient ID, or doctor ID
- **Multi-Entity Linking**: Links to sales orders, services, patients, doctors, nurses, and beauticians
- **Date Tracking**: Track document date and posted date separately

---

## Endpoints

### 1. Get All Treatments

Retrieves a list of all treatments with optional filtering and relationship preloading.

**Endpoint**: `GET /so/api/treatments`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| status_id | integer | No | Filter by treatment status ID |
| patient_id | integer | No | Filter by patient ID |
| doctor_id | integer | No | Filter by doctor ID |
| page | integer | No | Page number (default: 1) |
| limit | integer | No | Items per page (default: 10) |

**Response Codes**:
- `200 OK` - Successfully retrieved treatments
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatments retrieved successfully",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "location_id": 1,
      "customer_id": 1001,
      "sales_order_id": 500,
      "sales_order_detail_id": 501,
      "sales_order_service_id": 502,
      "service_id": 50,
      "patient_id": 2001,
      "doctor_id": 301,
      "nurse_id": 302,
      "beautician_id": 303,
      "doc_number": "TRT-2025-001",
      "doc_date": "2025-01-15",
      "posted_date": "2025-01-15",
      "service_text": "Facial treatment session",
      "note": "Patient responded well to treatment",
      "status_id": 1,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z",
      "details": [
        {
          "id": "660e8400-e29b-41d4-a716-446655440001",
          "treatment_id": "550e8400-e29b-41d4-a716-446655440000",
          "item_id": 101,
          "unit_id": 1,
          "quantity": 2,
          "created_by": 1001,
          "created_at": "2025-01-15T10:30:00Z"
        }
      ]
    }
  ]
}
```

**Example Requests**:

Get all treatments:
```bash
curl -X GET "http://localhost:8080/so/api/treatments" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by patient:
```bash
curl -X GET "http://localhost:8080/so/api/treatments?patient_id=2001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by doctor:
```bash
curl -X GET "http://localhost:8080/so/api/treatments?doctor_id=301" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by status:
```bash
curl -X GET "http://localhost:8080/so/api/treatments?status_id=1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Treatment by ID

Retrieves a single treatment by its unique UUID, including all relationships (details).

**Endpoint**: `GET /so/api/treatments/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The unique UUID of the treatment |

**Response Codes**:
- `200 OK` - Successfully retrieved the treatment
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - Treatment not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment retrieved successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "sales_order_id": 500,
    "sales_order_detail_id": 501,
    "sales_order_service_id": 502,
    "service_id": 50,
    "patient_id": 2001,
    "doctor_id": 301,
    "nurse_id": 302,
    "beautician_id": 303,
    "doc_number": "TRT-2025-001",
    "doc_date": "2025-01-15",
    "posted_date": "2025-01-15",
    "service_text": "Facial treatment session",
    "note": "Patient responded well to treatment",
    "status_id": 1,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z",
    "details": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "treatment_id": "550e8400-e29b-41d4-a716-446655440000",
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
        "id": "660e8400-e29b-41d4-a716-446655440002",
        "treatment_id": "550e8400-e29b-41d4-a716-446655440000",
        "item_id": 102,
        "unit_id": 1,
        "quantity": 1,
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
  "message": "Invalid treatment ID",
  "error": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "Treatment not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/treatments/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Create Treatment

Creates a new treatment with optional details in a single transactional operation.

**Endpoint**: `POST /so/api/treatments`

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
  "sales_order_id": 500,
  "sales_order_detail_id": 501,
  "sales_order_service_id": 502,
  "service_id": 50,
  "patient_id": 2001,
  "doctor_id": 301,
  "nurse_id": 302,
  "beautician_id": 303,
  "doc_number": "TRT-2025-001",
  "doc_date": "2025-01-15",
  "posted_date": "2025-01-15",
  "service_text": "Facial treatment session",
  "note": "Patient responded well to treatment",
  "status_id": 1,
  "details": [
    {
      "item_id": 101,
      "unit_id": 1,
      "quantity": 2
    },
    {
      "item_id": 102,
      "unit_id": 1,
      "quantity": 1
    }
  ]
}
```

**Request Body Schema**:

**Main Treatment Fields**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| location_id | integer | No | The location/branch ID |
| customer_id | integer | No | The customer ID |
| sales_order_id | integer | No | Related sales order ID |
| sales_order_detail_id | integer | No | Related sales order detail ID |
| sales_order_service_id | integer | No | Related sales order service ID |
| service_id | integer | No | Service type ID |
| patient_id | integer | No | Patient ID receiving treatment |
| doctor_id | integer | No | Doctor ID performing treatment |
| nurse_id | integer | No | Nurse ID assisting treatment |
| beautician_id | integer | No | Beautician ID performing treatment |
| doc_number | string | No | Document/treatment number |
| doc_date | string | No | Document date (YYYY-MM-DD) |
| posted_date | string | No | Posted date (YYYY-MM-DD) |
| service_text | string | No | Service description text |
| note | string | No | Additional treatment notes |
| status_id | integer | No | Treatment status ID |
| details | array | No | Array of treatment detail line items |

**Detail Object Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| item_id | integer | No | Item/product ID used in treatment |
| unit_id | integer | No | Unit of measurement ID |
| quantity | integer | No | Quantity of item used |

**Transaction Flow**:
1. Database transaction begins
2. Treatment is created (UUID auto-generated)
3. All detail records are created
4. Transaction commits (or rolls back on error)
5. Full treatment with relationships is returned

**Response Codes**:
- `201 Created` - Treatment created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or transaction failure

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Treatment created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "patient_id": 2001,
    "doctor_id": 301,
    "doc_number": "TRT-2025-001",
    "doc_date": "2025-01-15",
    "service_text": "Facial treatment session",
    "status_id": 1,
    "created_by": 1001,
    "created_at": "2025-01-15T14:30:00Z",
    "details": [
      {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "treatment_id": "550e8400-e29b-41d4-a716-446655440000",
        "item_id": 101,
        "unit_id": 1,
        "quantity": 2,
        "created_by": 1001,
        "created_at": "2025-01-15T14:30:00Z"
      },
      {
        "id": "660e8400-e29b-41d4-a716-446655440002",
        "treatment_id": "550e8400-e29b-41d4-a716-446655440000",
        "item_id": 102,
        "unit_id": 1,
        "quantity": 1,
        "created_by": 1001,
        "created_at": "2025-01-15T14:30:00Z"
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
  "error": "detailed error information"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/treatments" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 2001,
    "doctor_id": 301,
    "doc_number": "TRT-2025-001",
    "doc_date": "2025-01-15",
    "service_text": "Facial treatment session",
    "status_id": 1,
    "details": [
      {
        "item_id": 101,
        "unit_id": 1,
        "quantity": 2
      }
    ]
  }'
```

---

### 4. Update Treatment

Updates an existing treatment header (does not update details - use their respective endpoints).

**Endpoint**: `PUT /so/api/treatments/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the treatment to update |

**Request Body**:
```json
{
  "location_id": 1,
  "customer_id": 1001,
  "sales_order_id": 500,
  "sales_order_detail_id": 501,
  "sales_order_service_id": 502,
  "service_id": 50,
  "patient_id": 2001,
  "doctor_id": 301,
  "nurse_id": 302,
  "beautician_id": 303,
  "doc_number": "TRT-2025-001-UPDATED",
  "doc_date": "2025-01-15",
  "posted_date": "2025-01-16",
  "service_text": "Facial treatment session - updated",
  "note": "Patient responded very well to treatment",
  "status_id": 2
}
```

**Response Codes**:
- `200 OK` - Treatment updated successfully
- `400 Bad Request` - Invalid request body or UUID format
- `404 Not Found` - Treatment not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment updated successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "patient_id": 2001,
    "doctor_id": 301,
    "doc_number": "TRT-2025-001-UPDATED",
    "service_text": "Facial treatment session - updated",
    "status_id": 2,
    "created_by": 1001,
    "updated_by": 1002,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T16:00:00Z",
    "details": []
  }
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/treatments/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 2001,
    "doctor_id": 301,
    "doc_number": "TRT-2025-001-UPDATED",
    "service_text": "Facial treatment session - updated",
    "status_id": 2
  }'
```

---

### 5. Delete Treatment

Soft deletes a treatment record.

**Endpoint**: `DELETE /so/api/treatments/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the treatment to delete |

**Response Codes**:
- `200 OK` - Treatment deleted successfully
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - Treatment not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Treatment deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/treatments/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### Treatment Object

| Field | Type | Description |
|-------|------|-------------|
| id | UUID | Unique identifier (Primary Key, auto-generated) |
| location_id | integer | Location/branch ID |
| customer_id | integer | Customer ID |
| sales_order_id | integer | Related sales order ID |
| sales_order_detail_id | integer | Related sales order detail ID |
| sales_order_service_id | integer | Related sales order service ID |
| service_id | integer | Service type ID |
| patient_id | integer | Patient ID receiving treatment |
| doctor_id | integer | Doctor ID performing treatment |
| nurse_id | integer | Nurse ID assisting treatment |
| beautician_id | integer | Beautician ID performing treatment |
| doc_number | string | Document/treatment number |
| doc_date | date | Document date |
| posted_date | date | Posted date |
| service_text | string | Service description text |
| note | string | Additional treatment notes |
| status_id | integer | Treatment status ID |
| created_by | integer | User ID who created |
| updated_by | integer | User ID who updated |
| deleted_by | integer | User ID who deleted |
| deleted_at | timestamp | Soft delete timestamp |
| created_at | timestamp | Creation timestamp |
| updated_at | timestamp | Update timestamp |
| details | array | Array of treatment detail items (preloaded) |

### TreatmentDetail Object

| Field | Type | Description |
|-------|------|-------------|
| id | UUID | Unique identifier (Primary Key, auto-generated) |
| treatment_id | UUID | Parent treatment ID (Foreign Key) |
| item_id | integer | Item/product ID used in treatment |
| unit_id | integer | Unit of measurement ID |
| quantity | integer | Quantity of item used |
| created_by | integer | User ID who created |
| updated_by | integer | User ID who updated |
| deleted_by | integer | User ID who deleted |
| deleted_at | timestamp | Soft delete timestamp |
| created_at | timestamp | Creation timestamp |
| updated_at | timestamp | Update timestamp |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, Delete functionality
- **UUID Primary Key**: Globally unique identifiers
- **Nested Creation**: Create treatment with details in one transaction
- **Transaction Support**: Automatic rollback on errors
- **Auto-Preload**: Automatically loads related details
- **Flexible Filtering**: Filter by patient, doctor, or status
- **Multi-Entity Linking**: Links to sales orders, services, patients, and staff
- **Soft Delete**: Maintains data integrity and audit trails
- **Multi-Tenant**: Tenant-specific data isolation
- **JWT Authentication**: Secure access control
- **Audit Trail**: Tracks who created, updated, and deleted records

---

## Business Logic

### Use Cases

1. **Treatment Recording**: Record medical/beauty treatment sessions with patient and staff details
2. **Item Tracking**: Track items and products used during treatment
3. **Staff Assignment**: Assign doctors, nurses, and beauticians to treatments
4. **Sales Order Integration**: Link treatments to sales orders and services
5. **Treatment Status**: Track treatment status from scheduled to completed
6. **Historical Records**: Maintain complete treatment history for patients

### Relationships

- Each treatment belongs to one patient (`patient_id`)
- Each treatment can have one doctor (`doctor_id`)
- Each treatment can have one nurse (`nurse_id`)
- Each treatment can have one beautician (`beautician_id`)
- Each treatment belongs to one customer (`customer_id`)
- Each treatment has one location (`location_id`)
- Each treatment can link to a sales order (`sales_order_id`)
- Each treatment can link to a sales order service (`sales_order_service_id`)
- Each treatment has one status (`status_id`)
- Each treatment can have multiple details (one-to-many with `TreatmentDetail`)

### Transaction Handling

When creating a treatment with nested data:
1. Transaction begins
2. Treatment created (UUID auto-generated)
3. All detail records created sequentially
4. If any step fails, entire transaction rolls back
5. On success, transaction commits
6. Full treatment with relationships loaded and returned

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
- `400 Bad Request` - Invalid input or UUID format
- `401 Unauthorized` - Missing or invalid JWT token
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server-side error or transaction failure

---

## Usage Examples

### JavaScript (Fetch API)

```javascript
const BASE_URL = 'http://localhost:8080/so/api/treatments';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all treatments for a patient
async function getTreatmentsByPatient(patientId) {
  const response = await fetch(
    `${BASE_URL}?patient_id=${patientId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Get treatment by UUID
async function getTreatmentById(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Create treatment with details
async function createTreatment(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update treatment
async function updateTreatment(uuid, data) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Delete treatment
async function deleteTreatment(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example: Create complete treatment
const newTreatment = await createTreatment({
  patient_id: 2001,
  doctor_id: 301,
  nurse_id: 302,
  doc_number: "TRT-2025-001",
  doc_date: "2025-01-15",
  service_text: "Facial treatment session",
  note: "First session",
  status_id: 1,
  details: [
    {
      item_id: 101,
      unit_id: 1,
      quantity: 2
    },
    {
      item_id: 102,
      unit_id: 1,
      quantity: 1
    }
  ]
});

console.log('Created Treatment:', newTreatment);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/treatments"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all treatments for a patient
def get_treatments_by_patient(patient_id):
    params = {"patient_id": patient_id}
    response = requests.get(BASE_URL, headers=headers, params=params)
    return response.json()

# Get treatment by UUID
def get_treatment_by_id(treatment_uuid):
    url = f"{BASE_URL}/{treatment_uuid}"
    response = requests.get(url, headers=headers)
    return response.json()

# Create treatment
def create_treatment(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update treatment
def update_treatment(treatment_uuid, data):
    url = f"{BASE_URL}/{treatment_uuid}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Delete treatment
def delete_treatment(treatment_uuid):
    url = f"{BASE_URL}/{treatment_uuid}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example: Create treatment with details
treatment_data = {
    "patient_id": 2001,
    "doctor_id": 301,
    "nurse_id": 302,
    "doc_number": "TRT-2025-001",
    "doc_date": "2025-01-15",
    "service_text": "Facial treatment session",
    "note": "First session",
    "status_id": 1,
    "details": [
        {
            "item_id": 101,
            "unit_id": 1,
            "quantity": 2
        },
        {
            "item_id": 102,
            "unit_id": 1,
            "quantity": 1
        }
    ]
}

# Create treatment
new_treatment = create_treatment(treatment_data)
print("Created Treatment:", new_treatment)

# Get all treatments for patient
patient_treatments = get_treatments_by_patient(2001)
print("Patient Treatments:", patient_treatments)
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
    baseURL    = "http://localhost:8080/so/api/treatments"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type TreatmentDetail struct {
    ItemID   *int `json:"item_id,omitempty"`
    UnitID   *int `json:"unit_id,omitempty"`
    Quantity *int `json:"quantity,omitempty"`
}

type TreatmentRequest struct {
    LocationID          *int               `json:"location_id,omitempty"`
    CustomerID          *int               `json:"customer_id,omitempty"`
    SalesOrderID        *int               `json:"sales_order_id,omitempty"`
    SalesOrderDetailID  *int               `json:"sales_order_detail_id,omitempty"`
    SalesOrderServiceID *int               `json:"sales_order_service_id,omitempty"`
    ServiceID           *int               `json:"service_id,omitempty"`
    PatientID           *int               `json:"patient_id,omitempty"`
    DoctorID            *int               `json:"doctor_id,omitempty"`
    NurseID             *int               `json:"nurse_id,omitempty"`
    BeauticianID        *int               `json:"beautician_id,omitempty"`
    DocNumber           *string            `json:"doc_number,omitempty"`
    DocDate             *string            `json:"doc_date,omitempty"`
    PostedDate          *string            `json:"posted_date,omitempty"`
    ServiceText         *string            `json:"service_text,omitempty"`
    Note                *string            `json:"note,omitempty"`
    StatusID            *int               `json:"status_id,omitempty"`
    Details             []TreatmentDetail  `json:"details,omitempty"`
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

// Get all treatments
func getAllTreatments() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get treatment by UUID
func getTreatmentByID(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("GET", url, nil)
}

// Create treatment
func createTreatment(treatment *TreatmentRequest) (*Response, error) {
    return makeRequest("POST", baseURL, treatment)
}

// Update treatment
func updateTreatment(uuid string, treatment *TreatmentRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("PUT", url, treatment)
}

// Delete treatment
func deleteTreatment(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create treatment with details
    patientID := 2001
    doctorID := 301
    nurseID := 302
    docNumber := "TRT-2025-001"
    docDate := "2025-01-15"
    serviceText := "Facial treatment session"
    note := "First session"
    statusID := 1
    itemID1 := 101
    unitID := 1
    qty1 := 2
    itemID2 := 102
    qty2 := 1

    newTreatment := &TreatmentRequest{
        PatientID:   &patientID,
        DoctorID:    &doctorID,
        NurseID:     &nurseID,
        DocNumber:   &docNumber,
        DocDate:     &docDate,
        ServiceText: &serviceText,
        Note:        &note,
        StatusID:    &statusID,
        Details: []TreatmentDetail{
            {
                ItemID:   &itemID1,
                UnitID:   &unitID,
                Quantity: &qty1,
            },
            {
                ItemID:   &itemID2,
                UnitID:   &unitID,
                Quantity: &qty2,
            },
        },
    }

    result, err := createTreatment(newTreatment)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Treatment: %+v\n", result)
}
```

---

## Best Practices

1. **Use Transactions**: Nested creation uses transactions automatically - ensure proper error handling
2. **Link to Sales Orders**: Link treatments to sales orders for billing integration
3. **Track Staff**: Always assign appropriate staff (doctor, nurse, beautician) to treatments
4. **UUID Format**: Always validate UUID format before API calls
5. **Update Strategy**: Update treatment header separately from details
6. **Status Management**: Use appropriate status transitions for treatment workflow
7. **Date Tracking**: Use doc_date for treatment date and posted_date for posting to system
8. **Item Tracking**: Record all items used during treatment in details
9. **Filter Large Datasets**: Use query parameters to reduce response size
10. **Handle Errors**: Implement retry logic for transaction failures

---

## Notes

1. **Authentication**: JWT token required in Authorization header
2. **Tenant Isolation**: `X-Tenant-Code` header required
3. **UUID Format**: Treatment IDs are UUIDs (e.g., `550e8400-e29b-41d4-a716-446655440000`)
4. **Soft Deletes**: Deleted records filtered out automatically
5. **Audit Trail**: System tracks creation, updates, and deletions
6. **Database Table**: Data stored in `alana.treatment` table
7. **Nested Creation**: Can create treatment with details in one transaction
8. **Update Limitation**: PUT updates header only, not nested details
9. **Auto-Preload**: Always loads details relationships
10. **All Fields Optional**: No required fields - provide only relevant data

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD and nested creation |
