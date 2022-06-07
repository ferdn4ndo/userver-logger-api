package health

import (
	"net/http"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
)

const HEALTH_STATUS_OK string = "OK"

type HealthData struct {
	Status          string `json:"status"`
	DbSize          int64  `json:"databaseSizeInBytes"`
	LogEntriesCount int64  `json:"logEntriesCount"`
}

func (healthData *HealthData) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type HealthService struct {
	DbService database.DatabaseServiceInterface
}

func (service HealthService) GetHealthData() *HealthData {
	service.DbService.AddHeartbeatLog()

	data := &HealthData{
		Status:          HEALTH_STATUS_OK,
		DbSize:          service.DbService.GetDatabaseFileSize(),
		LogEntriesCount: service.DbService.GetLogEntriesTotalCount(),
	}

	return data
}
