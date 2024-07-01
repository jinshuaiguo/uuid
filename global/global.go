package global

import (
	"github.com/spf13/viper"
	"uuid_srv/config"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
