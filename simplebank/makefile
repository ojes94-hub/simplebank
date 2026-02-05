postgres18:
	docker run --name postgres18 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=daniel@@1 -p 5432:5432 -d postgres:18.1-alpine3.23
rmpostgres18:
	docker rm -f postgres18
startdb:
	docker start postgres18
stopdb:
	docker stop postgres18
createdb:
	docker exec -it postgres18 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres18 dropdb simple_bank
dblog:
	docker logs postgres18
dblogin:
	docker exec -it postgres18 psql -U root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:daniel@@1@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:daniel@@1@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres18 stopdb createdb dropdb startdb rmpostgres18 dblog dblogin migrateup migratedown sqlc test