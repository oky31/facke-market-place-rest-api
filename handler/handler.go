package handler

import (
	"encoding/json"
	"fake-market/data"
	"net/http"
)

type Login struct{}

func (l Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var payload data.LoginPayload

	dec := json.NewDecoder(r.Body)
	dec.Decode(&payload)

	w.Write([]byte("This is login page"))
}
