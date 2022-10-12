package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	// test to see that we got the right path
	if r.URL.Path != "/health" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// only allow GET requests
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "")
}

func main() {

	http.HandleFunc("/health", healthHandler)

	//start the server
	fmt.Println("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
