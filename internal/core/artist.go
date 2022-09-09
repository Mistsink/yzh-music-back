package core

import (
	"fmt"

	kuwo_art "github.com/Mistsink/kuwo-api/internal/core/kuwo/artist"
	std_artist "github.com/Mistsink/kuwo-api/internal/core/standard/artist"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func ArtMusicList(c *gin.Context, param *service.ArtMusicListReq) (*std_artist.MusicListResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/artist/artistMusic"
	newUrl = fmt.Sprintf("%s?artistid=%d&pn=%d&rn=%d", newUrl, param.Id, param.P, param.N)

	rawRet := &kuwo_art.MusicListResp{}
	ret := &std_artist.MusicListResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func ArtAlbum(c *gin.Context, param *service.ArtMusicListReq) (*std_artist.AlbumResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/artist/artistAlbum"
	newUrl = fmt.Sprintf("%s?artistid=%d&pn=%d&rn=%d", newUrl, param.Id, param.P, param.N)

	rawRet := &kuwo_art.AlbumResp{}
	ret := &std_artist.AlbumResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}
