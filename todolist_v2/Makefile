include .env
export

run:
	go run main.go

postgres-up:
	docker compose up -d postgres

postgres-down:
	docker compose stop postgres

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(LOCAL_DATABASE_URL)" up $(n)

migrate-down:
	migrate -path migrations -database "$(LOCAL_DATABASE_URL)" down $(n)