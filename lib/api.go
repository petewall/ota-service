package lib

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	DeviceService   DeviceService
	FirmwareService FirmwareService
	LogOutput       io.Writer
}

func (a *API) handleUpdate(w http.ResponseWriter, r *http.Request) {
	return
}

func (a *API) GetMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/update", a.handleUpdate).Methods("GET")
	return handlers.LoggingHandler(a.LogOutput, r)
}
