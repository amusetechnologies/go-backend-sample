# Theatre Management System - API Documentation

## Overview

The Theatre Management System provides a comprehensive REST API for managing theatres, shows, locations, and their relationships. This API follows RESTful principles and returns JSON responses.

**Base URL**: `http://localhost:8080/api/v1`

## Authentication

Currently, the API does not require authentication. This can be added later using JWT tokens or API keys.

## Response Format

All API responses follow a consistent format:

### Success Response

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "error": "Detailed error information"
}
```

## Pagination

List endpoints support pagination with query parameters:

- `limit`: Number of results per page (default: 20, max: 100)
- `offset`: Number of results to skip (default: 0)

Example: `GET /api/v1/theatres?limit=10&offset=20`

## Geographic Queries

The API supports geographic queries using latitude, longitude, and radius:

- `latitude`: Latitude coordinate (-90 to 90)
- `longitude`: Longitude coordinate (-180 to 180)
- `radius`: Search radius in kilometers (default: 50)

Example: `GET /api/v1/theatres/nearby?latitude=40.7831&longitude=-73.9712&radius=10`

## Health Check Endpoints

### GET /health

Returns the overall health status of the API and its dependencies.

**Response:**

```json
{
  "success": true,
  "message": "OK",
  "data": {
    "status": "healthy",
    "timestamp": "2024-01-01T12:00:00Z",
    "version": "1.0.0",
    "environment": "development",
    "services": {
      "database": "healthy",
      "cache": "healthy"
    },
    "cache": {
      "item_count": 150,
      "size": 25
    }
  }
}
```

### GET /ready

Checks if the application is ready to serve requests.

### GET /live

Simple liveness check for container orchestration.

## Location Endpoints

### POST /api/v1/locations

Create a new location.

**Request Body:**

```json
{
  "name": "Manhattan",
  "city": "New York",
  "state": "New York",
  "country": "United States",
  "latitude": 40.7831,
  "longitude": -73.9712,
  "postal_code": "10019",
  "address": "Manhattan, NY",
  "description": "The heart of Broadway theater district",
  "is_active": true
}
```

**Response:** `201 Created`

```json
{
  "success": true,
  "message": "Location created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Manhattan",
    "city": "New York",
    "state": "New York",
    "country": "United States",
    "latitude": 40.7831,
    "longitude": -73.9712,
    "postal_code": "10019",
    "address": "Manhattan, NY",
    "description": "The heart of Broadway theater district",
    "is_active": true,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z",
    "theatres": []
  }
}
```

### GET /api/v1/locations

List all locations with pagination.

**Query Parameters:**

- `limit` (optional): Number of results per page
- `offset` (optional): Number of results to skip

### GET /api/v1/locations/:id

Get a specific location by ID.

### PATCH /api/v1/locations/:id

Update a location.

**Request Body:** Same as POST, all fields optional.

### DELETE /api/v1/locations/:id

Delete a location (soft delete).

### GET /api/v1/locations/active

Get all active locations.

### GET /api/v1/locations/nearby

Find locations within a radius of given coordinates.

**Query Parameters:**

- `latitude` (required): Latitude coordinate
- `longitude` (required): Longitude coordinate
- `radius` (optional): Search radius in kilometers

### GET /api/v1/locations/search

Search locations by name, city, or country.

**Query Parameters:**

- `q` (required): Search query string

## Theatre Type Endpoints

### POST /api/v1/theatre-types

Create a new theatre type.

**Request Body:**

```json
{
  "name": "Broadway",
  "description": "Professional theaters in Manhattan's Theater District with 500+ seats",
  "is_active": true
}
```

### GET /api/v1/theatre-types

List all theatre types with pagination.

### GET /api/v1/theatre-types/:id

Get a specific theatre type by ID.

### PATCH /api/v1/theatre-types/:id

Update a theatre type.

### DELETE /api/v1/theatre-types/:id

Delete a theatre type.

### GET /api/v1/theatre-types/active

Get all active theatre types.

### GET /api/v1/theatre-types/name/:name

Get a theatre type by name.

## Show Type Endpoints

### POST /api/v1/show-types

Create a new show type.

**Request Body:**

```json
{
  "name": "Musical",
  "description": "Theatrical productions featuring songs, spoken dialogue, acting, and dance",
  "is_active": true
}
```

### GET /api/v1/show-types

List all show types with pagination.

### GET /api/v1/show-types/:id

Get a specific show type by ID.

### PATCH /api/v1/show-types/:id

Update a show type.

### DELETE /api/v1/show-types/:id

Delete a show type.

### GET /api/v1/show-types/active

Get all active show types.

### GET /api/v1/show-types/name/:name

Get a show type by name.

## Theatre Endpoints

### POST /api/v1/theatres

Create a new theatre.

**Request Body:**

```json
{
  "name": "Majestic Theatre",
  "description": "Historic Broadway theater, home to The Phantom of the Opera for over 30 years",
  "capacity": 1681,
  "address": "245 W 44th St, New York, NY 10036",
  "phone": "(212) 239-6200",
  "email": "info@majestictheatre.com",
  "website": "https://majestictheatre.com",
  "image_url": "https://example.com/majestic.jpg",
  "is_featured": true,
  "is_active": true,
  "location_id": "123e4567-e89b-12d3-a456-426614174000",
  "theatre_type_id": "123e4567-e89b-12d3-a456-426614174001"
}
```

**Response:** `201 Created`

```json
{
  "success": true,
  "message": "Theatre created successfully",
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174002",
    "name": "Majestic Theatre",
    "description": "Historic Broadway theater, home to The Phantom of the Opera for over 30 years",
    "capacity": 1681,
    "address": "245 W 44th St, New York, NY 10036",
    "phone": "(212) 239-6200",
    "email": "info@majestictheatre.com",
    "website": "https://majestictheatre.com",
    "image_url": "https://example.com/majestic.jpg",
    "is_featured": true,
    "is_active": true,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z",
    "location_id": "123e4567-e89b-12d3-a456-426614174000",
    "theatre_type_id": "123e4567-e89b-12d3-a456-426614174001",
    "location": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "Manhattan",
      "city": "New York",
      "state": "New York",
      "country": "United States",
      "latitude": 40.7831,
      "longitude": -73.9712,
      "is_active": true
    },
    "theatre_type": {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "name": "Broadway",
      "description": "Professional theaters in Manhattan's Theater District with 500+ seats",
      "is_active": true
    },
    "shows": []
  }
}
```

### GET /api/v1/theatres

List all theatres with pagination.

### GET /api/v1/theatres/:id

Get a specific theatre by ID.

### PATCH /api/v1/theatres/:id

Update a theatre.

### DELETE /api/v1/theatres/:id

Delete a theatre.

### GET /api/v1/theatres/active

Get all active theatres.

### GET /api/v1/theatres/featured

Get all featured theatres.

### GET /api/v1/theatres/location/:locationId

Get theatres by location ID.

### GET /api/v1/theatres/type/:typeId

Get theatres by theatre type ID.

### GET /api/v1/theatres/nearby

Find theatres within a radius of given coordinates.

### GET /api/v1/theatres/search

Search theatres by name or description.

## Show Endpoints

### POST /api/v1/shows

Create a new show.

**Request Body:**

```json
{
  "title": "Hamilton",
  "description": "Lin-Manuel Miranda's revolutionary musical about Alexander Hamilton",
  "director": "Thomas Kail",
  "cast": "Lin-Manuel Miranda, Daveed Diggs, Leslie Odom Jr.",
  "duration": 165,
  "start_date": "2024-02-01T00:00:00Z",
  "end_date": "2024-11-30T00:00:00Z",
  "price": 200.0,
  "image_url": "https://example.com/hamilton.jpg",
  "trailer_url": "https://youtube.com/watch?v=hamilton",
  "is_featured": true,
  "is_active": true,
  "theatre_id": "123e4567-e89b-12d3-a456-426614174002",
  "show_type_id": "123e4567-e89b-12d3-a456-426614174003"
}
```

### GET /api/v1/shows

List all shows with pagination.

### GET /api/v1/shows/:id

Get a specific show by ID.

### PATCH /api/v1/shows/:id

Update a show.

### DELETE /api/v1/shows/:id

Delete a show.

### GET /api/v1/shows/active

Get all active shows.

### GET /api/v1/shows/featured

Get all featured shows.

### GET /api/v1/shows/current

Get shows that are currently running.

### GET /api/v1/shows/upcoming

Get shows that will start in the future.

### GET /api/v1/shows/theatre/:theatreId

Get shows by theatre ID.

### GET /api/v1/shows/type/:typeId

Get shows by show type ID.

### GET /api/v1/shows/search

Search shows by title, description, director, or cast.

## Error Codes

| HTTP Status | Error Code            | Description                             |
| ----------- | --------------------- | --------------------------------------- |
| 400         | Bad Request           | Invalid input data or malformed request |
| 404         | Not Found             | Requested resource not found            |
| 422         | Unprocessable Entity  | Validation failed                       |
| 500         | Internal Server Error | Server error                            |
| 503         | Service Unavailable   | Service temporarily unavailable         |

## Common Error Responses

### Validation Error (422)

```json
{
  "success": false,
  "message": "Validation failed",
  "error": "Field 'name' is required"
}
```

### Not Found Error (404)

```json
{
  "success": false,
  "message": "Theatre not found",
  "error": null
}
```

### Bad Request Error (400)

```json
{
  "success": false,
  "message": "Invalid UUID format",
  "error": "invalid UUID length: 35"
}
```

## Rate Limiting

Currently, no rate limiting is implemented. This can be added using middleware for production use.

## CORS

The API is configured to accept requests from any origin (`*`). In production, this should be restricted to specific domains.

## Data Types

### UUID

All entity IDs use UUID v4 format: `123e4567-e89b-12d3-a456-426614174000`

### Timestamps

All timestamps are in ISO 8601 format: `2024-01-01T12:00:00Z`

### Coordinates

- Latitude: Float between -90 and 90
- Longitude: Float between -180 and 180

### Capacity

Integer between 1 and 100,000

### Duration

Integer representing minutes (1 to 600)

### Price

Float representing currency amount (non-negative)

## Example Workflows

### Creating a Complete Theatre Setup

1. **Create Location**

```bash
curl -X POST http://localhost:8080/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{"name":"Broadway District","city":"New York","state":"New York","country":"United States","latitude":40.7589,"longitude":-73.9851}'
```

2. **Create Theatre Type**

```bash
curl -X POST http://localhost:8080/api/v1/theatre-types \
  -H "Content-Type: application/json" \
  -d '{"name":"Broadway","description":"Professional theaters in Manhattan"}'
```

3. **Create Theatre**

```bash
curl -X POST http://localhost:8080/api/v1/theatres \
  -H "Content-Type: application/json" \
  -d '{"name":"New Theatre","location_id":"[location-id]","theatre_type_id":"[type-id]","capacity":1500}'
```

4. **Create Show Type**

```bash
curl -X POST http://localhost:8080/api/v1/show-types \
  -H "Content-Type: application/json" \
  -d '{"name":"Musical","description":"Musical theater production"}'
```

5. **Create Show**

```bash
curl -X POST http://localhost:8080/api/v1/shows \
  -H "Content-Type: application/json" \
  -d '{"title":"New Musical","theatre_id":"[theatre-id]","show_type_id":"[show-type-id]","duration":120,"price":85.00}'
```

### Searching for Shows Near a Location

```bash
# Find theatres near Times Square
curl "http://localhost:8080/api/v1/theatres/nearby?latitude=40.7589&longitude=-73.9851&radius=5"

# Search for musicals
curl "http://localhost:8080/api/v1/shows/search?q=musical"

# Get current shows
curl "http://localhost:8080/api/v1/shows/current"
```

This API provides a comprehensive solution for managing theatre-related data with proper validation, error handling, and geographic capabilities.
