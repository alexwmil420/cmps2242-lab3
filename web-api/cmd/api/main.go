package main

import (
	"net/http"
	"time"
)

//home route
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API\n"))
}

//Health check route
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running\n"))
}

//About route
func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alex Milian\n"))
}

//time route
func currentTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC1123)
	w.Write([]byte(now + "\n"))
}

//Custom route
func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to the Shapes API\n"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/time", currentTime)
	mux.HandleFunc("/greeting", greeting)

	http.ListenAndServe(":4000", mux)
}
