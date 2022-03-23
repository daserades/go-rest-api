package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JimySheepman/go-rest-api/config/db"
	"github.com/JimySheepman/go-rest-api/config/env"
	"github.com/JimySheepman/go-rest-api/internal/handler"
	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	_, err := env.LoadEnvironmentConfigure("../../.env")
	if err != nil {
		log.Fatal("Loading .env file failed")
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		fmt.Println("PATH: /")

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).GetHandler()

	router.HandleFunc("/getData", handler.GetDataHandler(database))

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		fmt.Println("PATH: /api/health")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}