# Exchange Rate API

In this hands-on project, you should apply what we learnt about http web server, contexts, databases and file handling in Go.

You must submit two systems in Go:

- client.go
- server.go

# Requirements

- `client.go` should send Http Request to `server.go` requesting USD exchange rate.
- `server.go` should consume following endpoint `https://economia.awesomeapi.com.br/json/last/USD-BRL`, the return its response in JSON to the client.
- `server.go` should save exchange rate to SQLite database, and the timeout for calling exchange rate API is 200ms, and timeout for database persistance is 10ms.
- `client.go` should receive current exchange rate ("bid" property in JSON). Using context package, set timeout as 300ms when calling server.go endpoint.
- All 3 contexts should return error as a log in case timeout is reached.
- `client.go` should save current exchange rate in a file called `rate.txt` with format {
  USD: value
  }

- The endpoint for server.go for this assignment is `/exchange-rate` with port 8080.

# Testing

- To run `server`, change directory to `projects/exchange_rate/` and run `make run-server`
- To run `client`, change directory to `projects/exchange_rate/` and run `make run-client`

- Note:
  - Running `go run server.go` won't execute the program. You must specify all files with `package main` like `go run .` or `go run server.go handler.go database.go`
