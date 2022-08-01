package setting

import "time"

type AppSettingS struct {
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	DefaultContextTimeout time.Duration
	Cookie                string
	Host                  string
	CSRF                  string
	Referer               string
	UserAgent             string
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
