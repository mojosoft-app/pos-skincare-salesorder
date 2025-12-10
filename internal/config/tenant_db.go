package config

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// TenantDBManager manages database connections for multiple tenants
type TenantDBManager struct {
	connections map[string]*gorm.DB
	mu          sync.RWMutex
}

var (
	tenantDBManager *TenantDBManager
	once            sync.Once
)

// GetTenantDBManager returns singleton instance of TenantDBManager
func GetTenantDBManager() *TenantDBManager {
	once.Do(func() {
		tenantDBManager = &TenantDBManager{
			connections: make(map[string]*gorm.DB),
		}
	})
	return tenantDBManager
}

// InitializeTenantConnections creates database connections for all tenants
func (m *TenantDBManager) InitializeTenantConnections(cfg *Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	logrus.Infof("Initializing database connections for %d tenants: %v", len(cfg.TenantCodes), cfg.TenantCodes)

	for _, tenantCode := range cfg.TenantCodes {
		schemaName := fmt.Sprintf("%s", tenantCode)

		// Create connection with specific search_path
		db, err := initTenantDB(cfg, schemaName)
		if err != nil {
			return fmt.Errorf("failed to initialize DB for tenant %s: %w", tenantCode, err)
		}

		m.connections[tenantCode] = db
		logrus.Infof("✓ Tenant '%s' connected to schema '%s'", tenantCode, schemaName)
	}

	logrus.Infof("All tenant database connections initialized successfully")
	return nil
}

// GetTenantDB returns the database connection for a specific tenant
func (m *TenantDBManager) GetTenantDB(tenantCode string) (*gorm.DB, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	db, exists := m.connections[tenantCode]
	if !exists {
		return nil, fmt.Errorf("tenant '%s' not found. Available tenants: %v", tenantCode, m.GetAvailableTenants())
	}

	return db, nil
}

// GetAvailableTenants returns list of available tenant codes
func (m *TenantDBManager) GetAvailableTenants() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tenants := make([]string, 0, len(m.connections))
	for tenantCode := range m.connections {
		tenants = append(tenants, tenantCode)
	}
	return tenants
}

// Close closes all tenant database connections
func (m *TenantDBManager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for tenantCode, db := range m.connections {
		sqlDB, err := db.DB()
		if err != nil {
			logrus.Errorf("Failed to get DB instance for tenant %s: %v", tenantCode, err)
			continue
		}

		if err := sqlDB.Close(); err != nil {
			logrus.Errorf("Failed to close DB connection for tenant %s: %v", tenantCode, err)
		} else {
			logrus.Infof("✓ Closed connection for tenant '%s'", tenantCode)
		}
	}

	return nil
}

// initTenantDB creates a database connection with specific search_path
func initTenantDB(cfg *Config, schemaName string) (*gorm.DB, error) {
	// Create DSN with search_path
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
		cfg.Database.TimeZone,
		schemaName,
	)

	return connectToDatabase(dsn)
}
