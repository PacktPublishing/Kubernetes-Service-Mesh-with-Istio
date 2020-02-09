package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/message/index", indexMsgHandler)
	http.HandleFunc("/api/message/hello", helloMsgHandler)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func indexMsgHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Kubernetes Service Meshes with Istio by Packt Publising")
}

func helloMsgHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Istio from Golang.")
}
