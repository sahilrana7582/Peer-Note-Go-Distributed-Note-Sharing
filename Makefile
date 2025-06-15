include .env
export

dirpath=./internal/db/migrations

run:
	go run ./cmd/main.go

migrate-up:
	migrate -path $(dirpath) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(dirpath) -database "$(DB_URL)" down

migrate-drop:
	migrate -path $(dirpath) -database "$(DB_URL)" drop -f

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(dirpath) -seq $$name

