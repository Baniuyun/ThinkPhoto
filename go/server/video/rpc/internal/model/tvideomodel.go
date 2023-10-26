package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TVideoModel = (*customTVideoModel)(nil)

type (
	// TVideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVideoModel.
	TVideoModel interface {
		tVideoModel
	}

	customTVideoModel struct {
		*defaultTVideoModel
	}
)

// NewTVideoModel returns a model for the database table.
func NewTVideoModel(conn sqlx.SqlConn) TVideoModel {
	return &customTVideoModel{
		defaultTVideoModel: newTVideoModel(conn),
	}
}
