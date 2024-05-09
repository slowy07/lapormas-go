dev:
	go run ./cmd/web/main.go

build:
	go build -o ./bin/web ./cmd/web/main.go

migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/lapormas-go?sslmode=disable" -path ./db/migrations up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/lapormas-go?sslmode=disable" -path ./db/migrations down

.PHONY: start migrate-up migrate-down
