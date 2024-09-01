migrateCreate:
	migrate create -ext postgres -dir db/migration -seq init

migrationUp:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/rinha_backend_2023?sslmode=disable" -verbose up

migrationDown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/rinha_backend_2023?sslmode=disable" -verbose down

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

server:
	go run main.go

.PHONY: migrateCreate migrationUp migrationDown sqlc-gen server