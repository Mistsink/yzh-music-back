package core

// import (
// 	"fmt"

// 	"github.com/Mistsink/kuwo-api/internal/core/kuwo"
import (
	"fmt"

	kuwo_pl "github.com/Mistsink/kuwo-api/internal/core/kuwo/playlist"
	std_pl "github.com/Mistsink/kuwo-api/internal/core/standard/playlist"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/gin-gonic/gin"
)

func PLDefault(c *gin.Context, param *service.PLDefaultReq) (*std_pl.DefaultResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/classify/playlist/getRcmPlayList"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&order=%s",
		newUrl, param.P, param.N, param.Order)

	rawRet := &kuwo_pl.DefaultResp{}
	ret := &std_pl.DefaultResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLRecommend(c *gin.Context, param *service.PLRecommendReq) (*std_pl.RecResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/rcm/index/playlist"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&httpsStatus=1",
		newUrl, param.P, param.N)
	rawRet := &kuwo_pl.RecResp{}
	ret := &std_pl.RecResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLTags(c *gin.Context, _ any) (*std_pl.TagResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/playlist/getTagList"

	rawRet := &kuwo_pl.TagResp{}
	ret := &std_pl.TagResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLWithTag(c *gin.Context, param *service.PLWithTagReq) (*std_pl.WithTagResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/classify/playlist/getTagPlayList"
	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&id=%d",
		newUrl, param.P, param.N, param.Id)
	rawRet := &kuwo_pl.WithTagResp{}
	ret := &std_pl.WithTagResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}

func PLDetail(c *gin.Context, param *service.PLDetailReq) (*std_pl.DetailResp, error) {
	newUrl := "http://www.kuwo.cn/api/www/playlist/playListInfo"

	newUrl = fmt.Sprintf(
		"%s?pn=%d&rn=%d&pid=%d&httpsStatus=1",
		newUrl, param.P, param.N, param.Id)
	rawRet := &kuwo_pl.DetailResp{}
	ret := &std_pl.DetailResp{}

	err := getRespByNewUrl(c, newUrl, rawRet, ret)
	return ret, err
}
