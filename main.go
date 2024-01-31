package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gz4z2b/xinconf/types"
)

func main() {
	server := gin.Default()

	conf, err := NewXinConf(types.XinConfConf{
		ConfName:     "dev",
		ConfType:     types.ConfTypeViper,
		ConfFileType: "yaml",
		ConfPath:     []string{"conf"},
	})
	if err != nil {
		panic(err)
	}
	type Config struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	var config Config
	conf.ConfigKeyOnChange("db.mysql", config, func(after any) {
		fmt.Printf("after change: %+v", after)
	})
	conf.Unmarshal("db.mysql", &config)
	fmt.Printf("获取配置%+v", config)

	server.Run(":8081")
}
