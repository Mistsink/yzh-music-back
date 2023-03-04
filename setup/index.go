package setup

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Mistsink/kuwo-api/config"
	"github.com/Mistsink/kuwo-api/global"
	"github.com/Mistsink/kuwo-api/internal/proxy"
	"github.com/Mistsink/kuwo-api/pkg/logger"
	"github.com/Mistsink/kuwo-api/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting() err %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger() err %v", err)
	}

	proxy.InitProxyService()
}

func setupSetting() error {
	var (
		settings *setting.Setting
		err      error
	)

	configIn, err := config.FS.Open("config.yaml")
	defer func() { _ = configIn.Close() }()

	if err != nil {
		return err
	}

	settings, err = setting.NewSetting(configIn)
	if err != nil {
		return err
	}

	if err = settings.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second

	if err = settings.ReadSection("Proxy", &global.ProxySetting); err != nil {
		return err
	}
	global.ProxySetting.AvailableTags = make([]string, len(global.ProxySetting.ProxyItems))
	for i, proxyS := range global.ProxySetting.ProxyItems {
		global.ProxySetting.ProxyItems[i].Name = strings.ToLower(proxyS.Name)
		global.ProxySetting.AvailableTags[i] = global.ProxySetting.ProxyItems[i].Name
	}

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
