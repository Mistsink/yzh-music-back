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

// /	get processed resp from new url
// /	rawRet: rawResp	pointer
// /	ret: prcessedResp pointer
func getRespByNewUrl(c *gin.Context, newUrl string, rawRet any, ret any) (err error) {
	url, err := parseUrl(newUrl, c)
	if err != nil {
		return
	}
	rewriteUrl(c, url)

	rawResp, err := getRaw(c)
	if err != nil {
		return
	}
	defer rawResp.Body.Close()

	if err = unmarshal(c, rawResp, rawRet); err != nil {
		return err
	}

	copy.CopyByTag(ret, rawRet, "yzh")

	return
}
