package api

import (
	"github.com/Mistsink/kuwo-api/internal/core"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/convert"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// import (
// 	"github.com/Mistsink/kuwo-api/internal/proxy"
// 	"github.com/Mistsink/kuwo-api/internal/service"
// 	"github.com/Mistsink/kuwo-api/pkg/app"
// 	"github.com/Mistsink/kuwo-api/pkg/convert"
// 	"github.com/gin-gonic/gin"
// )

type Playlist struct{}

func NewPlaylist() *Playlist {
	return &Playlist{}
}

func (p *Playlist) Default(c *gin.Context) {
	param := &service.PLDefaultReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.PLDefault(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (p *Playlist) Recommend(c *gin.Context) {
	param := &service.PLRecommendReq{}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.PLRecommend(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (p *Playlist) GetTags(c *gin.Context) {
	response := app.NewResponse(c)

	standardRes := core.PLTags(c, nil)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (p *Playlist) WithTag(c *gin.Context) {
	param := &service.PLWithTagReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.PLWithTag(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (p *Playlist) GetDetail(c *gin.Context) {
	param := &service.PLDetailReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.PLDetail(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}
