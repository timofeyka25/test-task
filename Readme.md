## Test Task

This project is a small REST application written in Go. It uses PostgreSQL as the database, along with migrations for schema management, and is containerized using Docker.



### Project Structure
```
├── cmd
│   └── main.go              // Application entry point
├── compose.yaml             // Docker Compose file (brings up the database and the app)
├── config.example.yaml      // Example configuration file
├── Dockerfile               // Multi-stage Dockerfile for building the application
├── go.mod
├── go.sum
├── internal                 // Core application logic
│   ├── config
│   │   ├── config.go        // Application configuration
│   │   └── module.go
│   ├── container
│   │   └── di.go            // Dependency Injection container setup
│   ├── dto
│   │   ├── auth.go          // DTO models for authentication
│   │   └── records.go       // DTO models for records
│   ├── entities
│   │   ├── record.go        // Record entity definition
│   │   └── user.go          // User entity definition
│   ├── repository
│   │   ├── interfaces.go    // Repository interfaces
│   │   └── pgsql
│   │       ├── auth.go      // PostgreSQL implementation for auth repository
│   │       ├── module.go
│   │       └── record.go    // PostgreSQL implementation for records repository
│   ├── services
│   │   ├── auth.go          // Authentication service
│   │   ├── module.go
│   │   └── record.go        // Records service
│   └── transport
│       └── http
│           ├── config.go
│           ├── handler.go
│           ├── handlers
│           │   ├── auth.go  // HTTP handlers for authentication
│           │   ├── meta.go  // Health-check and other meta endpoints
│           │   ├── module.go
│           │   └── record.go // HTTP handlers for records
│           ├── middleware
│           │   └── auth.go  // Middleware for token authentication
│           ├── module.go
│           └── server.go    // HTTP server setup and initialization
├── migrations
│   └── pgsql
│       ├── 20250222102758_create_users_table.sql    // Migration to create users table
│       └── 20250222102810_create_records_table.sql  // Migration to create records table
└── pkg
    ├── jwt
    │   └── jwt.go         // JWT token generation and validation
    ├── pgsql
    │   ├── config.go
    │   ├── module.go
    │   └── pgsql.go       // PostgreSQL connection initialization
    ├── utils
    │   └── hash.go        // Functions for hashing and verifying passwords
    └── validator
        └── validator.go   // Example input validation functions

```

### Endpoints

GET ``/health``
Health-check endpoint to verify that the application is running.

POST `/sign-up`
User registration endpoint. Request body should include username and password.

POST `/sign-in`
User authentication endpoint. Request body should include username and password.

GET `/records/all`
Endpoint to retrieve records and test middleware authentication.

### How to Run the Project
1. Clone the repository: 
```bash 
git clone git@github.com:timofeyka25/test-task.git
cd test-task
```
2. Launch Docker Compose:
```bash
docker compose up -d
```
This command will build the application image and start both the PostgreSQL database and the app.

3. Access the Application:
the app will be available on port 8000. For example, check the health endpoint:
```bash 
curl http://localhost:8000/health
```