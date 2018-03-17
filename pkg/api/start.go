package api

import (
	"io"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "\"pong\"")
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
