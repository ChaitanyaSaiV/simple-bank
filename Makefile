postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres

postgrescmd:
	docker exec -it postgres psql -U root

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres drop simple_bank

migrateup:
	migrate -path internal/db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

build:
	go build -o simple-bank ./cmd/main.go

run:
	go run ./cmd/.

.PHONY: postgres postgrescmd createdb dropdb migrateup migratedown sqlc build run