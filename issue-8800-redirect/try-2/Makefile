all: stop start traefik resources stat

start:
	k3d cluster create mycluster \
		--api-port 6550 \
		--timestamps \
		-p 80:80@loadbalancer \
		-p 8080:8080@loadbalancer \
		-p 443:443@loadbalancer \
		--k3s-arg '--no-deploy=traefik@server:0' # --k3s-server-arg '--no-deploy=traefik' \ # --k3s-arg '--no-deploy=traefik@server:0'
		-i rancher/k3s:v1.21.4-k3s1
		# -i rancher/k3s:v1.20.2-k3s1
		# -i rancher/k3s:v1.17.17-k3s1
		# -i rancher/k3s:v1.16.14-k3s1
	k3d image import traefik:v2.6 -c mycluster
	k3d image import traefik/traefik:latest -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

traefik:
	-helm repo add traefik https://helm.traefik.io/traefik
	-helm repo update
	helm install -f values.yml traefik traefik/traefik --replace

traefik-dashboard:
	xdg-open http://localhost:9000/dashboard/
	kubectl port-forward $(shell kubectl get pods --selector "app.kubernetes.io/name=traefik" --output=name) 9000:9000

resources:
	kubectl apply -f whoami.yml
	kubectl apply -f resources.yml

stat:
	kubectl get all

stop:
	k3d cluster delete mycluster
