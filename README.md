# Theatre Management System API

A comprehensive Go-based REST API for managing theatres, locations, shows, and their relationships. Built with a layered architecture following best practices for maintainability and scalability.

## 🏗️ Architecture

The project follows a clean layered architecture:

```
├── main.go                 # Application entry point
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
├── Dockerfile             # Container configuration
├── docker-compose.yml     # Multi-service orchestration
├── init.sql              # Database initialization with sample data
├── src/
│   ├── controllers/      # HTTP request handlers
│   ├── business/         # Business logic layer
│   ├── dao/             # Data access objects
│   ├── dto/             # Data transfer objects
│   ├── models/          # Database models
│   ├── constants/       # Application constants
│   ├── mappers/         # Object mapping utilities
│   └── interfaces/      # Service interfaces
└── dev-tools/
    └── docker-compose.yml # Development database setup
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL with PostGIS (handled by Docker)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd theatre-management-system
```

### 2. Start Database

```bash
docker-compose up -d
```

This will start:

- PostgreSQL with PostGIS on port 5433
- pgAdmin on port 8888 (admin@theatre.com / admin)

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

### 5. Import Postman Collection (Optional)

For easy API testing, import the provided Postman collection:

```bash
# The collection file is located at:
Theatre_Management_API.postman_collection.json
```

**To import in Postman:**

1. Open Postman
2. Click "Import" in the top left
3. Select the `Theatre_Management_API.postman_collection.json` file
4. The collection includes all endpoints with sample requests and proper query parameters
5. Set the `baseUrl` variable to `http://localhost:8080` if it's not already set

The collection includes:

- All CRUD operations for all entities
- Health check endpoints
- Search and filtering examples
- Geographic queries with sample coordinates
- Pagination examples
- Sample request bodies with realistic data

## 📊 Database Schema

### Core Entities

- **Locations**: Geographic locations for theatres
- **Theatre Types**: Categories (Broadway, Off-Broadway, Regional, etc.)
- **Show Types**: Categories (Musical, Opera, Play, etc.)
- **Theatres**: Theatre venues with location and type relationships
- **Shows**: Productions playing at theatres

### Key Features

- UUID-based primary keys
- Soft deletes using GORM
- Geographic queries using PostGIS
- Proper foreign key relationships
- Timestamps for all entities

## 🔗 API Endpoints

### Health Check

- `GET /health` - API health status

### Locations

- `POST /api/v1/locations` - Create location
- `GET /api/v1/locations` - List locations (paginated)
- `GET /api/v1/locations/:id` - Get location by ID
- `PATCH /api/v1/locations/:id` - Update location
- `DELETE /api/v1/locations/:id` - Delete location
- `GET /api/v1/locations/active` - Get active locations
- `GET /api/v1/locations/nearby?latitude=40.7831&longitude=-73.9712&radius=50` - Find nearby locations
- `GET /api/v1/locations/search?q=manhattan` - Search locations

### Theatre Types

- `POST /api/v1/theatre-types` - Create theatre type
- `GET /api/v1/theatre-types` - List theatre types (paginated)
- `GET /api/v1/theatre-types/:id` - Get theatre type by ID
- `PATCH /api/v1/theatre-types/:id` - Update theatre type
- `DELETE /api/v1/theatre-types/:id` - Delete theatre type
- `GET /api/v1/theatre-types/active` - Get active theatre types
- `GET /api/v1/theatre-types/name/:name` - Get theatre type by name

### Show Types

- `POST /api/v1/show-types` - Create show type
- `GET /api/v1/show-types` - List show types (paginated)
- `GET /api/v1/show-types/:id` - Get show type by ID
- `PATCH /api/v1/show-types/:id` - Update show type
- `DELETE /api/v1/show-types/:id` - Delete show type
- `GET /api/v1/show-types/active` - Get active show types
- `GET /api/v1/show-types/name/:name` - Get show type by name

### Theatres

- `POST /api/v1/theatres` - Create theatre
- `GET /api/v1/theatres` - List theatres (paginated)
- `GET /api/v1/theatres/:id` - Get theatre by ID
- `PATCH /api/v1/theatres/:id` - Update theatre
- `DELETE /api/v1/theatres/:id` - Delete theatre
- `GET /api/v1/theatres/active` - Get active theatres
- `GET /api/v1/theatres/featured` - Get featured theatres
- `GET /api/v1/theatres/location/:locationId` - Get theatres by location
- `GET /api/v1/theatres/type/:typeId` - Get theatres by type
- `GET /api/v1/theatres/nearby?latitude=40.7831&longitude=-73.9712&radius=50` - Find nearby theatres
- `GET /api/v1/theatres/search?q=broadway` - Search theatres

### Shows

- `POST /api/v1/shows` - Create show
- `GET /api/v1/shows` - List shows (paginated)
- `GET /api/v1/shows/:id` - Get show by ID
- `PATCH /api/v1/shows/:id` - Update show
- `DELETE /api/v1/shows/:id` - Delete show
- `GET /api/v1/shows/active` - Get active shows
- `GET /api/v1/shows/featured` - Get featured shows
- `GET /api/v1/shows/current` - Get currently running shows
- `GET /api/v1/shows/upcoming` - Get upcoming shows
- `GET /api/v1/shows/theatre/:theatreId` - Get shows by theatre
- `GET /api/v1/shows/type/:typeId` - Get shows by type
- `GET /api/v1/shows/search?q=hamilton` - Search shows

## 📝 API Examples

### Create a Location

```bash
curl -X POST http://localhost:8080/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Broadway District",
    "city": "New York",
    "state": "New York",
    "country": "United States",
    "latitude": 40.7589,
    "longitude": -73.9851,
    "postal_code": "10036",
    "address": "Times Square, NYC",
    "description": "Heart of Broadway theater"
  }'
```

### Create a Theatre

```bash
curl -X POST http://localhost:8080/api/v1/theatres \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ambassador Theatre",
    "description": "Historic Broadway venue",
    "capacity": 1125,
    "address": "219 W 49th St, New York, NY",
    "phone": "(212) 239-6200",
    "email": "info@ambassador.com",
    "website": "https://ambassador.com",
    "location_id": "location-uuid-here",
    "theatre_type_id": "theatre-type-uuid-here"
  }'
```

### Search Shows

```bash
curl "http://localhost:8080/api/v1/shows/search?q=hamilton"
```

### Find Nearby Theatres

```bash
curl "http://localhost:8080/api/v1/theatres/nearby?latitude=40.7831&longitude=-73.9712&radius=10"
```

## 🔧 Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)
- `DATABASE_URL` - PostgreSQL connection string (default: postgres://postgres:postgres@localhost:5433/theatre_api?sslmode=disable)

### Pagination

- `limit` - Number of results per page (default: 20, max: 100)
- `offset` - Number of results to skip (default: 0)

## 🧪 Sample Data

The application includes comprehensive sample data:

- **8 Theatre Types**: Broadway, Off-Broadway, Regional, Community, etc.
- **10 Show Types**: Musical, Play, Opera, Ballet, Comedy, etc.
- **5 Locations**: Manhattan, West End London, Chicago, Las Vegas, Toronto
- **8 Theatres**: Famous venues like Majestic Theatre, Gershwin Theatre
- **8 Shows**: Popular productions like Hamilton, The Lion King, Chicago

## 🏢 Technology Stack

- **Framework**: Gin (HTTP router)
- **ORM**: GORM with PostgreSQL driver
- **Database**: PostgreSQL with PostGIS extension
- **Validation**: go-playground/validator
- **UUID**: Google UUID library
- **CORS**: Gin CORS middleware
- **Containerization**: Docker & Docker Compose

## 🔍 Advanced Features

### Geographic Queries

- Find locations and theatres within a radius using PostGIS
- Coordinate validation and distance calculations

### Caching Ready

- Structure prepared for Redis caching integration
- Cache keys defined in constants

### Comprehensive Error Handling

- Structured error responses
- Validation error details
- HTTP status code mapping

### Business Logic Validation

- Date range validation for shows
- Foreign key relationship validation
- Duplicate name prevention for types

## 🚀 Production Deployment

### Docker Build

```bash
docker build -t theatre-api .
docker run -p 8080:8080 -e DATABASE_URL="your-db-url" theatre-api
```

### Environment Setup

1. Set up PostgreSQL with PostGIS extension
2. Run database migrations (handled automatically)
3. Configure environment variables
4. Deploy container or binary

## 📋 API Response Format

All API responses follow a consistent format:

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... },
  "error": null
}
```

Error responses:

```json
{
  "success": false,
  "message": "Error message",
  "data": null,
  "error": "Detailed error information"
}
```

## 🧪 Testing

The architecture supports comprehensive testing:

- Unit tests for business logic
- Integration tests for API endpoints
- Repository tests with test database
- Mock interfaces for isolated testing

### Quick API Testing

For immediate testing and evaluation, use the provided **Postman collection** (`Theatre_Management_API.postman_collection.json`):

- **Complete coverage**: All 50+ endpoints included
- **Ready-to-use**: Sample data and proper query parameters
- **Geographic testing**: Coordinates for NYC Broadway district
- **Pagination examples**: Limit/offset parameters configured
- **Search functionality**: Example queries for all search endpoints

Perfect for candidates to quickly explore the API functionality without writing test code.

## 📚 Development

### Adding New Features

1. Define models in `src/models/`
2. Create DTOs in `src/dto/`
3. Implement repository interface in `src/interfaces/`
4. Create repository in `src/dao/`
5. Implement business service in `src/business/`
6. Create controller in `src/controllers/`
7. Add routes in `main.go`

### Code Structure Guidelines

- Use dependency injection throughout
- Implement interfaces for testability
- Follow repository pattern for data access
- Separate business logic from HTTP handling
- Use DTOs for API contracts

This theatre management system provides a robust foundation for managing theatrical venues, shows, and their relationships with proper geographic capabilities and a clean, maintainable architecture.
