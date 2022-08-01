package service

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
}

func NewService(c *gin.Context) *Service {
	return &Service{c}
}

func (svc *Service) RedirectUrl(host, u string) error {
	var err error
	svc.ctx.Request.Host = host
	svc.ctx.Request.URL, err = url.Parse(u)
	return err
}

func RewriteUrl(req *http.Request, u *url.URL) {
	req.Host = u.Host
	req.URL = u
}



