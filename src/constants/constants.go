package constants

// HTTP Status Messages
const (
	StatusOK                  = "OK"
	StatusCreated             = "Created"
	StatusBadRequest          = "Bad Request"
	StatusNotFound            = "Not Found"
	StatusInternalServerError = "Internal Server Error"
	StatusUnprocessableEntity = "Unprocessable Entity"
)

// Error Messages
const (
	ErrorInvalidUUID         = "Invalid UUID format"
	ErrorInvalidInput        = "Invalid input data"
	ErrorLocationNotFound    = "Location not found"
	ErrorTheatreNotFound     = "Theatre not found"
	ErrorShowNotFound        = "Show not found"
	ErrorTheatreTypeNotFound = "Theatre type not found"
	ErrorShowTypeNotFound    = "Show type not found"
	ErrorDatabaseConnection  = "Database connection error"
	ErrorDuplicateEntry      = "Duplicate entry"
	ErrorValidationFailed    = "Validation failed"
	ErrorInternalServerError = "Internal Server Error"
)

// Success Messages
const (
	MessageLocationCreated    = "Location created successfully"
	MessageLocationUpdated    = "Location updated successfully"
	MessageLocationDeleted    = "Location deleted successfully"
	MessageTheatreCreated     = "Theatre created successfully"
	MessageTheatreUpdated     = "Theatre updated successfully"
	MessageTheatreDeleted     = "Theatre deleted successfully"
	MessageShowCreated        = "Show created successfully"
	MessageShowUpdated        = "Show updated successfully"
	MessageShowDeleted        = "Show deleted successfully"
	MessageTheatreTypeCreated = "Theatre type created successfully"
	MessageTheatreTypeUpdated = "Theatre type updated successfully"
	MessageTheatreTypeDeleted = "Theatre type deleted successfully"
	MessageShowTypeCreated    = "Show type created successfully"
	MessageShowTypeUpdated    = "Show type updated successfully"
	MessageShowTypeDeleted    = "Show type deleted successfully"
)

// Default Values
const (
	DefaultLimit  = 20
	DefaultOffset = 0
	MaxLimit      = 100
	DefaultRadius = 50.0 // kilometers
)

// Database Constants
const (
	DefaultDatabaseURL = "postgres://postgres:postgres@localhost:5433/theatre_api?sslmode=disable"
	TestDatabaseURL    = "postgres://postgres:postgres@localhost:5434/theatre_api_test?sslmode=disable"
)

// Cache Constants
const (
	CacheDefaultExpiration = 300 // 5 minutes
	CacheCleanupInterval   = 600 // 10 minutes
	CacheKeyLocations      = "locations"
	CacheKeyTheatres       = "theatres"
	CacheKeyShows          = "shows"
	CacheKeyTheatreTypes   = "theatre_types"
	CacheKeyShowTypes      = "show_types"
)
