package middlerware

import (
	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/pkg/app"
	"github.com/Mistsink/kuwo-api/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				app.NewResponse(c).ToErrResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
