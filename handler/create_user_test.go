package handler

import (
	"bytes"
	"encoding/json"
	"fake-market/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	payload := data.CreateUserPayload{
		FirstName: "Oky",
		LastName:  "Saputra",
		Email:     "oky@gmail.com",
		Address:   "Jl. Perjalanan panjang",
		UserName:  "oky55",
		Password:  "12345",
	}

	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("Error in marshal json %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/user/create", bytes.NewReader(bytesPayload))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	createUserHandler := CreateUserHandler{}
	createUserHandler.ServeHTTP(res, req)

	assertHttpStatus(t, res.Code, http.StatusCreated)
}
