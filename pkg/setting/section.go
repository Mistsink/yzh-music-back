package setting

import "time"

type AppSettingS struct {
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	DefaultContextTimeout time.Duration
}

type ProxySettingS struct {
	AvailableTags []string
	ProxyItems    []ProxySettingItem
}

type ProxySettingItem struct {
	Name      string
	Cookie    string
	CSRF      string
	Host      string
	Referer   string
	UserAgent string
	ProxyAddr string
}

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Setting) ReadSection(key string, setting interface{}) error {
	return s.vp.UnmarshalKey(key, setting)
}
