package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configPath ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	if configPath == nil {
		vp.AddConfigPath("config/")
	} else {
		vp.AddConfigPath(configPath[0])
	}
	err := vp.ReadInConfig()
	return &Setting{vp}, err
}
