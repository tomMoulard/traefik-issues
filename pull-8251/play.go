package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("/api/%08bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed URL: %v\n", u)
	fmt.Printf("Parsed URL.Path: %v\n", u.Path)
	fmt.Printf("Parsed URL.RawPath: %v\n", u.RawPath)
	fmt.Printf("Parsed URL.String: %v\n", u.String())
}
