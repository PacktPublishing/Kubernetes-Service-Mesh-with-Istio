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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/v1/hello", helloHandler)

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

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Istio Golang v1")
}

type hello struct {
	Message string
	Version string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	sleeps, ok := r.URL.Query()["sleep"]
	if ok {
		i, err := strconv.Atoi(sleeps[0])
		if err == nil {
			time.Sleep(time.Duration(i) * time.Second)
		}
	}

	m := hello{getMessage("hello"), "v1"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func getMessage(msg string) string {
	resp, err := http.Get("http://hello-message:8080/api/message/" + msg)
	if err != nil {
		return "Error getting message from backend."
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading message from backend."
	}

	return string(body)
}
