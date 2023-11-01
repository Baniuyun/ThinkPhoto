package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TFollowModel = (*customTFollowModel)(nil)

type (
	// TFollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFollowModel.
	TFollowModel interface {
		tFollowModel
	}

	customTFollowModel struct {
		*defaultTFollowModel
	}
)

// NewTFollowModel returns a model for the database table.
func NewTFollowModel(conn sqlx.SqlConn) TFollowModel {
	return &customTFollowModel{
		defaultTFollowModel: newTFollowModel(conn),
	}
}
