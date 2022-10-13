package main

import (
	"log"
	"net"
	"net/http"
)

// Create random port toi listen on
func createListener() (l net.Listener, close func()) {
    l, err := net.Listen("tcp", ":0")
    if err != nil {
        panic(err)
    }

    return l, func() {
        _ = l.Close()
    }
}

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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))

}

func main() {
	l, close := createListener()
    defer close()

	http.HandleFunc("/health", healthHandler)

	//start the server
	log.Println("listening at", l.Addr().(*net.TCPAddr).Port)
	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}

}
