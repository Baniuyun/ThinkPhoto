package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TFollowCountModel = (*customTFollowCountModel)(nil)

type (
	// TFollowCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFollowCountModel.
	TFollowCountModel interface {
		tFollowCountModel
	}

	customTFollowCountModel struct {
		*defaultTFollowCountModel
	}
)

// NewTFollowCountModel returns a model for the database table.
func NewTFollowCountModel(conn sqlx.SqlConn) TFollowCountModel {
	return &customTFollowCountModel{
		defaultTFollowCountModel: newTFollowCountModel(conn),
	}
}
