#!make
include ./.env
export $(shell sed 's/=.*//' .env)


build:
	go build -o app cmd/server/main.go

run:
	docker compose up --build

test:
	go test -v ./...

integration-test:
	docker compose --env-file ./.env up -d book-db
	go test --tags=integration -v ./...
