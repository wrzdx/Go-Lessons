include .env
export

run:
	@go run main.go

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up $(n)

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down $(n)