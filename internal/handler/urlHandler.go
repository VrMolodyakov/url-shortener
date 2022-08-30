package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/internal/errs"
	"github.com/VrMolodyakov/url-shortener/pkg/logging"
	"github.com/gorilla/mux"
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

func (h *handler) EncodeUrl(w http.ResponseWriter, r *http.Request) {
	var request UrlRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Debugf("try to encode url = %v", request)
	shortUrl, err := h.urlService.CreateUrl(request.Url)
	if err != nil {
		errorResponse(w, err)
		return
	}
	response := UrlResponse{ShortUrl: shortUrl, FullUrl: request.Url}
	jsonReponce, err := json.MarshalIndent(response, prefix, indent)
	if err != nil {
		errorResponse(w, err)
		return
	}
	h.logger.Debug(jsonReponce)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReponce)

}

func (h *handler) EncodeCustomUrl(w http.ResponseWriter, r *http.Request) {
	var request CustomUrlRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.Debugf("try to encode custom = %v,full = %v", request.Custom, request.Url)
	shortUrl, err := h.urlService.CreateCustomUrl(request.Custom, request.Url)
	if err != nil {
		errorResponse(w, err)
		return
	}
	response := UrlResponse{ShortUrl: shortUrl, FullUrl: request.Url}
	jsonReponce, err := json.MarshalIndent(response, prefix, indent)
	if err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonReponce)

}

func (h *handler) DecodeUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := mux.Vars(r)["shortUrl"]
	h.logger.Debugf("try to encode url = %v", shortUrl)

	fullUrl, err := h.urlService.GetUrl(shortUrl)
	if err != nil {
		errorResponse(w, err)
		return
	}
	http.Redirect(w, r, fullUrl, http.StatusSeeOther)
}

// func (h *handler) DecodeCustomUrl(w http.ResponseWriter, r *http.Request) {
// 	shortUrl := mux.Vars(r)["shortUrl"]
// 	h.logger.Debugf("try to encode url = %v", shortUrl)

// 	fullUrl, err := h.urlService.GetByCustomUrl(shortUrl)
// 	if err != nil {
// 		errorResponse(w, err)
// 		return
// 	}
// 	http.Redirect(w, r, fullUrl, http.StatusSeeOther)
// }

func errorResponse(w http.ResponseWriter, err error) {
	if errors.Is(err, errs.ErrUrlNotSaved) || errors.Is(err, errs.ErrUrlNotFound) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
