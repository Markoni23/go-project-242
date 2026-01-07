test:
	go mod tidy
	go test -v ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

install:
	go install

lint:
	golangci-lint run ./...