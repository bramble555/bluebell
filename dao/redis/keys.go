package redis

// redis key, 注意使用命名，方便查询和拆分

const (
	KeyPrefix          = "bluebell:"
	KeyZSetPostTime    = "post:time"   // 帖子及发帖时间
	KeyZSetPostScore   = "post:score"  // 帖子及投票分数
	KeyZSetPostVotedPF = "post:voted:" // 记录用户及投票类型；参数是post_id
)
