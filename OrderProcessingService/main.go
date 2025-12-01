package main

import (
	"handle_order/api"
	"handle_order/cfg"
	"log"
	"net/http"
)

func main() {

	config := cfg.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/order", api.HandleOrder)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", config.Port)
	log.Printf("Auth service URL: %s", config.AuthURL)
	log.Printf("Audit service URL: %s", config.AuditURL)
	log.Fatal(srv.ListenAndServe())
}
