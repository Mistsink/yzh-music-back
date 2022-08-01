package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/pkg/logger"

	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type RespJSON struct {
	Code    int      `json:"code" yzh:"code"`
	Result  any      `json:"result"`
	Msg     string   `json:"msg" yzh:"msg"`
	Details []string `json:"details"`
	ReqID   string   `json:"req_id" yzh:"req_id"`
}

type Response struct {
	ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{c}
}

func (r Response) ToResponseWithReader(reader io.Reader) {
	res := new(RespJSON)
	all, _ := ioutil.ReadAll(reader)
	jsonFormat := global.Logger.JSONFormat(logger.LevelInfo, string(all))
	fmt.Printf("%+v\n", jsonFormat)
	deco := json.NewDecoder(reader)
	if err := deco.Decode(res); err != nil {
		global.Logger.Errorf(r.ctx, "Decoding reader err: %v", err)
		r.ToErrResponse(errcode.ServerWithMsg(fmt.Sprintf("Decoding reader err: %v", err)))
		return
	}
	r.ctx.JSON(errcode.Success.StatusCode(), res)
}

func (r *Response) ToResponse(data *RespJSON) {
	if data == nil {
		data = &RespJSON{}
	}
	r.ctx.JSON(errcode.Success.StatusCode(), data)
}

func (r *Response) ToErrResponse(err *errcode.Error) {
	res := RespJSON{
		Code: err.Code,
		Msg:  err.Msg,
	}
	if len(err.Details) != 0 {
		res.Details = err.Details
	}
	r.ctx.JSON(err.StatusCode(), &res)
}
