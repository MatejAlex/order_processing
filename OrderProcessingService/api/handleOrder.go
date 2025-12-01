package api

import (
	"fmt"
	"net/http"
)

func HandleOrder(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("Received request")
	_, err := w.Write([]byte("Order stored for processing"))
	if err != nil {
		w.WriteHeader(500)
	}
	w.WriteHeader(http.StatusOK)
}
