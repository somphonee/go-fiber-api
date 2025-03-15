Go Fiber REST API
A clean architecture REST API built with Go Fiber and PostgreSQL.
Features

Clean Architecture (models, repositories, services, handlers)
PostgreSQL with GORM
JWT Authentication
Request Validation
Pagination
Search Functionality
Docker Support

Setup
Docker
bashCopydocker-compose up -d
Manual

Setup PostgreSQL
Configure configs/config.yml
Run: go run cmd/api/main.go

API Endpoints
Auth

POST /api/auth/register - Register
POST /api/auth/login - Login

Resources

GET /api/resources - List all (paginated)
GET /api/resources/:id - Get one
POST /api/resources - Create (protected)
PUT /api/resources/:id - Update (protected)
DELETE /api/resources/:id - Delete (protected)
GET /api/resources/search - Search

Configuration
Edit configs/config.yml for server, database, and JWT settings.
Development
Follow clean architecture principles when adding new resources.