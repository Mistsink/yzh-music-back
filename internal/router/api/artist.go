package api

import (
	"github.com/Mistsink/kuwo-api/internal/proxy"
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

	res, err := proxy.ArtMusicList(param)
	if err != nil {
		response.ToErrResponse(errcode.ServerWithMsg(err.Error()))
		return
	}

	resJson := app.RespJSON{
		Code:   transCode(res.Code),
		Result: res.Result,
		Msg:    res.Msg,
		ReqID:  res.ReqID,
	}
	response.ToResponse(&resJson)
}

func (a *Artist) Album(c *gin.Context) {
	param := &service.ArtMusicListReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.ArtAlbum(c, param)
	if err != nil {
		response.ToErrResponse(errcode.ServerWithMsg(err.Error()))
		return
	}

	resJson := app.RespJSON{
		Code:   transCode(res.Code),
		Result: res.Result,
		Msg:    res.Msg,
		ReqID:  res.ReqID,
	}
	response.ToResponse(&resJson)
}
