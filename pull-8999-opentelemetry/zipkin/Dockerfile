FROM golang:1.17.4

WORKDIR /go/src/github.com/tommoulard/app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

EXPOSE 80

COPY . .

RUN go build -o app .
ENTRYPOINT ["./app"]
