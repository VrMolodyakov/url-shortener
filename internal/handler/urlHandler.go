package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/internal/errs"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
)

const (
	prefix string = ""
	indent string = "   "
)

type handler struct {
	logger     *logging.Logger
	urlService UrlService
}

func NewUrlHandler(logger *logging.Logger, urlService UrlService) *handler {
	return &handler{logger: logger, urlService: urlService}
}

func (h *handler) Encode(w http.ResponseWriter, r *http.Request) {
	var url UrlRequest
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Debugf("try to encode url = %v", url)
	shortUrl, err := h.urlService.CreateShortUrl(url.FullUrl)
	if err != nil {
		errorResponse(w, err)
		return
	}
	response := UrlResponse{ShortUrl: shortUrl, FullUrl: url.FullUrl}
	jsonReponce, err := json.MarshalIndent(response, prefix, indent)
	if err != nil {
		errorResponse(w, err)
		return
	}
	h.logger.Debug(jsonReponce)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReponce)

}

func errorResponse(w http.ResponseWriter, err error) {
	if errors.Is(err, errs.ErrUrlNotSaved) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
