build:
	go build -o app cmd/server/main.go

run:
	docker compose up --build

test:
	go test -v ./...
