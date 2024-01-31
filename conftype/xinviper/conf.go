package xinviper

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gz4z2b/xinconf/types"
	logertypes "github.com/gz4z2b/xinlog/types"
	"github.com/spf13/viper"
)

type ViperConf struct {
	viper     *viper.Viper
	watchKeys map[string][]onChangeDeal
	logger    logertypes.XinLogger
}

type onChangeDeal struct {
	dataType any
	deal     func(after any)
}

func NewViperConf(viper *viper.Viper, logger logertypes.XinLogger) types.XinConf {
	return &ViperConf{
		viper:     viper,
		watchKeys: make(map[string][]onChangeDeal),
		logger:    logger,
	}
}

// Unmarshal 对配置解析成对象
func (v *ViperConf) Unmarshal(key string, val any) error {
	err := v.viper.UnmarshalKey(key, val)
	return err
}

// ConfigOnChange key配置变更时对应的操作
func (v *ViperConf) ConfigKeyOnChange(key string, dataType any, deal func(after any)) {
	if _, exist := v.watchKeys[key]; exist {
		v.watchKeys[key] = append(v.watchKeys[key], onChangeDeal{
			dataType: dataType,
			deal:     deal,
		})
	} else {
		v.watchKeys[key] = []onChangeDeal{
			{
				dataType: dataType,
				deal:     deal,
			},
		}
	}

	v.viper.OnConfigChange(func(in fsnotify.Event) {
		for key, deals := range v.watchKeys {
			for _, deal := range deals {
				err := v.viper.UnmarshalKey(key, &deal.dataType)
				if err != nil {
					v.logger.Error("配置变更会写数据格式错误", logertypes.NewField("err", err))
				}
				deal.deal(deal.dataType)
			}
		}
	})
}
