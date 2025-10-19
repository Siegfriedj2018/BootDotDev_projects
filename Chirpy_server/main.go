package main

import (
	"log"
	"net/http"
)


func healthCheck(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte{'O', 'K'})
	if err != nil {
		log.Printf("error writing to response header: %v", err)
	}
}


func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app" ,http.FileServer(http.Dir("."))))
	mux.Handle("/assets", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/healthz", healthCheck)
	serv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}


	err := serv.ListenAndServe();
	if err != nil {
		log.Fatalf("error serving server: %v", err)
	}
}