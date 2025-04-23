# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
# Echo app
APP_ENV=development
TIMEZONE=UTC
LOCAL_TIMEZONE=Asia/Bangkok

# Postgres
DATABASE_HOST=postgres-shop-product
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=shop_product_db
DATABASE_HOST_PORT=5433
DATABASE_DOCKER_PORT=5432

# MinIO
MINIO_ENDPOINT=minio-shop-product:9000
MINIO_ROOT_USER=admin
MINIO_ROOT_PASSWORD=password
MINIO_BROWSER_REDIRECT_URL=http://localhost/minio-sp
MINIO_API_URL=http://localhost/minio-sp-api
MINIO_PUBLIC_BUCKET_NAME=public-bucket

# Docker
COMPOSE_PROJECT_NAME=demo-shop-product-service
APP_BUILD_CONTEXT=../../
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
pgcli postgres://postgres:postgres@127.0.0.1:5433/shop_product_db
```
