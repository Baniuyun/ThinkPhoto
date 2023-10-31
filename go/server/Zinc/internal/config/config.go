package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	ZincSearch struct {
		UserName string
		Password string
		Addr     string
		Index    string
	}
}
