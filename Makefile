postgres:
	docker run --name postgres_db -p 5500:5432 -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it postgres_db createdb --username=postgres --owner=postgres gobank

dropdb:
	docker exec -it postgres_db dropdb gobank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5500/gobank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5500/gobank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/suryadharma5/gobank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock