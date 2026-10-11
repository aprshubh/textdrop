package server

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TextDrop server is healthy"))
	})
	return mux
}