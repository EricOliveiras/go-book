package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/greet", greet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, boa noite %s", time.Now().Local())
}
