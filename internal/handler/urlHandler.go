package handler

import (
	"encoding/json"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
)

type handler struct {
	logger *logging.Logger
}

func NewUrlHandler(logger *logging.Logger) *handler {
	return &handler{logger: logger}
}

func (h *handler) Encode(w http.ResponseWriter, r *http.Request) {
	var url UrlRequest
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Debug(url)
}
