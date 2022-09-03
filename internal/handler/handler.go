package handler

import (
	"net/http"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/gorilla/mux"
)

type handler struct {
	logger     *logging.Logger
	urlService UrlService
}

func NewUrlHandler(logger *logging.Logger, urlService UrlService) *handler {
	return &handler{logger: logger, urlService: urlService}
}

func (h *handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/encode", h.EncodeUrl).Methods(http.MethodPost)
	router.HandleFunc("/{shortUrl}", h.DecodeUrl).Methods(http.MethodGet)
	router.HandleFunc("/encode/custom", h.EncodeCustomUrl).Methods(http.MethodPost)
}
