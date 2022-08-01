CONSUL := docker-compose exec consul consul

all: namespaces services

namespaces:
	${CONSUL} namespace write /json/namespace1.json
	${CONSUL} namespace write /json/namespace2.json

services:
	${CONSUL} services register /json/whoami1.hcl
	${CONSUL} services register /json/whoami2.hcl
