DOCKER_IMAGE=temperature-api
PORT=8081
.DEFAULT_GOAL := help

.PHONY: build
build: ## build docker image
	docker build \
		--load \
		-t "$(DOCKER_IMAGE)" \
		-f Dockerfile \
		.

.PHONY: test
test: ## build and run unit tests
	docker build \
		-f Dockerfile \
		-t "$(DOCKER_IMAGE)" \
		--progress plain \
		--no-cache \
		--target tester \
		.

.PHONY: run
run: build ## run container with application
	docker run --rm \
		-p $(PORT):8081 \
		--name $(DOCKER_IMAGE) \
		$(DOCKER_IMAGE)

.PHONY: help
help: ## show available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
