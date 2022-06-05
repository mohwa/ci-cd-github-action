#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

find_files() {
  find . -name '*.go'
}

find_files | xargs gofmt -s -w
find_files | xargs goimports -w

go mod tidy
#go generate gqlgen.go
