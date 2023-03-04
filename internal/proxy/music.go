package proxy

import (
	"fmt"

	kuwo "github.com/Mistsink/kuwo-api/internal/model/kuwo/music"
	netease "github.com/Mistsink/kuwo-api/internal/model/netease/music"
	std "github.com/Mistsink/kuwo-api/internal/model/standard/music"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

// /	get url by type
// /	type: [ music, mv ]
func urlByTypeKuwo(c *gin.Context, param *service.MusicUrlReq, typeStr string) string {
	newUrl := "http://www.kuwo.cn/api/v1/www/music/playUrl"
	newUrl = fmt.Sprintf("%s?mid=%d&type=%s&httpsStatus=1",
		newUrl, param.Id, typeStr)
	return newUrl
}

// MusicUrl /	get music url
func MusicUrl(c *gin.Context, param *service.MusicUrlReq) (*std.UrlResp, error) {
	var (
		err    error
		tag    = PlatformTag(param.Tag)
		url    string
		rawRet any
		ret    = &std.UrlResp{}
	)
	ret.Result.Tag = string(tag)

	// parse url & raw response with tag
	switch tag {
	case Tkuwo:
		url = urlByTypeKuwo(c, param, "music")
		rawRet = &kuwo.UrlResp{}
	case Tnetease:
		url = fmt.Sprintf("/song/url/v1?id=%d&level=exhigh", param.Id)
		rawRet = &netease.UrlResp{}
	}

	// send proxy request
	err = sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, rawRet, ret)

	return ret, err
}

// MusicMVUrl /	get mv url
func MusicMVUrl(c *gin.Context, param *service.MusicUrlReq) (*std.UrlResp, error) {
	url := urlByTypeKuwo(c, param, "mv")
	rawRet := &kuwo.UrlResp{}
	ret := &std.UrlResp{}

	err := sendAndFormat(&ProxyReqOpt{
		tag:    Tkuwo,
		rawUrl: url,
	}, rawRet, ret)

	return ret, err
}

// MusicInfo / get music info
func MusicInfo(c *gin.Context, param *service.MusicUrlReq) (*std.InfoResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.InfoResp{}
		tag = PlatformTag(param.Tag)
	)
	ret.Result.Tag = string(tag)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/music/musicInfo"
		url = fmt.Sprintf("%s?mid=%d&httpsStatus=1",
			url, param.Id)
		raw = &kuwo.InfoResp{}
	case Tnetease:
		url = fmt.Sprintf("/song/detail?ids=%d", param.Id)
		raw = &netease.InfoResp{}
	}

	err = sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	return ret, err
}

// MusicLyric / get music lyric
func MusicLyric(c *gin.Context, param *service.MusicUrlReq) (*std.LyricResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.LyricResp{}
		tag = PlatformTag(param.Tag)
	)
	ret.Result.Tag = string(tag)

	switch tag {
	case Tkuwo:
		url = "http://m.kuwo.cn/newh5/singles/songinfoandlrc"
		url = fmt.Sprintf("%s?musicId=%d&httpsStatus=1",
			url, param.Id)
		raw = &kuwo.LyricResp{}
	case Tnetease:
		url = fmt.Sprintf("/lyric?id=%d", param.Id)
		raw = &netease.LyricResp{}
	}

	err = sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	return ret, err
}
