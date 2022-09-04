package handler

import (
	"net/http"

	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/gorilla/mux"
)

type handler struct {
	logger     *logging.Logger
	urlService UrlService
	host       string
	port       string
}

func NewUrlHandler(logger *logging.Logger, urlService UrlService, host string, port string) *handler {
	return &handler{logger: logger, urlService: urlService, host: host, port: port}
}

func (h *handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/encode", h.EncodeUrl).Methods(http.MethodPost)
	router.HandleFunc("/{shortUrl}", h.DecodeUrl).Methods(http.MethodGet)
	router.HandleFunc("/encode/custom", h.EncodeCustomUrl).Methods(http.MethodPost)
}
