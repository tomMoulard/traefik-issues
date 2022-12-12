test:
	curl without-whoami.localhost -v 2>&1| grep '< HTTP/1.1 200 OK'
	curl with-whoami.localhost -v 2>&1| grep '< HTTP/1.1 200 OK'

export:
	curl localhost:8080/metrics > ${EXPORT}.log
