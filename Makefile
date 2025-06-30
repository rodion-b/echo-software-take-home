run:
	docker-compose up  --remove-orphans --build

# Generate mocks for interfaces
generate-mocks:
	mockgen -source=internal/app/service/interfaces.go -destination=internal/app/service/mocks.go -package=service
	mockgen -source=internal/app/client/fireblocks/fireblocks_client.go -destination=internal/app/client/fireblocks/mock_fireblocks_client.go -package=fireblocks

# Regenerate mocks (removes old ones first)
regenerate-mocks:
	./scripts/regenerate_mocks.sh

# Install mockgen if not available
install-mockgen:
	go install github.com/golang/mock/mockgen@latest

# Generate all mocks (install mockgen first if needed)
mocks: install-mockgen generate-mocks

# Run tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Install dependencies
deps:
	go mod download
	go mod tidy
	
# Install golangci-lint if not present
install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
lint:
	golangci-lint run ./...