all: down up sleep test

test:
	curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -d 542 -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
	# curl -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 542 status code 542'
	sleep 15s
	# curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 503 Service Unavailable'
	curl -d 200 -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
	# curl -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
	sleep 15s

sleep:
	sleep 2s

%.log:
	curl http://localhost:8080/metrics > $@

up:
	docker-compose up -d

down:
	docker-compose down
