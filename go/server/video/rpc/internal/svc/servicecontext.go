package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"rpc/internal/config"
	"rpc/internal/model"
)

type ServiceContext struct {
	Config      config.Config
	TVideoModel model.TVideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		TVideoModel: model.NewTVideoModel(mysqlConn),
	}
}
