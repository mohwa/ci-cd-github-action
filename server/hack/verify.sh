#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

find_files() {
  find . -name '*.go'
}

diff=$(find_files | xargs gofmt -d -s 2>&1) || true

if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run ./update.sh" >&2
  exit 1
fi

diff=$(find_files | xargs goimports -d 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run ./update.sh" >&2
  exit 1
fi

go mod tidy
diff=$(git diff -- go.mod go.sum 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run ./update.sh and commit diff" >&2
  exit 1
fi

diff=$(git diff -- api/graphql 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}" >&2
  echo >&2
  echo "Run 'go generate gqlgen.go' and commit diff" >&2
  exit 1
fi
