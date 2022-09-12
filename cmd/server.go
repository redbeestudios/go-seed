package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(config *Config, router *mux.Router) {
	srv := &http.Server{
		Addr:         config.Server.Address,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for signal to shut down server
	shutDownChannel := make(chan os.Signal, 1)
	signal.Notify(shutDownChannel, os.Interrupt)
	<-shutDownChannel
}
