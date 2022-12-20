all: stop start def whoami traefik resources stat

start:
	k3d cluster create mycluster \
		--api-port 6550 \
		--timestamps \
		-p 80:80@loadbalancer \
		-p 8080:8080@loadbalancer \
		-p 443:443@loadbalancer \
		-i rancher/k3s:v1.25.4-k3s1 \
		--k3s-arg '--disable=traefik@server:0' # --k3s-server-arg '--no-deploy=traefik' \ # --k3s-arg '--no-deploy=traefik@server:0'
		# -i rancher/k3s:v1.21.4-k3s1
		# -i rancher/k3s:v1.20.2-k3s1
		# -i rancher/k3s:v1.17.17-k3s1
		# -i rancher/k3s:v1.16.14-k3s1
	k3d image import traefik:v3.0 -c mycluster
	k3d image import traefik/traefik:latest -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

def:
	kubectl apply \
		-f ./00-definitions-v1-CRD.yml \
		-f ./01-RBAC-CRD-gateway.yml \
		-f ./02-definitions-v1-gateway.yml

whoami:
	kubectl apply -f ./10-whoami.yml

traefik:
	kubectl apply -f ./11-traefik.yml

resources:
	kubectl apply \
		-f ./20-ingressroute.yml \
		-f ./21-ingress.yml \
		-f ./22-gateway.yml

stat:
	kubectl get all

stop:
	k3d cluster delete mycluster
