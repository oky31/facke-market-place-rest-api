package main

import (
	"database/sql"
	"fake-market/users"
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	logJsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJsonHandler)
	mux := http.NewServeMux()

	// setup sqlite connection
	db, err := sql.Open("sqlite3", "file:fake-marketplace.db")
	if err != nil {
		logger.Error("error open db connection", err)
	}

	createUserHandler := users.NewCreateUserHandler(logger, db)
	mux.Handle("/v1/users/create", createUserHandler)

	logger.Info("Listen in port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Error("error In start http ", err)
	}
}
