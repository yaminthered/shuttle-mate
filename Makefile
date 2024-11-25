DB_URL=postgresql://root:secret@localhost:5432/shuttle-mate?sslmode=disable

server:
	go run main.go

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

.PHONY: server new_migration migrateup migratedown sqlc 