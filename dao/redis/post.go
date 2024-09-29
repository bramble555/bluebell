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
	ErrorRepeateVote  = errors.New("请不要重复投票")
)

// CreatePost 创建帖子的时候向redis里面增加 KeyZSetPostTime 和 KeyZSetPostScore
func CreatePost(postID, communityID int) error {
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
	// 把 postID 加入到 相应 KeySetCommuntiyPF 里面
	pipe.SAdd(getKeyName(KeySetCommuntiyPF+strconv.Itoa(communityID)), postID)
	_, err := pipe.Exec()
	return err
}

func GetPostIDList(ppl *models.ParamPostList) (res []string, err error) {
	var strat int64 = int64((ppl.Page - 1) * ppl.Size) // 2 1 start: 1*1=1 end: 1+1-1
	var end int64 = strat + int64(ppl.Size) - 1
	if ppl.Order == OrderByScore {
		res, err = global.RDB.ZRevRange(getKeyName(KeyZSetPostScore), strat, end).Result()
	} else {
		res, err = global.RDB.ZRevRange(getKeyName(KeyZSetPostTime), strat, end).Result()
	}
	global.Log.Debugln("id 分别为", res)
	return res, err
}
func GetCommuntiyPostIDList(ppl *models.ParamPostList) (res []string, err error) {
	
	// 默认按照时间
	orderKey := getKeyName(KeyZSetPostTime) // post:time
	if ppl.Order == OrderByScore {
		orderKey = getKeyName(KeyZSetPostScore)
	}
	// 社区的key
	cKey := getKeyName(KeySetCommuntiyPF) + strconv.Itoa(ppl.ID) // community:id
	key := getKeyName(orderKey) + ":" + strconv.Itoa(ppl.ID)     // post:time:id
	// 60秒存在的话直接取，不存在重新计算
	if global.RDB.Exists(key).Val() < 1 {
		pipe := global.RDB.Pipeline()
		pipe.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, orderKey, cKey)
		// 60 秒超时
		pipe.Expire(key, 60*time.Second)
		_, err := pipe.Exec()
		if err != nil {
			return nil, err
		}
	}
	res, err = global.RDB.SMembers(cKey).Result()
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
	if curPos == prePos { //重复投票
		return ErrorRepeateVote
	}
	var temp float64
	if curPos > prePos {
		temp = 1
	} else {
		temp = -1
	}
	// 更新分数和 3 要进行 事务处理
	pipe := global.RDB.Pipeline()
	// 更新分数
	_, err := pipe.ZIncrBy(getKeyName(KeyZSetPostScore), pos*temp*perVoteScore, pi).Result()
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

func GetPostApproNum(idList []string) ([]int64, error) {
	// 使用 pipeline 减少请求次数，减少 RTT
	pipe := global.RDB.Pipeline()
	for _, id := range idList {
		pipe.ZCount(getKeyName(KeyZSetPostVotedPF+id), "1", "1").Result()
	}
	cmders, err := pipe.Exec()
	data := make([]int64, 0, len(cmders))
	if err != nil {
		return data, err
	}
	for _, comder := range cmders {
		v := comder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, err
}
