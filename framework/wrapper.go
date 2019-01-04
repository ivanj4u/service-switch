package framework

import (
	"net/http"
	"strings"
)

func getOnly(handler CustomHandler) CustomHandler {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler(w, r)
			return
		}
		http.Error(w, "GET only", http.StatusMethodNotAllowed)
	}
}

func postOnly(handler CustomHandler) CustomHandler  {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]

			err := checkToken(reqToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			handler(w, r)
			return
		}
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
	}
}