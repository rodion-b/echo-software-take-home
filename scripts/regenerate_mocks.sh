#!/bin/bash

# Regenerate mocks for interfaces
echo "Regenerating mocks..."

# Install mockgen if not available
if ! command -v mockgen &> /dev/null; then
    echo "Installing mockgen..."
    go install github.com/golang/mock/mockgen@latest
fi

# Remove existing mock files
echo "Removing existing mock files..."
rm -f internal/app/service/mocks.go
rm -f internal/app/client/fireblocks/mock_fireblocks_client.go
rm -f internal/app/httpserver/mocks.go

# Generate mocks for service interfaces
echo "Generating service mocks..."
mockgen -source=internal/app/service/interfaces.go -destination=internal/app/service/mocks.go -package=service

# Generate mocks for Fireblocks client
echo "Generating Fireblocks client mocks..."
mockgen -source=internal/app/client/fireblocks/fireblocks_client.go -destination=internal/app/client/fireblocks/mock_fireblocks_client.go -package=fireblocks

# Generate mocks for HTTP server interfaces
echo "Generating HTTP server mocks..."
mockgen -source=internal/app/httpserver/interfaces.go -destination=internal/app/httpserver/mocks.go -package=httpserver

echo "Mock regeneration complete!" 