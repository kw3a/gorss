.PHONY: migrate-up migrate-down build run start sqlc-generate

migrate-up:
	goose -dir sql/schema/ postgres "postgres://postgres:postgres@localhost:5432/gorss?sslmode=disable" up

migrate-down:
	goose -dir sql/schema/ postgres "postgres://postgres:postgres@localhost:5432/gorss?sslmode=disable" down

compile:
	go build -o out

run:
	./out

start: compile run

sqlc-generate:
	sqlc generate

