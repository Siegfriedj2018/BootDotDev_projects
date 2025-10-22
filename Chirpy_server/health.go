package main

import (
	"net/http"
	"log"
)

func healthCheck(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte{'O', 'K', '\n'})
	if err != nil {
		log.Printf("error writing to response header: %v\n", err)
	}
}
