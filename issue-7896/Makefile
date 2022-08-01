
start:
	@echo "Starting stack"
	-k3d cluster delete mycluster
	k3d cluster create mycluster --api-port 6550 \
		-p 80:80@loadbalancer \
		-p 81:81@loadbalancer \
		-p 8080:8080@loadbalancer \
		-p 8081:8081@loadbalancer \
		-p 443:443@loadbalancer \
		-p 444:444@loadbalancer \
		--k3s-server-arg \
		'--no-deploy=traefik'\
		-i rancher/k3s:v1.20.2-k3s1
	# 	-i rancher/k3s:v1.16.14-k3s1
	# 	-i rancher/k3s:v1.17.17-k3s1
	k3d image import traefik:2.4.2 -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

def:
	@echo "Stack"
	-kubectl delete -f stack/
	kubectl apply -f stack/

traefik-internal:
	@echo "Traefik internal"
	-kubectl delete -f traefik-internal.yml
	kubectl apply -f traefik-internal.yml

traefik-external:
	@echo "Traefik external"
	-kubectl delete -f traefik-external.yml
	kubectl apply -f traefik-external.yml

resources:
	@echo "Resources"
	-kubectl delete -f resources.yaml
	kubectl apply -f resources.yaml

stop:
	@echo "Stopping stack"
	k3d cluster delete mycluster

get_pods:
	kubectl get pods

get_nodes:
	kubectl get nodes

stat: get_nodes
stat: get_pods

all: start
all: stack
all: def
# all: traefik-internal
all: traefik-external
all: resources
all: stat
