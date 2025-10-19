package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func healthCheck(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte{'O', 'K'})
	if err != nil {
		log.Printf("error writing to response header: %v\n", err)
	}
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	log.Println("incrementing...")
	cfg.fileserverHits.Add(1)
	log.Println("hits: ", cfg.fileserverHits.Load())
	return next
}

func (cfg *apiConfig) hitMetrics(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	hits := cfg.fileserverHits.Load()
	respText := fmt.Sprintf("Hits: %d", hits)

	_, err := res.Write([]byte(respText))
	if err != nil {
		log.Printf("error writing fileserverhits: %v\n", err)
	}
}

func (cfg *apiConfig) resetMetrics(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)
}
 
func main() {
	apiCfg := apiConfig{}
	mux := http.NewServeMux()
	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.Handle("/assets", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/healthz", healthCheck)
	mux.HandleFunc("/metrics", apiCfg.hitMetrics)
	mux.HandleFunc("/reset", apiCfg.resetMetrics)
	serv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Server starting...")
	err := serv.ListenAndServe();
	if err != nil {
		log.Fatalf("error serving server: %v", err)
	}
}