package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"

	"./files"
	"./handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
var port = ":9090"

//var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
//var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images-api",
			Level: hclog.LevelFromString("debug"),
		},
	)

	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	stor, err := files.NewLocal("./imagestore", 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	fh := handlers.NewFiles(stor, l)
	mw := handlers.GzipHandler{}

	sm := mux.NewRouter()

	ph := sm.Get(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir("./imagestore"))))
	gh.Use(mw.GzipMiddleware)

	s := &http.Server{
		Addr:         port,
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		sl.Println("Starting server on port ", port)

		err := s.ListenAndServe()
		if err != nil {
			sl.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	sl.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
