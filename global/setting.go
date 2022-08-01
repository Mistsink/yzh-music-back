package global

import (
	"github.com/Mistsink/kuwo-api/pkg/logger"
	"github.com/Mistsink/kuwo-api/pkg/setting"
)

var (
	AppSetting    *setting.AppSettingS
	ServerSetting *setting.ServerSettingS
	Logger        *logger.Logger
)
