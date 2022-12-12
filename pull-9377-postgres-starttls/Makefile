GOROOT=$(shell go env GOROOT)
DOMAIN ?= postgres

certs:
	mkdir -p certs
	go run ${GOROOT}/src/crypto/tls/generate_cert.go \
		--rsa-bits 4096 --host ${DOMAIN} --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
	chmod 600 certs/key.pem
	chown 999:999 certs/key.pem
	chmod 666 certs/cert.pem
	chown 0:0 certs/cert.pem
