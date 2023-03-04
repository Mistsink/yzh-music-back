package proxy

import (
	"fmt"
	kuwo "github.com/Mistsink/kuwo-api/internal/model/kuwo/artist"
	netease "github.com/Mistsink/kuwo-api/internal/model/netease/artist"
	std "github.com/Mistsink/kuwo-api/internal/model/standard/artist"

	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func ArtMusicList(param *service.ArtMusicListReq) (*std.MusicListResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.MusicListResp{}
		tag = PlatformTag(param.Tag)
	)
	ret.Result.Tag = string(tag)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/artist/artistMusic"
		url = fmt.Sprintf("%s?artistid=%d&pn=%d&rn=%d",
			url, param.Id, param.P, param.N)
		raw = &kuwo.MusicListResp{}
	case Tnetease:
		url = fmt.Sprintf("/artist/songs?id=%d&order=%s&limit=%d&offset=%d",
			param.Id, param.Order, param.N, (param.P-1)*param.N)
		raw = &netease.MusicListResp{}
	}

	err = sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	return ret, err
}

func ArtAlbum(c *gin.Context, param *service.ArtMusicListReq) (*std.AlbumResp, error) {
	var (
		err error
		url string
		raw any
		ret = &std.AlbumResp{}
		tag = PlatformTag(param.Tag)
	)
	ret.Result.Tag = string(tag)

	switch tag {
	case Tkuwo:
		url = "http://www.kuwo.cn/api/www/artist/artistAlbum"
		url = fmt.Sprintf("%s?artistid=%d&pn=%d&rn=%d",
			url, param.Id, param.P, param.N)
		raw = &kuwo.AlbumResp{}
	case Tnetease:
		url = fmt.Sprintf("/artist/album?id=%d&limit=%d&offset=%d",
			param.Id, param.N, (param.P-1)*param.N)
		raw = &netease.AlbumResp{}
	}

	err = sendAndFormat(&ProxyReqOpt{
		tag:    tag,
		rawUrl: url,
	}, raw, ret)
	return ret, err
}
