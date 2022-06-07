package application

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Auth struct {
		Username string
		Password string
	}
	Routes chi.Router
	Server *http.Server
}

func (app *Application) Start() {
	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", 5555),
		Handler:      app.Routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Opening server port %s", app.Server.Addr)
	listener, error := net.Listen("tcp", app.Server.Addr)
	if error != nil {
		log.Fatalf("Error occurred when opening '%s': %s", app.Server.Addr, error.Error())
	}

	log.Printf("Listening connections...")
	go func() {
		app.Server.Serve(listener)
	}()
	defer app.Stop()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-channel))
	log.Println("Stopping API server.")
}

func (app *Application) Stop() {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if error := app.Server.Shutdown(context); error != nil {
		log.Printf("Could not shut down server correctly: %v\n", error)
		os.Exit(1)
	}
}
