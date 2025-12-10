package middleware

import (
	"net/http"

	"pos-mojosoft-so-service/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const TenantCodeHeader = "X-Tenant-Code"
const TenantCodeKey = "tenant_code"
const TenantDBKey = "tenant_db"

// TenantMiddleware extracts tenant code from header, validates it, and injects tenant DB into context
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantCode := c.GetHeader(TenantCodeHeader)

		if tenantCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "tenant code is required in X-Tenant-Code header",
			})
			c.Abort()
			return
		}

		// Get tenant database connection from manager
		dbManager := config.GetTenantDBManager()
		tenantDB, err := dbManager.GetTenantDB(tenantCode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "invalid tenant code",
				"tenant_code": tenantCode,
			})
			c.Abort()
			return
		}

		// Store tenant code and DB connection in context
		c.Set(TenantCodeKey, tenantCode)
		c.Set(TenantDBKey, tenantDB)
		c.Next()
	}
}

// GetTenantCode retrieves tenant code from context
func GetTenantCode(c *gin.Context) string {
	tenantCode, exists := c.Get(TenantCodeKey)
	if !exists {
		return ""
	}
	return tenantCode.(string)
}

// GetTenantDB retrieves tenant database connection from context
func GetTenantDB(c *gin.Context) (*gorm.DB, error) {
	db, exists := c.Get(TenantDBKey)
	if !exists {
		return nil, ErrTenantDBNotFound
	}
	return db.(*gorm.DB), nil
}

var ErrTenantDBNotFound = &TenantDBError{Message: "tenant database not found in context"}

type TenantDBError struct {
	Message string
}

func (e *TenantDBError) Error() string {
	return e.Message
}
