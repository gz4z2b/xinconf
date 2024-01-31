# xinconf
一个统一go配置解决方案

## 快速接入

```go
import (
    "github.com/gz4z2b/xinconf"
	"github.com/gz4z2b/xinconf/types"
)

func main() {
    conf, err := xinconf.NewXinConf(types.XinConfConf{
        // 配置文件名
		ConfName:     "dev",
        // 配置框架类型，目前仅支持viper
		ConfType:     types.ConfTypeViper,
        // 配置文件后缀
		ConfFileType: "yaml",
        // 配置所在路径
		ConfPath:     []string{"conf"},
        // 日志记录,默认路径在runtime/config.log
	    // Logger types.XinLogger
	})

    // 目前仅支持用Unmarshal来结构化返回结果，所以需要先定义结果的结构体
    type Config struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	var config Config
    // key所对应值变更需要的回调
	conf.ConfigKeyOnChange("db.mysql", config, func(after any) {
        // 这里替换需要进行的操作
		fmt.Printf("after change: %+v", after)
	})
    // 获取配置值
	conf.Unmarshal("db.mysql", &config)
	fmt.Printf("获取配置%+v", config)
}
```