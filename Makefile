export PWD := $(shell pwd)

generate-mock:
	@go install github.com/golang/mock/mockgen@v1.6.0
	@PROJECT_DIR=${PWD} go generate ./...

.PHONY: swagger-discover
swagger-discover:
	@echo "Generating discover Swagger documentation..." \
	&& swag init -g cmd/discover-api/main.go -o docs/swag/discover \
	--parseDependency \
    --parseInternal \
    --parseVendor \
    --parseDepth 5
	@echo "Discover Swagger docs generated successfully!"

.PHONY: swagger-fmt
swagger-fmt:
	swag fmt

.PHONY: swagger-clean
swagger-clean:
	rm -rf docs/swag/*

.PHONY: swagger-all
swagger-all: swagger-clean swagger-discover

# Path to golangci-lint binary (override with `make GOLANGCI_LINT=/path/to/bin lint`)
GOLANGCI_LINT := golangci-lint

# Run all linters using the config file
lint:
	$(GOLANGCI_LINT) run ./...
