package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cfg "github.com/aswinkarthik93/ingress-consul-register/pkg/config"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "\"pong\"")
}

func config(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(cfg.Cfg)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/config", config)
	port := fmt.Sprintf(":%d", cfg.ApiPort())
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
