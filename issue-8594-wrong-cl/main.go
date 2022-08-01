package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println(r)
		io.WriteString(rw, "oops something went wrong")
	})

	http.ListenAndServe(":8090", nil)
}
