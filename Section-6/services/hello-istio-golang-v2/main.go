package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	serverHTTP.HandleFunc("/api/hello", helloHandler)
	serverHTTP.HandleFunc("/api/v2/hello", helloHandler)

	go func() {
		fmt.Println("Listen and Serve Hello Istio Golang v2")
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

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Istio Golang v2")
}

type hello struct {
	Message string
	Version string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler func /api/hello called.")

	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	s, code := getMessage("hello", r.Header)
	m := hello{s, "v2"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

var client = http.Client{Timeout: 3 * time.Second}

func getMessage(msg string, h http.Header) (string, int) {

	req, _ := http.NewRequest("GET", "http://hello-message:8080/api/message/"+msg, nil)

	// perform header propagation
	req.Header.Set("x-request-id", h.Get("x-request-id"))
	req.Header.Set("x-b3-traceid", h.Get("x-b3-traceid"))
	req.Header.Set("x-b3-spanid", h.Get("x-b3-spanid"))
	req.Header.Set("x-b3-parentspanid", h.Get("x-b3-parentspanid"))
	req.Header.Set("x-b3-sampled", h.Get("x-b3-sampled"))
	req.Header.Set("x-b3-flags", h.Get("x-b3-flags"))
	req.Header.Set("x-ot-span-context", h.Get("x-ot-span-context"))

	resp, err := client.Do(req)
	if err != nil {
		message := "Error getting message from backend."
		fmt.Println(message)
		return message, http.StatusServiceUnavailable
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		message := "Error reading message from backend."
		fmt.Println(message)
		return message, http.StatusBadGateway
	}

	return string(body), resp.StatusCode
}
