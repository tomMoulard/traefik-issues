all: down up

up:
	docker swarm init --advertise-addr 2a01:cb1c:81fa:d500:6fe7:3dd:33be:ba55/64
	# docker swarm init # --advertise-addr 192.168.1.27
	docker network create \
		--driver overlay \
		--ipv6 \
		--subnet=2a01:cb1c:81fa:d500:6fe7:3dd:33be:ba55/64 \
		--attachable ext_net
	# docker network create -d overlay traefik_network --attachable
	docker stack deploy -c docker-compose.yml demo

down:
	docker swarm leave --force; \
		docker service rm $(docker service ls -q); \
		docker network prune -f

ps:
	docker service ps demo_whoami

update:
	docker service update demo_whoami
