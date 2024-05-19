DB_NAME=anki
DB_URL=postgresql://root:secret@localhost:5432/"$(DB_NAME)"?sslmode=disable


# docker
postgres:
	 docker run --name postgres_foogo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e TZ=Asia/Shanghai -d postgres:alpine3.18
createdb:
	docker exec -it postgres_foogo createdb --username=root --owner=root "$(DB_NAME)"
dropdb:
	docker exec -it postgres_foogo dropdb $(DB_NAME)
psql:
	docker exec -it postgres_foogo psql -U root "$(DB_NAME)"
stop:
	docker stop postgres_foogo
start:
	docker start postgres_foogo

# migrate up or down
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# sqlc
sqlc:
	sqlc generate

# mock
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/FeiNiaoBF/simplebank_backend_learn/db/sqlc Store
# go
test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: createdb dropdb postgres migratedown \
		migrateup sqlc stop start test server mock \
		psql