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
├── redis
│   └── client.go
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
