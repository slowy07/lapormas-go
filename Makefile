dev:
	go run ./cmd/web/main.go

build:
	go build -o ./bin/web ./cmd/web/main.go

migrate-up:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

.PHONY: start migrate-up migrate-down
