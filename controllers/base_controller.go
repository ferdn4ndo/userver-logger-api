package controllers

import (
	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

type BaseController struct {
	DbService     database.DatabaseService
	RenderService render.RenderService
}
