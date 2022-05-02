package application

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/ferdn4ndo/userver-logger-api/services/environment"
)

type Application struct {
	Auth struct {
		Username string
		Password string
	}
	Port   int
	Routes chi.Router
	Server *http.Server
}

func (app *Application) Start() {
	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.getServerPort()),
		Handler:      app.Routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Opening server port %s", app.Server.Addr)
	listener, error := net.Listen("tcp", app.Server.Addr)
	if error != nil {
		log.Fatalf("Error occurred when opening port %d: %s", app.getServerPort(), error.Error())
	}

	log.Printf("Starting server...")
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

func (app *Application) getServerPort() int {
	port, err := strconv.Atoi(environment.GetEnvKey("SERVER_PORT"))
	if err != nil {
		log.Fatal("Unable to determine application port!")
	}

	return port
}
