package proxy

import (
	"fmt"
	"github.com/Mistsink/kuwo-api/internal/model/kuwo/rank"
	std_rank2 "github.com/Mistsink/kuwo-api/internal/model/standard/rank"

	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func RankPlaylist(c *gin.Context, _ any) (*std_rank2.PlaylistResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/bang/bang/bangMenu"
	newUrl = fmt.Sprintf(
		"%s?httpsStatus=1",
		newUrl)

	rawRet := &rank.PlaylistResp{}
	ret := &std_rank2.PlaylistResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func RankPlDetail(c *gin.Context, param *service.RankDetailReq) (*std_rank2.PlDetailResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/bang/bang/musicList"
	newUrl = fmt.Sprintf(
		"%s?bangId=%d&pn=%d&rn=%d&httpsStatus=1",
		newUrl, param.Id, param.P, param.N)

	rawRet := &rank.PlDetailResp{}
	ret := &std_rank2.PlDetailResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}
