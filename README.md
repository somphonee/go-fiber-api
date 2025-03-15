Go Fiber REST API
A robust REST API built with Go Fiber framework implementing Clean Architecture principles.
Features


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
