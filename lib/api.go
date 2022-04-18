package lib

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	Updater   Updater
	LogOutput io.Writer
}

func (a *API) handleUpdate(w http.ResponseWriter, r *http.Request) {
	mac := r.Header.Get("x-esp8266-sta-mac")
	if mac == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("X-ESP8266-STA-MAC is not set"))
		return
	}

	currentType := r.URL.Query().Get("firmware")
	if currentType == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("current firmware type was not sent"))
		return
	}

	currentVersion := r.URL.Query().Get("version")
	if currentVersion == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("current firmware version was not sent"))
		return
	}

	firmware, err := a.Updater.Update(mac, &Firmware{
		Type:    currentType,
		Version: currentVersion,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("failed to get update: %s", err.Error())))
		return
	}

	if firmware == nil {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(firmware.Data)
}

func (a *API) GetMux() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/update", a.handleUpdate).Methods("GET")
	return handlers.LoggingHandler(a.LogOutput, r)
}
