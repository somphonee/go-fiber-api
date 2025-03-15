Go Fiber REST API
A robust REST API built with Go Fiber framework implementing Clean Architecture principles.
Features

Clean Architecture Structure - Organized in layers (models, repositories, services, handlers) for better maintainability
PostgreSQL Integration - Using GORM as ORM for database operations
JWT Authentication - Complete register and login system with JWT-protected routes
Request Validation - Data validation before database operations
Pagination Support - Efficient data retrieval with pagination
Search Functionality - Built-in search feature
Docker Ready - Docker and Docker Compose configuration for easy setup and deployment



Go 1.21+
Docker and Docker Compose
PostgreSQL

Quick Start
Using Docker

Clone the repository:

bashCopygit clone https://github.com/somphonee/go-fiber-api.git
cd go-fiber-api

Start the application:

bashCopydocker-compose up -d
The API will be available at http://localhost:8080.
Manual Setup

Clone the repository:

bashCopygit clone https://github.com/somphonee/go-fiber-api.git
cd go-fiber-api

Set up PostgreSQL database.
Update the configuration in configs/config.yml.
Run the application:

bashCopygo run cmd/api/main.go
API Endpoints
Authentication

POST /api/auth/register - Register a new user
POST /api/auth/login - Login and get JWT token

Resources

GET /api/resources - Get all resources (with pagination)
GET /api/resources/:id - Get a specific resource
POST /api/resources - Create a new resource (protected)
PUT /api/resources/:id - Update a resource (protected)
DELETE /api/resources/:id - Delete a resource (protected)
GET /api/resources/search - Search resources

Configuration
Configuration is stored in configs/config.yml:
yamlCopyserver:
  port: 8080
  
database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  name: go_fiber_api
  
jwt:
  secret: your_jwt_secret
  expiration: 24h
Development
Adding a New Resource

Create a model in internal/models
Implement repository in internal/repositories
Create service in internal/services
Add handlers in internal/handlers
Register routes in cmd/api/main.go

Testing
Run tests:
bashCopygo test ./...
License
MIT