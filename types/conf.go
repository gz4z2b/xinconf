package types

import "github.com/gz4z2b/xinlog/types"

const ConfTypeViper = "viper"

type XinConfConf struct {
	// 配置文件找寻路径
	ConfPath []string
	// 配置文件名
	ConfName string
	// 配置框架类型
	ConfType string
	// 配置文件后缀
	ConfFileType string
	// 日志记录
	Logger types.XinLogger
	// 远程配置中心类型
	RemoteConfType string
	// 远程配置中心地址
	RemoteConfAdd string
}
