# all: stop start def traefik resources stat
all: stop start def traefik stat

start:
	k3d cluster create mycluster \
		--api-port 6550 \
		--timestamps \
		-p 80:80@loadbalancer \
		-p 8080:8080@loadbalancer \
		-p 443:443@loadbalancer \
		-p 8000:8000@loadbalancer \
		-p 8443:8443@loadbalancer \
		--k3s-server-arg '--no-deploy=traefik' \
		-i rancher/k3s:v1.21.4-k3s1
		# -i rancher/k3s:v1.20.2-k3s1
		# -i rancher/k3s:v1.16.14-k3s1
		# -i rancher/k3s:v1.17.17-k3s1
	k3d image import traefik:v2.5 -c mycluster
	k3d image import traefik/traefik:latest -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

def:
	kubectl apply -f stack/

traefik:
	kubectl apply -f traefik.yml

resources:
	kubectl apply -f resources.yml

stat:
	kubectl get all

stop:
	k3d cluster delete mycluster
