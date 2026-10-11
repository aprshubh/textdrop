package main

import (
	"log"
	"net/http"

	appserver "github.com/aprshubh/textdrop/server/internal/server"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TextDrop server is healthy"))
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}