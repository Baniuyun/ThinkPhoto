package svc

import (
	"Thinkphoto/server/video/rpc/internal/config"
	"Thinkphoto/server/video/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	TVideoModel           model.TVideoModel
	TVideoLikeCountModel  model.LikeCountModel
	TVideoLikeRecordModel model.LikeRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		TVideoModel:           model.NewTVideoModel(mysqlConn),
		TVideoLikeCountModel:  model.NewLikeCountModel(mysqlConn),
		TVideoLikeRecordModel: model.NewLikeRecordModel(mysqlConn),
	}
}
