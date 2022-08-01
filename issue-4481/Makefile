all: test

test:
	./test.sh

run:
	go run app.go

get_errors_500:
	ag 500 ~/tmp/tail.log | sed \
		-e 's/[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*/IP/g' \
		-e 's/time="[0-9A-Z:-]*"/time="..."/' \
		-e 's/.*http_/http_/g' \
		| sort -u | grep --color 500

get_errors_502:
	ag 502 ~/tmp/tail.log | sed -e 's/[0-9]*\.[0-9]*\.[0-9]*\.[0-9]*:[0-9]*/IP/g' -e 's/.*level/level/g' | sort -u

MAKEFLAGS += -j
traefik: traefik-vanilla
traefik: traefik-maxidleconnsperhost

traefik-vanilla:
	traefik \
		--api.insecure=true \
		--api.debug=true \
		--log.level=DEBUG \
		--entryPoints.web.address=:8000 \
		--entryPoints.traefik.address=:8080 \
		--providers.docker=true \
		--providers.file.directory=${PWD}
		# --providers.file.filename=${PWD}/traefik.dynamic.yml
		# --configFile=${PWD}/traefik.yml

traefik-maxidleconnsperhost:
	traefik \
		--api.insecure=true \
		--api.debug=true \
		--log.level=DEBUG \
		--entryPoints.web.address=:8001 \
		--entryPoints.traefik.address=:8081 \
		--serverstransport.maxidleconnsperhost=-1 \
		--providers.docker=true \
		--providers.file.directory=${PWD}

