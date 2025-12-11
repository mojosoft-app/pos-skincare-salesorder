# Sales Order Status API Documentation

## Overview
API untuk mengelola status sales order dalam sistem POS Mojosoft.

## Base URL
```
/so/api/sales-order-status
```

## Endpoints

### 1. Get All Sales Order Statuses
Mengambil daftar semua status sales order.

**Endpoint:** `GET /so/api/sales-order-status`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer <token>
X-Tenant-ID: <tenant_id>
```

**Response Success (200):**
```json
{
  "status": "success",
  "message": "Sales order statuses retrieved successfully",
  "data": [
    {
      "id": 1,
      "status_name": "Pending",
      "description": "Order is pending",
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-01T10:00:00Z"
    },
    {
      "id": 2,
      "status_name": "Processing",
      "description": "Order is being processed",
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-01T10:00:00Z"
    }
  ]
}
```

**Response Error (500):**
```json
{
  "status": "error",
  "message": "Failed to retrieve sales order statuses",
  "errors": null
}
```

**Response Error (500 - Database Not Found):**
```json
{
  "status": "error",
  "message": "Database connection not found",
  "errors": null
}
```

---

### 2. Get Sales Order Status by ID
Mengambil detail status sales order berdasarkan ID.

**Endpoint:** `GET /so/api/sales-order-status/:id`

**Headers:**
```
Content-Type: application/json
Authorization: Bearer <token>
X-Tenant-ID: <tenant_id>
```

**Path Parameters:**
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | ID status sales order |

**Example Request:**
```
GET /so/api/sales-order-status/1
```

**Response Success (200):**
```json
{
  "status": "success",
  "message": "Sales order status retrieved successfully",
  "data": {
    "id": 1,
    "status_name": "Pending",
    "description": "Order is pending",
    "created_at": "2025-01-01T10:00:00Z",
    "updated_at": "2025-01-01T10:00:00Z"
  }
}
```

**Response Error (400):**
```json
{
  "status": "error",
  "message": "Invalid status ID",
  "errors": null
}
```

**Response Error (404):**
```json
{
  "status": "error",
  "message": "Sales order status not found",
  "errors": null
}
```

**Response Error (500):**
```json
{
  "status": "error",
  "message": "Failed to retrieve sales order status",
  "errors": null
}
```

**Response Error (500 - Database Not Found):**
```json
{
  "status": "error",
  "message": "Database connection not found",
  "errors": null
}
```

---

## Data Model

### SalesOrderStatus
| Field | Type | Description |
|-------|------|-------------|
| id | integer | ID unik status (Primary Key) |
| status_name | string | Nama status |
| description | string | Deskripsi status |
| created_at | timestamp | Waktu pembuatan record |
| updated_at | timestamp | Waktu update terakhir |
| deleted_at | timestamp | Waktu soft delete (null jika belum dihapus) |

---

## Common Status Codes
| Status Code | Description |
|-------------|-------------|
| 200 | Success - Request berhasil |
| 400 | Bad Request - Parameter tidak valid |
| 404 | Not Found - Resource tidak ditemukan |
| 500 | Internal Server Error - Kesalahan server |

---

## Notes
- Semua endpoint memerlukan autentikasi menggunakan Bearer token
- Tenant ID wajib disertakan dalam header untuk multi-tenant support
- Soft delete digunakan, sehingga data yang dihapus tidak akan muncul dalam query GET
- Timestamp menggunakan format ISO 8601

---

## Example Usage

### cURL Example - Get All Statuses
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-status" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -H "X-Tenant-ID: tenant123"
```

### cURL Example - Get Status by ID
```bash
curl -X GET "http://localhost:8080/so/api/sales-order-status/1" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_token_here" \
  -H "X-Tenant-ID: tenant123"
```

---

## Typical Status Flow
Status sales order yang umum digunakan:
1. **Pending** - Order baru dibuat
2. **Processing** - Order sedang diproses
3. **Completed** - Order selesai
4. **Cancelled** - Order dibatalkan
5. **On Hold** - Order ditahan sementara

---

## Error Handling
API menggunakan format error response yang konsisten:
```json
{
  "status": "error",
  "message": "Descriptive error message",
  "errors": null
}
```

Untuk validasi error, field `errors` dapat berisi detail error spesifik.
