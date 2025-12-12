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
	salesOrderDetailHandler := handlers.NewSalesOrderDetailHandler(healthCheckDB)
	remindedHandler := handlers.NewRemindedHandler(healthCheckDB)
	arReceiptHandler := handlers.NewARReceiptHandler(healthCheckDB)
	arReceiptDetailHandler := handlers.NewARReceiptDetailHandler(healthCheckDB)
	treatmentHandler := handlers.NewTreatmentHandler(healthCheckDB)
	treatmentDetailHandler := handlers.NewTreatmentDetailHandler(healthCheckDB)
	summaryByTransactionTypeHandler := handlers.NewSummaryByTransactionTypeHandler(healthCheckDB)
	summaryByPaymentMethodHandler := handlers.NewSummaryByPaymentMethodHandler(healthCheckDB)
	summaryByTransactionTypeAndPaymentMethodHandler := handlers.NewSummaryByTransactionTypeAndPaymentMethodHandler(healthCheckDB)
	bookkeepingHandler := handlers.NewBookkeepingHandler(healthCheckDB)
	bookkeepingDetailHandler := handlers.NewBookkeepingDetailHandler(healthCheckDB)
	bookkeepingStatusHandler := handlers.NewBookkeepingStatusHandler(healthCheckDB)
	bookTransactionTypeHandler := handlers.NewBookTransactionTypeHandler(healthCheckDB)
	bookTransactionCategoryHandler := handlers.NewBookTransactionCategoryHandler(healthCheckDB)
	paymentMethodHandler := handlers.NewPaymentMethodHandler(healthCheckDB)

	// Setup Gin router
	router := setupRouter(cfg, jwtUtil, healthHandler, salesOrderStatusHandler, salesOrderHandler, salesOrderServiceHandler, salesOrderDetailHandler, remindedHandler, arReceiptHandler, arReceiptDetailHandler, treatmentHandler, treatmentDetailHandler, summaryByTransactionTypeHandler, summaryByPaymentMethodHandler, summaryByTransactionTypeAndPaymentMethodHandler, bookkeepingHandler, bookkeepingDetailHandler, bookkeepingStatusHandler, bookTransactionTypeHandler, bookTransactionCategoryHandler, paymentMethodHandler)

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
	salesOrderDetailHandler *handlers.SalesOrderDetailHandler,
	remindedHandler *handlers.RemindedHandler,
	arReceiptHandler *handlers.ARReceiptHandler,
	arReceiptDetailHandler *handlers.ARReceiptDetailHandler,
	treatmentHandler *handlers.TreatmentHandler,
	treatmentDetailHandler *handlers.TreatmentDetailHandler,
	summaryByTransactionTypeHandler *handlers.SummaryByTransactionTypeHandler,
	summaryByPaymentMethodHandler *handlers.SummaryByPaymentMethodHandler,
	summaryByTransactionTypeAndPaymentMethodHandler *handlers.SummaryByTransactionTypeAndPaymentMethodHandler,
	bookkeepingHandler *handlers.BookkeepingHandler,
	bookkeepingDetailHandler *handlers.BookkeepingDetailHandler,
	bookkeepingStatusHandler *handlers.BookkeepingStatusHandler,
	bookTransactionTypeHandler *handlers.BookTransactionTypeHandler,
	bookTransactionCategoryHandler *handlers.BookTransactionCategoryHandler,
	paymentMethodHandler *handlers.PaymentMethodHandler,
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

		// Sales Order Detail CRUD endpoints (JWT required)
		salesOrderDetails := api.Group("/sales-order-details")
		salesOrderDetails.Use(middleware.AuthMiddleware(jwtUtil))
		{
			salesOrderDetails.GET("", salesOrderDetailHandler.GetAll)
			salesOrderDetails.GET("/:id", salesOrderDetailHandler.GetByID)
			salesOrderDetails.GET("/by-sales-order/:sales_order_id", salesOrderDetailHandler.GetBySalesOrderID)
			salesOrderDetails.POST("", salesOrderDetailHandler.Create)
			salesOrderDetails.PUT("/:id", salesOrderDetailHandler.Update)
			salesOrderDetails.DELETE("/:id", salesOrderDetailHandler.Delete)
		}

		// Reminded endpoints (JWT required, read-only)
		reminded := api.Group("/reminded")
		reminded.Use(middleware.AuthMiddleware(jwtUtil))
		{
			reminded.GET("", remindedHandler.GetAll)
			reminded.GET("/:id", remindedHandler.GetByID)
		}

		// AR Receipt CRUD endpoints (JWT required)
		arReceipts := api.Group("/ar-receipts")
		arReceipts.Use(middleware.AuthMiddleware(jwtUtil))
		{
			arReceipts.GET("", arReceiptHandler.GetAll)
			arReceipts.GET("/:id", arReceiptHandler.GetByID)
			arReceipts.POST("", arReceiptHandler.Create)
			arReceipts.PUT("/:id", arReceiptHandler.Update)
			arReceipts.DELETE("/:id", arReceiptHandler.Delete)
		}

		// AR Receipt Detail CRUD endpoints (JWT required)
		arReceiptDetails := api.Group("/ar-receipt-details")
		arReceiptDetails.Use(middleware.AuthMiddleware(jwtUtil))
		{
			arReceiptDetails.GET("", arReceiptDetailHandler.GetAll)
			arReceiptDetails.GET("/:id", arReceiptDetailHandler.GetByID)
			arReceiptDetails.GET("/by-ar-receipt/:ar_receipt_id", arReceiptDetailHandler.GetByARReceiptID)
			arReceiptDetails.POST("", arReceiptDetailHandler.Create)
			arReceiptDetails.PUT("/:id", arReceiptDetailHandler.Update)
			arReceiptDetails.DELETE("/:id", arReceiptDetailHandler.Delete)
		}

		// Treatment CRUD endpoints (JWT required)
		treatments := api.Group("/treatments")
		treatments.Use(middleware.AuthMiddleware(jwtUtil))
		{
			treatments.GET("", treatmentHandler.GetAll)
			treatments.GET("/:id", treatmentHandler.GetByID)
			treatments.POST("", treatmentHandler.Create)
			treatments.PUT("/:id", treatmentHandler.Update)
			treatments.DELETE("/:id", treatmentHandler.Delete)
		}

		// Treatment Detail CRUD endpoints (JWT required)
		treatmentDetails := api.Group("/treatment-details")
		treatmentDetails.Use(middleware.AuthMiddleware(jwtUtil))
		{
			treatmentDetails.GET("", treatmentDetailHandler.GetAll)
			treatmentDetails.GET("/:id", treatmentDetailHandler.GetByID)
			treatmentDetails.GET("/by-treatment/:treatment_id", treatmentDetailHandler.GetByTreatmentID)
			treatmentDetails.POST("", treatmentDetailHandler.Create)
			treatmentDetails.PUT("/:id", treatmentDetailHandler.Update)
			treatmentDetails.DELETE("/:id", treatmentDetailHandler.Delete)
		}

		// Summary By Transaction Type CRUD endpoints (JWT required)
		summaryByTransactionType := api.Group("/summary-by-transaction-type")
		summaryByTransactionType.Use(middleware.AuthMiddleware(jwtUtil))
		{
			summaryByTransactionType.GET("", summaryByTransactionTypeHandler.GetAll)
			summaryByTransactionType.GET("/:id", summaryByTransactionTypeHandler.GetByID)
			summaryByTransactionType.GET("/by-bookkeeping/:bookkeeping_id", summaryByTransactionTypeHandler.GetByBookkeepingID)
			summaryByTransactionType.POST("", summaryByTransactionTypeHandler.Create)
			summaryByTransactionType.PUT("/:id", summaryByTransactionTypeHandler.Update)
			summaryByTransactionType.DELETE("/:id", summaryByTransactionTypeHandler.Delete)
		}

		// Summary By Payment Method CRUD endpoints (JWT required)
		summaryByPaymentMethod := api.Group("/summary-by-payment-method")
		summaryByPaymentMethod.Use(middleware.AuthMiddleware(jwtUtil))
		{
			summaryByPaymentMethod.GET("", summaryByPaymentMethodHandler.GetAll)
			summaryByPaymentMethod.GET("/:id", summaryByPaymentMethodHandler.GetByID)
			summaryByPaymentMethod.GET("/by-bookkeeping/:bookkeeping_id", summaryByPaymentMethodHandler.GetByBookkeepingID)
			summaryByPaymentMethod.POST("", summaryByPaymentMethodHandler.Create)
			summaryByPaymentMethod.PUT("/:id", summaryByPaymentMethodHandler.Update)
			summaryByPaymentMethod.DELETE("/:id", summaryByPaymentMethodHandler.Delete)
		}

		// Summary By Transaction Type And Payment Method CRUD endpoints (JWT required)
		summaryByTransactionTypeAndPaymentMethod := api.Group("/summary-by-transaction-type-and-payment-method")
		summaryByTransactionTypeAndPaymentMethod.Use(middleware.AuthMiddleware(jwtUtil))
		{
			summaryByTransactionTypeAndPaymentMethod.GET("", summaryByTransactionTypeAndPaymentMethodHandler.GetAll)
			summaryByTransactionTypeAndPaymentMethod.GET("/:id", summaryByTransactionTypeAndPaymentMethodHandler.GetByID)
			summaryByTransactionTypeAndPaymentMethod.GET("/by-bookkeeping/:bookkeeping_id", summaryByTransactionTypeAndPaymentMethodHandler.GetByBookkeepingID)
			summaryByTransactionTypeAndPaymentMethod.POST("", summaryByTransactionTypeAndPaymentMethodHandler.Create)
			summaryByTransactionTypeAndPaymentMethod.PUT("/:id", summaryByTransactionTypeAndPaymentMethodHandler.Update)
			summaryByTransactionTypeAndPaymentMethod.DELETE("/:id", summaryByTransactionTypeAndPaymentMethodHandler.Delete)
		}

		// Bookkeeping CRUD endpoints (JWT required)
		bookkeeping := api.Group("/bookkeeping")
		bookkeeping.Use(middleware.AuthMiddleware(jwtUtil))
		{
			bookkeeping.GET("", bookkeepingHandler.GetAll)
			bookkeeping.GET("/:id", bookkeepingHandler.GetByID)
			bookkeeping.GET("/by-location/:location_id", bookkeepingHandler.GetByLocationID)
			bookkeeping.POST("", bookkeepingHandler.Create)
			bookkeeping.PUT("/:id", bookkeepingHandler.Update)
			bookkeeping.DELETE("/:id", bookkeepingHandler.Delete)
		}

		// Bookkeeping Detail CRUD endpoints (JWT required)
		bookkeepingDetail := api.Group("/bookkeeping-detail")
		bookkeepingDetail.Use(middleware.AuthMiddleware(jwtUtil))
		{
			bookkeepingDetail.GET("", bookkeepingDetailHandler.GetAll)
			bookkeepingDetail.GET("/:id", bookkeepingDetailHandler.GetByID)
			bookkeepingDetail.GET("/by-bookkeeping/:bookkeeping_id", bookkeepingDetailHandler.GetByBookkeepingID)
			bookkeepingDetail.POST("", bookkeepingDetailHandler.Create)
			bookkeepingDetail.PUT("/:id", bookkeepingDetailHandler.Update)
			bookkeepingDetail.DELETE("/:id", bookkeepingDetailHandler.Delete)
		}

		// Bookkeeping Status READ-ONLY endpoints (JWT required)
		bookkeepingStatus := api.Group("/bookkeeping-status")
		bookkeepingStatus.Use(middleware.AuthMiddleware(jwtUtil))
		{
			bookkeepingStatus.GET("", bookkeepingStatusHandler.GetAll)
			bookkeepingStatus.GET("/:id", bookkeepingStatusHandler.GetByID)
		}

		// Book Transaction Type CRUD endpoints (JWT required)
		bookTransactionType := api.Group("/book-transaction-type")
		bookTransactionType.Use(middleware.AuthMiddleware(jwtUtil))
		{
			bookTransactionType.GET("", bookTransactionTypeHandler.GetAll)
			bookTransactionType.GET("/:id", bookTransactionTypeHandler.GetByID)
			bookTransactionType.POST("", bookTransactionTypeHandler.Create)
			bookTransactionType.PUT("/:id", bookTransactionTypeHandler.Update)
			bookTransactionType.DELETE("/:id", bookTransactionTypeHandler.Delete)
		}

		// Book Transaction Category CRUD endpoints (JWT required)
		bookTransactionCategory := api.Group("/book-transaction-category")
		bookTransactionCategory.Use(middleware.AuthMiddleware(jwtUtil))
		{
			bookTransactionCategory.GET("", bookTransactionCategoryHandler.GetAll)
			bookTransactionCategory.GET("/:id", bookTransactionCategoryHandler.GetByID)
			bookTransactionCategory.POST("", bookTransactionCategoryHandler.Create)
			bookTransactionCategory.PUT("/:id", bookTransactionCategoryHandler.Update)
			bookTransactionCategory.DELETE("/:id", bookTransactionCategoryHandler.Delete)
		}

		// Payment Method CRUD endpoints (JWT required)
		paymentMethod := api.Group("/payment-method")
		paymentMethod.Use(middleware.AuthMiddleware(jwtUtil))
		{
			paymentMethod.GET("", paymentMethodHandler.GetAll)
			paymentMethod.GET("/:id", paymentMethodHandler.GetByID)
			paymentMethod.POST("", paymentMethodHandler.Create)
			paymentMethod.PUT("/:id", paymentMethodHandler.Update)
			paymentMethod.DELETE("/:id", paymentMethodHandler.Delete)
		}
	}

	return router
}
