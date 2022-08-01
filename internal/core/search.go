package core

import (
	"fmt"

	kuwo_search "github.com/Mistsink/kuwo-api/internal/core/kuwo/search"
	std_search "github.com/Mistsink/kuwo-api/internal/core/standard/search"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func SearMusic(c *gin.Context, param *service.SearComReq) *std_search.MusicResp {
	newUrl := "http://www.kuwo.cn/api/www/search/searchMusicBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&key=%s&httpsStatus=1",
		newUrl, param.P, param.N, param.Key)

	rawRet := &kuwo_search.MusicResp{}
	ret := &std_search.MusicResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

func SearHint(c *gin.Context, param *service.SearHintReq) *std_search.HintResp {
	newUrl := "http://www.kuwo.cn/api/www/search/searchKey"
	newUrl = fmt.Sprintf(
		"%s?key=%s&httpsStatus=1",
		newUrl, param.Key)

	rawRet := &kuwo_search.HintResp{}
	ret := &std_search.HintResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

func SearAlbum(c *gin.Context, param *service.SearComReq) *std_search.AlbumResp {

	newUrl := "http://www.kuwo.cn/api/www/search/searchAlbumBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo_search.AlbumResp{}
	ret := &std_search.AlbumResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

func SearMV(c *gin.Context, param *service.SearComReq) *std_search.MVResp {

	newUrl := "http://www.kuwo.cn/api/www/search/searchMvBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo_search.MVResp{}
	ret := &std_search.MVResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

func SearArtist(c *gin.Context, param *service.SearComReq) *std_search.ArtistResp {
	newUrl := "http://www.kuwo.cn/api/www/search/searchArtistBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo_search.ArtistResp{}
	ret := &std_search.ArtistResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}

func SearPlaylist(c *gin.Context, param *service.SearComReq) *std_search.PlaylistResp {
	newUrl := "http://www.kuwo.cn/api/www/search/searchPlayListBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo_search.PlaylistResp{}
	ret := &std_search.PlaylistResp{}

	getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret
}
