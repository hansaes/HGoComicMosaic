run:
	go run ./cmd/api

test:
	go test ./...

fmt:
	go fmt ./...

deps:
	go mod tidy