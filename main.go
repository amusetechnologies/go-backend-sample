package main

import (
	"log"
	"os"
	"theatre-management-system/src/business"
	"theatre-management-system/src/constants"
	"theatre-management-system/src/controllers"
	"theatre-management-system/src/dao"
	"theatre-management-system/src/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Database connection
	db := connectToDB()

	// Auto-migrate database schema
	if err := migrateDB(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup Gin router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize repositories
	locationRepo := dao.NewLocationRepository(db)
	theatreTypeRepo := dao.NewTheatreTypeRepository(db)
	showTypeRepo := dao.NewShowTypeRepository(db)
	theatreRepo := dao.NewTheatreRepository(db)
	showRepo := dao.NewShowRepository(db)

	// Initialize services
	locationService := business.NewLocationService(locationRepo)
	theatreTypeService := business.NewTheatreTypeService(theatreTypeRepo)
	showTypeService := business.NewShowTypeService(showTypeRepo)
	theatreService := business.NewTheatreService(theatreRepo, locationRepo, theatreTypeRepo)
	showService := business.NewShowService(showRepo, theatreRepo, showTypeRepo)

	// Initialize controllers
	locationController := controllers.NewLocationController(locationService)
	theatreTypeController := controllers.NewTheatreTypeController(theatreTypeService)
	showTypeController := controllers.NewShowTypeController(showTypeService)
	theatreController := controllers.NewTheatreController(theatreService)
	showController := controllers.NewShowController(showService)

	// Setup routes
	setupRoutes(r, locationController, theatreTypeController, showTypeController, theatreController, showController)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// connectToDB establishes database connection
func connectToDB() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = constants.DefaultDatabaseURL
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
	return db
}

// migrateDB runs database migrations
func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Location{},
		&models.TheatreType{},
		&models.ShowType{},
		&models.Theatre{},
		&models.Show{},
	)
}

// setupRoutes configures all API routes
func setupRoutes(
	r *gin.Engine,
	locationController *controllers.LocationController,
	theatreTypeController *controllers.TheatreTypeController,
	showTypeController *controllers.ShowTypeController,
	theatreController *controllers.TheatreController,
	showController *controllers.ShowController,
) {
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Theatre Management API is running",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")

	// Location routes
	locations := v1.Group("/locations")
	{
		locations.POST("", locationController.CreateLocation)
		locations.GET("", locationController.GetAllLocations)
		locations.GET("/:id", locationController.GetLocationByID)
		locations.PATCH("/:id", locationController.UpdateLocation)
		locations.DELETE("/:id", locationController.DeleteLocation)
		locations.GET("/active", locationController.GetActiveLocations)
		locations.GET("/nearby", locationController.GetLocationsByCoordinates)
		locations.GET("/search", locationController.SearchLocations)
	}

	// Theatre Type routes
	theatreTypes := v1.Group("/theatre-types")
	{
		theatreTypes.POST("", theatreTypeController.CreateTheatreType)
		theatreTypes.GET("", theatreTypeController.GetAllTheatreTypes)
		theatreTypes.GET("/:id", theatreTypeController.GetTheatreTypeByID)
		theatreTypes.PATCH("/:id", theatreTypeController.UpdateTheatreType)
		theatreTypes.DELETE("/:id", theatreTypeController.DeleteTheatreType)
		theatreTypes.GET("/active", theatreTypeController.GetActiveTheatreTypes)
		theatreTypes.GET("/name/:name", theatreTypeController.GetTheatreTypeByName)
	}

	// Show Type routes
	showTypes := v1.Group("/show-types")
	{
		showTypes.POST("", showTypeController.CreateShowType)
		showTypes.GET("", showTypeController.GetAllShowTypes)
		showTypes.GET("/:id", showTypeController.GetShowTypeByID)
		showTypes.PATCH("/:id", showTypeController.UpdateShowType)
		showTypes.DELETE("/:id", showTypeController.DeleteShowType)
		showTypes.GET("/active", showTypeController.GetActiveShowTypes)
		showTypes.GET("/name/:name", showTypeController.GetShowTypeByName)
	}

	// Theatre routes
	theatres := v1.Group("/theatres")
	{
		theatres.POST("", theatreController.CreateTheatre)
		theatres.GET("", theatreController.GetAllTheatres)
		theatres.GET("/:id", theatreController.GetTheatreByID)
		theatres.PATCH("/:id", theatreController.UpdateTheatre)
		theatres.DELETE("/:id", theatreController.DeleteTheatre)
		theatres.GET("/active", theatreController.GetActiveTheatres)
		theatres.GET("/featured", theatreController.GetFeaturedTheatres)
		theatres.GET("/location/:locationId", theatreController.GetTheatresByLocationID)
		theatres.GET("/type/:typeId", theatreController.GetTheatresByTheatreTypeID)
		theatres.GET("/nearby", theatreController.GetNearbyTheatres)
		theatres.GET("/search", theatreController.SearchTheatres)
	}

	// Show routes
	shows := v1.Group("/shows")
	{
		shows.POST("", showController.CreateShow)
		shows.GET("", showController.GetAllShows)
		shows.GET("/:id", showController.GetShowByID)
		shows.PATCH("/:id", showController.UpdateShow)
		shows.DELETE("/:id", showController.DeleteShow)
		shows.GET("/active", showController.GetActiveShows)
		shows.GET("/featured", showController.GetFeaturedShows)
		shows.GET("/current", showController.GetCurrentShows)
		shows.GET("/upcoming", showController.GetUpcomingShows)
		shows.GET("/theatre/:theatreId", showController.GetShowsByTheatreID)
		shows.GET("/type/:typeId", showController.GetShowsByShowTypeID)
		shows.GET("/search", showController.SearchShows)
	}
}
