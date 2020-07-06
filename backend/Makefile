BASE_PKG := github.com/awanku/awanku
DOCKER_COMPOSE_TEST := docker-compose -f docker-compose.test.yml

build:
	go build -o ./dist/core-api $(BASE_PKG)/cmd/core-api

run:
	go run $(BASE_PKG)/cmd/core-api

run-dev:
	watchexec --watch . --exts go --signal SIGKILL --restart make run

test:
	./scripts/test.sh

docker-build: build
	docker build -f docker/Production.dockerfile -t awanku/core-api:latest .

docker-push:
	docker tag awanku/core-api:latest docker.awanku.id/awanku/core-api:latest
	docker push docker.awanku.id/awanku/core-api:latest

docker-test:
	$(DOCKER_COMPOSE_TEST) build core-api
	$(DOCKER_COMPOSE_TEST) up postgres-main -d
	$(DOCKER_COMPOSE_TEST) run core-api sh /app/core-api/scripts/docker-test.sh
	$(DOCKER_COMPOSE_TEST) stop
