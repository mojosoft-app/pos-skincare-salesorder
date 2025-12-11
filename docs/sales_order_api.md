# Sales Order API Documentation

## Overview
The Sales Order API provides full CRUD (Create, Read, Update, Delete) operations for managing sales orders. Sales orders represent customer orders and can contain multiple detail line items (products) and service records in a single transactional operation. This is the main entity for managing customer transactions in the system.

**Base URL**: `/so/api/sales-orders`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **UUID Primary Key**: Uses globally unique identifiers for sales orders
- **Nested Creation**: Create sales order with details and services in one API call
- **Transaction Support**: All nested creations use database transactions for data integrity
- **Auto-Preload**: Automatically loads status, details, and services relationships
- **Flexible Filtering**: Query by customer ID, status ID, or get all orders
- **Complex Data Model**: Supports payment tracking, vouchers, delivery costs, and more

---

## Endpoints

### 1. Get All Sales Orders

Retrieves a list of all sales orders with optional filtering and relationship preloading.

**Endpoint**: `GET /so/api/sales-orders`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Query Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| status_id | integer | No | Filter by sales order status ID |
| customer_id | integer | No | Filter by customer ID |
| page | integer | No | Page number (default: 1) |
| limit | integer | No | Items per page (default: 10) |

**Response Codes**:
- `200 OK` - Successfully retrieved sales orders
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales orders retrieved successfully",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "location_id": 1,
      "customer_id": 1001,
      "doc_date": "2025-01-15",
      "inv_number": "INV-2025-001",
      "address": "123 Main St, City",
      "delivery_cost": 50000.00,
      "total_amount": 500000.00,
      "total_payment": 300000.00,
      "outstanding": 200000.00,
      "total_voucher": 0,
      "voucher_number": null,
      "posted_date": "2025-01-15",
      "additional_cost": 25000.00,
      "previous_payment": 0,
      "fully_paid": false,
      "note": "Customer order for January",
      "status_id": 1,
      "created_by": 1001,
      "updated_by": null,
      "deleted_by": null,
      "deleted_at": null,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z",
      "status": {
        "id": 1,
        "name": "Pending"
      },
      "details": [
        {
          "id": 1,
          "sales_order_id": null,
          "item_id": 101,
          "item_name": "Product A",
          "quantity": 2,
          "price": 150000.00,
          "item_total": 300000.00,
          "discount_pct": 0
        }
      ],
      "services": [
        {
          "id": 1,
          "sales_order_detail_id": 1,
          "service_id": 50,
          "service_name": "Service A",
          "treated": false,
          "schedule": "2025-01-20T14:00:00Z"
        }
      ]
    }
  ]
}
```

**Example Requests**:

Get all orders:
```bash
curl -X GET "http://localhost:8080/so/api/sales-orders" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by customer:
```bash
curl -X GET "http://localhost:8080/so/api/sales-orders?customer_id=1001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by status:
```bash
curl -X GET "http://localhost:8080/so/api/sales-orders?status_id=1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Sales Order by ID

Retrieves a single sales order by its unique UUID, including all relationships (status, details, services).

**Endpoint**: `GET /so/api/sales-orders/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The unique UUID of the sales order |

**Response Codes**:
- `200 OK` - Successfully retrieved the sales order
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - Sales order not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order retrieved successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "doc_date": "2025-01-15",
    "inv_number": "INV-2025-001",
    "address": "123 Main St, City",
    "delivery_cost": 50000.00,
    "total_amount": 500000.00,
    "total_payment": 300000.00,
    "outstanding": 200000.00,
    "total_voucher": 0,
    "voucher_number": null,
    "posted_date": "2025-01-15",
    "additional_cost": 25000.00,
    "previous_payment": 0,
    "fully_paid": false,
    "note": "Customer order for January",
    "status_id": 1,
    "created_by": 1001,
    "updated_by": null,
    "deleted_by": null,
    "deleted_at": null,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z",
    "status": {
      "id": 1,
      "name": "Pending",
      "created_at": "2025-01-01T00:00:00Z"
    },
    "details": [
      {
        "id": 1,
        "item_id": 101,
        "item_name": "Product A",
        "quantity": 2,
        "price": 150000.00,
        "item_total": 300000.00,
        "discount_pct": 0,
        "used_sessions": 0
      },
      {
        "id": 2,
        "item_id": 102,
        "item_name": "Product B",
        "quantity": 1,
        "price": 200000.00,
        "item_total": 200000.00,
        "discount_pct": 0,
        "used_sessions": 0
      }
    ],
    "services": [
      {
        "id": 1,
        "sales_order_detail_id": 1,
        "service_id": 50,
        "service_name": "Service A",
        "treated": false,
        "schedule": "2025-01-20T14:00:00Z",
        "reminded_id": 1
      }
    ]
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "status": "error",
  "message": "Invalid sales order ID",
  "error": null
}
```

**Error Response** (404 Not Found):
```json
{
  "status": "error",
  "message": "Sales order not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/sales-orders/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Create Sales Order

Creates a new sales order with optional details and services in a single transactional operation.

**Endpoint**: `POST /so/api/sales-orders`

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
  "doc_date": "2025-01-15",
  "inv_number": "INV-2025-001",
  "address": "123 Main St, City",
  "delivery_cost": 50000.00,
  "total_amount": 500000.00,
  "total_payment": 300000.00,
  "outstanding": 200000.00,
  "total_voucher": 0,
  "voucher_number": null,
  "posted_date": "2025-01-15",
  "additional_cost": 25000.00,
  "previous_payment": 0,
  "fully_paid": false,
  "note": "Customer order for January",
  "status_id": 1,
  "details": [
    {
      "item_id": 101,
      "unit_id": 1,
      "promoter_id": 50,
      "item_name": "Product A",
      "quantity": 2,
      "price": 150000.00,
      "item_total": 300000.00,
      "discount_pct": 0,
      "used_sessions": 0
    },
    {
      "item_id": 102,
      "unit_id": 1,
      "item_name": "Product B",
      "quantity": 1,
      "price": 200000.00,
      "item_total": 200000.00,
      "discount_pct": 0,
      "used_sessions": 0
    }
  ],
  "services": [
    {
      "sales_order_detail_id": null,
      "service_id": 50,
      "treatment_id": 10,
      "reminded_id": 1,
      "service_name": "Service A",
      "treated": false,
      "schedule": "2025-01-20T14:00:00Z"
    }
  ]
}
```

**Request Body Schema**:

**Main Sales Order Fields**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| location_id | integer | No | The location/branch ID |
| customer_id | integer | **Yes** | The customer ID (required) |
| doc_date | string | No | Document date (YYYY-MM-DD) |
| inv_number | string | No | Invoice/order number |
| address | string | No | Delivery/billing address |
| delivery_cost | number | No | Delivery/shipping cost |
| total_amount | number | No | Total order amount |
| total_payment | number | No | Total payment received |
| outstanding | number | No | Outstanding balance |
| total_voucher | number | No | Total voucher amount |
| voucher_number | string | No | Voucher number/code |
| posted_date | string | No | Posted date (YYYY-MM-DD) |
| additional_cost | number | No | Additional costs |
| previous_payment | number | No | Previous payment amount |
| fully_paid | boolean | No | Whether order is fully paid |
| note | string | No | Additional notes |
| status_id | integer | No | Order status ID |
| details | array | No | Array of order detail line items |
| services | array | No | Array of service records |

**Detail Object Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| item_id | integer | No | Item/product ID |
| unit_id | integer | No | Unit of measurement ID |
| promoter_id | integer | No | Promoter/salesperson ID |
| item_name | string | No | Item name/description |
| quantity | integer | No | Quantity ordered |
| price | number | No | Unit price |
| item_total | number | No | Line total amount |
| discount_pct | integer | No | Discount percentage |
| used_sessions | integer | No | Sessions used |

**Service Object Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| sales_order_detail_id | integer | No | Related detail ID (can be null) |
| service_id | integer | No | Service ID |
| treatment_id | integer | No | Treatment ID |
| message_log_detail_id | string | No | Message log detail ID |
| reminded_id | integer | No | Reminder type ID |
| service_name | string | No | Service name/description |
| treated | boolean | No | Whether service is completed |
| schedule | string | No | Scheduled date/time (ISO 8601) |

**Transaction Flow**:
1. Database transaction begins
2. Sales order is created (UUID auto-generated)
3. All detail records are created
4. All service records are created
5. Transaction commits (or rolls back on error)
6. Full order with relationships is returned

**Response Codes**:
- `201 Created` - Sales order created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or transaction failure

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Sales order created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "location_id": 1,
    "customer_id": 1001,
    "total_amount": 500000.00,
    "status_id": 1,
    "created_by": 1001,
    "created_at": "2025-01-15T14:30:00Z",
    "status": {
      "id": 1,
      "name": "Pending"
    },
    "details": [
      {
        "id": 1,
        "item_name": "Product A",
        "quantity": 2,
        "price": 150000.00,
        "item_total": 300000.00
      }
    ],
    "services": [
      {
        "id": 1,
        "service_name": "Service A",
        "treated": false
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
  "error": "Key: 'CreateSalesOrderRequest.CustomerID' Error:Field validation for 'CustomerID' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/sales-orders" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1001,
    "inv_number": "INV-2025-001",
    "total_amount": 500000.00,
    "status_id": 1,
    "details": [
      {
        "item_id": 101,
        "item_name": "Product A",
        "quantity": 2,
        "price": 150000.00,
        "item_total": 300000.00
      }
    ],
    "services": [
      {
        "service_id": 50,
        "service_name": "Service A",
        "treated": false,
        "schedule": "2025-01-20T14:00:00Z"
      }
    ]
  }'
```

---

### 4. Update Sales Order

Updates an existing sales order header (does not update details or services - use their respective endpoints).

**Endpoint**: `PUT /so/api/sales-orders/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the sales order to update |

**Request Body**:
```json
{
  "location_id": 1,
  "customer_id": 1001,
  "doc_date": "2025-01-15",
  "inv_number": "INV-2025-001-UPDATED",
  "address": "456 New St, City",
  "delivery_cost": 75000.00,
  "total_amount": 550000.00,
  "total_payment": 400000.00,
  "outstanding": 150000.00,
  "fully_paid": false,
  "note": "Updated order notes",
  "status_id": 2
}
```

**Response Codes**:
- `200 OK` - Sales order updated successfully
- `400 Bad Request` - Invalid request body or UUID format
- `404 Not Found` - Sales order not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order updated successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "customer_id": 1001,
    "inv_number": "INV-2025-001-UPDATED",
    "total_amount": 550000.00,
    "status_id": 2,
    "created_by": 1001,
    "updated_by": 1002,
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T16:00:00Z",
    "status": {
      "id": 2,
      "name": "Confirmed"
    },
    "details": [],
    "services": []
  }
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/so/api/sales-orders/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1001,
    "inv_number": "INV-2025-001-UPDATED",
    "total_amount": 550000.00,
    "status_id": 2
  }'
```

---

### 5. Delete Sales Order

Soft deletes a sales order record.

**Endpoint**: `DELETE /so/api/sales-orders/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | UUID | Yes | The UUID of the sales order to delete |

**Response Codes**:
- `200 OK` - Sales order deleted successfully
- `400 Bad Request` - Invalid UUID format
- `404 Not Found` - Sales order not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/sales-orders/550e8400-e29b-41d4-a716-446655440000" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### SalesOrder Object

| Field | Type | Description |
|-------|------|-------------|
| id | UUID | Unique identifier (Primary Key, auto-generated) |
| location_id | integer | Location/branch ID |
| customer_id | integer | Customer ID (required) |
| doc_date | date | Document date |
| inv_number | string | Invoice/order number |
| address | string | Delivery/billing address |
| delivery_cost | number | Delivery/shipping cost |
| total_amount | number | Total order amount |
| total_payment | number | Total payment received |
| outstanding | number | Outstanding balance amount |
| total_voucher | number | Total voucher/discount amount |
| voucher_number | string | Voucher number/code |
| posted_date | date | Posted date |
| additional_cost | number | Additional costs |
| previous_payment | number | Previous payment amount |
| fully_paid | boolean | Payment status flag |
| note | string | Additional notes |
| status_id | integer | Order status ID |
| created_by | integer | User ID who created |
| updated_by | integer | User ID who updated |
| deleted_by | integer | User ID who deleted |
| deleted_at | timestamp | Soft delete timestamp |
| created_at | timestamp | Creation timestamp |
| updated_at | timestamp | Update timestamp |
| status | object | Status object (preloaded) |
| details | array | Array of detail line items (preloaded) |
| services | array | Array of service records (preloaded) |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, Delete functionality
- **UUID Primary Key**: Globally unique identifiers
- **Nested Creation**: Create order with details and services in one transaction
- **Transaction Support**: Automatic rollback on errors
- **Auto-Preload**: Automatically loads related status, details, and services
- **Flexible Filtering**: Filter by customer or status
- **Payment Tracking**: Track payments, outstanding balances, and vouchers
- **Complex Data Model**: Support for delivery costs, additional costs, and more
- **Soft Delete**: Maintains data integrity and audit trails
- **Multi-Tenant**: Tenant-specific data isolation
- **JWT Authentication**: Secure access control
- **Audit Trail**: Tracks who created, updated, and deleted records

---

## Business Logic

### Use Cases

1. **Order Creation**: Create complete orders with products and services in one transaction
2. **Order Management**: Track order status from pending to completed
3. **Payment Tracking**: Monitor payments, outstanding balances, and fully paid status
4. **Service Scheduling**: Schedule and track services related to order items
5. **Order Modification**: Update order details as requirements change
6. **Financial Reporting**: Track delivery costs, vouchers, and additional charges

### Relationships

- Each sales order belongs to one customer (`customer_id`)
- Each sales order has one location (`location_id`)
- Each sales order has one status (`status_id`)
- Each sales order can have multiple details (one-to-many with `SalesOrderDetail`)
- Each sales order can have multiple services (one-to-many with `SalesOrderService`)

### Transaction Handling

When creating a sales order with nested data:
1. Transaction begins
2. Sales order created (UUID auto-generated)
3. All detail records created sequentially
4. All service records created sequentially
5. If any step fails, entire transaction rolls back
6. On success, transaction commits
7. Full order with relationships loaded and returned

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
const BASE_URL = 'http://localhost:8080/so/api/sales-orders';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all sales orders for a customer
async function getOrdersByCustomer(customerId) {
  const response = await fetch(
    `${BASE_URL}?customer_id=${customerId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Get order by UUID
async function getOrderById(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'GET',
    headers: headers
  });
  return await response.json();
}

// Create complete order with details and services
async function createOrder(data) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Update order
async function updateOrder(uuid, data) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'PUT',
    headers: headers,
    body: JSON.stringify(data)
  });
  return await response.json();
}

// Delete order
async function deleteOrder(uuid) {
  const response = await fetch(`${BASE_URL}/${uuid}`, {
    method: 'DELETE',
    headers: headers
  });
  return await response.json();
}

// Example: Create complete order
const newOrder = await createOrder({
  customer_id: 1001,
  inv_number: "INV-2025-001",
  total_amount: 500000.00,
  status_id: 1,
  details: [
    {
      item_id: 101,
      item_name: "Product A",
      quantity: 2,
      price: 150000.00,
      item_total: 300000.00
    },
    {
      item_id: 102,
      item_name: "Product B",
      quantity: 1,
      price: 200000.00,
      item_total: 200000.00
    }
  ],
  services: [
    {
      service_id: 50,
      service_name: "Service A",
      treated: false,
      schedule: "2025-01-20T14:00:00Z"
    }
  ]
});

console.log('Created Order:', newOrder);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/sales-orders"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all orders for a customer
def get_orders_by_customer(customer_id):
    params = {"customer_id": customer_id}
    response = requests.get(BASE_URL, headers=headers, params=params)
    return response.json()

# Get order by UUID
def get_order_by_id(order_uuid):
    url = f"{BASE_URL}/{order_uuid}"
    response = requests.get(url, headers=headers)
    return response.json()

# Create complete order
def create_order(data):
    response = requests.post(BASE_URL, headers=headers, json=data)
    return response.json()

# Update order
def update_order(order_uuid, data):
    url = f"{BASE_URL}/{order_uuid}"
    response = requests.put(url, headers=headers, json=data)
    return response.json()

# Delete order
def delete_order(order_uuid):
    url = f"{BASE_URL}/{order_uuid}"
    response = requests.delete(url, headers=headers)
    return response.json()

# Example: Create complete order with nested data
order_data = {
    "customer_id": 1001,
    "inv_number": "INV-2025-001",
    "address": "123 Main St, City",
    "total_amount": 500000.00,
    "total_payment": 300000.00,
    "outstanding": 200000.00,
    "status_id": 1,
    "details": [
        {
            "item_id": 101,
            "item_name": "Product A",
            "quantity": 2,
            "price": 150000.00,
            "item_total": 300000.00,
            "discount_pct": 0
        },
        {
            "item_id": 102,
            "item_name": "Product B",
            "quantity": 1,
            "price": 200000.00,
            "item_total": 200000.00,
            "discount_pct": 0
        }
    ],
    "services": [
        {
            "service_id": 50,
            "service_name": "Service A",
            "treated": False,
            "schedule": "2025-01-20T14:00:00Z"
        }
    ]
}

# Create order
new_order = create_order(order_data)
print("Created Order:", new_order)

# Get all orders for customer
customer_orders = get_orders_by_customer(1001)
print("Customer Orders:", customer_orders)
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
    baseURL    = "http://localhost:8080/so/api/sales-orders"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type SalesOrderDetail struct {
    ItemID       *int     `json:"item_id,omitempty"`
    UnitID       *int     `json:"unit_id,omitempty"`
    PromoterID   *int     `json:"promoter_id,omitempty"`
    ItemName     *string  `json:"item_name,omitempty"`
    Quantity     *int     `json:"quantity,omitempty"`
    Price        *float64 `json:"price,omitempty"`
    ItemTotal    *float64 `json:"item_total,omitempty"`
    DiscountPct  *int     `json:"discount_pct,omitempty"`
    UsedSessions *int     `json:"used_sessions,omitempty"`
}

type SalesOrderService struct {
    SalesOrderDetailID *int    `json:"sales_order_detail_id,omitempty"`
    ServiceID          *int    `json:"service_id,omitempty"`
    TreatmentID        *int    `json:"treatment_id,omitempty"`
    RemindedID         *int    `json:"reminded_id,omitempty"`
    ServiceName        *string `json:"service_name,omitempty"`
    Treated            *bool   `json:"treated,omitempty"`
    Schedule           *string `json:"schedule,omitempty"`
}

type SalesOrderRequest struct {
    LocationID      *int                `json:"location_id,omitempty"`
    CustomerID      *int                `json:"customer_id"`
    DocDate         *string             `json:"doc_date,omitempty"`
    InvNumber       *string             `json:"inv_number,omitempty"`
    Address         *string             `json:"address,omitempty"`
    DeliveryCost    *float64            `json:"delivery_cost,omitempty"`
    TotalAmount     *float64            `json:"total_amount,omitempty"`
    TotalPayment    *float64            `json:"total_payment,omitempty"`
    Outstanding     *float64            `json:"outstanding,omitempty"`
    TotalVoucher    *float64            `json:"total_voucher,omitempty"`
    VoucherNumber   *string             `json:"voucher_number,omitempty"`
    PostedDate      *string             `json:"posted_date,omitempty"`
    AdditionalCost  *float64            `json:"additional_cost,omitempty"`
    PreviousPayment *float64            `json:"previous_payment,omitempty"`
    FullyPaid       *bool               `json:"fully_paid,omitempty"`
    Note            *string             `json:"note,omitempty"`
    StatusID        *int                `json:"status_id,omitempty"`
    Details         []SalesOrderDetail  `json:"details,omitempty"`
    Services        []SalesOrderService `json:"services,omitempty"`
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

// Get all orders
func getAllOrders() (*Response, error) {
    return makeRequest("GET", baseURL, nil)
}

// Get order by UUID
func getOrderByID(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("GET", url, nil)
}

// Create order
func createOrder(order *SalesOrderRequest) (*Response, error) {
    return makeRequest("POST", baseURL, order)
}

// Update order
func updateOrder(uuid string, order *SalesOrderRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("PUT", url, order)
}

// Delete order
func deleteOrder(uuid string) (*Response, error) {
    url := fmt.Sprintf("%s/%s", baseURL, uuid)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create complete order
    customerID := 1001
    invNumber := "INV-2025-001"
    totalAmount := 500000.00
    statusID := 1
    itemName1 := "Product A"
    qty1 := 2
    price1 := 150000.00
    total1 := 300000.00
    serviceName := "Service A"
    treated := false
    schedule := "2025-01-20T14:00:00Z"

    newOrder := &SalesOrderRequest{
        CustomerID:  &customerID,
        InvNumber:   &invNumber,
        TotalAmount: &totalAmount,
        StatusID:    &statusID,
        Details: []SalesOrderDetail{
            {
                ItemName:  &itemName1,
                Quantity:  &qty1,
                Price:     &price1,
                ItemTotal: &total1,
            },
        },
        Services: []SalesOrderService{
            {
                ServiceName: &serviceName,
                Treated:     &treated,
                Schedule:    &schedule,
            },
        },
    }

    result, err := createOrder(newOrder)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Order: %+v\n", result)
}
```

---

## Best Practices

1. **Use Transactions**: Nested creation uses transactions automatically - ensure proper error handling
2. **Validate Customer**: Verify customer exists before creating order
3. **Calculate Totals**: Ensure detail totals match order total_amount
4. **UUID Format**: Always validate UUID format before API calls
5. **Update Strategy**: Update order header separately from details/services
6. **Status Management**: Use appropriate status transitions
7. **Payment Tracking**: Keep total_payment, outstanding, and fully_paid in sync
8. **Nested Creation**: Prefer creating complete orders in one call when possible
9. **Filter Large Datasets**: Use query parameters to reduce response size
10. **Handle Errors**: Implement retry logic for transaction failures

---

## Notes

1. **Authentication**: JWT token required in Authorization header
2. **Tenant Isolation**: `X-Tenant-Code` header required
3. **UUID Format**: Order IDs are UUIDs (e.g., `550e8400-e29b-41d4-a716-446655440000`)
4. **Soft Deletes**: Deleted records filtered out automatically
5. **Audit Trail**: System tracks creation, updates, and deletions
6. **Database Table**: Data stored in `alana.sales_order` table
7. **Nested Creation**: Can create order with details and services in one transaction
8. **Update Limitation**: PUT updates header only, not nested details/services
9. **Auto-Preload**: Always loads status, details, and services relationships
10. **Required Field**: Only `customer_id` is required

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD and nested creation |
