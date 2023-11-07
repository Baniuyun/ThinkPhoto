package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TVideoModel = (*customTVideoModel)(nil)

type (
	// TVideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTVideoModel.
	TVideoModel interface {
		tVideoModel
		FindListAll(ctx context.Context) ([]*TVideo, error)
		FindListById(ctx context.Context, id int64) ([]*TVideo, error)
		FindListByTag(ctx context.Context, tag int64) ([]*TVideo, error)
		UpdateCommentCount(ctx context.Context, id int64) (int64, error)
		UpdateFavoriteCountByStatus(ctx context.Context, status, id int64) (int64, error)
		FindListByAuthorId(ctx context.Context, id int64) ([]*TVideo, error)
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

func (m *defaultTVideoModel) FindListAll(ctx context.Context) ([]*TVideo, error) {
	query := fmt.Sprintf("select %s from %s", tVideoRows, m.table)
	var resp []*TVideo
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTVideoModel) FindListById(ctx context.Context, id int64) ([]*TVideo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ?", tVideoRows, m.table)
	var resp []*TVideo
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

func (m *defaultTVideoModel) FindListByTag(ctx context.Context, tag int64) ([]*TVideo, error) {
	query := fmt.Sprintf("select %s from %s where `tag` = ?", tVideoRows, m.table)
	var resp []*TVideo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, tag)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTVideoModel) UpdateCommentCount(ctx context.Context, id int64) (int64, error) {
	update := fmt.Sprintf("update %s set comment_count = comment_count + 1 where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, update, id)
	switch err {
	case nil:
		return 1, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultTVideoModel) UpdateFavoriteCountByStatus(ctx context.Context, status, id int64) (int64, error) {
	var update string
	if status == 0 {
		update = fmt.Sprintf("update %s set favorite_count = favorite_count + 1 where `id` = ?", m.table)
	} else {
		update = fmt.Sprintf("update %s set favorite_count = favorite_count - 1 where `id` = ?", m.table)
	}
	_, err := m.conn.ExecCtx(ctx, update, id)
	switch err {
	case nil:
		return 1, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultTVideoModel) FindListByAuthorId(ctx context.Context, id int64) ([]*TVideo, error) {
	query := fmt.Sprintf("select %s from %s where `author_id` = ?", tVideoRows, m.table)
	var resp []*TVideo
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
