all: down up sleep test

up:
	docker-compose up -d

down:
	docker-compose down

sleep:
	sleep 20s

PORT ?= 443

test:
	@echo "" | openssl s_client -connect localhost:${PORT} -servername localhost 2>&1 | grep Verification:
