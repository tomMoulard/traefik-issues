package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

// Worload API socket path.
const socketPath = "unix:///tmp/agent.sock"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a `/` resource handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		_, _ = io.WriteString(w, "Success!!!")
	})

	log.Println("connect to the Workload API")

	// Create a `workloadapi.X509Source`, it will connect to Workload API using
	// provided socket. If socket path is not defined using
	// `workloadapi.SourceOption`, value from environment variable
	// `SPIFFE_ENDPOINT_SOCKET` is used.
	source, err := workloadapi.NewX509Source(ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))
	if err != nil {
		log.Fatalf("Unable to create X509Source: %v", err)
	}
	defer source.Close()

	log.Println("create a TLS config with the X509Source")

	// Allowed SPIFFE ID
	clientID := spiffeid.RequireFromString("spiffe://example.org/traefik")

	// Create a `tls.Config` to allow mTLS connections, and verify that
	// presented certificate has SPIFFE ID `spiffe://example.org/client`
	tlsConfig := tlsconfig.MTLSServerConfig(source, source, tlsconfig.AuthorizeID(clientID))
	server := &http.Server{
		Addr:              ":8081",
		TLSConfig:         tlsConfig,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Serving on https://localhost:8081")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("Error on serve: %v", err)
	}
}
