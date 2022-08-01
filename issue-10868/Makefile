all: start def traefik resources stat
# all: start stack def resources stat

start:
	@echo "Starting stack"
	-k3d cluster delete mycluster
	k3d cluster create mycluster --api-port 6550 -p 80:80@loadbalancer -p 8080:8080@loadbalancer -p 443:443@loadbalancer --k3s-server-arg '--no-deploy=traefik'\
 	-i rancher/k3s:v1.20.2-k3s1
# 	-i rancher/k3s:v1.16.14-k3s1
# 	-i rancher/k3s:v1.17.17-k3s1
	k3d image import traefik:v2.4 -c mycluster
	# k3d image import traefik/traefik:latest -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

def:
	@echo "Stack"
	-kubectl delete -f stack/
	kubectl apply -f stack/

traefik:
	@echo "Traefik"
	-kubectl delete -f traefik.yml
	kubectl apply -f traefik.yml

resources:
	@echo "Resources"
	-kubectl delete -f resources.yaml
	kubectl apply -f resources.yaml

stat:
	watch kubectl get all

stop:
	@echo "Stopping stack"
	k3d cluster delete mycluster
