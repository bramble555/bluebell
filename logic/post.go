package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/global"
	"bluebell/models"
	"errors"
)

func CreatePost(p *models.Post) error {
	// 根据雪花算法获取postID
	p.PostID = global.Snflk.GetID()
	err := mysql.CreatePost(p)
	if err != nil {
		global.Log.Errorln("mysql create post error", err.Error())
		return err
	}
	err = redis.CreatePost(p.PostID, p.CommunityID)
	if err != nil {
		global.Log.Errorln("redis create post error", err.Error())
		return err
	}
	return nil
}

func GetPostDetail(p *models.Post) (*models.PostDetail, error) {
	err := mysql.GetPostDetail(p)
	if err != nil {
		global.Log.Errorf("logic GetPostDetail err post_id是 %d: %v", p.PostID, err)
		return nil, err
	}
	pd := &models.PostDetail{
		Post: p, // 显式初始化 Post 字段
	}
	// 获取社区详情
	cd, err := mysql.GetCommunityDetailByID(p.CommunityID)
	if err != nil {
		global.Log.Errorf("logic GetCommunityDetailByID err community_id是 %d: %v", p.CommunityID, err)
		return nil, err
	}
	pd.CommunityIntroduction = cd.Introduction
	pd.CommunityName = cd.Name
	// 获取用户详情
	ud, err := mysql.GetUserDetail(p.UserID)
	if err != nil {
		global.Log.Errorf("logic GetUserDetail err user_id是 %d: %v", p.UserID, err)
		return nil, err
	}
	pd.Username = ud.Username

	return pd, nil
}

func GetPostList(ppl *models.ParamPostList) ([]*models.PostDetail, error) {
	idList, err := redis.GetPostIDList(ppl)
	if err != nil {
		global.Log.Errorf("redis GetPostIDList2 error %s\n", err.Error())
		return nil, err
	}
	// 获取帖子赞同票数
	approvalNum, err := redis.GetPostApproNum(idList)
	if err != nil {
		global.Log.Errorf("redis GetPostApproNum error %s\n", err.Error())
		return nil, err
	}
	posts, err := mysql.GetPostList(idList)
	pds := make([]*models.PostDetail, 0, len(idList))
	global.Log.Debugln("idList长度为", len(idList))
	if err != nil {
		global.Log.Errorf("mysql GetPostIDList2 error %s\n", err.Error())
		return nil, err
	}
	global.Log.Debugln("posts为", posts)
	for i, p := range posts {
		pd := &models.PostDetail{
			Username:              "",
			Post:                  p,
			ApprovalNum:           approvalNum[i],
			CommunityName:         "",
			CommunityIntroduction: "",
		}
		// 获取社区详情
		cd, err := mysql.GetCommunityDetailByID(p.CommunityID)
		if err != nil {
			global.Log.Errorf("logic GetCommunityDetailByID err community_id是 %d: %v", p.CommunityID, err)
			return nil, err
		}
		pd.CommunityIntroduction = cd.Introduction
		pd.CommunityName = cd.Name
		// 获取用户详情
		ud, err := mysql.GetUserDetail(p.UserID)
		if err != nil {
			global.Log.Errorf("logic GetUserDetail err user_id是 %d: %v", p.UserID, err)
			return nil, err
		}
		pd.Username = ud.Username
		global.Log.Debugf("详细信息为%+v", pd)
		pds = append(pds, pd)
	}
	return pds, nil
}
func GetPostCommunityList(pcpl *models.ParamCommunityPostList) ([]*models.PostDetail, error) {
	idList, err := redis.GetCommuntiyPostIDList(pcpl)
	if err != nil {
		global.Log.Errorf("redis GetPostIDList2 error %s\n", err.Error())
		return nil, err
	}
	// 获取帖子赞同票数
	approvalNum, err := redis.GetPostApproNum(idList)
	if err != nil {
		global.Log.Errorf("redis GetPostApproNum error %s\n", err.Error())
		return nil, err
	}
	posts, err := mysql.GetPostList(idList)
	pds := make([]*models.PostDetail, 0, len(idList))
	global.Log.Debugln("idList长度为", len(idList))
	if err != nil {
		global.Log.Errorf("mysql GetPostIDList2 error %s\n", err.Error())
		return nil, err
	}
	global.Log.Debugln("posts为", posts)
	for i, p := range posts {
		pd := &models.PostDetail{
			Username:              "",
			Post:                  p,
			ApprovalNum:           approvalNum[i],
			CommunityName:         "",
			CommunityIntroduction: "",
		}
		// 获取社区详情
		cd, err := mysql.GetCommunityDetailByID(p.CommunityID)
		if err != nil {
			global.Log.Errorf("logic GetCommunityDetailByID err community_id是 %d: %v", p.CommunityID, err)
			return nil, err
		}
		pd.CommunityIntroduction = cd.Introduction
		pd.CommunityName = cd.Name
		// 获取用户详情
		ud, err := mysql.GetUserDetail(p.UserID)
		if err != nil {
			global.Log.Errorf("logic GetUserDetail err user_id是 %d: %v", p.UserID, err)
			return nil, err
		}
		pd.Username = ud.Username
		global.Log.Debugf("详细信息为%+v", pd)
		pds = append(pds, pd)
	}
	return pds, nil
}

// 投票有很多算法
// 本项目采用简单算法，投一票就 + 432 分，86400/200 -> 200张赞成票可以让你的帖子续一天 《redis实战》
/*
	投票的几种情况
	direction = 1：	之前没有投票，现在点赞	+ 432
			之前反对，现在点赞	+ 432 * 2
	direction = 0: 	之前点赞，现在取消	- 432
			之前反对，现在取消	+ 432
	direction = -1:	之前没有投票，现在反对 	-432
			之前点赞，现在反对	-432*2
*/
// curDirection > preDirection + 绝对值(curDirection - preDirection)
// curDirection < preDirection - 绝对值(curDirection - preDirection)

// 投票限制：每个帖子自发表之日后一星期，就不允许点赞和反对了
// 到期之后存入 MySQL ，然后删除 redis 保存的KeyZSetPostVotedPF

func VoteForPost(userID int, pv *models.ParamPostVote) error {
	// 判断pid是否存在
	ok, err := mysql.CheckPIDExist(pv.PostID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("查询的post_id 不存在")
	}
	global.Log.Debugln("uid", userID, "pid", pv.PostID, "正在投票，方向是", pv.Direction)
	return redis.VoteForPost(userID, pv.PostID, float64(pv.Direction))
}
