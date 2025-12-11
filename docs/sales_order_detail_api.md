# Sales Order Detail API Documentation

## Overview
The Sales Order Detail API provides full CRUD (Create, Read, Update, Delete) operations for managing sales order line items. Each detail record represents a single item/product line within a sales order, including quantity, price, discounts, and automatic total calculations.

**Base URL**: `/so/api/sales-order-details`

**Authentication**: JWT Token required (via Authorization header)

**Content Type**: `application/json`

---

## Key Features

- **Automatic Calculation**: Item totals are automatically calculated from quantity × price with discount application
- **Flexible Filtering**: Query by sales order ID, item ID, or get all details
- **Custom Lookup**: Dedicated endpoint to retrieve all details for a specific sales order
- **Discount Support**: Percentage-based discounts automatically applied to totals
- **Session Tracking**: Track used sessions for service-based items

---

## Endpoints

### 1. Get All Sales Order Details

Retrieves a list of all sales order details with optional filtering.

**Endpoint**: `GET /so/api/sales-order-details`

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
| item_id | integer | No | Filter by item/product ID |

**Response Codes**:
- `200 OK` - Successfully retrieved sales order details
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order details retrieved successfully",
  "data": [
    {
      "id": 1,
      "sales_order_id": 5001,
      "item_id": 101,
      "unit_id": 1,
      "promoter_id": 50,
      "item_name": "Product A",
      "quantity": 2,
      "price": 150000.00,
      "item_total": 270000.00,
      "discount_pct": 10,
      "used_sessions": 0,
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
      "item_id": 102,
      "unit_id": 1,
      "promoter_id": null,
      "item_name": "Product B",
      "quantity": 1,
      "price": 200000.00,
      "item_total": 200000.00,
      "discount_pct": 0,
      "used_sessions": 0,
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
curl -X GET "http://localhost:8080/so/api/sales-order-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by sales order:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-details?sales_order_id=5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

Filter by item:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-details?item_id=101" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 2. Get Sales Order Detail by ID

Retrieves a single sales order detail by its unique ID.

**Endpoint**: `GET /so/api/sales-order-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The unique ID of the sales order detail |

**Response Codes**:
- `200 OK` - Successfully retrieved the sales order detail
- `400 Bad Request` - Invalid detail ID format
- `404 Not Found` - Sales order detail not found
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order detail retrieved successfully",
  "data": {
    "id": 1,
    "sales_order_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "promoter_id": 50,
    "item_name": "Product A",
    "quantity": 2,
    "price": 150000.00,
    "item_total": 270000.00,
    "discount_pct": 10,
    "used_sessions": 0,
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
  "message": "Sales order detail not found",
  "error": null
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 3. Get Sales Order Details by Sales Order ID

Retrieves all detail records for a specific sales order.

**Endpoint**: `GET /so/api/sales-order-details/by-sales-order/{sales_order_id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| sales_order_id | integer | Yes | The sales order ID to get details for |

**Response Codes**:
- `200 OK` - Successfully retrieved sales order details
- `400 Bad Request` - Invalid sales order ID format
- `500 Internal Server Error` - Database connection error or query failure

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order details retrieved successfully",
  "data": [
    {
      "id": 1,
      "sales_order_id": 5001,
      "item_id": 101,
      "unit_id": 1,
      "promoter_id": 50,
      "item_name": "Product A",
      "quantity": 2,
      "price": 150000.00,
      "item_total": 270000.00,
      "discount_pct": 10,
      "used_sessions": 0,
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
      "item_id": 102,
      "unit_id": 1,
      "promoter_id": null,
      "item_name": "Product B",
      "quantity": 1,
      "price": 200000.00,
      "item_total": 200000.00,
      "discount_pct": 0,
      "used_sessions": 0,
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
curl -X GET "http://localhost:8080/so/api/sales-order-details/by-sales-order/5001" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

### 4. Create Sales Order Detail

Creates a new sales order detail record with automatic total calculation.

**Endpoint**: `POST /so/api/sales-order-details`

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
  "item_id": 101,
  "unit_id": 1,
  "promoter_id": 50,
  "item_name": "Product A",
  "quantity": 2,
  "price": 150000.00,
  "discount_pct": 10,
  "used_sessions": 0
}
```

**Request Body Schema**:
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| sales_order_id | integer | No | The sales order ID this detail belongs to |
| item_id | integer | No | The item/product ID |
| unit_id | integer | No | The unit of measurement ID |
| promoter_id | integer | No | The promoter/salesperson ID |
| item_name | string | No | Name or description of the item |
| quantity | integer | **Yes** | Quantity ordered (required) |
| price | number | **Yes** | Unit price (required) |
| item_total | number | No | Total amount (auto-calculated if omitted) |
| discount_pct | integer | No | Discount percentage (0-100) |
| used_sessions | integer | No | Number of sessions used (for service items) |

**Calculation Logic**:
If `item_total` is not provided, it will be automatically calculated as:
```
item_total = (quantity × price) - ((quantity × price) × discount_pct / 100)
```

**Response Codes**:
- `201 Created` - Sales order detail created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Database error or server error

**Success Response** (201 Created):
```json
{
  "status": "success",
  "message": "Sales order detail created successfully",
  "data": {
    "id": 3,
    "sales_order_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "promoter_id": 50,
    "item_name": "Product A",
    "quantity": 2,
    "price": 150000.00,
    "item_total": 270000.00,
    "discount_pct": 10,
    "used_sessions": 0,
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
  "error": "Key: 'SalesOrderDetailRequest.Quantity' Error:Field validation for 'Quantity' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/so/api/sales-order-details" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "sales_order_id": 5001,
    "item_id": 101,
    "item_name": "Product A",
    "quantity": 2,
    "price": 150000.00,
    "discount_pct": 10
  }'
```

---

### 5. Update Sales Order Detail

Updates an existing sales order detail record with automatic total recalculation.

**Endpoint**: `PUT /so/api/sales-order-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the sales order detail to update |

**Request Body**:
```json
{
  "sales_order_id": 5001,
  "item_id": 101,
  "unit_id": 1,
  "promoter_id": 50,
  "item_name": "Product A - Updated",
  "quantity": 3,
  "price": 150000.00,
  "discount_pct": 15,
  "used_sessions": 1
}
```

**Calculation Logic**:
If `item_total` is not provided in the request, it will be automatically recalculated based on the updated quantity, price, and discount.

**Response Codes**:
- `200 OK` - Sales order detail updated successfully
- `400 Bad Request` - Invalid request body or ID format
- `404 Not Found` - Sales order detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order detail updated successfully",
  "data": {
    "id": 1,
    "sales_order_id": 5001,
    "item_id": 101,
    "unit_id": 1,
    "promoter_id": 50,
    "item_name": "Product A - Updated",
    "quantity": 3,
    "price": 150000.00,
    "item_total": 382500.00,
    "discount_pct": 15,
    "used_sessions": 1,
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
curl -X PUT "http://localhost:8080/so/api/sales-order-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json" \
  -d '{
    "sales_order_id": 5001,
    "item_id": 101,
    "item_name": "Product A - Updated",
    "quantity": 3,
    "price": 150000.00,
    "discount_pct": 15
  }'
```

---

### 6. Delete Sales Order Detail

Soft deletes a sales order detail record.

**Endpoint**: `DELETE /so/api/sales-order-details/{id}`

**Headers**:
```
Authorization: Bearer <jwt_token>
X-Tenant-Code: <tenant_code>
Content-Type: application/json
```

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | The ID of the sales order detail to delete |

**Response Codes**:
- `200 OK` - Sales order detail deleted successfully
- `400 Bad Request` - Invalid ID format
- `404 Not Found` - Sales order detail not found
- `500 Internal Server Error` - Database error or server error

**Success Response** (200 OK):
```json
{
  "status": "success",
  "message": "Sales order detail deleted successfully",
  "data": null
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/so/api/sales-order-details/1" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "X-Tenant-Code: TENANT001" \
  -H "Content-Type: application/json"
```

---

## Data Model

### SalesOrderDetail Object

| Field | Type | Description |
|-------|------|-------------|
| id | integer | Unique identifier for the sales order detail (Primary Key, Auto-increment) |
| sales_order_id | integer | Foreign key referencing the sales order |
| item_id | integer | Foreign key referencing the item/product |
| unit_id | integer | Foreign key referencing the unit of measurement |
| promoter_id | integer | Foreign key referencing the promoter/salesperson |
| item_name | string | Name or description of the item |
| quantity | integer | Quantity ordered (required) |
| price | number | Unit price per item (required) |
| item_total | number | Total line amount (auto-calculated) |
| discount_pct | integer | Discount percentage applied (0-100) |
| used_sessions | integer | Number of sessions used (for service-based items) |
| created_by | integer | User ID who created the record |
| updated_by | integer | User ID who last updated the record |
| deleted_by | integer | User ID who deleted the record (null if not deleted) |
| deleted_at | timestamp | Timestamp when the record was soft deleted (null if active) |
| created_at | timestamp | Timestamp when the record was created |
| updated_at | timestamp | Timestamp when the record was last updated |

---

## Features

- **Full CRUD Operations**: Complete Create, Read, Update, and Delete functionality
- **Automatic Calculations**: Item totals automatically calculated from quantity, price, and discount
- **Flexible Querying**: Filter by sales order ID or item ID
- **Custom Lookup**: Dedicated endpoint for retrieving all details of a specific sales order
- **Discount Support**: Percentage-based discounts with automatic calculation
- **Session Tracking**: Track used sessions for service-based items
- **Soft Delete Support**: Deleted records are marked rather than physically removed
- **Multi-Tenant**: Supports tenant-specific data isolation
- **JWT Authentication**: All endpoints require valid JWT authentication
- **Audit Trail**: Tracks who created, updated, and deleted each record

---

## Business Logic

### Automatic Calculation Example

**Example 1: With Discount**
```
Input:
  quantity = 2
  price = 150,000
  discount_pct = 10

Calculation:
  subtotal = 2 × 150,000 = 300,000
  discount = 300,000 × 10 / 100 = 30,000
  item_total = 300,000 - 30,000 = 270,000
```

**Example 2: Without Discount**
```
Input:
  quantity = 1
  price = 200,000
  discount_pct = 0 (or omitted)

Calculation:
  item_total = 1 × 200,000 = 200,000
```

### Use Cases

1. **Adding Line Items**: Add products/services to a sales order
2. **Order Modification**: Update quantities, prices, or discounts for existing items
3. **Discount Application**: Apply percentage-based discounts to line items
4. **Session Management**: Track service sessions used for appointment-based services
5. **Order Analysis**: Query all items in a specific order for reporting

### Relationships

- Each sales order detail belongs to one sales order (`sales_order_id`)
- Each sales order detail references one item/product (`item_id`)
- Each sales order detail has one unit of measurement (`unit_id`)
- Each sales order detail may have one promoter (`promoter_id`)
- A sales order can have multiple detail records (one-to-many)

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
const BASE_URL = 'http://localhost:8080/so/api/sales-order-details';
const JWT_TOKEN = 'YOUR_JWT_TOKEN';
const TENANT_CODE = 'TENANT001';

const headers = {
  'Authorization': `Bearer ${JWT_TOKEN}`,
  'X-Tenant-Code': TENANT_CODE,
  'Content-Type': 'application/json'
};

// Get all details for a sales order
async function getDetailsBySalesOrder(salesOrderId) {
  const response = await fetch(
    `${BASE_URL}/by-sales-order/${salesOrderId}`,
    { method: 'GET', headers: headers }
  );
  return await response.json();
}

// Create new detail (with automatic calculation)
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

// Example: Create detail with automatic total calculation
const newDetail = await createDetail({
  sales_order_id: 5001,
  item_id: 101,
  item_name: "Product A",
  quantity: 2,
  price: 150000.00,
  discount_pct: 10
  // item_total will be auto-calculated as 270,000
});

console.log('Created Detail:', newDetail);

// Get all details for an order
const orderDetails = await getDetailsBySalesOrder(5001);
console.log('Order Details:', orderDetails);
```

### Python (Requests)

```python
import requests

BASE_URL = "http://localhost:8080/so/api/sales-order-details"
JWT_TOKEN = "YOUR_JWT_TOKEN"
TENANT_CODE = "TENANT001"

headers = {
    "Authorization": f"Bearer {JWT_TOKEN}",
    "X-Tenant-Code": TENANT_CODE,
    "Content-Type": "application/json"
}

# Get all details for a sales order
def get_details_by_sales_order(sales_order_id):
    url = f"{BASE_URL}/by-sales-order/{sales_order_id}"
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

# Calculate total before sending (optional - server does this automatically)
def calculate_total(quantity, price, discount_pct=0):
    subtotal = quantity * price
    discount = subtotal * discount_pct / 100
    return subtotal - discount

# Example usage
detail_data = {
    "sales_order_id": 5001,
    "item_id": 101,
    "item_name": "Product A",
    "quantity": 2,
    "price": 150000.00,
    "discount_pct": 10
}

# Create detail (server calculates total automatically)
new_detail = create_detail(detail_data)
print("Created Detail:", new_detail)
print(f"Auto-calculated total: {new_detail['data']['item_total']}")

# Get all details for the order
order_details = get_details_by_sales_order(5001)
print("Order Details:", order_details)
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
    baseURL    = "http://localhost:8080/so/api/sales-order-details"
    jwtToken   = "YOUR_JWT_TOKEN"
    tenantCode = "TENANT001"
)

type SalesOrderDetailRequest struct {
    SalesOrderID *int     `json:"sales_order_id,omitempty"`
    ItemID       *int     `json:"item_id,omitempty"`
    UnitID       *int     `json:"unit_id,omitempty"`
    PromoterID   *int     `json:"promoter_id,omitempty"`
    ItemName     *string  `json:"item_name,omitempty"`
    Quantity     *int     `json:"quantity"`
    Price        *float64 `json:"price"`
    ItemTotal    *float64 `json:"item_total,omitempty"`
    DiscountPct  *int     `json:"discount_pct,omitempty"`
    UsedSessions *int     `json:"used_sessions,omitempty"`
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

// Get details by sales order ID
func getDetailsBySalesOrderID(salesOrderID int) (*Response, error) {
    url := fmt.Sprintf("%s/by-sales-order/%d", baseURL, salesOrderID)
    return makeRequest("GET", url, nil)
}

// Create detail
func createDetail(detail *SalesOrderDetailRequest) (*Response, error) {
    return makeRequest("POST", baseURL, detail)
}

// Update detail
func updateDetail(id int, detail *SalesOrderDetailRequest) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("PUT", url, detail)
}

// Delete detail
func deleteDetail(id int) (*Response, error) {
    url := fmt.Sprintf("%s/%d", baseURL, id)
    return makeRequest("DELETE", url, nil)
}

func main() {
    // Example: Create detail with automatic calculation
    salesOrderID := 5001
    itemID := 101
    itemName := "Product A"
    quantity := 2
    price := 150000.00
    discountPct := 10

    newDetail := &SalesOrderDetailRequest{
        SalesOrderID: &salesOrderID,
        ItemID:       &itemID,
        ItemName:     &itemName,
        Quantity:     &quantity,
        Price:        &price,
        DiscountPct:  &discountPct,
    }

    result, err := createDetail(newDetail)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Created Detail: %+v\n", result)

    // Get all details for the sales order
    details, err := getDetailsBySalesOrderID(5001)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Order Details: %+v\n", details)
}
```

---

## Best Practices

1. **Let Server Calculate**: Omit `item_total` in requests - let the server calculate it automatically
2. **Validate Quantity**: Always ensure quantity is greater than 0
3. **Discount Range**: Keep discount percentage between 0-100
4. **Use Filters**: When querying many details, use filters to reduce response size
5. **Batch Operations**: Use the sales order endpoint for creating orders with multiple details in one transaction
6. **Update Strategy**: Update individual details separately from the parent sales order
7. **Session Tracking**: Use `used_sessions` for service-based items requiring session management
8. **Error Handling**: Always validate required fields (quantity, price) before sending requests

---

## Calculation Formula

The automatic calculation follows this formula:

```
IF item_total is not provided:
    subtotal = quantity × price
    IF discount_pct > 0:
        discount_amount = subtotal × (discount_pct ÷ 100)
        item_total = subtotal - discount_amount
    ELSE:
        item_total = subtotal
```

---

## Notes

1. **Authentication**: All requests must include a valid JWT token
2. **Tenant Isolation**: `X-Tenant-Code` header is required
3. **Required Fields**: `quantity` and `price` are required for create/update
4. **Auto-Calculation**: Server automatically calculates `item_total` if not provided
5. **Soft Deletes**: Deleted records are filtered out automatically
6. **Audit Trail**: System tracks who created, updated, and deleted records
7. **Database Table**: Data stored in `alana.sales_order_detail` table
8. **Foreign Keys**: Ensure referenced sales orders and items exist

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2025-01-15 | Initial release with full CRUD and automatic calculations |
