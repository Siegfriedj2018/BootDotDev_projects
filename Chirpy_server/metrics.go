package main

import (
	"log"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cfg.fileserverHits.Add(1)
		log.Println("hits: ", cfg.fileserverHits.Load())
		next.ServeHTTP(res, req)
	})
}

func (cfg *apiConfig) hitMetrics(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html;")
	res.WriteHeader(http.StatusOK)
	hits := cfg.fileserverHits.Load()
	respText := fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", hits)

	_, err := res.Write([]byte(respText))
	if err != nil {
		log.Printf("error writing fileserverhits: %v\n", err)
	}
}
