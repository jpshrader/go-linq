# starts the api server
run:
	go run ./cmd/main.go

test:
	go test -cover ./...