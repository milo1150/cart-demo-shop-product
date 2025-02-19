# Guide

## Database CLI

```bash
pgcli postgres://postgres:postgres@127.0.0.1:5432/cartdb
```

## Folder Structure

```bash
/project-root
│── /cmd/                  # Main application entry point
│── /config/               # Configuration files
│── /internal/
│   │── /database/         # Database connection setup
│   │   ├── postgres.go    # Initializes GORM with PostgreSQL
│   │── /models/           # GORM Models
│   │   ├── user.go        # User struct
│   │── /repositories/     # Database interactions
│   │   ├── user_repository.go
│   │── /services/         # Business logic
│   │   ├── user_service.go
│── /pkg/                  # Utility functions (helpers)
│── /api/                  # HTTP Handlers (Controllers)
│   ├── user_handler.go    # Handles user requests
│   ├── auth_handler.go    # Handles authentication
│   ├── product_handler.go # Handles product requests
│── /routes/               # Defines routes separately
│   ├── user_routes.go     # User-related routes
│   ├── auth_routes.go     # Auth-related routes
│   ├── product_routes.go  # Product-related routes
│   ├── router.go          # Main router setup
│── /migrations/           # SQL migration scripts
│── /main.go               # Entry point of the app
```
