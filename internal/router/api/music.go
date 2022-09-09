package api

import (
	"github.com/Mistsink/kuwo-api/internal/core"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/convert"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Music struct{}

func NewMusic() *Music {
	return &Music{}
}

func (m *Music) Url(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := core.MusicUrl(c, param)
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

func (m *Music) MV(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := core.MusicUrl(c, param)
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

func (m *Music) Lyric(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := core.MusicLyric(c, param)
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

func (m *Music) Info(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	res, err := core.MusicInfo(c, param)
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
