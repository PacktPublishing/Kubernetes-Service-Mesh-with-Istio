package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
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

	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func indexMsgHandler(w http.ResponseWriter, r *http.Request) {
	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Kubernetes Service Meshes with Istio by Packt Publising")
}

func helloMsgHandler(w http.ResponseWriter, r *http.Request) {
	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Istio from Golang.")
}
