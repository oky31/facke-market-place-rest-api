package handler

import (
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

type CreateUserHandler struct{}

func (cuh *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") == "" {
		message := "Content-Type header required"
		http.Error(w, message, http.StatusUnsupportedMediaType)
		return
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			message := "Content-Type header is not application/json"
			http.Error(w, message, http.StatusUnsupportedMediaType)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
