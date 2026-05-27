run:
	go run cmd/api/main.go

tidy:
	go mod tidy
	
migrate:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/log_service?sslmode=disable" up