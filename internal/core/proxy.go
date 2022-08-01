package core

import (
	"github.com/Mistsink/kuwo-api/internal/core/copy"
	"github.com/gin-gonic/gin"
)

// 	fill headers:
//		cookie
//		reqId
//		csrf
//		host
//		referer
//		content-encoding
//		httpStatus

///	get processed resp from new url
///	rawRet: rawResp	pointer
///	ret: prcessedResp pointer
func getRespByNewUrl(c *gin.Context, newUrl string, rawRet any, ret any) {
	url := parseUrl(newUrl, c)
	RewriteUrl(c, url)

	rawResp := GetRaw(c)
	defer rawResp.Body.Close()

	unmarshal(c, rawResp, rawRet)
	copy.CopyByTag(ret, rawRet, "yzh")
}
