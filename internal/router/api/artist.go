package api

import (
	"fmt"

	"github.com/Mistsink/kuwo-api/internal/core"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/convert"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Artist struct{}

func NewArtist() *Artist {
	return &Artist{}
}

func (a *Artist) Music(c *gin.Context) {
	param := &service.ArtMusicListReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.ArtMusicList(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (a *Artist) Album(c *gin.Context) {
	param := &service.ArtMusicListReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.ArtAlbum(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

///	no support
func (a *Artist) Recommend(c *gin.Context) {
	param := &service.ArtRecReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	_u := "http://www.kuwo.cn/api/www/artist/artistInfo"
	_u = fmt.Sprintf("%s?category=%d&pn=%d&rn=%d", _u, param.Category, param.P, param.N)
	u := parseUrl(_u, response)

	RewriteUrl(c, u)
}

func (a *Artist) WithTag(c *gin.Context) {
	param := &service.ArtWithTagReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	_u := "http://www.kuwo.cn/api/www/artist/artistInfo"
	_u = fmt.Sprintf("%s?category=%d&pn=%d&rn=%d", _u, param.Category, param.P, param.N)
	u := parseUrl(_u, response)

	RewriteUrl(c, u)
}
