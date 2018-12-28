package framework

import (
	"log"
	"net/http"
)

func mappingUrl() {
	// Register Handler
	http.HandleFunc("/", getOnly(handlerUniversal))
	http.HandleFunc("/reload", getOnly(reload))

	for k := range UrlLoader {
		log.Println("Mapping URL :", k)
		http.HandleFunc(k, postOnly(requestHandler))
	}

	log.Println("Application Running on ports : " + applicationProperties["application.port"])
	http.ListenAndServe(":" + applicationProperties["application.port"], nil)
}