package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("html"))
	http.Handle("/", fs)

	addr := ":" + port
	log.Printf("Serving ./html on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
