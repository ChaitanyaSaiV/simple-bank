postgrespull:
	docker pull postgres:latest

postgres:
	docker run -d --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres

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
	go build -o main ./cmd/main.go

server:
	go run ./cmd/.

image:
	docker build -t simplebank:latest .

container:
	docker run --name simplebank --network bank-network -p 8080:8080 -e DB_SOURCE="postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable" simplebank:latest

test:
	go test -v -cover ./...

vendor:
	go mod vendor

.PHONY: postgrespull postgres postgrescmd createdb dropdb migrateup migratedown sqlc build server image container test vendor