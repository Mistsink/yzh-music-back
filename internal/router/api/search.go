package api

import (
	"github.com/Mistsink/kuwo-api/internal/proxy"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Search struct{}

func NewSearch() *Search {
	return &Search{}
}

func (s Search) Music(c *gin.Context) {
	param := &service.SearComReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearMusic(c, param)
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

func (s Search) Hint(c *gin.Context) {
	param := &service.SearHintReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearHint(c, param)
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

func (s Search) Album(c *gin.Context) {
	param := &service.SearComReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearAlbum(c, param)
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

func (s Search) MV(c *gin.Context) {
	param := &service.SearComReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearMV(c, param)
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

func (s Search) Artist(c *gin.Context) {
	param := &service.SearComReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearArtist(c, param)
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

func (s Search) Playlist(c *gin.Context) {
	param := &service.SearComReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := proxy.SearPlaylist(c, param)
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
