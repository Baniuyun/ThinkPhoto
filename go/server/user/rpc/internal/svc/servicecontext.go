package svc

import (
	"Thinkphoto/server/user/rpc/internal/config"
	"Thinkphoto/server/user/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TUserModel model.TUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TUserModel: model.NewTUserModel(mysqlConn),
	}
}
