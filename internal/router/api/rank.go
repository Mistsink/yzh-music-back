package api

import (
	"github.com/Mistsink/kuwo-api/internal/core"
	"github.com/Mistsink/kuwo-api/internal/service"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/convert"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Rank struct{}

func NewRank() *Rank {
	return &Rank{}
}

func (r *Rank) Rank(c *gin.Context) {
	response := app.NewResponse(c)

	standardRes := core.RankPlaylist(c, nil)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}

func (r *Rank) Detail(c *gin.Context) {
	param := &service.RankDetailReq{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	execBindAndValid(c, response, param)

	standardRes := core.RankPlDetail(c, param)

	resJson := app.RespJSON{
		Code:   errcode.Success.Code,
		Result: standardRes.Result,
		Msg:    standardRes.Msg,
		ReqID:  standardRes.ReqID,
	}
	response.ToResponse(&resJson)
}
