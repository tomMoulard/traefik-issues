FROM golang:1.17
WORKDIR /go/src/github.com/tommoulard/app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN go build -o app app.go
EXPOSE 8080
ENTRYPOINT ["./app"]
