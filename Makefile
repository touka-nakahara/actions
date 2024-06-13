.PHONY: test

test:
	go test ./cmd/

.PHONY: build

build:
	go build -o ./build/main ./cmd/main.go
