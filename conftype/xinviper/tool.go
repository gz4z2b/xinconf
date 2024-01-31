package xinviper

import (
	"github.com/gz4z2b/xinconf/types"
	xinlogger "github.com/gz4z2b/xinlog"
	loggertypes "github.com/gz4z2b/xinlog/types"
	"github.com/spf13/viper"
)

func NewViper(conf types.XinConfConf) types.XinConf {
	viper := viper.New()
	viper.SetConfigName(conf.ConfName)
	viper.SetConfigType(conf.ConfFileType)
	for _, v := range conf.ConfPath {
		viper.AddConfigPath(v)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	if conf.Logger == nil {
		logger, err := xinlogger.NewLogger(loggertypes.Conf{
			LogPath:     "runtime/conf.log",
			Type:        xinlogger.ZAP_TYPE,
			MaxSize:     10,
			MaxBackups:  30,
			MaxAge:      10,
			EnableLevel: loggertypes.InfoLevel,
		})
		if err != nil {
			panic(err)
		}
		conf.Logger = logger
	}
	return NewViperConf(viper, conf.Logger)
}
