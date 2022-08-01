package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "8000", "give me a port number")
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	_, _ = fmt.Fprintln(w, "Hostname:", hostname)
	_, _ = fmt.Fprintln(w, "RemoteAddr:", req.RemoteAddr)
	for i, header := range req.Header {
		_, _ = fmt.Fprintf(w, "%s: %+v\n", i, header)
	}
	log.Printf("Query : URL: '%s', Header: '%+v'", req.URL.String(), req.Header)
}

func custom(w http.ResponseWriter, req *http.Request) {
	u, _ := url.Parse(req.URL.String())
	queryParams := u.Query()
	for i := 0; i < len(queryParams["hk"]); i++ {
		w.Header().Set(queryParams["hk"][i], queryParams["hv"][i])
		log.Println("Header", queryParams["hk"][i], queryParams["hv"][i])
	}
	defaultHandler(w, req)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/custom", custom)
	fmt.Println("Starting up on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
