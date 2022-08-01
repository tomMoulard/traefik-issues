FROM golang:1.15

RUN go get github.com/traefik/traefik
RUN apt update
RUN apt upgrade -y
RUN apt install -y graphviz

COPY . .

EXPOSE 8081


CMD go tool pprof -http=0.0.0.0:8081 ./*.pprof
