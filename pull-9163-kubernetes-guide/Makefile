all: stop start traefik whoami stat

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
	k3d image import traefik:v2.8 -c mycluster
	k3d image import traefik/traefik:latest -c mycluster
	k3d image import traefik/whoami:latest -c mycluster

traefik:
	kubectl apply \
		-f 00-role.yml \
		-f 00-account.yml \
		-f 01-role-binding.yml \
		-f 02-traefik.yml \
		-f 02-traefik-services.yml

whoami:
	kubectl apply \
		-f 03-whoami.yml \
		-f 03-whoami-services.yml \
		-f 04-whoami-ingresses.yml

stat:
	kubectl get all

stop:
	k3d cluster delete mycluster
