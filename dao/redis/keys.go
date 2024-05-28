package redis

//redis key 尽量使用命名空间的方式，方便业务的查询和拆分

const (
	Prefix             = "bluebell:"
	KeyPostTimeZSet    = "post:time"   //zset:帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  //zset：帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" //zset 记录用户和投票类型 参数是post id
	KeyCommunitySetPF  = "community:"  //set;保存分区下帖子的id
)

// 给rediskey加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
