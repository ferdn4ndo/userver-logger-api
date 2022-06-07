package render

import (
	"encoding/json"
	"fmt"
	"net/http"

	chiRender "github.com/go-chi/render"

	"github.com/ferdn4ndo/userver-logger-api/services/handler"
	"github.com/ferdn4ndo/userver-logger-api/services/logging"
)

type RenderInterface interface {
	Render(writer http.ResponseWriter, request *http.Request) error
}

type RenderServiceInterface interface {
	Render(writer http.ResponseWriter, request *http.Request, data RenderInterface)
}

type RenderService struct{}

func (service RenderService) Render(writer http.ResponseWriter, request *http.Request, data RenderInterface) {
	error := chiRender.Render(writer, request, data)
	if error != nil {
		renderError := chiRender.Render(writer, request, handler.ServerErrorRenderer(error))

		if renderError != nil {
			logging.Error(fmt.Sprintf("Error when rendering: %s", renderError))
		}

		return
	}
}

type MockedRenderService struct{}

func (service MockedRenderService) Render(writer http.ResponseWriter, request *http.Request, data RenderInterface) {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(bytes))
	if err := data.Render(writer, request); err != nil {
		logging.Errorf("Error rendering: %s", err)
	}
}
