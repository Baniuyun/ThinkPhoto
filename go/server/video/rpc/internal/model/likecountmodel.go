package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeCountModel = (*customLikeCountModel)(nil)

type (
	// LikeCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeCountModel.
	LikeCountModel interface {
		likeCountModel
		//	func (m *defaultLikeCountModel) FindLikeListById(ctx context.Context) ([]*LikeCount, error)
		UpdateFavoriteCount(ctx context.Context, status, linking_id int64) (int64, error)
		DeleteByVideoId(ctx context.Context, id int64) error
	}

	customLikeCountModel struct {
		*defaultLikeCountModel
	}
)

// NewLikeCountModel returns a model for the database table.
func NewLikeCountModel(conn sqlx.SqlConn) LikeCountModel {
	return &customLikeCountModel{
		defaultLikeCountModel: newLikeCountModel(conn),
	}
}

func (m *defaultLikeCountModel) UpdateFavoriteCount(ctx context.Context, status, linking_id int64) (int64, error) {
	var update string
	if status == 0 {
		update = fmt.Sprintf("update %s set like_num = like_num + 1 where `linking_id` = ?", m.table)
	} else {
		update = fmt.Sprintf("update %s set like_num = like_num - 1 where `linking_id` = ?", m.table)
	}
	_, err := m.conn.ExecCtx(ctx, update, linking_id)
	switch err {
	case nil:
		return 1, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultLikeCountModel) DeleteByVideoId(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `linking_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
