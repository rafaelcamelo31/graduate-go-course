DOCKER_IMAGE=cloudrun-temperature-api
PORT=8080
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
		-p $(PORT):8080 \
		--name $(DOCKER_IMAGE) \
		$(DOCKER_IMAGE)

.PHONY: cloudrun-image-tag
cloudrun-image-tag: build
	docker tag cloudrun-temperature-api:latest us-west1-docker.pkg.dev/fullcycle-478823/cloudrun-temperature-api/cloudrun-temperature-api

.PHONY: cloudrun-image-push
cloudrun-image-push: cloudrun-image-tag
	docker push us-west1-docker.pkg.dev/fullcycle-478823/cloudrun-temperature-api/cloudrun-temperature-api

.PHONY: cloudrun-deploy
cloudrun-deploy:
	gcloud run deploy temperature-api --image us-west1-docker.pkg.dev/fullcycle-478823/cloudrun-temperature-api/cloudrun-temperature-api

.PHONY: help
help: ## show available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
