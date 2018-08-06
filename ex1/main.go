package main

import (
	"fmt"
	"log"
	"net/http"
)

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("log=remote:%s,url:%s", r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", withLogging(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the app that does nothing!")
	}))

	http.HandleFunc("/health", withLogging(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	}))

	log.Fatal(http.ListenAndServe(":4444", nil))
}
