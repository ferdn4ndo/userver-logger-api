package application

import (
	"log"

	"github.com/ferdn4ndo/userver-logger-api/services/database"
	"github.com/ferdn4ndo/userver-logger-api/services/log_entry"
	"github.com/ferdn4ndo/userver-logger-api/services/log_file"
	"github.com/ferdn4ndo/userver-logger-api/services/router"
)

type ApplicationStartService struct {
	GenerateDocs bool
	DryRun       bool
}

func (service ApplicationStartService) StartApplication() {
	log.Println("Initializing uServer Logger")

	dbService := service.StartDatabase()

	logEntryDbService := log_entry.LogEntryDatabaseService{
		DbService: dbService,
	}

	logFileScannerService := log_file.LogFileScannerService{
		LogEntryDbService: logEntryDbService,
	}

	go logFileScannerService.ForeverScanLogFiles()

	app := Application{}
	app.Routes = router.CreateRouter(dbService, service.GenerateDocs)
	app.Start()
}

func (service ApplicationStartService) StartDatabase() database.DatabaseServiceInterface {
	dbService := database.InitializeDatabase()

	if service.DryRun {
		dbService = database.MockedDatabaseService{}
	}

	return dbService
}
