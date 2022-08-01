#!/bin/bash

PIPE=mypipe

# set -x

function curl_test() {
    printf %"$(tput cols)"s |tr " " "-"

    # echo "Testing with curl: $1: ${URL}"
    # curl -v ${URL} ${URL} 1>/dev/null 2>${PIPE} \
        # | grep --color -i -e '#0' -e "< HTTP" ${PIPE}

    echo "Testing with vegeta: $1: ${URL}"
    echo "GET ${URL}" \
        | vegeta attack -rate=0 -max-workers=12 -duration=20s -http2  \
        | tee ~/tmp/$(date +%d-%H:%M:%S)_results.bin \
        | vegeta report

    echo "Testing with vegeta and body: $1: ${URL}"
    echo "POST ${URL}" \
        | vegeta attack -rate=0 -max-workers=12 -duration=20s -http2  -body 'Makefile' \
        | tee ~/tmp/$(date +%d-%H:%M:%S)_results.bin \
        | vegeta report

    # echo "Testing with ab: $1: ${URL}"
    # ab -n 5000 ${URL}/

    # echo "Testing with hey: $1: ${URL}"
    # hey ${URL} -d "toto"
}

[ -p "${PIPE}" ] || mkfifo ${PIPE}

URL=http://localhost:3030 curl_test "no ka"
URL=http://localhost:3031 curl_test "ka"

# URL=http://nka.whoami.localhost:8000 curl_test "no ka with traefik"
# URL=http://ka.whoami.localhost:8000 curl_test "ka with traefik"

# URL=http://nka.whoami.localhost:8001 curl_test "no ka with traefik maxidleconnsperhost"
# URL=http://ka.whoami.localhost:8001 curl_test "ka with traefik maxidleconnsperhost"


URL=http://localhost:3030/closing curl_test "no ka closing request body"
URL=http://localhost:3031/closing curl_test "ka closing request body"

# URL=http://nka.whoami.localhost:8000/closing curl_test "no ka closing with traefik"
# URL=http://ka.whoami.localhost:8000/closing curl_test "ka closing with traefik"

# URL=http://nka.whoami.localhost:8001/closing curl_test "no ka closing with traefik maxidleconnsperhost"
# URL=http://ka.whoami.localhost:8001/closing curl_test "ka closing with traefik maxidleconnsperhost"


# URL=http://nka.backend.localhost:8000 curl_test "nka uwsgi with traefik"
# URL=http://backend.localhost:8000 curl_test "ka uwsgi with traefik"

# URL=http://nka.backend.localhost:8001 curl_test "nka uwsgi with traefik maxidleconnsperhost"
# URL=http://backend.localhost:8001 curl_test "ka uwsgi with traefik maxidleconnsperhost"

# Reproducible case

# URL=http://localhost:9090 curl_test "nka uwsgi direct"
# URL=http://localhost:9091 curl_test "ka uwsgi direct"

# URL=http://nka.backend.localhost:8888 curl_test "nka uwsgi with traefik"
# URL=http://backend.localhost:8888 curl_test "ka uwsgi with traefik"

rm ${PIPE}
