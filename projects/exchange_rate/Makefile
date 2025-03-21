# Define project variables
SERVER_DIR=./server
CLIENT_DIR=./client
DB_CONTAINER_NAME=sqlite-go-container
DB_IMAGE_NAME=sqlite-go
DB_FILE=exchangerate.db

# Build and Run application
# ----------------------------------------------
build-server:
	go build -o server $(SERVER_DIR)

build-client:
	go build -o client $(CLIENT_DIR)

run-server: build-server
	$(SERVER_DIR)/server

run-client: build-client
	$(CLIENT_DIR)/client

# Build and Run Docker container
# ----------------------------------------------
build-db:
	docker build -t $(DB_IMAGE_NAME) .

run-db: build-db
	@touch $(DB_FILE)
	docker run --name $(DB_CONTAINER_NAME) -d -v $(PWD)/$(DB_FILE):/db/$(DB_FILE) $(DB_IMAGE_NAME)

# Stop and remove Docker container
# ----------------------------------------------
stop-db:
	docker stop $(DB_CONTAINER_NAME)
	docker rm $(DB_CONTAINER_NAME)

# Connect to SQLite shell
# ----------------------------------------------
sqlite-shell:
	docker exec -it $(DB_CONTAINER_NAME) sqlite3
