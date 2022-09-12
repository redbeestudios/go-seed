package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func StartServer(config *Config, r *mux.Router) {
	srv := &http.Server{
		Addr:         config.Server.Address,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second * 15,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second * 15,
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for terminate signal to shut down server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
