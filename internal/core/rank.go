package core

import (
	"fmt"

	kuwo_rank "github.com/Mistsink/kuwo-api/internal/core/kuwo/rank"
	std_rank "github.com/Mistsink/kuwo-api/internal/core/standard/rank"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func RankPlaylist(c *gin.Context, _ any) (*std_rank.PlaylistResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/bang/bang/bangMenu"
	newUrl = fmt.Sprintf(
		"%s?httpsStatus=1",
		newUrl)

	rawRet := &kuwo_rank.PlaylistResp{}
	ret := &std_rank.PlaylistResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func RankPlDetail(c *gin.Context, param *service.RankDetailReq) (*std_rank.PlDetailResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/bang/bang/musicList"
	newUrl = fmt.Sprintf(
		"%s?bangId=%d&pn=%d&rn=%d&httpsStatus=1",
		newUrl, param.Id, param.P, param.N)

	rawRet := &kuwo_rank.PlDetailResp{}
	ret := &std_rank.PlDetailResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}
