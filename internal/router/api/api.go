package api

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func execBindAndValid(c *gin.Context, response *app.Response, param interface{}) {
	if err := c.ShouldBind(param); err != nil {
		response.ToErrResponse(errcode.ServerWithMsg(fmt.Sprintf("c.ShouldBind(param) err: %v", err)))
		log.Println("c.ShouldBind(param) err:", err)
		panic(fmt.Sprintf("c.ShouldBind(param) err: %v", err))
	}
}

func parseUrl(_url string, response *app.Response) *url.URL {
	u, err := url.Parse(_url)
	if err != nil {
		response.ToErrResponse(errcode.ServerWithMsg(fmt.Sprintf("url.Parse err: %v", err)))
		log.Println("url.Parse err:", err)
		panic(fmt.Sprintf("c.ShouldBind(param) err: %v", err))
	}

	return u
}

func RewriteUrl(c *gin.Context, u *url.URL) {
	c.Request.Host = u.Host
	c.Request.URL = u
}
