package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const lorem string = "\"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func closing(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// helloworld(w, r, nil)

	// w.WriteHeader(http.StatusAlreadyReported)

	// Check that the rw can be hijacked.
	hj, ok := w.(http.Hijacker)

	// The rw can't be hijacked, return early.
	if !ok {
		http.Error(w, "can't hijack rw", http.StatusInternalServerError)
		return
	}

	// Hijack the rw.
	conn, _, err := hj.Hijack()
	if err != nil {
		// handle error
	}

	// Close the hijacked raw tcp connection.
	if err := conn.Close(); err != nil {
		http.Error(w, "could not close rw", http.StatusInternalServerError)
	}
}

func helloworld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	for i := 0; i < 2<<1; i++ {
		fmt.Fprintf(w, "i: %d -> %q\n", i, lorem)
	}
}

func serve(url string, ka bool) {
	router := httprouter.New()
	router.GET("/", helloworld)
	router.GET("/closing", closing)
	server := &http.Server{Addr: url, Handler: router}
	server.SetKeepAlivesEnabled(ka)
	log.Printf("Listening on %q with SetKeepAlivesEnabled set to '%t'", url, ka)
	log.Fatal(server.ListenAndServe())
}

func main() {
	go serve(":3030", false)
	serve(":3031", true)
}
