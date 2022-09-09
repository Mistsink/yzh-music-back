package api

import (
	"fmt"
	"log"

	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func execBindAndValid(c *gin.Context, response *app.Response, param interface{}) {
	if err := c.ShouldBind(param); err != nil {
		response.ToErrResponse(errcode.ServerWithMsg(fmt.Sprintf("c.ShouldBind(param) err: %v", err)))
		log.Println("c.ShouldBind(param) err:", err)
		panic(fmt.Sprintf("c.ShouldBind(param) err: %v", err))
	}
}

func transCode(rawCode int) (ret int) {
	switch rawCode {
	case 200:
		ret = errcode.Success.Code
	default:
		ret = errcode.ServerError.Code
	}

	return
}
