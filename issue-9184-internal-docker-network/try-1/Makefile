TRAEFIK_CONTAINER := $(shell docker ps | grep traefik: | cut -d' ' -f 1)
TRAEFIK_IP := $(shell docker inspect ${TRAEFIK_CONTAINER} | jq -r .[0].NetworkSettings.Networks.foo.IPAddress)

curl:
	curl http://${TRAEFIK_IP}:8080/api/rawdata | jq '.services."whoami@docker",.routers."whoami@docker"'
