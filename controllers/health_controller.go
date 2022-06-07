package controllers

import (
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/health"
	"github.com/ferdn4ndo/userver-logger-api/services/render"
)

const STATUS_OK string = "OK"

type HealthController struct {
	DbService     database.DatabaseServiceInterface
	RenderService render.RenderServiceInterface
}

type HealthData struct {
	Status          string `json:"status"`
	DbSize          int64  `json:"databaseSizeInBytes"`
	LogEntriesCount int64  `json:"logEntriesCount"`
}

func (healthData *HealthData) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (controller HealthController) GetHealthState(writer http.ResponseWriter, request *http.Request) {
	healthService := health.HealthService{DbService: controller.DbService}

	data := healthService.GetHealthData()

	controller.RenderService.Render(writer, request, data)
}
