all: down up sleep test down

up:
	docker-compose up -d --build

down:
	docker-compose down

sleep:
	sleep 5s

test:
	curl -v http://localhost:8000/not/404
