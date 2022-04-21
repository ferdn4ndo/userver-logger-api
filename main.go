package main

import (
	"fmt"
	"log"

	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/application"
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/router"
)

func main() {
	log.Println("Initializing uServer Logger API database...")
	_, err := database.InitializeDatabase()
	if err != nil {
		panic(fmt.Sprintf("Error initializing database: %s", err))
	}

	app := application.GetBaseApplication()
	app.Routes = router.CreateRouter()
	app.Start()
}

// This is entirely optional, but I wanted to demonstrate how you could easily
// add your own logic to the render.Respond method.
func init() {
	render.Respond = handler.AddCustomErrorHandlerfunc
}
