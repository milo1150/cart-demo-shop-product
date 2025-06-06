#!/bin/bash

# Check if migration name is provided
if [ -z "$1" ]; then
	echo "Usage: $0 <migration_name>"
	exit 1
fi

# Set migration directory
MIGRATION_DIR="../internal/database/migrations"

# Run golang-migrate CLI with the provided migration name
migrate create -ext sql -dir "$MIGRATION_DIR" -seq "$1"

# Output success message
echo "Migration '$1' created successfully in $MIGRATION_DIR."
