.PHONY :  run-gateway run-tidy run-redis down-redis

run-gateway : 
		go run cmd/gateway/main.go

run-tidy : 
		go mod tidy

run-redis:
	docker-compose up

down-redis:
	docker-compose down