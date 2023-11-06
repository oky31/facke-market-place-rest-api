package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/exp/slog"

	"github.com/golang/gddo/httputil/header"
)

const size1MB = 1048576

const ContentTypeJson = "application/json"

type ResBody struct {
	ErrorRes interface{} `json:"errors,omitempty"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

func NewCreateUserHandler(logger *slog.Logger, db *sql.DB) http.Handler {
	return &CreateUserHandler{
		db:     db,
		logger: logger,
	}
}

type CreateUserHandler struct {
	db     *sql.DB
	logger *slog.Logger
}

func (cuh *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cuh.logger.Info("request", "method", r.Method, "path", r.URL.Path)

	if r.Header.Get("Content-Type") == "" {
		message := "Content-Type header required"
		http.Error(w, message, http.StatusUnsupportedMediaType)
		cuh.logger.Error(message)
		return
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			message := "Content-Type header is not application/json"
			http.Error(w, message, http.StatusUnsupportedMediaType)
			cuh.logger.Error(message)
			return
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, size1MB)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var payload CreateUserPayload
	err := dec.Decode(&payload)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {

		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			http.Error(w, msg, http.StatusBadRequest)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		default:
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	ctx := r.Context()
	err = CreateNewUser(ctx, cuh.db, payload)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
