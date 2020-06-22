docker-build:
	docker build -f docker/Production.dockerfile -t awanku/console-webui:latest .

docker-push:
	docker tag awanku/console-webui:latest docker.awanku.id/awanku/console-webui:latest
	docker push docker.awanku.id/awanku/console-webui:latest
