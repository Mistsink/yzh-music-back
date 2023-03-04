package setting

import (
	"io"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configIn any) (*Setting, error) {
	var err error

	vp := viper.New()

	vp.SetConfigType("yaml")
	vp.SetConfigName("config")

	switch configIn := configIn.(type) {
	case string:
		vp.AddConfigPath(configIn)
		err = vp.ReadInConfig()
	case nil:
		vp.AddConfigPath("config/")
		err = vp.ReadInConfig()
	case io.Reader:
		err = vp.ReadConfig(configIn)
	}

	return &Setting{vp}, err
}
