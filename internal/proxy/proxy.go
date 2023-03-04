package proxy

import (
	"github.com/Mistsink/kuwo-api/internal/model"
	"net/url"

	"github.com/Mistsink/kuwo-api/internal/proxy/copy"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func InitProxyService() {
	initProxyClient()
}

type PlatformTag string

const (
	Tkuwo    PlatformTag = "kuwo"
	Tkugou   PlatformTag = "kugou"
	Tqq      PlatformTag = "qq"
	Tnetease PlatformTag = "netease"
)

type ProxyReqOpt struct {
	tag    PlatformTag
	method string
	rawUrl string
	url    *url.URL //	optional	初始化时最好不用这个
}

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

// TODO: prepare request -> send -> get response -> unmarshal response -> copy to std_response
func sendAndFormat(opt *ProxyReqOpt, raw, dst any) (err error) {
	if u, err := prepareReq(opt.rawUrl); err != nil {
		return err
	} else {
		opt.url = u
	}

	if err = sendProxy(opt, raw); err != nil {
		return err
	}

	if formable, ok := raw.(model.Formable); ok {
		raw, err = formable.Format()
		if err != nil {
			return err
		}
	}

	return copy.CopyByTag(dst, raw, "yzh")
}

func prepareReq(newUrl string) (*url.URL, error) {
	u, err := url.Parse(newUrl)
	if err != nil {
		return u, err
	}

	query := u.Query()
	query.Add("reqId", uuid.NewV4().String())

	u.RawQuery = query.Encode()
	return u, nil
}
