build:
	go build ./cmd/main.go

run:
	docker-compose up

test:
	go test ./...

gen_docs:
	swag init -g cmd/main.go