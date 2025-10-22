package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	ip := "localhost"
	port := "8080"
	apiCfg := apiConfig{}
	mux := http.NewServeMux()

	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	mux.Handle("/assets", http.FileServer(http.Dir(".")))
	
	mux.HandleFunc("GET /api/healthz", healthCheck)
	mux.HandleFunc("GET /admin/metrics", apiCfg.hitMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handlervalidateChirp)
	
	serv := &http.Server{
		Addr: ip + ":" + port,
		Handler: mux,
	}

	log.Printf("Server starting on %s on port %s...", ip, port)
	err := serv.ListenAndServe();
	if err != nil {
		log.Fatalf("error serving server: %v", err)
	}
}