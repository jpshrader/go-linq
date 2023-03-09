run:
	go run ./cmd/main.go

benchmark:
	go test -bench=. -benchmem ./...

test:
	go test -cover ./...