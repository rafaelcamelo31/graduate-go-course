APP_NAME=auction
DOCKER_COMPOSE=docker compose
GO=go
GOFLAGS=-v
BUILD_DIR=bin
MAIN_PATH=./cmd

COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m

help: ## Show this help message
	@echo '$(COLOR_BOLD)Available targets:$(COLOR_RESET)'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(COLOR_GREEN)%-15s$(COLOR_RESET) %s\n", $$1, $$2}'

test: ## Run tests
	@echo "$(COLOR_YELLOW)Running tests...$(COLOR_RESET)"
	$(GO) test -v -race
	@echo "$(COLOR_GREEN)Tests complete$(COLOR_RESET)"

docker-up: ## Start Docker containers
	@echo "$(COLOR_YELLOW)Starting Docker containers...$(COLOR_RESET)"
	$(DOCKER_COMPOSE) up -d --build
	@echo "$(COLOR_GREEN)Containers started$(COLOR_RESET)"

docker-down: ## Stop Docker containers
	@echo "$(COLOR_YELLOW)Stopping Docker containers...$(COLOR_RESET)"
	$(DOCKER_COMPOSE) down -v
	@echo "$(COLOR_GREEN)Containers stopped$(COLOR_RESET)"

docker-logs: ## Show Docker container logs
	$(DOCKER_COMPOSE) logs -f

docker-restart: docker-down docker-up ## Restart Docker containers

# API Testing targets
API_URL=http://localhost:8080

create-auction: ## Create a test auction
	@echo "$(COLOR_YELLOW)Creating auction...$(COLOR_RESET)"
	@curl -X POST $(API_URL)/auction \
		-H "Content-Type: application/json" \
		-d '{ "product_name": "Lightsaber", "category": "Weapon", "description": "Blue colored jedi knights lightsaber", "condition": 0 }' | jq .
	@echo ""

list-active: ## List auctions with status=0 (active)
	@echo "$(COLOR_YELLOW)Listing active auctions (status=0)...$(COLOR_RESET)"
	@curl -X GET $(API_URL)/auction?status=0 | jq .
	@echo ""

list-closed: ## List auctions with status=1 (closed)
	@echo "$(COLOR_YELLOW)Listing closed auctions (status=1)...$(COLOR_RESET)"
	@curl -X GET "$(API_URL)/auction?status=1" | jq .
	@echo ""

test-api: ## Run full API test flow
	@echo "$(COLOR_GREEN)========== Running Full API Test Flow ==========$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_BOLD)1. Creating auction...$(COLOR_RESET)"
	@make create-auction
	@sleep 1
	@echo ""
	@echo "$(COLOR_BOLD)2. Listing all active auction...$(COLOR_RESET)"
	@make list-active
	@sleep 1
	@echo "$(COLOR_BOLD)3. Wait 15 seconds until auction is closed..."
	@sleep 15
	@echo ""
	@echo "$(COLOR_BOLD)4. Listing closed auction...$(COLOR_RESET)"
	@make list-closed
	@echo ""
	@sleep 1
	@echo "$(COLOR_BOLD)5. Make sure there are no more active auction..."
	@make list-active
	@sleep 1
	@echo "$(COLOR_GREEN)========== Test Flow Complete ==========$(COLOR_RESET)"