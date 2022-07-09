DB_SOURCE: postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable

network:
	docker network create travel-network

dockerbuild:
	docker build -t gomcommerce:latest .

dockerun:
	docker run --name gomcommerce --network travel-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://postgres:0@postgres14.2:5432/mcommerce?sslmode=disable" gomcommerce:latest

postgres:
	docker run --name postgres14.2 --network travel-network -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=0  -d postgres:14.2-alpine

createdb: 
	docker exec -it postgres14.2 createdb --username=postgres --owner=postgres mcommerce

dropdb: 
	docker exec -it postgres14.2 dropdb --username=postgres mcommerce

migrateup: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose up

migrateup1: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose up 1

migratedown: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose down

migratedown1: 
	migrate --path db/migration -database "postgresql://postgres:0@localhost:5432/mcommerce?sslmode=disable" -verbose down 1

migratecreate:
	migrate create -ext sql -dir db/migration -seq add_mobil

createmigrate: 
	migrate create -ext sql -dir db/migration -seq init_schema

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

run: 
	go run main.go

mock: 
	mockgen -package mockdb -destination  db/mock/store.go github.com/capungkoneng/gomcommerce/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown createmigrate sqlc test