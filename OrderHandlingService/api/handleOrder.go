package api

import (
	"net/http"
)

func HandleOrder(w http.ResponseWriter, req *http.Request) {

	_, err := w.Write([]byte("Order stored for processing"))
	if err != nil {
		w.WriteHeader(500)
	}
	w.WriteHeader(http.StatusOK)
}
