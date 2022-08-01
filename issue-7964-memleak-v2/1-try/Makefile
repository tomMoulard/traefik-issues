all: start

start:
	-kubectl delete namespaces ingresses
	go run main.go

setup: dc k3d

dc:
	docker-compose up -d

k3d:
	./start.sh
