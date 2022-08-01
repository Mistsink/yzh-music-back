package api

import (
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/gin-gonic/gin"
)

//	轮播图
func Carousel(c *gin.Context) {
	response := app.NewResponse(c)

	u := parseUrl("http://www.kuwo.cn/api/www/banner/index/bannerList",
		response)

	RewriteUrl(c, u)
}
