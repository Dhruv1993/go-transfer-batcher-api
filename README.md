## Transfer Money API

This is also a task for Upwork.The Transfer Money API is a Go-based application designed to mimic money transfers while incorporating a batching mechanism to enhance concurrency and performance.

## Project Structure

The project follows a modular structure to promote maintainability and readability. Here's an overview of the project structure:

````markdown
```go
.
├── Makefile
├── cmd
│   ├── gateway
│   │   └── main.go
│   └── test
├── concurency.go
├── docker-compose.yml
├── errors.go
├── genericHandler.go
├── go.mod
├── go.sum
├── services
│   ├── gateway
│   │   ├── routes.go
│   │   └── service.go
│   ├── microBatchService
│   │   ├── api.go
│   │   ├── domain.go
│   │   ├── handler.go
│   │   ├── microbatchService.go
│   │   ├── routes.go
│   │   ├── service.go
│   │   ├── simpleBatcher.go
│   │   └── swagger.yaml
│   └── service.go
└── utils
    ├── commons.go
    └── slices.go

```
````

cmd: Contains the main application entry points, such as the gateway and test entry points.
services: Contains various services, including the gateway service and the micro-batch service.
utils: Houses utility functions for common tasks.

## How to Run

Follow these steps to run the Transfer Money API:

1. Ensure you have Go installed on your machine.
2. Also ensure that docker-compose is installed as this api uses redis.
3. Clone the repository: `git clone <repository-url>`
4. Navigate to the project root: `cd <project-directory>`

To run the docker-compose, use the following command:

```bash
make run-tidy
### download dependencies
```

```bash
make run-redis
### Run redis and redis-insight
```

To run the API gateway, use the following command:

```bash
make run-gateway
### Run Gateway

```
