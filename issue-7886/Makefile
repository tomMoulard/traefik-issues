all: down up test

up:
	docker swarm init --advertise-addr 192.168.1.95
	docker network create -d overlay traefik_network --attachable
	docker stack deploy -c docker-stack.yml demo

test: update ab

update:
	docker service update --force demo_app &

ab:
	echo "GET http://localhost/" | vegeta attack -duration=10s | tee results.bin | vegeta report

down:
	docker swarm leave --force; \
		docker service rm $(docker service ls -q); \
		docker network prune -f
