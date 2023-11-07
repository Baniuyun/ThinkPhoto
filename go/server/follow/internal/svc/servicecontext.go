package svc

import (
	"Thinkphoto/server/follow/internal/config"
	"Thinkphoto/server/follow/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	TFollowModel      model.TFollowModel
	Conn              sqlx.SqlConn
	TFollowCountModel model.TFollowCountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:            c,
		Conn:              mysqlConn,
		TFollowModel:      model.NewTFollowModel(mysqlConn),
		TFollowCountModel: model.NewTFollowCountModel(mysqlConn),
	}

}
