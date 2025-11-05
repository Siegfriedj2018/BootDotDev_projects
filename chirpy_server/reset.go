package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) resetMetrics(res http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(res, http.StatusForbidden, "You do not have access to this endpoint", fmt.Errorf("forbidden access"))
		return
	}
	
	err := cfg.databaseQ.DeleteUser(req.Context())
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Error deleting users", err)
		return
	}
	
	res.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)
}
