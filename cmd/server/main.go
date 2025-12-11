package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"pos-mojosoft-so-service/internal/config"
	"pos-mojosoft-so-service/internal/handlers"
	"pos-mojosoft-so-service/internal/middleware"
	"pos-mojosoft-so-service/internal/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Setup logging
	setupLogging(cfg.Logging)

	// Initialize tenant database connections
	tenantDBManager := config.GetTenantDBManager()
	if err := tenantDBManager.InitializeTenantConnections(cfg); err != nil {
		logrus.Fatal("Failed to initialize tenant database connections:", err)
	}

	// Ensure all tenant connections are closed on shutdown
	defer func() {
		if err := tenantDBManager.Close(); err != nil {
			logrus.Error("Failed to close tenant database connections:", err)
		}
	}()

	// Initialize JWT utility
	jwtUtil := utils.NewJWTUtil(&cfg.JWT)

	// Get a sample tenant DB for health check (use first tenant)
	var healthCheckDB *gorm.DB
	if len(cfg.TenantCodes) > 0 {
		healthCheckDB, _ = tenantDBManager.GetTenantDB(cfg.TenantCodes[0])
	}

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(healthCheckDB)
	salesOrderStatusHandler := handlers.NewSalesOrderStatusHandler(healthCheckDB)
	salesOrderHandler := handlers.NewSalesOrderHandler(healthCheckDB)
	salesOrderServiceHandler := handlers.NewSalesOrderServiceHandler(healthCheckDB)

	// Setup Gin router
	router := setupRouter(cfg, jwtUtil, healthHandler, salesOrderStatusHandler, salesOrderHandler, salesOrderServiceHandler)

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		logrus.Infof("Starting SO service on port %s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Error("Server forced to shutdown:", err)
	}

	logrus.Info("Server exited")
}

func setupLogging(config config.LoggingConfig) {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	if config.Format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	if config.FilePath != "" {
		if err := os.MkdirAll(config.FilePath, 0755); err != nil {
			logrus.Error("Failed to create log directory:", err)
			return
		}

		logFileName := fmt.Sprintf("so-service-%s.log", time.Now().Format("2006-01-02"))
		logFilePath := filepath.Join(config.FilePath, logFileName)

		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("Failed to open log file:", err)
			return
		}

		multiWriter := io.MultiWriter(os.Stdout, logFile)
		logrus.SetOutput(multiWriter)

		logrus.Infof("Logging to file: %s", logFilePath)
	}
}

func setupRouter(
	cfg *config.Config,
	jwtUtil *utils.JWTUtil,
	healthHandler *handlers.HealthHandler,
	salesOrderStatusHandler *handlers.SalesOrderStatusHandler,
	salesOrderHandler *handlers.SalesOrderHandler,
	salesOrderServiceHandler *handlers.SalesOrderServiceHandler,
) *gin.Engine {
	// Set Gin mode
	if cfg.Logging.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Global middleware
	router.Use(middleware.CORSMiddleware(&cfg.CORS))
	router.Use(middleware.SecurityHeadersMiddleware())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorLoggerMiddleware())
	router.Use(middleware.RateLimitMiddleware(&cfg.RateLimit))

	// Health check endpoint (no auth required)
	router.GET("/health", healthHandler.Check)

	// Apply TenantMiddleware to all /so/api routes
	api := router.Group("/so/api")
	api.Use(middleware.TenantMiddleware())
	{
		// Sales Order Status endpoints (JWT required)
		statusGroup := api.Group("/sales-order-status")
		statusGroup.Use(middleware.AuthMiddleware(jwtUtil))
		{
			statusGroup.GET("", salesOrderStatusHandler.GetAll)
			statusGroup.GET("/:id", salesOrderStatusHandler.GetByID)
		}

		// Sales Order CRUD endpoints (JWT required)
		salesOrders := api.Group("/sales-orders")
		salesOrders.Use(middleware.AuthMiddleware(jwtUtil))
		{
			salesOrders.GET("", salesOrderHandler.GetAll)
			salesOrders.GET("/:id", salesOrderHandler.GetByID)
			salesOrders.POST("", salesOrderHandler.Create)
			salesOrders.PUT("/:id", salesOrderHandler.Update)
			salesOrders.DELETE("/:id", salesOrderHandler.Delete)
		}

		// Sales Order Service CRUD endpoints (JWT required)
		salesOrderServices := api.Group("/sales-order-services")
		salesOrderServices.Use(middleware.AuthMiddleware(jwtUtil))
		{
			salesOrderServices.GET("", salesOrderServiceHandler.GetAll)
			salesOrderServices.GET("/:id", salesOrderServiceHandler.GetByID)
			salesOrderServices.POST("", salesOrderServiceHandler.Create)
			salesOrderServices.PUT("/:id", salesOrderServiceHandler.Update)
			salesOrderServices.DELETE("/:id", salesOrderServiceHandler.Delete)
			salesOrderServices.PATCH("/:id/mark-treated", salesOrderServiceHandler.MarkAsTreated)
		}
	}

	return router
}
