package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/internal/router"
	"github.com/Mistsink/kuwo-api/setup"
	"github.com/gin-gonic/gin"
)

func init() {
	setup.Init()
}

func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	r := router.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 10 << 20,
		Handler:        r,
	}
	fmt.Println("Service running on port:", global.ServerSetting.HttpPort)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln("Fail to ListenAndServe():", err)
	}
}
