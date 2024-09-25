package redis

import (
	"bluebell/global"
	"bluebell/models"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds         = 7 * 24 * 3600 * 3600
	perVoteScore     float64 = 432.0 // 每一票值多少分
	OrderByTime              = "time"
	OrderByScore             = "score"
)

var (
	ErrVoteTimeExprie = errors.New("投票时间已经过去")
)

// CreatePost 创建帖子的时候向redis里面增加 KeyZSetPostTime 和 KeyZSetPostScore
func CreatePost(postID int) error {
	pipe := global.RDB.Pipeline()
	// 下面增加需要 事务处理
	pipe.ZAdd(getKeyName(KeyZSetPostTime), redis.Z{
		Member: postID,
		Score:  float64(time.Now().Unix()),
	})
	// 每个帖子原始分数是 创建的时间戳
	pipe.ZAdd(getKeyName(KeyZSetPostScore), redis.Z{
		Member: postID,
		Score:  float64(time.Now().Unix()),
	})
	_, err := pipe.Exec()
	return err
}

func GetPostIDList2(ppl *models.ParamPostList) (res []string, err error) {
	var strat int64 = int64((ppl.Page - 1) * ppl.Size) // 3 4 start: 2*4=8 end: 8+4-1=11
	var end int64 = strat + int64(ppl.Size) - 1
	if ppl.Order == OrderByScore {
		res, err = global.RDB.ZRange(getKeyName(KeyZSetPostScore), strat, end).Result()
	} else {
		res, err = global.RDB.ZRange(getKeyName(KeyZSetPostTime), strat, end).Result()
	}
	global.Log.Debugln("id 分别为", res)
	return res, err
}
func VoteForPost(userID int, postID int, curPos float64) error {
	ui := strconv.Itoa(userID) // userID String 类型
	pi := strconv.Itoa(postID) // PostID String 类型
	// 1.判断投票限制
	// redis取帖子发布时间
	postTime := global.RDB.ZScore(getKeyName(KeyZSetPostTime), ui).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExprie
	}
	// 2.更新帖子分数
	// 先获取当前是赞成还是反对
	prePos := global.RDB.ZScore(getKeyName(KeyZSetPostVotedPF+pi), ui).Val()
	pos := math.Abs(prePos - curPos)
	var temp float64
	if curPos > prePos {
		temp = 1
	} else {
		temp = -1
	}
	// 更新分数和 3 要进行 事务处理
	pipe := global.RDB.Pipeline()
	// 更新分数
	_, err := pipe.ZIncrBy(getKeyName(KeyZSetPostVotedPF+pi), pos*temp*perVoteScore, pi).Result()
	if err != nil {
		return err
	}
	// 3.记录用户为该帖子的投票数据
	if curPos == 0 {
		pipe.ZRem(getKeyName(KeyZSetPostVotedPF+pi), ui)
	} else {
		pipe.ZAdd(getKeyName(KeyZSetPostVotedPF+pi), redis.Z{
			Member: ui,
			Score:  curPos,
		})
	}
	_, err = pipe.Exec()
	return err

}
