package setup

import (
	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/pkg/logger"
	"github.com/Mistsink/kuwo-api/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	os.Chdir(getDir())
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting() err %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger() err %v", err)
	}
}

func setupSetting() error {
	var (
		settings *setting.Setting
		err      error
	)
	settings, err = setting.NewSetting("config/")
	if err != nil {
		return err
	}

	if err = settings.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second

	if err = settings.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func getDir() string {
	dir, _ := os.Getwd()
	i := strings.Index(dir, "kuwo-api")
	return dir[:i] + "kuwo-api/"
}
