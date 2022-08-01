FROM golang:1.17.3 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/app .

FROM scratch
COPY --from=build /go/bin/app /app
ENTRYPOINT ["/app"]

