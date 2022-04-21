package docs

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"

	"github.com/ferdn4ndo/userver-logger-api/services/file"
)

var DOC_JSON_FILE_PATH = "/data/routes.json"

func getApiDocsJsonFile() string {
	dataFolder := file.GetDataFolder()

	return fmt.Sprintf("%s/routes.json", dataFolder)
}

func ExportApiDocumentation(router chi.Router) {
	documentation := []byte(docgen.JSONRoutesDoc(router))
	outputJsonFile := getApiDocsJsonFile()

	error := ioutil.WriteFile(outputJsonFile, documentation, 0644)
	if error != nil {
		log.Fatal(error)
	}

	log.Printf("Exported API documentation to %s\n", outputJsonFile)
}
