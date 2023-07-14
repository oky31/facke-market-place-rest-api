package handler

import (
	"bytes"
	"encoding/json"
	"fake-market/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCanLogin(t *testing.T) {
	login := Login{}

	payload := data.LoginPayload{
		Username: "oky31",
		Password: "12345",
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("fail convert payload to json %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(payloadJson))
	res := httptest.NewRecorder()

	login.ServeHTTP(res, req)

}
