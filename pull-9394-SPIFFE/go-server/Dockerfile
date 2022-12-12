FROM golang:1.19

WORKDIR /go/src/github.com/tommoulard/siffe-server

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify
RUN go mod tidy

COPY . .

RUN go build -o /go/bin/spiffe-server ./...

EXPOSE 8081

CMD ["/go/bin/spiffe-server"]
