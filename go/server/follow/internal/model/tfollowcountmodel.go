package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TFollowCountModel = (*customTFollowCountModel)(nil)

type (
	// TFollowCountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTFollowCountModel.
	TFollowCountModel interface {
		tFollowCountModel
		FollowCountWithSession(session sqlx.Session, followerId int64, followingId int64, isFollow int64) error
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

// 关注/取关 更新关注者和被关注者的关注数，粉丝数
func (m *defaultTFollowCountModel) FollowCountWithSession(session sqlx.Session, followerId int64, followingId int64, isFollow int64) error {
	followCountModel := m.withSession(session)

	// 更新 关注者/取关者的 关注数
	follower, err := followCountModel.FindOneByUserId(context.Background(), followerId)
	if err != nil {
		logx.Errorf("t_follow_count select fail:%v", err)
	}
	if isFollow == 1 {
		follower.FollowingCount += 1
	} else {
		follower.FollowingCount -= 1
	}

	// 更新操作
	err = followCountModel.Update(context.Background(), follower)
	if err != nil {
		logx.Errorf("t_follow_count update fail:%v", err)
	}

	// 更新 被关注者/被取关者的 粉丝数
	following, err := followCountModel.FindOneByUserId(context.Background(), followingId)
	if err != nil {
		logx.Errorf("t_follow_count select fail:%v", err)
	}
	if isFollow == 0 {
		following.FollowerCount -= 1
	} else {
		following.FollowerCount += 1
	}

	// 更新操作
	err = followCountModel.Update(context.Background(), following)
	if err != nil {
		logx.Errorf("t_follow_count update fail:%v", err)
	}
	return nil
}
