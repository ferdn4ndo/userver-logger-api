package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/handler"
)

// GET /health
func GetHealthState(writer http.ResponseWriter, request *http.Request) {
	db, err := database.GetDatabaseService()
	if err != nil {
		render.Render(writer, request, handler.ServerErrorRenderer(err))

		return
	}

	database.AddHeartbeatLog(db)
	fmt.Fprintf(writer, "It's just fine.")
}
