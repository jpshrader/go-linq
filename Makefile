run:
	go run ./cmd/examples/main.go

benchmark:
	go test -bench=. -benchmem ./...

test:
	go test -cover ./...

update:
	go get -u
	go mod tidy