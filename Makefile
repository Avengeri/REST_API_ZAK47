include .env
export
start:
	docker compose up -d

stop:
	docker compose down

restart: stop start

up:
	docker compose up -d
down:
	docker compose down
swag:
	swag init -g ./cmd/main.go

migrate_up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable" up 1
migrate_down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://${DB_USER}:${DB_PASSWORD}@localhost:5432/postgres?sslmode=disable" down 1

build:
	@go build -o .bin/app.exe cmd/main.go
run: build
	@.bin/app.exe
make migrate:
	migrate create -ext sql -dir migrations add_users_table

