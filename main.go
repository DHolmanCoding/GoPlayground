package main

import (
	"fmt"
	"net/http"
)

// curl 127.0.0.1:6666/basic-endpoint
func basicEndpoint(writer http.ResponseWriter, request* http.Request) {
	fmt.Fprintln(writer, "This is the response test")
}


// curl -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:6666/header-endpoint
func headerEndpoint(w http.ResponseWriter, req* http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "Header -- %v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/basic-endpoint", basicEndpoint)
	http.HandleFunc("/header-endpoint", headerEndpoint)
	http.ListenAndServe(":6666", nil)
}
