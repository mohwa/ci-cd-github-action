.PHONY: default verify update lint build test run fmt diff gen imports clean

TARGET = server

default: run

verify:
	./hack/verify.sh

update:
	./hack/update.sh

lint:
	${HOME}/go/bin/golangci-lint run ./...

build:
	go build -o $(TARGET) ./cmd/server

test:
	go test -v ./...

run:
	go run ./cmd/server

fmt:
	go fmt ./...

diff:
	git diff */*.go

gen:
	go generate hack/gqlgen.go

imports:
	goimports -w **/*.go

clean:
	rm -f $(TARGET)
