package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DB struct {
		Host         string
		Port         string
		Database     string
		Username     string
		Password     string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	zrpc.RpcServerConf
}
