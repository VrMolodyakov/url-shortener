package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/VrMolodyakov/url-shortener/internal/errs"
	"github.com/gorilla/mux"
)

const (
	prefix string = ""
	indent string = "   "
)

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
	shortUrl = fmt.Sprintf("http://%v:%v/%v", h.host, h.port, shortUrl)
	response := UrlResponse{ShortUrl: shortUrl, FullUrl: request.Url}
	jsonReponce, err := json.MarshalIndent(response, prefix, indent)
	if err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
	err = h.urlService.CreateCustomUrl(request.Custom, request.Url)
	if err != nil {
		errorResponse(w, err)
		return
	}
	shortUrl := fmt.Sprintf("http://%v:%v/%v", h.host, h.port, request.Custom)
	response := UrlResponse{ShortUrl: shortUrl, FullUrl: request.Url}
	jsonReponce, err := json.MarshalIndent(response, prefix, indent)
	if err != nil {
		errorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
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

func errorResponse(w http.ResponseWriter, err error) {
	if errors.Is(err, errs.ErrUrlIsEmpty) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
