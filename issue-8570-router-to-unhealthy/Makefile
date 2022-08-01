all: up test down

up:
	docker-compose up -d

down:
	docker-compose down

test:
	sleep 1s
	curl localhost
	curl -X POST -d 500 blue.localhost/health
	sleep 2s
	curl localhost
