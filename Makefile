postgres:
	docker run --name postgres14.2 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=0  -d postgres:14.2-alpine

createdb: 
	docker exec -it postgres14.2 createdb --username=postgres --owner=postgres mcommerce

dropdb: 
	docker exec -it postgres14.2 dropdb --username=postgres mcommerce

migrateup: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose up

migratedown: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose down

createmigrate: 
	migrate create -ext sql -dir db/migration -seq init_schema

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

run: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown createmigrate sqlc test