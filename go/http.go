package main

import (
	"net/http"
	"fmt"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shoutHelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shoutHelloHandler)
	http.ListenAndServe("localhost:9999", nil)
}