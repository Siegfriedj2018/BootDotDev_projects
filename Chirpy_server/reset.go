package main

import (
	"net/http"
)

func (cfg *apiConfig) resetMetrics(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)
}
