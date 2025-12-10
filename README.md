# POS Mojosoft Sales Order Service

Sales Order service for POS Mojosoft system. This service handles sales order operations with multi-tenant support.

## Features

- Multi-tenant support with separate database connections per tenant
- JWT-based authentication
- Rate limiting
- CORS support
- Structured logging
- Health check endpoint
- Graceful shutdown

## Requirements

- Go 1.25.1+
- PostgreSQL database

## Installation

1. Clone the repository or copy this template
2. Copy `.env.example` to `.env` and update the configuration
3. Install dependencies:
   ```bash
   go mod download
   ```

## Configuration

The service uses environment variables for configuration. See `.env.example` for available options.

### Important Configuration

- `SERVER_PORT`: Server port (default: 8082)
- `DB_HOST`, `DB_PORT`, etc.: PostgreSQL connection details
- `JWT_SECRET`: Secret key for JWT token signing
- `TENANT_CODES`: Comma-separated list of tenant codes

## Running the Service

### Development

```bash
# Run the service
go run cmd/server/main.go
```

### Production

```bash
# Build the binary
go build -o so-service cmd/server/main.go

# Run the binary
./so-service
```

## API Endpoints

### Health Check

```
GET /health
```

### Sales Order API (Protected)

All sales order API endpoints are protected with JWT authentication and tenant middleware.

```
GET /so/api/sales-orders
```

## Architecture

```
service-so/
├── cmd/server/           # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models (to be added)
│   ├── services/        # Business logic (to be added)
│   ├── utils/           # Utility functions
│   └── repositories/    # Data access layer (to be added)
├── migrations/          # Database migrations
├── docs/               # Documentation
├── test-http/          # HTTP test files
└── logs/               # Log files
```

## Multi-Tenant Support

The service supports multi-tenancy by maintaining separate database connections for each tenant. Tenant is identified via the `X-Tenant-Code` header in HTTP requests.

## Authentication

Protected endpoints require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <jwt_token>
```

## Next Steps

1. Define your Sales Order models in `internal/models/`
2. Create repositories in `internal/repositories/`
3. Implement business logic in `internal/services/`
4. Add HTTP handlers in `internal/handlers/`
5. Register routes in `cmd/server/main.go`

## Template Features

This template includes:
- ✅ Project structure
- ✅ Configuration management (multi-tenant support)
- ✅ Middleware (Auth, CORS, Logging, Rate Limit, Security, Tenant)
- ✅ Utilities (JWT, Password, Response, Validator)
- ✅ Health check endpoint
- ✅ Graceful shutdown
- ⏳ Models (ready to be implemented based on your design)
- ⏳ Services (ready to be implemented)
- ⏳ Handlers (ready to be implemented)

## License

[Add your license information here]
