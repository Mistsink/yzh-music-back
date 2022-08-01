package test

import (
	"encoding/json"
	"github.com/Mistsink/kuwo-api/internal/router"
	"github.com/Mistsink/kuwo-api/pkg/app"
	_ "github.com/Mistsink/kuwo-api/setup"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	r   *gin.Engine
	req *http.Request
	w   *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	r = router.NewRouter()
	w = httptest.NewRecorder()
	m.Run()
}

func TestProxyPlaylist(t *testing.T) {
	req, _ = http.NewRequest("GET", "/playlist/", nil)
	VerifyReq(req, t)
}

func VerifyReq(req *http.Request, t *testing.T) {
	r.ServeHTTP(w, req)

	resp := new(app.RespJSON)
	_ = json.Unmarshal(w.Body.Bytes(), resp)

	if resp.Code != 0 {
		t.Error("response.code is not 0. url:", req.URL, "err:", resp.Msg)
	}
	if w.Code != 200 {
		t.Error("response.statusCode is:", w.Code)
	}
}
