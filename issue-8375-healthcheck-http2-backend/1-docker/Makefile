all:
	docker-compose up --build

build:
	CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o whoami



