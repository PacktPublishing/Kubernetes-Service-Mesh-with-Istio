package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	finish := make(chan bool)

	serverProbe := http.NewServeMux()
	serverProbe.HandleFunc("/", indexHandler)

	go func() {
		http.ListenAndServe(":8081", serverProbe)
	}()

	serverHTTP := http.NewServeMux()
	serverHTTP.HandleFunc("/", indexHandler)
	serverHTTP.HandleFunc("/api/message/index", indexMsgHandler)
	serverHTTP.HandleFunc("/api/message/hello", helloMsgHandler)

	go func() {
		fmt.Println("Listen and Serve Hello Message Golang v1")
		http.ListenAndServe(port(), serverHTTP)
	}()

	<-finish
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
	fmt.Println("Handler func /api/message/index called.")

	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Default Index Message from Golang (v1).")
}

func helloMsgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler func /api/message/hello called.")

	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Istio Message from Golang (v1).")
}
