package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserModel = (*customTUserModel)(nil)

type (
	// TUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserModel.
	TUserModel interface {
		tUserModel
		FindOneByUsername(ctx context.Context, username string) (*TUser, error)
	}

	customTUserModel struct {
		*defaultTUserModel
	}
)

// NewTUserModel returns a model for the database table.
func NewTUserModel(conn sqlx.SqlConn) TUserModel {
	return &customTUserModel{
		defaultTUserModel: newTUserModel(conn),
	}
}

// FindOneByUsername 通过用户名查找用户
func (c customTUserModel) FindOneByUsername(ctx context.Context, username string) (*TUser, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", tUserRows, c.table)
	var resp TUser
	err := c.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		logx.Errorf("FindOneByUsername User Position Model Model err, err=%v", err)
		return nil, err
	}
}
