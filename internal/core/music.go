package core

import (
	"fmt"

	kuwo_music "github.com/Mistsink/kuwo-api/internal/core/kuwo/music"
	std_music "github.com/Mistsink/kuwo-api/internal/core/standard/music"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

///	get url by type
///	type: [ music, mv ]
func urlByType(c *gin.Context, param *service.MusicUrlReq, typeStr string) *std_music.UrlResp {
	newUrl := "http://www.kuwo.cn/api/v1/www/music/playUrl"
	newUrl = fmt.Sprintf("%s?mid=%d&type=%s&httpsStatus=1",
		newUrl, param.Id, typeStr)

	rawRet := &kuwo_music.UrlResp{}
	ret := &std_music.UrlResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

///	get music url
func MusicUrl(c *gin.Context, param *service.MusicUrlReq) *std_music.UrlResp {
	return urlByType(c, param, "music")
}

///	get mv url
func MusicMVUrl(c *gin.Context, param *service.MusicUrlReq) *std_music.UrlResp {
	return urlByType(c, param, "mv")
}

/// get music info
func MusicInfo(c *gin.Context, param *service.MusicUrlReq) *std_music.InfoResp {
	newUrl := "http://www.kuwo.cn/api/www/music/musicInfo"
	newUrl = fmt.Sprintf("%s?mid=%d&httpsStatus=1",
		newUrl, param.Id)

	rawRet := &kuwo_music.InfoResp{}
	ret := &std_music.InfoResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

/// get music lyric
func MusicLyric(c *gin.Context, param *service.MusicUrlReq) *std_music.LyricResp {

	newUrl := "http://m.kuwo.cn/newh5/singles/songinfoandlrc"
	newUrl = fmt.Sprintf("%s?musicId=%d&httpsStatus=1",
		newUrl, param.Id)

	rawRet := &kuwo_music.LyricResp{}
	ret := &std_music.LyricResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}
