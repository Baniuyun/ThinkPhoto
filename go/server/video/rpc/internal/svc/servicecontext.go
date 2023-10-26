package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"rpc/internal/config"
	"rpc/internal/model"
)

type ServiceContext struct {
	Config      config.Config
	TVideoModel model.TVideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	addr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.DB.Username,
		c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
	conn := sqlx.NewMysql(addr)
	return &ServiceContext{
		Config:      c,
		TVideoModel: model.NewTVideoModel(conn),
	}
}
