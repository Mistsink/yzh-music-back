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

	standardRes := core.MusicUrl(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (m *Music) MV(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.MusicUrl(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

// func (m *Music) Comment(c *gin.Context) {
// 	param := &service.MusicCommentReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
// 	response := app.NewResponse(c)
// 	execBindAndValid(c, response, param)

// 	_u := "http://www.kuwo.cn/comment"
// 	_u = fmt.Sprintf("%s?type=%s&f=web&page=%d&rows=%d&digest=%d&sid=%d&prod=newWeb&uid=%d",
// 		_u, param.Type, param.P, param.N, param.Digest, param.Id, 0)
// 	u := parseUrl(_u, response)

// 	RewriteUrl(c, u)
// }

func (m *Music) Lyric(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.MusicLyric(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (m *Music) Info(c *gin.Context) {
	param := &service.MusicUrlReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.MusicInfo(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}
