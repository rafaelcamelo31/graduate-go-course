# Automated Auction App

## Requirements

Implement:

A function that calculates the auction duration based on parameters defined via environment variables.

A goroutine responsible for:

Checking whether an auction has expired.

Updating its status to closed when the time is exceeded.

A test that validates the auction is being closed automatically.

## Deliverable

Full source code of the implementation.

Documentation explaining how to run the project in a development environment.

Use Docker / Docker Compose so the application can be tested easily.

## Full API Test Flow

To run unit test:

> make test

**Docker commands**

Start the application:

> make docker-up

Stop and remove the application:

> make docker-down

Restart containers:

> make docker-restart

View application logs?

> make docker-logs

**API Test Flow**

Address: http://localhost:8080

You can find each commands in [Makefile](./Makefile).

1. Create an auction

2. List active auctions

3. Wait for the auction to close automatically

4. List closed auctions

5. Confirm no active auctions remain

You can run **_`make test-api`_** to run test flow above.
