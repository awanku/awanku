staging-up:
	docker-compose -f docker-compose.staging.yml up

nomad:
	NOMAD_ADDR=http://nomad.service.consul:4646 nomad run awanku-systems.hcl
