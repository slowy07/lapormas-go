dev:
	go run ./cmd/web/main.go

build:
	go build -o ./bin/web ./cmd/web/main.go

start:
	./bin/web

migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/lapormas-go?sslmode=disable" -path ./db/migrations up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/lapormas-go?sslmode=disable" -path ./db/migrations down

.PHONY: dev build start migrate-up migrate-down
