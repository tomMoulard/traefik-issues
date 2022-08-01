FROM golang:1-alpine as builder

WORKDIR /go/app

COPY app.go .

RUN CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags="-s" -o app app.go

# Create a minimal container to run a Golang static binary
FROM scratch

COPY --from=builder /go/app/app .

ENTRYPOINT ["/app", "-port", "80"]
EXPOSE 80
