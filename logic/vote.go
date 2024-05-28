package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能
// 投一票就加423分 86400/200 ->200张赞成票才能给你的帖子续一天

// VoteForPost 为帖子投票函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
