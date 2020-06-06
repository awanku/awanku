run:
	docker-compose up

run-build:
	docker-compose up --build

db-migrate:
	docker-compose exec core-api ./database/up.sh

db-clean:
	docker-compose exec core-api ./database/clean.sh

db-version:
	docker-compose exec core-api ./database/version.sh

db-nuke:
	docker-compose exec core-api ./database/nuke.sh

test-backend:
	docker-compose exec core-api make test

psql-main:
	docker-compose exec db-main su postgres -c "psql awanku"
