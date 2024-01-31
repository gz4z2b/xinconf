package main

import (
	"errors"

	"github.com/gz4z2b/xinconf/types"
	"github.com/gz4z2b/xinconf/xinviper"
)

var Err_NotSupport = errors.New("not support")

func NewXinConf(conf types.XinConfConf) (types.XinConf, error) {
	switch conf.ConfType {
	case types.ConfTypeViper:
		return xinviper.NewViper(conf), nil
	default:
		return nil, Err_NotSupport
	}
}
