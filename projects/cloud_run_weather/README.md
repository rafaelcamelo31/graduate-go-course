# GCP Cloud Run - Temperature API

## Goal

The system in Go receives CEP as query parameter, identifies the city and returns the current weather (temperature in C, F, K).

The system must be deployed on Google Cloud Run.

## How to test

**Testing locally with Docker**

```sh
Example endpoint: http://localhost:8080/api/temperature?cep=13083970

# Build Go binary for execution
$ make build

# Run on docker container, listening on port 8080
$ make run

# Run unit tests on docker container
$ make test
```

**Testing deployed API**

```sh
Cloud Run: https://temperature-api-961628580524.us-west1.run.app/api/temperature?cep=13083970
```
