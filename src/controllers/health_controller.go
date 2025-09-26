package controllers

import (
	"net/http"
	"theatre-management-system/src/business"
	"theatre-management-system/src/constants"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthController handles health check and system status endpoints
type HealthController struct {
	db           *gorm.DB
	cacheService *business.CacheService
}

// NewHealthController creates a new health controller
func NewHealthController(db *gorm.DB, cacheService *business.CacheService) *HealthController {
	return &HealthController{
		db:           db,
		cacheService: cacheService,
	}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status      string            `json:"status"`
	Timestamp   time.Time         `json:"timestamp"`
	Version     string            `json:"version"`
	Environment string            `json:"environment"`
	Services    map[string]string `json:"services"`
	Cache       CacheStats        `json:"cache"`
}

// CacheStats represents cache statistics
type CacheStats struct {
	ItemCount int `json:"item_count"`
	Size      int `json:"size"`
}

// HealthCheck handles GET /health
func (ctrl *HealthController) HealthCheck(c *gin.Context) {
	response := HealthResponse{
		Status:      "healthy",
		Timestamp:   time.Now(),
		Version:     "1.0.0",
		Environment: "development",
		Services:    make(map[string]string),
	}

	// Check database connection
	sqlDB, err := ctrl.db.DB()
	if err != nil {
		response.Status = "unhealthy"
		response.Services["database"] = "error: " + err.Error()
	} else {
		err = sqlDB.Ping()
		if err != nil {
			response.Status = "unhealthy"
			response.Services["database"] = "error: " + err.Error()
		} else {
			response.Services["database"] = "healthy"
		}
	}

	// Check cache stats
	if ctrl.cacheService != nil {
		itemCount, size := ctrl.cacheService.GetStats()
		response.Cache = CacheStats{
			ItemCount: itemCount,
			Size:      size,
		}
		response.Services["cache"] = "healthy"
	}

	statusCode := http.StatusOK
	if response.Status == "unhealthy" {
		statusCode = http.StatusServiceUnavailable
	}

	SuccessResponse(c, statusCode, constants.StatusOK, response)
}

// ReadinessCheck handles GET /ready
func (ctrl *HealthController) ReadinessCheck(c *gin.Context) {
	// Check if the application is ready to serve requests
	ready := true
	services := make(map[string]string)

	// Check database readiness
	sqlDB, err := ctrl.db.DB()
	if err != nil {
		ready = false
		services["database"] = "not ready: " + err.Error()
	} else {
		err = sqlDB.Ping()
		if err != nil {
			ready = false
			services["database"] = "not ready: " + err.Error()
		} else {
			services["database"] = "ready"
		}
	}

	response := map[string]interface{}{
		"ready":     ready,
		"timestamp": time.Now(),
		"services":  services,
	}

	statusCode := http.StatusOK
	if !ready {
		statusCode = http.StatusServiceUnavailable
	}

	SuccessResponse(c, statusCode, "Readiness check completed", response)
}

// LivenessCheck handles GET /live
func (ctrl *HealthController) LivenessCheck(c *gin.Context) {
	// Simple liveness check - if we can respond, we're alive
	response := map[string]interface{}{
		"alive":     true,
		"timestamp": time.Now(),
		"uptime":    time.Since(time.Now().Add(-time.Hour)), // Placeholder uptime
	}

	SuccessResponse(c, http.StatusOK, "Liveness check completed", response)
}
