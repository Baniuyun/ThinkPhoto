package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ TFollowModel = (*customTFollowModel)(nil)

type (
	// TFollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFollowModel.
	TFollowModel interface {
		tFollowModel
		FollowWithSession(session sqlx.Session, followerId int64, followingId int64, isFollow int64) error
		FindFansList(ctx context.Context, followingId int64, cursor int64, pageSize int64) ([]*TFollow, error)
		FindFollowList(ctx context.Context, followerId int64, cursor int64, pageSize int64) ([]*TFollowItem, error)
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

// 关注 此前未关注过：插入数据到关注关系表 关注过：更新 isFollow
// 取关 更新 isFollow
func (m *defaultTFollowModel) FollowWithSession(session sqlx.Session, followerId int64, followingId int64, isFollow int64) error {
	followModel := m.withSession(session)

	follow, err := followModel.FindOneByFollowerIdFollowingId(context.Background(), followerId, followingId)

	// 判断是否关注过
	if err != nil {
		// 未关注过, 执行 insert
		tFollow := &TFollow{FollowerId: followerId, FollowingId: followingId, IsFollow: isFollow}
		_, err = followModel.Insert(context.Background(), tFollow)
		if err != nil {
			logx.Errorf("t_follow insert err:%v", err)
			return err
		}
	} else {
		// 关注过，执行 update
		follow.IsFollow = isFollow
		err = followModel.Update(context.Background(), follow)
		if err != nil {
			logx.Errorf("t_follow update err:%v", err)
			return err
		}
	}

	return nil
}

func (m *defaultTFollowModel) FindFansList(ctx context.Context, followingId int64, cursor int64, pageSize int64) ([]*TFollow, error) {
	query := fmt.Sprintf("select %s from %s where `following_id` = ? order by `create_time` desc limit ?, ?", tFollowRows, m.table)
	var resp []*TFollow

	err := m.conn.QueryRowsCtx(ctx, &resp, query, followingId, (cursor-1)*pageSize, pageSize)

	// 异常由上层处理
	return resp, err
}

// 封装返回参数
type TFollowItem struct {
	Id         int64     `db:"id"`
	FollowerId int64     `db:"following_id"`
	FansCount  int64     `db:"follower_count"`
	CreateTime time.Time `db:"create_time"`
}

func (m *defaultTFollowModel) FindFollowList(ctx context.Context, followerId int64, cursor int64, pageSize int64) ([]*TFollowItem, error) {
	// select `tf.id`, `tf.following_id`, `tf.create_time`, `tfc.follower_count` from `t_follow` tf left join `t_follow_count` tfc on `tf.following_id` = `tfc.user_id` order by `tf.create_time` desc limit ?, ?
	query := fmt.Sprintf("select tf.id, tf.following_id, tfc.follower_count, tf.create_time from t_follow as tf left join t_follow_count as tfc on tf.following_id = tfc.user_id where tf.follower_id = ? order by tf.create_time desc limit ?, ?")
	var resp []*TFollowItem

	err := m.conn.QueryRowsCtx(ctx, &resp, query, followerId, (cursor-1)*pageSize, pageSize)

	// 异常由上层处理
	return resp, err
}
