BASE_PKG := github.com/awanku/awanku

build:
	go build -o ./dist/core-api $(BASE_PKG)/cmd/core-api

run:
	go run $(BASE_PKG)/cmd/core-api

run-dev:
	watchexec --watch . --exts go --signal SIGKILL --restart make run

test: export DB_URL = postgres://postgres:@localhost/awanku?sslmode=disable
test:
	./database/nuke.sh
	./database/up.sh
	go test -v -race ./...

docker-build: build
	docker build -f docker/Production.dockerfile -t awanku/core-api:latest .

docker-push:
	docker tag awanku/core-api:latest docker.awanku.id/awanku/core-api:latest
	docker push docker.awanku.id/awanku/core-api:latest
