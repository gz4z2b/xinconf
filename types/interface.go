package types

type XinConf interface {
	// Unmarshal 对配置解析成对象
	Unmarshal(key string, val any) error
	// ConfigOnChange key配置变更时对应的操作
	ConfigKeyOnChange(key string, dataType any, deal func(after any))
}
