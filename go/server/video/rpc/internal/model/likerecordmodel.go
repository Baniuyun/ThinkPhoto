package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeRecordModel = (*customLikeRecordModel)(nil)

type (
	// LikeRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeRecordModel.
	LikeRecordModel interface {
		likeRecordModel
		FindLikeListById(ctx context.Context, id int64) ([]*LikeRecord, error)
		FindLikeRelation(ctx context.Context, likeing_id, liker_id int64) (*LikeRecord, error)
		DeleteByLinkIdAndLikerId(ctx context.Context, likeing_id, liker_id int64) error
	}

	customLikeRecordModel struct {
		*defaultLikeRecordModel
	}
)

// NewLikeRecordModel returns a model for the database table.
func NewLikeRecordModel(conn sqlx.SqlConn) LikeRecordModel {
	return &customLikeRecordModel{
		defaultLikeRecordModel: newLikeRecordModel(conn),
	}
}

func (m *defaultLikeRecordModel) FindLikeListById(ctx context.Context, id int64) ([]*LikeRecord, error) {
	query := fmt.Sprintf("select %s from %s where `liker_id` = ? ", likeRecordRows, m.table)
	var resp []*LikeRecord
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeRecordModel) FindLikeRelation(ctx context.Context, likeing_id, liker_id int64) (*LikeRecord, error) {
	query := fmt.Sprintf("select %s from %s where `linking_id` = ? and  `liker_id` = ?", likeRecordRows, m.table)
	var resp *LikeRecord
	err := m.conn.QueryRowCtx(ctx, &resp, query, likeing_id, liker_id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeRecordModel) DeleteByLinkIdAndLikerId(ctx context.Context, likeing_id, liker_id int64) error {
	query := fmt.Sprintf("delete from %s where `linking_id` = ? and `liker_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, likeing_id, liker_id)
	return err
}
