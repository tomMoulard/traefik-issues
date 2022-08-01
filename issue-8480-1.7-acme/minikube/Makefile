all: stop start apply

stop:
	minikube delete

start:
	minikube start --kubernetes-version=v1.16.0
	minikube image load traefik/traefik:latest

apply:
	minikube kubectl -- apply -f https://raw.githubusercontent.com/traefik/traefik/v1.7/examples/k8s/traefik-rbac.yaml
	minikube kubectl -- apply -f https://raw.githubusercontent.com/traefik/traefik/v1.7/examples/k8s/traefik-deployment.yaml

