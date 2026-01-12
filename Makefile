test:
	go mod tidy
	go test -v ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

install:
	go install


ARGS?=tests/testdata
run:
	go run ./cmd/hexlet-path-size/main.go $(ARGS)

lint:
	golangci-lint run ./...