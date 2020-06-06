BASE_PKG := github.com/awanku/awanku/backend

build:
	go build -o ./dist/core-api $(BASE_PKG)/cmd/core-api

run:
	go run $(BASE_PKG)/cmd/core-api

run-dev:
	watchexec --watch . --exts go --signal SIGKILL --restart make run

test:
	go test -v -race ./...
