FROM golang:1.17
WORKDIR /go/src/github.com/tommoulard/whoamigrpc
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o app ./cmd/server
EXPOSE 50051
ENTRYPOINT ["./app"]

