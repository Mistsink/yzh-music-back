package main

import (
	"encoding/json"
	"github.com/Mistsink/kuwo-api/internal/router"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRPing(t *testing.T) {
	r := router.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r.ServeHTTP(w, req)

	resp := new(app.RespJSON)
	_ = json.Unmarshal(w.Body.Bytes(), resp)

	if resp.Code != 0 {
		t.Error("response.code is not 0. url:", req.URL)
	}
	if w.Code != 200 {
		t.Error("response.statusCode is:", w.Code)
	}
}
