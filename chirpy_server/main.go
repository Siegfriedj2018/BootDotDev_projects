package main

import (
	"bootdev_projects/chirpy_server/internal/database"
	"bootdev_projects/chirpy_server/internal/auth"
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits 	atomic.Int32
	databaseQ 			*database.Queries
	platform				string
	jwtSecret				string
}

func main() {
	wrote, err := auth.CreateJWTSecret()
	if err != nil {
		log.Fatalf("could not create jwt secret: %v", err)
	}
	log.Printf("Wrote to .env: %v", wrote)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file: %v\n", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	dbQueries := database.New(db)

	platEnv := os.Getenv("PLATFORM")
	jwtSecret := os.Getenv("JWT_SECRET")
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		databaseQ: 			dbQueries,
		platform: 			platEnv,
		jwtSecret: 			jwtSecret,
	}

	const ip = "localhost"
	const port = "8080"
	
	mux := http.NewServeMux()

	handler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(handler))
	// mux.Handle("/assets", http.FileServer(http.Dir(".")))
	
	mux.HandleFunc("GET /api/healthz", healthCheck)
	mux.HandleFunc("GET /admin/metrics", apiCfg.hitMetrics)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerGetChirp)

	mux.HandleFunc("POST /admin/reset", apiCfg.resetMetrics)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlervalidateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)

	mux.HandleFunc("PUT /api/users", apiCfg.handlerUpdatePassword)
	
	serv := &http.Server{
		Addr: ip + ":" + port,
		Handler: mux,
	}

	log.Printf("Server starting on %s on port %s...", ip, port)
	err = serv.ListenAndServe();
	if err != nil {
		log.Fatalf("error serving server: %v", err)
	}
}