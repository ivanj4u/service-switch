package framework

import (
	"net/http"
)

func getOnly(handler CustomHandler) CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler(w, r)
			return
		}
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
	}
}

func postOnly(handler CustomHandler) CustomHandler  {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler(w, r)
			return
		}
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
	}
}