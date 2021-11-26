build:
	go build ./cmd/main.go

run:
	docker-compose up

test:
	go test ./...