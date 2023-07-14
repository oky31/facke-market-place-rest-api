package main

import (
	"fake-market/handler"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	loginHandler := handler.Login{}

	mux.Handle("/login", loginHandler)

	http.ListenAndServe(":8080", mux)
}
