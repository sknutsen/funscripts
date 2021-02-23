package main

import (
	"log"
	"net/http"
	"os"

	"./home"
	"./server"
)

func main() {
	logger := log.New(os.Stdout, "srv", log.LstdFlags|log.Lshortfile)

	h := home.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux)

	logger.Println("Server starting...")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
