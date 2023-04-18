package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	fmt.Println("Server listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(res, "Hello from Load Balancer")
	url := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(res, req)
}

func getServer() *url.URL {
	currentIdx := (lastServerIndex + 1) % 5
	urlToForward := serverList[currentIdx]
	fmt.Println("Current url:", urlToForward)
	url, _ := url.Parse(urlToForward)
	lastServerIndex = currentIdx
	return url
}

var (
	serverList = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
		"http://127.0.0.1:5003",
		"http://127.0.0.1:5004",
	}

	lastServerIndex = 4
)
