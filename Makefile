build:
	docker-compose up -d
	go mod tidy
run:
	go run main.go
down:
	docker-compose down
