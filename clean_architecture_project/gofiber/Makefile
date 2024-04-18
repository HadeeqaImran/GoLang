.PHONY: help run run-hot swag-init

help: ## Show available targets
	@echo "Available targets:"
	@echo "  run                 Run the application without hot reload"
	@echo "  run-hot             Run the application with nodemon for hot reload"
	@echo "  swag-init           Initialize Swagger documentation"

run: main.go ## Run the application without hot reload
	go run main.go

run-hot: main.go ## Run the application with nodemon for hot reload
	nodemon --watch './**/*.go' --ignore './docs/*' --signal SIGTERM --exec go run main.go

swag-init: ## Initialize Swagger documentation
	swag init
