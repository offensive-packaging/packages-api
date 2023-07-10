package main

import (
	"log"
	"net/http"
	h "packages-api/handlers"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "https://api.parrotsec.org")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		value := w.Header().Get("Content-Type")
		if value != "" && value != "application/json" {
			http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 512000)

		next.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/packages", h.GetPackage)

	err := http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Endpoint /packages set")
	log.Println("Try curl --request GET 'http://localhost:8080/packages?branch=main&arch=amd64&package=nginx'")
}
