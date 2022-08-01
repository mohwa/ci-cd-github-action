#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

find_files() {
  find . -name '*.go'
}

find_files | xargs gofmt -s -w
find_files | xargs goimports -w

# 사용하지않는 의존성을 제거한다.
# https://soyoung-new-challenge.tistory.com/130
go mod tidy
go generate gqlgen.go
