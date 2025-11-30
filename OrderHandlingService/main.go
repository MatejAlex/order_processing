package main

import (
	"handle_order/api"
	"log"
	"net/http"
)

func main() {

	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/order", api.HandleOrder)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Listening on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
