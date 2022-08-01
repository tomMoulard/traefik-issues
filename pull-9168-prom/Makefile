test:
	curl -v whoami.localhost 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -d 542 -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 542 status code 542'
	curl -d 200 -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
	curl -v whoami.localhost/health 2>&1 | grep '< HTTP/1.1 200 OK'
