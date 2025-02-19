# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
APP_ENV=development
DATABASE_HOST=postgres-cart-demo
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=cartdb
DATABASE_HOST_PORT=5432
DATABASE_DOCKER_PORT=5432
TIMEZONE=UTC
LOCAL_TIMEZONE=Asis/Shanghai
```

- For first time.

```bash
cd scripts && chmod +x dev-start.sh && ./dev-start.sh
```

- Later

```bash
cd scripts && ./dev-start.sh
```

## Database CLI

```bash
pgcli postgres://postgres:postgres@127.0.0.1:5432/cartdb
```

## Debug docker build (Dev)

```bash
docker-compose -f internal/deployments/dev/docker-compose.yaml build --progress=plain --no-cache
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
