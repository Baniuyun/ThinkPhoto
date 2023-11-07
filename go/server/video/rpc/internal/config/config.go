package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	Qiniu struct {
		Addr            string
		AccessKeyId     string
		SecretAccessKey string
		VideoBuckets    string
		VideoPath       string
	}
}
