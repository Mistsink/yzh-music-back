package proxy

import (
	"fmt"
	"github.com/Mistsink/kuwo-api/global"
	kuwo "github.com/Mistsink/kuwo-api/internal/model/kuwo/search"
	netease "github.com/Mistsink/kuwo-api/internal/model/netease/search"
	std "github.com/Mistsink/kuwo-api/internal/model/standard/search"
	"github.com/pkg/errors"
	"sync"

	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func SearMusic(c *gin.Context, param *service.SearComReq) (*std.MusicResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.MusicResp{}

		wg    = &sync.WaitGroup{}
		retCh = make(chan *std.MusicResp)
	)

	for _, _tag := range global.ProxySetting.AvailableTags {

		tag := PlatformTag(_tag)
		switch tag {
		case Tkuwo:
			url = "http://www.kuwo.cn/api/www/search/searchMusicBykeyWord"
			url = fmt.Sprintf("%s?pn=%d&rn=%d&key=%s&httpsStatus=1",
				url, param.P, param.N, param.Key)
			raw = &kuwo.MusicResp{}
		case Tnetease:
			url = fmt.Sprintf("/cloudsearch?keywords=%s&limit=%d&offset=%d&type=1",
				param.Key, param.N, (param.P-1)*param.N)
			raw = &netease.MusicResp{}
		}

		wg.Add(1)
		go func(opt *ProxyReqOpt, raw any) {
			defer wg.Done()

			_ret := &std.MusicResp{}
			e := sendAndFormat(opt, raw, _ret)
			if e != nil {
				err = errors.Wrap(err, e.Error())
				return
			}

			for i := range _ret.Result.List {
				_ret.Result.List[i].Tag = _tag
			}

			retCh <- _ret
		}(&ProxyReqOpt{
			tag:    tag,
			rawUrl: url,
		}, raw)
	}

	return ret, err
}

func SearHint(c *gin.Context, param *service.SearHintReq) (*std.HintResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/search/searchKey"
	newUrl = fmt.Sprintf(
		"%s?key=%s&httpsStatus=1",
		newUrl, param.Key)

	rawRet := &kuwo.HintResp{}
	ret := &std.HintResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func SearAlbum(c *gin.Context, param *service.SearComReq) (*std.AlbumResp, error) {

	newUrl := "http://www.kuwo.cn/api/www/search/searchAlbumBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo.AlbumResp{}
	ret := &std.AlbumResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func SearMV(c *gin.Context, param *service.SearComReq) (*std.MVResp, error) {

	newUrl := "http://www.kuwo.cn/api/www/search/searchMvBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo.MVResp{}
	ret := &std.MVResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func SearArtist(c *gin.Context, param *service.SearComReq) (*std.ArtistResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/search/searchArtistBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo.ArtistResp{}
	ret := &std.ArtistResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func SearPlaylist(c *gin.Context, param *service.SearComReq) (*std.PlaylistResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/search/searchPlayListBykeyWord"
	newUrl = fmt.Sprintf(
		"%s?key=%s&pn=%d&rn=%d",
		newUrl, param.Key, param.P, param.N)

	rawRet := &kuwo.PlaylistResp{}
	ret := &std.PlaylistResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}
