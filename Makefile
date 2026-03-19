.PHONY: build run test lint fmt clean install-hooks

BINARY=schoty
MAIN_PATH=./cmd/schoty
VERSION?=dev
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

build:
	go build $(LDFLAGS) -o $(BINARY) $(MAIN_PATH)

run: build
	./$(BINARY)

test:
	go test -v ./...

lint:
	go vet ./...
	golangci-lint run

fmt:
	go fmt ./...

clean:
	rm -f $(BINARY)

install-hooks:
	git config core.hooksPath .githooks
