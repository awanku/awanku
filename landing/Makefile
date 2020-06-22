docker-build:
	docker build -f docker/Production.dockerfile -t awanku/landing-webui:latest .

docker-push:
	docker tag awanku/landing-webui:latest docker.awanku.id/awanku/landing-webui:latest
	docker push docker.awanku.id/awanku/landing-webui:latest
