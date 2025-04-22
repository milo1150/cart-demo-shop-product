#!/bin/sh

# Load environment
. ./load-env.sh

# Run Docker Compose with debug
docker-compose -f ../deployments/dev/docker-compose.yaml build --no-cache --progress=plain
